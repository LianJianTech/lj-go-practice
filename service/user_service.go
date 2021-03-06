package service

import (
	"github.com/LianJianTech/lj-go-common/errno"
	"lj-go-practice/dao"
	"lj-go-practice/model"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{dao.NewUserDao()}
}

func (service *UserService) QueryUsers(req *model.UserReq) (*model.UserResp, *errno.Errno) {
	users, total, err := service.userDao.GetUsers(req)
	if err != nil {
		return nil, errno.HandleError
	}
	resp := &model.UserResp{
		Rows:  users,
		Total: total,
	}
	return resp, nil
}

func (service *UserService) AddUser(form *model.UserForm) *errno.Errno {
	if err := service.userDao.AddUser(form); err != nil {
		return errno.HandleError
	}
	return nil
}

func (service *UserService) UpdateUser(form *model.UserForm) *errno.Errno {
	if err := service.userDao.UpdateUser(form); err != nil {
		return errno.HandleError
	}
	return nil
}
