package stdreq

//ListStdRequest holds the request for standard list request
type ListStdRequest struct {
	Search *string `json:"search"`
	Limit  *int    `json:"limit"`
	Page   *int    `json:"page"`
	Order  *string `json:"order"`
	StdRequest
}
