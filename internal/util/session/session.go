package session

import (
	"strconv"

	"github.com/ZYallers/fine/os/fctx"
	"github.com/ZYallers/fine/os/fsession"
	"github.com/ZYallers/fine/util/fphp"
	"github.com/gin-gonic/gin"
)

func LoginUserId(ctx *gin.Context) int {
	var vars map[string]interface{}
	if value, ok := ctx.Get(fctx.SessionDataKey); ok && value != nil {
		if s, ok := value.(string); ok && s != "" {
			vars = fphp.Unserialize(s)
		}
	} else {
		if s := fsession.Data(ctx, nil); s != "" {
			vars = fphp.Unserialize(s)
		}
	}
	if vars != nil {
		if userInfo, ok := vars["userinfo"].(map[string]interface{}); ok {
			if s, ok := userInfo["userid"].(string); ok {
				i, _ := strconv.Atoi(s)
				return i
			}
		}
	}
	return 0
}
