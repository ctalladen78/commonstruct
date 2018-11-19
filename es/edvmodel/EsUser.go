package edvmodel

import (
	"time"
)

//EsUser holds the user in elastic search
type EsUser struct {
	SchoolPrefix *string    `json:"school_prefix,omitempty"`
	Student      *EsStudent `json:"student,omitempty"`
	Parent       *EsParent  `json:"parent,omitempty"`

	Username  string  `json:"username,omitempty"`
	Principal *string `json:"principal,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`

	Active    *bool          `json:"active,omitempty"`
	ParentExt *UserParentExt `json:"parent_ext,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

//EsParent holds the parent in elastic search
type EsParent struct {
	FirstName          *string `json:"first_name"`
	LastName           *string `json:"last_name"`
	Phone              *string `json:"phone"`
	Email              *string `json:"email"`
	Gender             *string `json:"gender"`
	Religion           *string `json:"religion"`
	DateOfBirth        *string `json:"date_of_birth"`
	PlaceOfBirth       *string `json:"place_of_birth"`
	Relationship       *string `json:"relationship"`
	PasswordNeedChange *bool   `json:"password_need_change"`
}

//EsStudent holds the student in elastic search
type EsStudent struct {
	FirstName          *string      `json:"first_name"`
	LastName           *string      `json:"last_name"`
	Phone              *string      `json:"phone"`
	Email              *string      `json:"email"`
	Gender             *string      `json:"gender"`
	Religion           *string      `json:"religion"`
	DateOfBirth        *string      `json:"date_of_birth"`
	PlaceOfBirth       *string      `json:"place_of_birth"`
	Relationship       *string      `json:"relationship"`
	PasswordNeedChange *bool        `json:"password_need_change"`
	Medical            *UserMedical `json:"medical"`
	Address            *UserAddress `json:"address"`
}

//UserMedical holds the medical condition of the student
type UserMedical struct {
	MedicalCondition *string `json:"medical_condition,omitempty"`
}

//UserAddress holds the primary address of the student
type UserAddress struct {
	Address  *string `json:"address,omitempty"`
	State    *string `json:"state,omitempty"`
	City     *string `json:"city,omitempty"`
	PostCode *string `json:"post_code,omitempty"`
	Country  *string `json:"country,omitempty"`
}

//UserParentExt holds additional parent info
type UserParentExt struct {
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	Email        *string `json:"email"`
	Phone        *string `json:"phone"`
	Gender       *string `json:"gender"`
	Religion     *string `json:"religion"`
	DateOfBirth  *string `json:"date_of_birth"`
	PlaceOfBirth *string `json:"place_of_birth"`
	Relationship *string `json:"relationship"`
}
