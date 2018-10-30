package stdreq

import (
	"math"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/jinzhu/gorm"
)

//DBPageStdQuery standard limit / offset pagination
func DBPageStdQuery(query *gorm.DB, request *ListStdRequest, count int) *gorm.DB {
	if request.Limit == nil {
		request.Limit = aws.Int(10)
	}
	if request.Page == nil {
		request.Page = aws.Int(1)
	}
	query = query.Limit(*request.Limit)
	if *request.Page > 1 {
		offset := (*request.Page - 1) * *request.Limit
		if count <= offset {
			offset = int(math.Floor(float64(count) / float64(*request.Limit)))
			*request.Page = int(math.Ceil(float64(count) / float64(*request.Limit)))
		}
		query = query.Offset(offset)
	}
	return query
}
