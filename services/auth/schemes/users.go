package schemes

import (
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/gen/proto/auth"
	"github.com/Drozd0f/ttto-go/services/auth/db"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidData = errors.New("invalid data")

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required, validation.Length(4, 30)),
		validation.Field(&u.Password, validation.Required, validation.Length(4, 30)),
	)
}

func (u *User) EncryptPassword() error {
	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("generate from password: %w", err)
	}

	u.Password = string(bs)
	return nil
}

func (u *User) CheckPassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(u.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrInvalidData
		}
		return fmt.Errorf("compare hash and password: %w", err)
	}

	return nil
}

func (u *User) UserToCreateUserParam() db.CreateUserParams {
	return db.CreateUserParams{
		ID:       uuid.New(),
		Username: u.Username,
		Password: u.Password,
	}
}

func UserFromCreateUserRequest(cur *auth.CreateUserRequest) *User {
	return &User{
		Username: cur.Username,
		Password: cur.Password,
	}
}

func UserFromLoginUserRequest(lur *auth.LoginUserRequest) *User {
	return &User{
		Username: lur.Username,
		Password: lur.Password,
	}
}
