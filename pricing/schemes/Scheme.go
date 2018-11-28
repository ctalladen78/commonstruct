package schemes

//SchemeConstant holds the scheme codes
type SchemeConstant struct {
	Iteacloud IteacloudScheme `json:"iteacloud"`
	Edvise    EdviseScheme    `json:"edvise"`
}

//Scheme predefined schemes
var Scheme = SchemeConstant{
	Iteacloud: IteacloudScheme{
		UserLogin:       "ITC:USER:LOGIN",
		UserActive:      "ITC:USER:ACTIVE",
		UserCreate:      "ITC:USER:CREATE",
		UserBatchCreate: "ITC:USER:CREATE/BATCH",
	},
	Edvise: EdviseScheme{
		UserLogin:       "EDV:USER:LOGIN",
		UserActive:      "EDV:USER:ACTIVE",
		UserCreate:      "EDV:USER:CREATE",
		UserBatchCreate: "EDV:USER:CREATE/BATCH",
	},
}
