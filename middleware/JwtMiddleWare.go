package middleware

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/helper"
	"github.com/Haroxa/Integrated_documentation/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AuthMiddleware(c *gin.Context) {
	// 获取
	token := c.GetHeader("Authorization")
	// 前缀
	prefix := "Bearer"
	// 检验格式
	if token == "" || !strings.HasPrefix(token, prefix) {
		c.JSON(http.StatusUnauthorized, helper.ApiReturn(common.CodeError, "token不存在", nil))
		c.Abort()
		return
	}
	// 提取有效部分
	token = token[len(prefix)+1:]
	// 解析，验证
	UserId, err := helper.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeExpires, "权限不足", nil))
		c.Abort()
		return
	}
	// 检验 id
	_, err = model.GetUserById(UserId)
	if err != nil {
		log.Errorf("Invalid User_id %+v", errors.WithStack(err))
		c.Abort()
		return
	}
	// 将用户 id 写入上下文，方便读取
	c.Set("user_id", UserId)
	c.Next()
}
