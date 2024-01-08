// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table gf_user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:gf_user, do:true"`
	Id        interface{} //
	Name      interface{} // 用户名
	Phone     interface{} // 手机号
	Age       interface{} // 年龄
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
