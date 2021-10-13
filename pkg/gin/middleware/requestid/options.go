package requestid

type Option func(*request)

type IdFunc func() string

type Cookie struct {
	MaxAge   int
	Path     string
	Host     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

type options struct {
	requestKey string
	idFunc     IdFunc
	cookie     Cookie
}

func WithRequestKey(key string) Option {
	return func(r *request) {
		r.options.requestKey = key
	}
}

func WithIdFunc(idFunc IdFunc) Option {
	return func(r *request) {
		r.options.idFunc = idFunc
	}
}

func WithCookie(cookie Cookie) Option {
	return func(r *request) {
		r.options.cookie = cookie
	}
}
