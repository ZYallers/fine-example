package valid

import (
	"reflect"
	"strconv"

	"github.com/ZYallers/fine/os/fctx"
	"github.com/ZYallers/fine/os/fsession"
	"github.com/ZYallers/fine/util/fphp"
	"github.com/ZYallers/fine/util/fvaild"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const (
	requiredLoginTagName   = "required_login"
	requiredLoginErrorText = "Please log in first"
)

func init() {
	v10 := fvaild.V10()
	_ = v10.Validate.RegisterValidation(requiredLoginTagName, requiredLoginValidator)
	_ = v10.Validate.RegisterTranslation(requiredLoginTagName, v10.Translator,
		func(u ut.Translator) error { return u.Add(requiredLoginTagName, requiredLoginErrorText, true) },
		func(u ut.Translator, fe validator.FieldError) string {
			t, _ := u.T(requiredLoginTagName, fe.Field())
			return t
		},
	)
}

func requiredLoginValidator(fl validator.FieldLevel) bool {
	ctxVal := fl.Top().Elem().FieldByName("Ctx")
	if ctxVal.Kind() == reflect.Invalid {
		return false
	}
	ctx := ctxVal.Interface().(*gin.Context)
	if ctx == nil {
		return false
	}
	loginUserId := getLoginUserId(ctx)
	if loginUserId == 0 {
		return false
	}
	if idVal := fl.Top().Elem().FieldByName("LoginUserId"); idVal.Kind() != reflect.Invalid {
		idVal.SetInt(int64(loginUserId))
	}
	return true
}

func getLoginUserId(ctx *gin.Context) int {
	var vars map[string]interface{}
	if value, ok := ctx.Get(fctx.SessionDataKey); ok && value != nil {
		if s, ok := value.(string); ok {
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
