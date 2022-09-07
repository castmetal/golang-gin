package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	_common "golang-gin/src/core/domains/common"
	_security "golang-gin/src/core/domains/common/security"
)

type User struct {
	gorm.Model
	_common.EntityBase `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ID                 uuid.UUID                    `json:"id" bson:"_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	FirstName          string                       `json:"first_name" gorm:"type:varchar(60);column:first_name"`
	LastName           string                       `json:"last_name" gorm:"type:varchar(140);column:last_name"`
	UserName           string                       `json:"user_name" gorm:"type:varchar(90);unique;uniqueIndex;collumn:user_name"`
	Email              string                       `json:"email" gorm:"type:varchar(150);unique;uniqueIndex;column:email"`
	EncryptedPassword  *_security.EncryptedPassword `json:"encrypted_password" gorm:"embedded"`
	CreatedAt          time.Time                    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          time.Time                    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt          gorm.DeletedAt               `json:"deleted_at" gorm:"column:deleted_at"`
}

type UserProps struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()

	u.ID = id

	return nil
}

func NewUserEntity(props UserProps) (*User, error) {
	var user *User

	abstractEntity := _common.NewAbstractEntity(props.ID)

	if _common.IsNullOrEmpty(props.UserName) {
		return nil, _common.IsNullOrEmptyError("user_name")
	}

	actualDate := time.Now()

	user = &User{
		FirstName:         props.FirstName,
		LastName:          props.LastName,
		UserName:          props.UserName,
		Email:             props.Email,
		EncryptedPassword: _security.NewEncryptedPassword(props.Password),
		CreatedAt:         actualDate,
		UpdatedAt:         actualDate,
	}

	user.ID = abstractEntity.ID

	return user, nil
}
