package Service

import (
	"GinStart/Repository"
	"GinStart/Service/sms"
	"context"
	"fmt"
	"math/rand"
)

type CodeService struct {
	repo Repository.CodeRepository
	sms  sms.Service
}

// biz代表一个业务，使用这个字段来区别不同业务使用这个服务
func (s *CodeService) Send(ctx context.Context, biz, phone string) error {
	code := s.generate()
	err := s.repo.Set(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	const tplID = "100000000"
	return s.sms.Send(ctx, tplID, []string{code}, phone)
}

func (s *CodeService) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	ok, err := s.repo.Verify(ctx, biz, phone, inputCode)
	if err == Repository.ErrVerifySendToMany {
		//对外屏蔽验证次数过多的错误，告诉调用者，就是不对
		return false, err
	}
	return ok, nil
}

func (s *CodeService) generate() string {
	//	生成0-999999的数字
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}