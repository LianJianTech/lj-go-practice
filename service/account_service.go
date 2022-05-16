package service

import (
	"github.com/LianJianTech/lj-go-common/errno"
	"github.com/LianJianTech/lj-go-common/log"
	"github.com/LianJianTech/lj-go-common/util"
	"lj-go-practice/dao"
	"lj-go-practice/model"
	"lj-go-practice/pkg"
)

type AccountService struct {
	accountDao *dao.AccountDao
}

func NewAccountService() *AccountService {
	return &AccountService{dao.NewAccountDao()}
}

func (service *AccountService) LoginAccount(form *model.LoginForm) (*model.AccountResp, *errno.Errno) {
	account, err := service.accountDao.GetAccountByName(form.Name)
	if err != nil {
		return nil, errno.HandleError
	}
	if account.Name == "" {
		return nil, errno.HandleError
	}
	salt := account.Salt
	pwdCrypt := util.MD5(salt + form.Password)
	if pwdCrypt != account.PwdCrypt {
		return nil, errno.HandleError
	}

	token, err := signApiToken(account.ID, account.Name)
	if err != nil {
		return nil, errno.HandleError
	}
	resp := &model.AccountResp{
		Name:  account.Name,
		Token: token,
	}

	return resp, nil
}

func (service *AccountService) UpdateAccountPwd(form *model.UpdatePwdForm) *errno.Errno {
	account, err := service.accountDao.GetAccountByName(form.Name)
	if err != nil {
		return errno.HandleError
	}
	if account.Name == "" {
		return errno.HandleError
	}
	oldSalt := account.Salt
	oldPwdCrypt := util.MD5(oldSalt + form.OldPwd)
	if account.PwdCrypt != oldPwdCrypt {
		return errno.HandleError
	}
	newSalt := util.GetRandomString(8)
	newPwdCrypt := util.MD5(newSalt + form.NewPwd)
	if err := service.accountDao.UpdateAccountPwd(form.Name, newSalt, newPwdCrypt); err != nil {
		return errno.HandleError
	}
	return nil
}

func signApiToken(id uint64, username string) (string, error) {
	token, err := pkg.Sign(pkg.Context{ID: id, Username: username})
	if err != nil {
		log.Errorf(err, "signApiToken error")
		return "", nil
	}

	return token, nil
}
