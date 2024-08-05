package db

import (
	"context"
	"testing"

	"github.com/VitorCdSouza/wattageMonitorApi/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		UserEmail:    util.RandomEmail(4),
		UserPassword: util.RandomPassword(6),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserEmail, user.UserEmail)
	require.Equal(t, arg.UserPassword, user.UserPassword)

	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.UserEmail, user2.UserEmail)
	require.Equal(t, user1.UserPassword, user2.UserPassword)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{ID: user1.ID, UserEmail: user1.UserEmail, UserPassword: user1.UserPassword}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.UserEmail, user2.UserEmail)
	require.Equal(t, user1.UserPassword, user2.UserPassword)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	_, err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUserParams{
		Limit: 5,
		Offset: 5,
	}

	users, err := testQueries.ListUser(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, users := range users {
		require.NotEmpty(t, users)
	}
}