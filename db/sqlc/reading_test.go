package db

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/VitorCdSouza/wattageMonitorApi/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomReading(t *testing.T) Reading {

	readingWattage := pgtype.Numeric{
		Int: big.NewInt(util.RandomInt(1000, 100000)),
		Exp: -2,
		NaN: false,
		Valid: true,
    }
	deviceId := pgtype.Int8{
		Int64: util.RandomInt(12, 21),
		Valid: true,
	}
	arg := CreateReadingParams{
		ReadingWattage: readingWattage,
		DeviceID:       deviceId,
	}

	reading, err := testQueries.CreateReading(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reading)

	require.Equal(t, arg.ReadingWattage, reading.ReadingWattage)
	require.Equal(t, arg.DeviceID, reading.DeviceID)

	require.NotZero(t, reading.ID)

	return reading
}

func TestCreateReading(t *testing.T) {
	createRandomReading(t)
}

func TestGetReading(t *testing.T) {
	reading1 := createRandomReading(t)
	reading2, err := testQueries.GetReading(context.Background(), reading1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, reading2)

	require.Equal(t, reading1.ID, reading2.ID)
	require.Equal(t, reading1.ReadingWattage, reading2.ReadingWattage)
	require.WithinDuration(t, reading1.ReadingHour, reading2.ReadingHour, time.Second)
	require.Equal(t, reading1.DeviceID, reading2.DeviceID)
}


func TestDeleteReading(t *testing.T) {
	reading1 := createRandomReading(t)
	_, err := testQueries.DeleteReading(context.Background(), reading1.ID)
	require.NoError(t, err)

	reading2, err := testQueries.GetReading(context.Background(), reading1.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, reading2)
}

func TestListReadings(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomReading(t)
	}

	arg := ListReadingParams{
		Limit:  5,
		Offset: 5,
	}

	readings, err := testQueries.ListReading(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, readings, 5)

	for _, reading := range readings {
		require.NotEmpty(t, reading)
	}
}
