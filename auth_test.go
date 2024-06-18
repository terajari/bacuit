package bacuit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Email:           "agus@gmail.com",
				Username:        "agus",
				Password:        "bastard423",
				ConfirmPassword: "bastard423",
			},
			err: nil,
		},
		{
			name: "email invalid",
			input: RegisterInput{
				Email:           "agus",
				Username:        "agus",
				Password:        "bastard423",
				ConfirmPassword: "bastard423",
			},
			err: ErrInvalidEmail,
		},
		{
			name: "username too short",
			input: RegisterInput{
				Email:           "agus@gmail.com",
				Username:        "a",
				Password:        "bastard423",
				ConfirmPassword: "bastard423",
			},
			err: ErrInvalidUsername,
		},
		{
			name: "password invalid",
			input: RegisterInput{
				Email:           "agus@gmail.com",
				Username:        "agus",
				Password:        "te",
				ConfirmPassword: "te",
			},
			err: ErrInvalidPassword,
		},
		{
			name: "password not confirmed",
			input: RegisterInput{
				Email:           "agus@gmail.com",
				Username:        "agus",
				Password:        "bastard423",
				ConfirmPassword: "bastars423",
			},
			err: ErrPasswordMismatch,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestRegisterInput_Sanitize(t *testing.T) {
	get := RegisterInput{
		Email:           "AGUS@GMAIL.COM      ",
		Username:        "AGUS ",
		Password:        "bastard423",
		ConfirmPassword: "bastard423",
	}

	want := RegisterInput{
		Email:           "agus@gmail.com",
		Username:        "AGUS",
		Password:        "bastard423",
		ConfirmPassword: "bastard423",
	}

	get.Sanitize()

	require.Equal(t, want, get)
	require.Equal(t, want, get)
}
