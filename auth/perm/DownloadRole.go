package perm

import (
	"encoding/json"

	"github.com/kkesley/commonstruct/s3lib"
)

//DownloadRole to s3 input bucket
func DownloadRole(bucket string, region string, folder string) ([]RawPolicy, error) {
	//download the bytes from s3
	roleByte, err := s3lib.GetS3DocumentDefault(region, bucket, folder+"/role.json")
	if err != nil {
		return make([]RawPolicy, 0), err
	}
	var policies []RawPolicy
	if err := json.Unmarshal(roleByte, &policies); err != nil {
		return make([]RawPolicy, 0), err
	}
	return policies, nil
}
