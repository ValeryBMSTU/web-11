package usecase

type Usecase struct {
	defaultMsg string

	p  Provider
	jp JWTProvider
}

func NewUsecase(defaultMsg string, p Provider, jp JWTProvider) *Usecase {
	return &Usecase{
		defaultMsg: defaultMsg,
		p:          p,
		jp:         jp,
	}
}
