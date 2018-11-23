package edvmodel

import "github.com/olivere/elastic"

//EsSchool holds the school in elasticsearch
type EsSchool struct {
	ClientARN    *string               `json:"client_arn,omitempty"`
	SchoolPrefix *string               `json:"school_prefix,omitempty"`
	Name         *string               `json:"name,omitempty"`
	Suggest      *elastic.SuggestField `json:"suggest,omitempty"`
}
