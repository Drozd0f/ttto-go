package schemes

import (
	"testing"

	"github.com/Drozd0f/ttto-go/pkg/rande"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	t.Run("User validate", func(t *testing.T) {
		type testCase struct {
			name            string
			user            *User
			isValidateError bool
		}

		testCases := []testCase{
			{
				name: "Valid user",
				user: &User{
					ID:       uuid.New(),
					Username: rande.String(5),
					Password: rande.String(5),
				},
				isValidateError: false,
			},
			{
				name: "No username user",
				user: &User{
					ID:       uuid.New(),
					Password: rande.String(5),
				},
				isValidateError: true,
			},
			{
				name: "No password user",
				user: &User{
					ID:       uuid.New(),
					Username: rande.String(5),
				},
				isValidateError: true,
			},
			{
				name: "Short username",
				user: &User{
					ID:       uuid.New(),
					Username: rande.String(3),
					Password: rande.String(5),
				},
				isValidateError: true,
			},
			{
				name: "Long username",
				user: &User{
					ID:       uuid.New(),
					Username: rande.String(31),
					Password: rande.String(5),
				},
				isValidateError: true,
			},
			{
				name: "short password",
				user: &User{
					ID:       uuid.New(),
					Username: rande.String(5),
					Password: rande.String(3),
				},
				isValidateError: true,
			},
			{
				name: "long password",
				user: &User{
					ID:       uuid.New(),
					Username: rande.String(5),
					Password: rande.String(31),
				},
				isValidateError: true,
			},
		}

		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				if test.isValidateError {
					require.Error(t, test.user.Validate())
					return
				}

				require.NoError(t, test.user.Validate())
			})
		}

	})
	t.Run("Encrypt password", func(t *testing.T) {
		type testCase struct {
			name string
			user *User
		}

		testCases := []testCase{
			{
				name: "Min len password",
				user: &User{
					Password: rande.String(4),
				},
			},
			{
				name: "Max len password",
				user: &User{
					Password: rande.String(30),
				},
			},
		}

		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				require.NoError(t, test.user.EncryptPassword())
			})
		}
	})

	t.Run("Check password", func(t *testing.T) {
		password := rande.String(15)
		u := User{
			Password: password,
		}
		u1 := User{
			Password: password,
		}
		require.NoError(t, u1.EncryptPassword())
		require.NoError(t, u.CheckPassword(u1.Password))

	})
}
