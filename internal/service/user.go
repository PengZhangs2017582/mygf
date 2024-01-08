// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mygf/internal/model"
)

type (
	IUser interface {
		AddUser(ctx context.Context, in model.UserCreateInput) (out model.UserCreateOutput, err error)
		GetUserList(ctx context.Context, in model.UserListGetInput) (out []model.UserGetListOutput, total int, err error)
		UpdateUser(ctx context.Context, in model.UserUpdateInput) error
		DelUser(ctx context.Context, id uint) (err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
