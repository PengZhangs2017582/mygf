// Package user -----------------------------
// @file      : user.go
// @author    : Allen
// @contact   : 364438619@qq.com
// @time      : 2024/1/4 21:36
// -------------------------------------------
package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"mygf/api/user"
	"mygf/internal/model"
	"mygf/internal/service"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) AddUser(ctx context.Context, req *user.AddReq) (res *user.AddRes, err error) {
	data := model.UserCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().AddUser(ctx, data)

	if err != nil {
		return nil, err
	}
	return &user.AddRes{Id: out.Id}, nil
}

func (c *Controller) GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error) {
	listRes, total, err := service.User().GetUserList(ctx, model.UserListGetInput{
		Id:          req.Id,
		Page:        req.Page,
		Size:        req.Size,
		OrderBy:     req.OrderBy,
		OrderByType: req.OrderByType,
	})

	if err != nil {
		return nil, err
	}
	return &user.GetListRes{
		List:  listRes,
		Page:  req.Page,
		Size:  req.Size,
		Total: total,
	}, nil
}

func (c *Controller) UpdateUser(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error) {
	data := model.UserUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	_ = service.User().UpdateUser(ctx, data)
	return &user.UpdateRes{Id: req.Id}, nil
}

func (c *Controller) DelUser(ctx context.Context, req *user.DelReq) (res *user.DelRes, err error) {
	err = service.User().DelUser(ctx, req.Id)
	return
}

func (c *Controller) Get(r *ghttp.Request) {
	md := g.Model("gf_user")
	user, err := md.Fields("id,name").Where("id", 3).All()
	if err != nil {
		return
	}
	r.Response.WriteJson(user)
}

func (c *Controller) Post(r *ghttp.Request) {
	r.Response.Writeln("Add")
}

func (c *Controller) Put(r *ghttp.Request) {
	r.Response.Writeln("Update")
}

func (c *Controller) Delete(r *ghttp.Request) {
	r.Response.Writeln("Delete")
}
