package testcases

import (
	"errors"

	"github.com/sapawarga/auth/model"
)

type ResponseActorLogin struct {
	Result *model.Actor
	Error  error
}

type GetCurrentLoginFromToken struct {
	Description              string
	UsecaseParam             string
	TokenRequest             string
	UsernameRequest          string
	ResponseAfterDecodeToken ResponseActorLogin
	ResponseAfterGetData     ResponseActorLogin
	UsecaseResponse          ResponseActorLogin
}

var result = &model.Actor{
	ID:        1,
	Username:  "actor",
	RoleLabel: "staffProv",
	LastLoginAt: map[string]interface{}{
		"hai": "test",
	},
}

// GetCurrentLoginData ...
var GetCurrentLoginData = []GetCurrentLoginFromToken{
	{
		Description:     "success_get_actor",
		UsecaseParam:    "token",
		TokenRequest:    "token",
		UsernameRequest: "actor",
		ResponseAfterDecodeToken: ResponseActorLogin{
			Result: result,
			Error:  nil,
		},
		ResponseAfterGetData: ResponseActorLogin{
			Result: result,
			Error:  nil,
		},
		UsecaseResponse: ResponseActorLogin{
			Result: result,
			Error:  nil,
		},
	}, {
		Description:     "failed_validate_token",
		UsecaseParam:    "token",
		TokenRequest:    "token",
		UsernameRequest: "token",
		ResponseAfterDecodeToken: ResponseActorLogin{
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
		ResponseAfterGetData: ResponseActorLogin{
			Result: nil,
			Error:  nil,
		},
		UsecaseResponse: ResponseActorLogin{
			Result: result,
			Error:  errors.New("something_went_wrong"),
		},
	}, {
		Description:     "failed_get_actor_from_db",
		UsecaseParam:    "token",
		TokenRequest:    "token",
		UsernameRequest: "actor",
		ResponseAfterDecodeToken: ResponseActorLogin{
			Result: result,
			Error:  nil,
		},
		ResponseAfterGetData: ResponseActorLogin{
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
		UsecaseResponse: ResponseActorLogin{
			Result: result,
			Error:  errors.New("something_went_wrong"),
		},
	},
}

// DescriptionGetCurrentLogin ...
func DescriptionGetCurrentLogin() []string {
	var arr = []string{}
	for _, data := range GetCurrentLoginData {
		arr = append(arr, data.Description)
	}
	return arr
}
