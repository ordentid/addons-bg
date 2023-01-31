package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	customAES "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/aes"
	authPb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/auth"

	svc "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
	svcConn       *svc.ServiceConnection
}

type UserClaims struct {
	jwt.StandardClaims
	UserType            string              `json:"user_type"`
	ProductRoles        []*ProductAuthority `json:"product_roles"`
	Authorities         []string            `json:"authorities"`
	EncryptedCompanyIDs string              `json:"company_ids"`
	P                   string              `json:"p"`
	E                   string              `json:"e"`
}

type CurrentUser struct {
	UserClaims
	FilterMe    string   `json:"filter_me"`
	StatusOrder []string `json:"status_order"`
	TaskFilter  string   `json:"task_filter"`
	UserID      string
	CompanyID   string
	CompanyIDs  []uint64
}

type VerifyTokenRes struct {
	IsValid      bool                `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	IsExpired    bool                `protobuf:"varint,2,opt,name=isExpired,proto3" json:"isExpired,omitempty"`
	UserID       uint64              `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
	Username     string              `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	UserType     string              `protobuf:"bytes,5,opt,name=userType,proto3" json:"userType,omitempty"`
	ProductRoles []*ProductAuthority `protobuf:"bytes,6,rep,name=productRoles,proto3" json:"productRoles,omitempty"`
}

type ProductAuthority struct {
	ProductName string   `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
	Authorities []string `protobuf:"bytes,2,rep,name=authorities,proto3" json:"authorities,omitempty"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration, svcConn *svc.ServiceConnection) *JWTManager {
	return &JWTManager{secretKey, tokenDuration, svcConn}
}

func (manager *JWTManager) Generate(username string, userID uint64, sessionID string, dateTime string) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (manager *JWTManager) GetMeFromJWT(ctx context.Context, accessToken string) (*CurrentUser, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md["authorization"]
		if len(values) > 0 {
			split := strings.Split(values[0], " ")
			accessToken = split[0]
			if len(split) > 1 {
				accessToken = split[1]
			}
		}

	}

	if accessToken == "" {
		logrus.Errorf("access token is empty")
		return nil, status.Error(codes.Unauthenticated, "session is empty")
	}

	userClaims, err := manager.Verify(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Session expired")
	}

	now := time.Now()

	fmt.Printf("token verify expired: %v|%v|%v", !(now.Unix() <= userClaims.ExpiresAt), now.String(), time.Unix(userClaims.ExpiresAt, 0).String())
	if !(now.Unix() <= userClaims.ExpiresAt) {
		return nil, status.Errorf(codes.Unauthenticated, "Session expired")
	}

	for _, v := range userClaims.ProductRoles {
		if v.ProductName == "User" {
			for _, j := range v.Authorities {
				data := strings.Split(j, ":")
				if data[1] != "-" {
					userClaims.Authorities = append(userClaims.Authorities, data[1])
					break
				}
			}
		}
	}

	currentUser := &CurrentUser{
		UserClaims: *userClaims,
	}
	// - Maker: 1. Draft, 2. Returned, 3. Pending, 4. Request for Delete, 5. Approved, 6. Rejected
	// - Signer: 1. Pending, 2. Request for Delete, 3. Approved, 4. Rejected
	if len(userClaims.Authorities) > 0 {
		switch strings.ToLower(userClaims.Authorities[0]) {
		case "maker":
			currentUser.StatusOrder = []string{"2", "3", "1", "6", "4", "5"}
			currentUser.FilterMe = "status:<>0,status:<>7"

		case "signer":
			currentUser.StatusOrder = []string{"1", "6", "4", "5"}
			currentUser.FilterMe = "status:<>0,status:<>2,status:<>3,status:<>7"

		default:
			return nil, status.Errorf(codes.PermissionDenied, "Authority Denied")
		}
	} else {
		return nil, status.Errorf(codes.PermissionDenied, "Authority Denied")
	}

	key := getEnv("JWT_AES_KEY", "Odj12345*12345678901234567890123")
	aes := customAES.NewCustomAES(key)

	currentUser.TaskFilter = ""
	if currentUser.UserType == "ca" || currentUser.UserType == "cu" {
		currentUser.TaskFilter = "data.user.companyID:"

		decrypted, err := aes.Decrypt(userClaims.EncryptedCompanyIDs)
		if err != nil {
			logrus.Errorf("[api.auth][func:VerifyToken][05] Failed to decrypt companyIDs: %v", err)
			return nil, status.Errorf(codes.Internal, "Server error")
		}

		if decrypted != "" {
			var ids []uint64
			err = json.Unmarshal([]byte(decrypted), &ids)
			if err != nil {
				logrus.Errorf("[api.auth][func:VerifyToken][06] Failed to unmarshal companyIDs: %v", err)
				return nil, status.Errorf(codes.Internal, "Server error")
			}

			currentUser.CompanyIDs = ids

			for i, v := range ids {
				if i == 0 {
					currentUser.TaskFilter = currentUser.TaskFilter + fmt.Sprintf("%d", v)
				} else {
					currentUser.TaskFilter = currentUser.TaskFilter + fmt.Sprintf(",%d", v)
				}
			}
		}
	}

	if userClaims.P != "" {
		currentUser.UserID, err = aes.Decrypt(userClaims.P)
		if err != nil {
			logrus.Errorf("[api.auth][func:VerifyToken][05] Failed to decrypt Principal: %v", err)
			return nil, status.Errorf(codes.Internal, "Server error")
		}
	}

	if userClaims.E != "" {
		currentUser.CompanyID, err = aes.Decrypt(userClaims.E)
		if err != nil {
			logrus.Errorf("[api.auth][func:VerifyToken][05] Failed to decrypt Entity: %v", err)
			return nil, status.Errorf(codes.Internal, "Server error")
		}
	}

	return currentUser, nil
}

func (manager *JWTManager) GetMeFromAuthService(ctx context.Context, accessToken string) (*VerifyTokenRes, error) {
	authClient := manager.svcConn.AuthServiceClient()

	dataUser, err := authClient.VerifyToken(ctx, &authPb.VerifyTokenReq{
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, err
	}
	if dataUser == nil {
		return nil, status.Errorf(codes.Aborted, "Failed To Get Data User")
	}

	user := &VerifyTokenRes{
		IsValid:   dataUser.IsValid,
		IsExpired: dataUser.IsExpired,
		UserID:    dataUser.UserID,
		Username:  dataUser.Username,
		UserType:  dataUser.UserType,
	}

	for _, v := range dataUser.ProductRoles {
		role := &ProductAuthority{
			ProductName: v.ProductName,
			Authorities: v.Authorities,
		}
		user.ProductRoles = append(user.ProductRoles, role)
	}
	return user, nil
}

func (manager *JWTManager) GetUserMD(ctx context.Context) (metadata.MD, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if len(md["user-userid"]) > 0 {
			return md, nil
		}

		ctx = metadata.NewOutgoingContext(context.Background(), md)

	}

	// Make RPC using the context with the metadata.
	var trailer metadata.MD

	authClient := manager.svcConn.AuthServiceClient()

	_, err := authClient.SetMe(ctx, &authPb.VerifyTokenReq{}, grpc.Trailer(&trailer))
	if err != nil {
		logrus.Errorln("[jwtManager][func:GetUserMD][01] Failed to get user from auth service: ", err)
		return nil, err
	}
	md = metadata.Join(md, trailer)

	return md, nil
}

type UserData struct {
	UserID         uint64   `json:"userID"`
	Username       string   `json:"username"`
	CompanyID      uint64   `json:"companyID"`
	CompanyName    string   `json:"companyName"`
	UserType       string   `json:"userType"`
	Authorities    []string `json:"authorities"`
	GroupIDs       []uint64 `json:"groupIDs"`
	RoleIDs        []uint64 `json:"roleIDs"`
	SessionID      string   `json:"sessionID"`
	DateTime       string   `json:"dateTime"`
	TokenCreatedAt string   `json:"tokenCreatedAt"`
	IdToken        string   `json:"idToken"`
}

func (manager *JWTManager) GetMeFromMD(ctx context.Context) (user *UserData, md metadata.MD, err error) {
	md, err = manager.GetUserMD(ctx)
	if err != nil {
		return nil, nil, err
	}

	user = &UserData{}
	user.UserID, err = strconv.ParseUint(md["user-userid"][0], 10, 64)
	if err != nil {
		logrus.Errorln("[jwtManager][func:GetMeFromMD][01] Failed to parse userID: ", err)
		return nil, nil, status.Errorf(codes.Internal, "Error Internal")
	}
	user.CompanyID, err = strconv.ParseUint(md["user-companyid"][0], 10, 64)
	if err != nil {
		logrus.Errorln("[jwtManager][func:GetMeFromMD][02] Failed to parse companyID: ", err)
		return nil, nil, status.Errorf(codes.Internal, "Error Internal")
	}

	user.Username = md["user-username"][0]
	user.CompanyName = md["user-companyname"][0]
	user.UserType = md["user-usertype"][0]

	user.Authorities = strings.Split(md["user-authorities"][0], ",")

	ids := strings.Split(md["user-groupids"][0], ",")
	for _, v := range ids {
		if len(v) > 0 {
			id, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				logrus.Errorf("[jwtManager][func:GetMeFromMD][03] Failed to parse groupID: %s, error: %v ", ids, err)
				return nil, nil, status.Errorf(codes.Internal, "Error Internal")
			}
			user.GroupIDs = append(user.GroupIDs, id)
		}
	}

	ids = strings.Split(md["user-roleids"][0], ",")
	for _, v := range ids {
		if len(v) > 0 {
			id, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				logrus.Errorf("[jwtManager][func:GetMeFromMD][04] Failed to parse roleID: %s, error: %v ", ids, err)
				return nil, nil, status.Errorf(codes.Internal, "Error Internal")
			}
			user.RoleIDs = append(user.RoleIDs, id)
		}
	}

	user.SessionID = md["user-sessionid"][0]
	user.DateTime = md["user-datetime"][0]
	user.TokenCreatedAt = md["user-tokencreatedat"][0]
	user.IdToken = md["user-idtoken"][0]
	// user.Fcm = md["user-fcm"][0]

	return user, md, nil
}

func (manager *JWTManager) GetProductAuthority(md metadata.MD, productName string) ([]string, error) {
	var authorities []string
	productName = strings.Replace(productName, ":", "_", -1)
	productName = strings.ToLower(productName)
	productName = fmt.Sprintf("user-product-%s", productName)

	if len(md[productName]) > 0 {
		result := strings.Split(md[productName][0], ",")
		if len(result) > 0 {
			authorities = result
		}
	}

	return authorities, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
