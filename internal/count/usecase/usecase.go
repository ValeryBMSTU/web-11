package usecase

type Usecase struct {
	defaultMsg int

	p Provider
}

func NewUsecase(defaultMsg int, p Provider) *Usecase {
	return &Usecase{
		defaultMsg: defaultMsg,
		p:          p,
	}
}
