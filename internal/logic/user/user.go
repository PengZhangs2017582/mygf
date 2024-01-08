package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"mygf/internal/dao"
	"mygf/internal/model"
	"mygf/internal/service"
)

type (
	sUser struct {
	}
)

func New() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(New())
}

func (s *sUser) AddUser(ctx context.Context, in model.UserCreateInput) (out model.UserCreateOutput, err error) {
	id, err := dao.User.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserCreateOutput{Id: uint(id)}, err

	// return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	// 	_, err = dao.User.Ctx(ctx).Data(do.User{
	// 		Name:  in.Name,
	// 		Phone: in.Phone,
	// 	}).Insert()
	// 	return err
	//
	// })

}
func (s *sUser) GetUserList(ctx context.Context, in model.UserListGetInput) (out []model.UserGetListOutput, total int, err error) {
	out = make([]model.UserGetListOutput, 0)
	m := dao.User.Ctx(ctx)
	// Where
	if in.Id > 0 {
		m = m.Where(dao.User.Columns().Id, in.Id)
	}
	// Order by
	if len(in.OrderBy) > 0 && in.OrderByType > 0 {
		if in.OrderByType == 1 {
			m = m.OrderAsc(in.OrderBy)
		} else if in.OrderByType == 1 {
			m = m.OrderDesc(in.OrderBy)
		}
	}
	err = m.Limit((in.Page-1)*in.Size, in.Size).ScanAndCount(&out, &total, true)
	return
}

func (s *sUser) UpdateUser(ctx context.Context, in model.UserUpdateInput) error {
	_, err := dao.User.Ctx(ctx).Data(in).FieldsEx(dao.User.Columns().Id).Where(dao.User.Columns().Id, in.Id).Update()
	return err
}

func (s *sUser) DelUser(ctx context.Context, id uint) (err error) {
	_, err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Id: id,
	}).Delete()

	if err != nil {
		return err
	}
	return
}
