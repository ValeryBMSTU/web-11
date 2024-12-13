package usecase

type Usecase struct {
	defaultCnt int

	p Provider
}

func NewUsecase(defaultCnt int, p Provider) *Usecase {
	return &Usecase{
		defaultCnt: defaultCnt,
		p:          p,
	}
}
