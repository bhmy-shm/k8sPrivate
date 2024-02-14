package controls

import (
	"github.com/bhmy-shm/gofks"
	"github.com/bhmy-shm/gofks/core/logx"
	"github.com/gin-gonic/gin"
)

type AccountCase struct {
}

func NewAccountController() *AccountCase {
	return &AccountCase{}
}

func (s *AccountCase) Build(gofk *gofks.Gofk) {

	account := gofk.Group("account")

	user := account.Group("user")
	user.Handle("POST", "/userDetail", s.UserDetail)
}

func (s *AccountCase) Name() string {
	return "userCase"
}

// ===========

func (s *AccountCase) UserDetail(ctx *gin.Context) {

	logx.Infof("this is user detail")
	logx.Info("user detail successful")
	ctx.JSON(200, gin.H{"message": "ok"})
}
