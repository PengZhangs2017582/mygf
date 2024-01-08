// Package user -----------------------------
// @file      : user.go
// @author    : Allen
// @contact   : 364438619@qq.com
// @time      : 2024/1/4 22:11
// -------------------------------------------
package user

import "github.com/gogf/gf/v2/frame/g"

type CommonAddUpdate struct {
	Name  string `json:"name" description:"用户名"`
	Phone string `json:"phone" description:"用户手机号"`
	Age   int    `json:"age" description:"用户年龄"`
}

// Add
type AddReq struct {
	g.Meta `path:"/add" method:"post"`
	Name   string `v:"required|length:5,20"`
	Phone  string `v:"required|length:11,11"`
}
type AddRes struct {
	Id uint
}

// Select
type GetListReq struct {
	g.Meta      `path:"/get-list" method:"get"`
	Id          uint
	Name        string
	Page        int `v:"min:0#分页号码错误" dc:"分页号码" d:"1"`
	Size        int `v:"max:50#分页数量最大100条" dc:"分页数量，最大100" d:"10"`
	OrderBy     string
	OrderByType int `v:"in:1,2#排序类型不合法" dc:"1-正序，2-逆序"`
}

type GetListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// Update
type UpdateReq struct {
	g.Meta `path:"/user-update" method:"post"`
	Id     uint `json:"id" v:"min:1#请选择需要修改的用户" dc:"用户id"`
	CommonAddUpdate
}
type UpdateRes struct {
	Id uint `json:"id"`
}

// Delete
type DelReq struct {
	g.Meta `path:"/user-del" method:"post"`
	Id     uint `json:"id" v:"min:1#请选择需要刪除的用户" dc:"用户id"`
}

type DelRes struct{}
