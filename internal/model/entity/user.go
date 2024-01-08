// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        description:""`
	Name      string      `json:"name"      description:"用户名"`
	Phone     string      `json:"phone"     description:"手机号"`
	Age       uint        `json:"age"       description:"年龄"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
