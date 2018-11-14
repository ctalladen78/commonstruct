package stdreq

import (
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

//GetListStdRequest get standard list from api gateway request
func GetListStdRequest(requestEvent events.APIGatewayProxyRequest, request *ListStdRequest) {
	if limit, ok := requestEvent.QueryStringParameters["limit"]; ok {
		limitInt, err := strconv.Atoi(limit)
		if err == nil {
			request.Limit = aws.Int(limitInt)
		}
	}

	if page, ok := requestEvent.QueryStringParameters["page"]; ok {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			request.Page = aws.Int(pageInt)
		}
	}

	tz := ""
	if val, ok := requestEvent.Headers["tz"]; ok {
		tz = val
		request.TZ = tz
	}

	if order, ok := requestEvent.QueryStringParameters["order"]; ok {
		request.Sort = make([]Sort, 0)
		for _, statement := range strings.Split(order, ",") {
			orderStack := strings.Split(statement, "=")
			if len(orderStack) > 1 {
				request.Sort = append(request.Sort, Sort{
					Field:     orderStack[0],
					Ascending: orderStack[1] == "asc",
				})
			}
		}
	}
	if search, ok := requestEvent.QueryStringParameters["search"]; ok {
		request.Search = aws.String(search)
	}
	if token, ok := requestEvent.QueryStringParameters["next_token"]; ok {
		request.NextToken = aws.String(token)
	}
}
