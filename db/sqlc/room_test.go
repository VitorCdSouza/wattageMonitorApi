package db

import (
	"context"
	"testing"

	"github.com/VitorCdSouza/wattageMonitorApi/util"
	"github.com/stretchr/testify/require"
)

func createRandomRoom(t *testing.T) Room {
	roomName := util.RandomString(int(util.RandomInt(4, 7)))
	room, err := testQueries.CreateRoom(context.Background(), roomName)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, roomName, room.RoomName)

	require.NotZero(t, room.ID)

	return room
}

func TestCreateRoom(t *testing.T) {
	createRandomRoom(t)
}

func TestGetRoom(t *testing.T) {
	room1 := createRandomRoom(t)
	room2, err := testQueries.GetRoom(context.Background(), room1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, room1.RoomName, room2.RoomName)
}

func TestUpdateRoom(t *testing.T) {
	room1 := createRandomRoom(t)

	arg := UpdateRoomParams{ID: room1.ID, RoomName: room1.RoomName}

	room2, err := testQueries.UpdateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, room1.RoomName, room2.RoomName)
}

func TestDeleteRoom(t *testing.T) {
	room1 := createRandomRoom(t)
	_, err := testQueries.DeleteRoom(context.Background(), room1.ID)
	require.NoError(t, err)

	room2, err := testQueries.GetRoom(context.Background(), room1.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, room2)
}

func TestListRooms(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomRoom(t)
	}

	arg := ListRoomParams{
		Limit:  5,
		Offset: 5,
	}

	rooms, err := testQueries.ListRoom(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, rooms, 5)

	for _, room := range rooms {
		require.NotEmpty(t, room)
	}
}
