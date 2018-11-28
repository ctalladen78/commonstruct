package edvmodel

import "time"

//EsClass holds the class in elasticsearch
type EsClass struct {
	ARN        *string `json:"arn,omitempty"`
	ClientARN  *string `json:"client_arn,omitempty"`
	Code       *string `json:"code,omitempty"`
	CustomCode *string `json:"custom_code,omitempty"`
	Name       *string `json:"name,omitempty"`

	IsBoundedByCourse *bool `json:"is_bounded_by_course,omitempty"` //class.code == course.code

	RetiredCodes     []string   `json:"retired_codes,omitempty"` //for original classes. To keep track the old codes
	IsRetired        *bool      `json:"is_retired,omitempty"`
	RetiredBatchTime *time.Time `json:"retired_batch_time,omitempty"`

	InheritedFromClassCode *string `json:"inherited_from_class_code,omitempty"` //for inherited class, this is the old class code. i.e. old class is retired

	Courses  []string `json:"courses,omitempty"`
	Students []string `json:"students,omitempty"`

	//CustomTeacher holds the teacher for a specific course. This will override all courses teachers
	CustomTeacher map[string][]string `json:"custom_teacher,omitempty"`

	//Teachers hold the user arn of teachers
	//This can be used for home teachers for this class. This field can also be used if every courses is taught per class basis,
	Teachers []string `json:"teachers,omitempty"`

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}
