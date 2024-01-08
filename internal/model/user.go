package model

import (
	"mygf/internal/model/entity"
)

type UserCreateUpdateBase struct {
	Name  string
	Phone string
	Age   int
}
type UserCreateInput struct {
	Name  string
	Phone string
}

type UserCreateOutput struct {
	Id uint
}

type UserListGetInput struct {
	Id          uint
	Page        int    // 分页号码
	Size        int    // 分页数量，最大50
	OrderBy     string // 排序字段
	OrderByType int    // 排序类型(1-正序，2-逆序)
}

//
// type UserGetListOutput struct {
// 	List  []UserItem `json:"list" description:"列表"`
// 	Page  int        `json:"page" description:"分页码"`
// 	Size  int        `json:"size" description:"分页数量"`
// 	Total int        `json:"total" description:"数据总数"`
// }

type UserGetListOutput struct {
	Id    uint
	Name  string `json:"name" description:"用戶名"`
	Phone string `json:"phone" description:"手机号"`
	Age   int    `json:"age" description:"年齡"`
}

type UserItem struct {
	entity.User
}

type UserUpdateInput struct {
	UserCreateUpdateBase
	Id uint
}

type UserDelInput struct {
	Id uint
}
