package usecase_test

import (
	"context"
	"fmt"
	"reflect"

	"github.com/sapawarga/auth/mocks"
	"github.com/sapawarga/auth/mocks/testcases"
	"github.com/sapawarga/auth/usecase"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usecase", func() {
	var (
		mockRepo    *mocks.MockAuthI
		mockDecoder *mocks.MockJWToken
		auth        usecase.Service
	)

	BeforeEach(func() {
		actor := "test"
		mockSvc := gomock.NewController(GinkgoT())
		mockSvc.Finish()
		mockRepo = mocks.NewMockAuthI(mockSvc)
		mockDecoder = mocks.NewMockJWToken(mockSvc)
		auth = usecase.NewAuth(mockRepo, mockDecoder, actor)
	})

	// DECLARE UNIT TEST FUNCTION

	var GetCurrentLoginFromTokenLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetCurrentLoginData[idx]
		mockDecoder.EXPECT().ParsingToken(ctx, data.TokenRequest).Return(data.ResponseAfterDecodeToken.Result, data.ResponseAfterDecodeToken.Error).Times(1)
		mockRepo.EXPECT().GetActorCurrentLoginByUsername(ctx, data.UsernameRequest).Return(data.ResponseAfterGetData.Result, data.ResponseAfterGetData.Error).Times(1)
		resp, err := auth.GetCurrenrLoginFromToken(ctx, data.UsecaseParam)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		}
	}

	var unitTestLogic = map[string]map[string]interface{}{
		"GetCurrentLoginFromToken": {"func": GetCurrentLoginFromTokenLogic, "test_case_count": len(testcases.GetCurrentLoginData), "desc": testcases.Description()},
	}

	for _, val := range unitTestLogic {
		s := reflect.ValueOf(val["desc"])
		var arr []TableEntry
		for i := 0; i < val["test_case_count"].(int); i++ {
			fmt.Println(s.Index(i).String())
			arr = append(arr, Entry(s.Index(i).String(), i))
		}
		DescribeTable("Function ", val["func"], arr...)
	}
})
