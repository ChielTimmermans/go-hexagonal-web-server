package middleware

type Servicer interface {
	AccessControl(token string, roles []uint8) (statusCode int, err error)
}
