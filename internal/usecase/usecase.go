package usecase

type UseCase interface {
	Execute() (interface{}, error)
}
