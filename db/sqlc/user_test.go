package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		UserEmail:    "aaa@gmail.com",
		UserPassword: "1234",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserEmail, user.UserEmail)
	require.Equal(t, arg.UserPassword, user.UserPassword)

	require.NotZero(t, user.ID)
}	
