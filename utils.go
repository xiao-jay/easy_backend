package main

import "github.com/gin-gonic/gin"

func loginData(c *gin.Context, account *Account)  {
	account.Phone  = c.PostForm("phone")
	account.Pass = c.PostForm("pass")
	if account.Phone == "" || account.Pass == ""{
		// 为了可能传json也能接受到值
		err := c.ShouldBindJSON(&account)
		if err != nil {
			return
		}

	}
}