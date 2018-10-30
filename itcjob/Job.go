package itcjob

//JobStatus holds the job status struct.
type JobStatus struct {
	JobID     string `json:"job_id"`    //id of the job
	Principal string `json:"principal"` //arn of the requester
	Status    string `json:"status"`    //status of the job
	JobClass  string `json:"job_class"` //where the job comes from (e.g. importing user from excel, making report, etc)
}
