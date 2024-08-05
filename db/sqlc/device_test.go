package db

import (
	"context"
	"testing"

	"github.com/VitorCdSouza/wattageMonitorApi/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomDevice(t *testing.T) Device {

	roomId := pgtype.Int8{
		Int64: util.RandomInt(10, 15),
		Valid: true,
	}
	userId := pgtype.Int8{
		Int64: util.RandomInt(10, 23),
		Valid: true,
	}
	arg := CreateDeviceParams{
		DeviceName: util.RandomString(int(util.RandomInt(4, 7))),
		RoomID:     roomId,
		UserID:     userId,	
	}

	device, err := testQueries.CreateDevice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, device)

	require.Equal(t, arg.DeviceName, device.DeviceName)
	require.Equal(t, arg.RoomID, device.RoomID)
	require.Equal(t, arg.UserID, device.UserID)

	require.NotZero(t, device.ID)

	return device
}

func TestCreateDevice(t *testing.T) {
	createRandomDevice(t)
}

func TestGetDevice(t *testing.T) {
	device1 := createRandomDevice(t)
	device2, err := testQueries.GetDevice(context.Background(), device1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, device2)

	require.Equal(t, device1.ID, device2.ID)
	require.Equal(t, device1.DeviceName, device2.DeviceName)
	require.Equal(t, device1.RoomID, device2.RoomID)
	require.Equal(t, device1.UserID, device2.UserID)
}

func TestUpdateDevice(t *testing.T) {
	device1 := createRandomDevice(t)

	arg := UpdateDeviceParams{ID: device1.ID, DeviceName: device1.DeviceName, RoomID: device1.RoomID, UserID: device1.UserID}

	device2, err := testQueries.UpdateDevice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, device2)

	require.Equal(t, device1.ID, device2.ID)
	require.Equal(t, device1.DeviceName, device2.DeviceName)
	require.Equal(t, device1.RoomID, device2.RoomID)
	require.Equal(t, device1.UserID, device2.UserID)
}

func TestDeleteDevice(t *testing.T) {
	device1 := createRandomDevice(t)
	_, err := testQueries.DeleteDevice(context.Background(), device1.ID)
	require.NoError(t, err)

	device2, err := testQueries.GetDevice(context.Background(), device1.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, device2)
}

func TestListDevices(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomDevice(t)
	}

	arg := ListDeviceParams{
		Limit:  5,
		Offset: 5,
	}

	devices, err := testQueries.ListDevice(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, devices, 5)

	for _, device := range devices {
		require.NotEmpty(t, device)
	}
}
