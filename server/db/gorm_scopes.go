package db

import (
	"fmt"
	"math"
	"strings"

	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"

	"gorm.io/gorm"
)

func Paginate(value interface{}, v *pb.PaginationResponse, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if v.Limit > 0 || v.Page > 0 {
		var totalRows int64
		db.Model(value).Count(&totalRows)

		v.TotalRows = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(v.Limit)))
		v.TotalPages = int32(totalPages)
	}

	return func(db *gorm.DB) *gorm.DB {
		if v.Limit < 1 || v.Page < 1 {
			return db
		}

		offset := (v.Page - 1) * v.Limit
		if v != nil {
			return db.Limit(int(v.Limit)).Offset(int(offset))
		}
		return db
	}
}

func Sort(v *pb.Sort) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == nil || v.Column == "" {
			return db
		}
		v.Column = columnNameBuilder(v.Column)
		if v != nil {
			return db.Order(v.Column + " " + v.Direction)
		}
		return db
	}
}

func QueryScoop(v string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		query := strings.Split(v, ":")
		if len(query) < 2 {
			return db
		}

		columns := strings.Split(query[0], ",")
		expression := ""
		value := query[1]

		if len(query[1]) > 2 {
			if string(query[1][0:1]) == "%" {
				expression = string(query[1][0:2])
			}
		}

		switch expression {
		case "%%":
			value := "%" + string(query[1][2:len(query[1])]) + "%"
			db = queryColumnsLoop(db, columns, "LIKE", value)

		case "%!":
			value := "%" + string(query[1][2:len(query[1])]) + "%"
			db = queryColumnsLoop(db, columns, "ILIKE", value)

		case "":
			db = queryColumnsLoop(db, columns, "=", value)
		}

		return db
	}
}

func queryColumnsLoop(db *gorm.DB, columns []string, expresion string, value string) *gorm.DB {
	for _, s := range columns {
		s = columnNameBuilder(s)
		db = db.Or(fmt.Sprintf("%s %s ?", s, expresion), value)
	}
	return db
}

func FilterScoope(v string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		filters := strings.Split(v, ",")
		if len(filters) < 1 {
			return db
		}

		for _, s := range filters {
			filter := strings.Split(s, ":")
			if len(filter) == 2 {
				keyword := filter[1]
				expression := ""
				if len(filter[1]) > 2 {
					if string(filter[1][0:1]) == "%" {
						expression = string(filter[1][0:2])
					} else if string(filter[1][0:1]) == ">" || string(filter[1][0:1]) == "<" {
						expression = string(filter[1][0:1])
					}
				}

				column := columnNameBuilder(filter[0])
				if expression == "%%" {
					value := "%" + string(keyword[2:len(filter[1])]) + "%"
					db = db.Where(fmt.Sprintf("%s LIKE ?", column), value)
				} else if expression == "%!" {
					value := "%" + string(keyword[2:len(filter[1])]) + "%"
					db = db.Where(fmt.Sprintf("%s ILIKE ?", column), value)
				} else if expression == ">" || expression == "<" {
					if expression == "<" && filter[1][1:2] == ">" {
						expression = string(filter[1][0:2])
						value := string(keyword[2:len(filter[1])])
						db = db.Where(fmt.Sprintf("%s %s ?", column, expression), value)
					}
					if filter[1][1:2] == "=" {
						expression = string(filter[1][0:2])
						value := string(keyword[2:len(filter[1])])
						db = db.Where(fmt.Sprintf("%s %s ?", column, expression), value)
					} else {
						value := string(keyword[1:len(filter[1])])
						db = db.Where(fmt.Sprintf("%s %s ?", column, expression), value)
					}

				} else {
					value := keyword
					db = db.Where(fmt.Sprintf("%s = ?", column), value)
				}
			}
		}

		return db
	}
}

func columnNameBuilder(s string) string {
	if strings.Contains(s, "->") {
		nested := strings.Split(s, "->")
		s = ""
		for i, t := range nested {
			if i == 0 {
				s = fmt.Sprintf("\"%s\"", t)
			} else if i == len(nested)-1 {
				s = s + fmt.Sprintf("->>'%s'", t)
			} else {
				s = s + fmt.Sprintf("->'%s'", t)
			}
		}
	} else {
		s = fmt.Sprintf("\"%s\"", s)
	}

	return s
}
