package ginuser

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/component/hasher"
	"finnal-exam/modules/user/userbiz"
	"finnal-exam/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResetPassword(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.Param("password")
		res := c.MustGet(common.CurrentUser).(common.Requester)

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewUpdatePasswordBiz(store, md5)

		if err := biz.ResetPassword(c.Request.Context(), res.GetUserId(), password); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
