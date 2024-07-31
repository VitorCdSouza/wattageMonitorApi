// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateDevice(ctx context.Context, arg CreateDeviceParams) (Device, error)
	CreateRoom(ctx context.Context, roomName string) (Room, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteDevice(ctx context.Context, id int64) (Device, error)
	DeleteRoom(ctx context.Context, id int64) (Room, error)
	DeleteUser(ctx context.Context, id int64) (User, error)
	GetDevice(ctx context.Context, id int64) (Device, error)
	GetRoom(ctx context.Context, id int64) (Room, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListDevice(ctx context.Context, arg ListDeviceParams) ([]Device, error)
	ListRoom(ctx context.Context, arg ListRoomParams) ([]Room, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]User, error)
	UpdateDevice(ctx context.Context, arg UpdateDeviceParams) (Device, error)
	UpdateRoom(ctx context.Context, arg UpdateRoomParams) (Room, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)