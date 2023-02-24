package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	w.Header().Set("Content-type", "application/json")
	grpcErrorCode := status.Code(err)
	httpErrorCode := runtime.HTTPStatusFromCode(grpcErrorCode)
	msg := status.Convert(err).Message()
	w.WriteHeader(httpErrorCode)

	if grpcErrorCode == codes.InvalidArgument {
		msg = err.Error()
	}

	log.Println("[Debug info] Error gateway: ", err)
	log.Println("[Debug info] Grpc Error code: ", grpcErrorCode)
	log.Println("[Debug info] Http Error code: ", httpErrorCode)

	body := &pb.ErrorBodyResponse{
		Error:   true,
		Code:    uint32(httpErrorCode),
		Message: msg,
	}

	jErr := json.NewEncoder(w).Encode(body)

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}

func httpResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code
	if vals := md.HeaderMD.Get("file-download"); len(vals) > 0 {

		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "file-download")

		w.Header().Set("Content-Disposition", md.HeaderMD.Get("Content-Disposition")[0])
		w.Header().Set("Content-Length", md.HeaderMD.Get("Content-Length")[0])

	}

	return nil
}
