package schemes

//SchemeConstant holds the scheme codes
type SchemeConstant struct {
	Iteacloud IteacloudScheme `json:"iteacloud"`
	Edvise    EdviseScheme    `json:"edvise"`
}

//Scheme predefined schemes
var Scheme = SchemeConstant{
	Iteacloud: IteacloudScheme{
		ClientCreate:          "ITC:CLIENT:CREATE",
		ClientClose:           "ITC:CLIENT:CLOSE",
		UserLogin:             "ITC:USER:LOGIN",
		UserLoginRefreshToken: "ITC:USER:LOGIN/REFRESH-TOKEN",
		UserActive:            "ITC:USER:ACTIVE",
		UserCreate:            "ITC:USER:CREATE",
		UserBatchCreate:       "ITC:USER:CREATE/BATCH",
		UserDelete:            "ITC:USER:DELETE",
		UserBatchDelete:       "ITC:USER:DELETE/BATCH",
	},
	Edvise: EdviseScheme{
		UserLogin:             "EDV:USER:LOGIN",
		UserLoginRefreshToken: "EDV:USER:LOGIN/REFRESH-TOKEN",
		UserActive:            "EDV:USER:ACTIVE",
		UserCreate:            "EDV:USER:CREATE",
		UserBatchCreate:       "EDV:USER:CREATE/BATCH",
		UserDelete:            "EDV:USER:DELETE",
		UserBatchDelete:       "EDV:USER:DELETE/BATCH",
	},
}
