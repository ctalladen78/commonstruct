package jwtidentity

//TokenRequest request object for generating JWT
type TokenRequest struct {
	IsRoot     bool     `json:"IsRoot"`
	UserARN    string   `json:"UserARN"`
	RoleARN    []string `json:"RoleARN"`
	ClientARN  string   `json:"ClientARN"`
	ClientName string   `json:"ClientName"`
	FirstName  *string  `json:"FirstName"`
	LastName   *string  `json:"LastName"`
	Username   string   `json:"Username"`
	Groups     []string `json:"Groups"`
	Device     string   `json:"Device"`
	IsExternal bool     `json:"IsExternal"`
	Service    string   `json:"Service"`
}
