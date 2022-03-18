package main

import (
	"github.com/gin-gonic/gin"
)
import "github.com/satori/go.uuid"

func test(c *gin.Context) {
	c.JSON(MakeSuccessReturn("success"))
}

func (s *Service) Login(c *gin.Context) {
	account := new(Account)
	loginData(c, account)
	if account.Phone == "" || account.Pass == "" {
		c.JSON(MakeErrorReturn("无账号或密码"))
		return
	}
	account1 := new(Account)
	s.DB.Where(Account{Phone: account.Phone}).Find(&account1)
	if account1.Phone == account.Phone && account1.Pass == account.Pass {
		uuid := uuid.NewV4()
		account.Token = uuid.String()
		s.DB.Save(account)
		c.JSON(MakeSuccessReturn(uuid))
	} else {
		c.JSON(MakeErrorReturn("账号或者密码错误"))
	}

}
