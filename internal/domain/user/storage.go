package user

type Storager interface {
	Register(u *Service) (statusCode int, err error)
}
