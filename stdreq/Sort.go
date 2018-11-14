package stdreq

//Sort holds the sort values
type Sort struct {
	Field     string `json:"field"`
	Ascending bool   `json:"ascending"`
}
