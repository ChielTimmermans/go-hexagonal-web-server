package user

type Servicer interface {
	Register(u *Service) (statusCode int, err error)
}
