package usecase

type Provider interface {
	GetCount() (int, error)
	AddCount(int) error
}
