package entities

import (
	"time"
)

type Auth struct {
	AuthId      int64  `json:"auth_id"`
	Lang        string `json:"lang"`
	Src         int8   `json:"src,omitempty"`
	ParentId    int64  `json:"parent_id,omitempty"`
	ManagerId   int64  `json:"manager_id,omitempty"`
	AccountType int8   `json:"acc_type,omitempty"`
	UserType    int8   `json:"user_type,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	UserTitle   string `json:"user_title"`
	UserName    string `json:"user_name"`
	IsDemo      bool   `json:"is_demo,omitempty"`
	UniqueId    string `json:"unique_id"`
	Status      int8   `json:"status,omitempty"`
}

type AuthHR struct {
	AuthId      int64     `json:"auth_id"`
	Lang        string    `json:"lang"`
	Src         int8      `json:"src,omitempty"`
	ParentId    int64     `json:"parent_id,omitempty"`
	AccountType int8      `json:"acc_type,omitempty"`
	UserType    int8      `json:"user_type,omitempty"`
	UserGender  int8      `json:"user_gender,omitempty"`
	CompanyName string    `json:"company_name,omitempty"`
	UserTitle   string    `json:"user_title"`
	UserName    string    `json:"user_name"`
	IsDemo      bool      `json:"is_demo,omitempty"`
	StaffId     int64     `json:"staff_id,omitempty"`
	UniqueId    string    `json:"unique_id"`
	Status      int8      `json:"status,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	ExpireAt    time.Time `json:"expire_at,omitempty"`
}

type AuthCandidate struct {
	AuthId     int64     `json:"auth_id"`
	Lang       string    `json:"lang"`
	Src        int8      `json:"src,omitempty"`
	ManagerId  int64     `json:"manager_id,omitempty"`
	UserType   int8      `json:"user_type,omitempty"`
	UserGender int8      `json:"user_gender,omitempty"`
	UserTitle  string    `json:"user_title"`
	UserName   string    `json:"user_name"`
	StaffId    int64     `json:"staff_id,omitempty"`
	UniqueId   string    `json:"unique_id"`
	Status     int8      `json:"status,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
	ExpireAt   time.Time `json:"expire_at,omitempty"`
}

type LoginReq struct {
	Src        int8   `json:"src" validate:"required,min=1,max=5"`
	UserType   int8   `json:"user_type" validate:"required,min=1,max=5"`
	UserName   string `json:"user_name" validate:"required,min=10,max=20"`
	UserPass   string `json:"user_pass" validate:"required,min=8,max=16"`
	VerifyCode int64  `json:"verify_code,omitempty"`
}

type RegisterReq struct {
	Src         int8   `json:"src" validate:"required,min=1,max=5"`
	UserType    int8   `json:"user_type" validate:"required,min=1,max=5"`
	UserTitle   string `json:"user_title" validate:"required"`
	CompanyName string `json:"company_name,omitempty"`
	UserName    string `json:"user_name" validate:"required,min=10,max=50"`
	UserPass    string `json:"user_pass,omitempty"`
	UserPhone   string `json:"user_phone" validate:"required,min=10,max=20"`
	VerifyCode  int64  `json:"verify_code,omitempty"`
}

type Forget struct {
	Src      int8   `json:"src" validate:"required"`
	UserType int8   `json:"user_type" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}

type UserChecker struct {
	Src      int8   `json:"src" validate:"required"`
	UserType int8   `json:"user_type" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}

type Verify struct {
	Src      int8   `json:"src"`
	UserName string `json:"user_name"`
}

type ContactData struct {
	AuthId      int64  `json:"auth_id,omitempty"`
	UserType    int8   `json:"user_type,omitempty"`
	TypeOf      string `json:"type_of"`
	Description string `json:"description,omitempty"`
	Definition  string `json:"defination"`
	IsDefault   bool   `json:"is_default,omitempty"`
	Status      int8   `json:"status,omitempty"`
}

type HandlerResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
