package itcjob

//Job holds the job
type Job struct {
	JobID       string `json:"job_id"`       //id of the job
	Principal   string `json:"principal"`    //arn of the requester
	Service     string `json:"service"`      //where the job comes from (e.g. importing user from excel, making report, etc)
	Bucket      string `json:"bucket"`       //bucket of event
	BucketKey   string `json:"bucket_key"`   //bucket key
	AccessToken string `json:"access_token"` //Access Token of the user
}
