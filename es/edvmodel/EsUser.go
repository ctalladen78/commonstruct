package edvmodel

import (
	"time"

	"bitbucket.org/edvise/edv-platform-api/db/model/dynamodel"
)

//EsUser holds the user in elastic search
type EsUser struct {
	SchoolPrefix *string    `json:"school_prefix,omitempty"`
	Student      *EsStudent `json:"student,omitempty"`
	Parent       *EsParent  `json:"parent,omitempty"`

	Username  string  `json:"username,omitempty"`
	Principal *string `json:"principal,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`

	Active    *bool                    `json:"active,omitempty"`
	ParentExt *dynamodel.UserParentExt `json:"parent_ext,omitempty"`

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
	FirstName          *string                `json:"first_name"`
	LastName           *string                `json:"last_name"`
	Phone              *string                `json:"phone"`
	Email              *string                `json:"email"`
	Gender             *string                `json:"gender"`
	Religion           *string                `json:"religion"`
	DateOfBirth        *string                `json:"date_of_birth"`
	PlaceOfBirth       *string                `json:"place_of_birth"`
	Relationship       *string                `json:"relationship"`
	PasswordNeedChange *bool                  `json:"password_need_change"`
	Medical            *dynamodel.UserMedical `json:"medical"`
	Address            *dynamodel.UserAddress `json:"address"`
}
