package testcases

import (
	"errors"

	"github.com/sapawarga/auth/model"
)

type ResponseGetUserDetail struct {
	Result *model.UserDetail
	Error  error
}

type GetUserDetailByUsername struct {
	Description      string
	UsecaseParams    string
	RepositoryParams string
	MockUsesace      ResponseGetUserDetail
	MockRepository   ResponseGetUserDetail
}

var username = "john_doe"
var userDetail = &model.UserDetail{
	ID:         1,
	Username:   "john_doe",
	Email:      "john_doe@email.com",
	RoleLabel:  "staffRW",
	VillageID:  2454,
	DistrictID: 344,
	RegencyID:  22,
}

// GetUserDetailByUsernameData ...
var GetUserDetailByUsernameData = []GetUserDetailByUsername{
	{
		Description:      "success_get_user_detail",
		UsecaseParams:    username,
		RepositoryParams: username,
		MockUsesace: ResponseGetUserDetail{
			Result: userDetail,
			Error:  nil,
		},
		MockRepository: ResponseGetUserDetail{
			Result: userDetail,
			Error:  nil,
		},
	}, {
		Description:      "failed_get_user_detail",
		UsecaseParams:    username,
		RepositoryParams: username,
		MockRepository: ResponseGetUserDetail{
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
		MockUsesace: ResponseGetUserDetail{
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
	},
}

// DescriptionGetUserDetail ...
func DescriptionGetUserDetail() []string {
	var arr = []string{}
	for _, data := range GetUserDetailByUsernameData {
		arr = append(arr, data.Description)
	}
	return arr
}
