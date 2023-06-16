package auth

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"gorm.io/gorm"

	"global-intf/internal/dto"
	"global-intf/internal/intf"
	"global-intf/internal/intfmocks"
	"global-intf/internal/model"
	"global-intf/pkg/server"
	"global-intf/pkg/util"
)

func TestHandlerLogin(t *testing.T) {
}

func TestServiceLogin(t *testing.T) {
	// define in function to prevent sharing type between test cases.
	type Input struct {
		ctx context.Context
		req *dto.LoginRequest

		// service dependencies
		userRepo      intf.UserRepo
		cryptoService intf.CryptoService

		// to prepair service dependecies base on input each times.
		provision func(t *testing.T, input *Input)
	}

	type Output struct {
		res *dto.LoginResponse
		err error
	}

	tcs := []*util.TestCase[Input, Output]{
		{
			Name: "it should login",
			Input: &Input{
				ctx: context.Background(),
				req: &dto.LoginRequest{
					Username: aws.String("hieutran@gmail.com"),
					Password: aws.String("123456"),
				},
				provision: func(t *testing.T, input *Input) {
					salt := "salt of password"
					password := []byte("123456")
					mockUserRepo := intfmocks.NewUserRepo(t)
					mockUserRepo.On("FindByEmail", input.ctx, *input.req.Username).Return(&model.User{
						Email:    *input.req.Username,
						Password: password,
						Salt:     []byte(salt),
					}, nil)

					mockCryptoService := intfmocks.NewCryptoService(t)
					mockCryptoService.On("ComparePassword", password, []byte(salt), []byte(*input.req.Password)).Return(nil)

					input.userRepo = mockUserRepo
					input.cryptoService = mockCryptoService
				},
			},
			Expected: &Output{
				res: &dto.LoginResponse{Success: true},
				err: nil,
			},
		},
		{
			Name: "it should return db error if userRepo get trouble when finding user by email",
			Input: &Input{
				ctx: context.Background(),
				req: &dto.LoginRequest{
					Username: aws.String("hieutran@gmail.com"),
					Password: aws.String("123456"),
				},
				provision: func(t *testing.T, input *Input) {
					mockUserRepo := intfmocks.NewUserRepo(t)
					mockUserRepo.On("FindByEmail", input.ctx, *input.req.Username).Return(nil, gorm.ErrPreloadNotAllowed)
					input.userRepo = mockUserRepo
				},
			},
			Expected: &Output{
				res: nil,
				err: server.ErrInternalDatabaseError.WithInternal(gorm.ErrPreloadNotAllowed),
			},
		},
		{
			Name: "it should return invalid credential error if userRepo return record not found error",
			Input: &Input{
				ctx: context.Background(),
				req: &dto.LoginRequest{
					Username: aws.String("hieutran@gmail.com"),
					Password: aws.String("123456"),
				},
				provision: func(t *testing.T, input *Input) {
					mockUserRepo := intfmocks.NewUserRepo(t)
					mockUserRepo.On("FindByEmail", input.ctx, *input.req.Username).Return(nil, gorm.ErrRecordNotFound)
					input.userRepo = mockUserRepo
				},
			},
			Expected: &Output{
				res: nil,
				err: ErrInvalidCredential,
			},
		},
		{
			Name: "it should return invalid	credential error if ComparePassword return error",
			Input: &Input{
				ctx: context.Background(),
				req: &dto.LoginRequest{
					Username: aws.String("hieutran@gmail.com"),
					Password: aws.String("123456"),
				},
				provision: func(t *testing.T, input *Input) {
					mockUserRepo := intfmocks.NewUserRepo(t)
					mockUserRepo.On("FindByEmail", input.ctx, *input.req.Username).Return(&model.User{
						Email:    *input.req.Username,
						Password: []byte("123456"),
						Salt:     []byte("salt of password"),
					}, nil)

					mockCryptoService := intfmocks.NewCryptoService(t)
					mockCryptoService.On("ComparePassword", []byte("123456"), []byte("salt of password"), []byte(*input.req.Password)).Return(errors.New("bad cipher text"))

					input.userRepo = mockUserRepo
					input.cryptoService = mockCryptoService
				},
			},
			Expected: &Output{
				res: nil,
				err: ErrInvalidCredential,
			},
		},
	}

	util.DoTestCases(t, tcs, func(t *testing.T, tc *util.TestCase[Input, Output]) {
		tc.Input.provision(t, tc.Input)

		authSrv := NewService(tc.Input.userRepo, tc.Input.cryptoService)
		res, err := authSrv.Login(tc.Input.ctx, tc.Input.req)

		assert.Equal(t, tc.Expected.res, res)
		assert.Equal(t, tc.Expected.err, err)

		// FIXME: not safe for concurrency testing.
		// cleanAfterEach(t, tc)
	})
}
