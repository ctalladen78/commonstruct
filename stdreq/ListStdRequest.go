package stdreq

//ListStdRequest holds the request for standard list request
type ListStdRequest struct {
	Search    *string `json:"search"`
	Limit     *int    `json:"limit"`
	Page      *int    `json:"page"`
	NextToken *string `json:"next_token"`
	Sort      []Sort  `json:"sort"`
	StdRequest
}
