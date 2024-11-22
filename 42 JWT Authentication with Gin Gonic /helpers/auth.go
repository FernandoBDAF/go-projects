package helpers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckUserType(ctx *gin.Context, role string) (err error) {
	userType := ctx.GetString("user_type")
	fmt.Println("userType:", userType)
	fmt.Println("ctx:", ctx.GetString("user_type"))
	err = nil
	if userType != role {
		return errors.New("unauthorized to access this resource: " + "userType:" + userType + " " + "role:" + role)
	}
	return err
}

func MatchUserTypeToUid(ctx *gin.Context, userId string) (err error) {
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		return errors.New("unauthorized to access this resource")
	}

	err = CheckUserType(ctx, userType)

	return err
}
