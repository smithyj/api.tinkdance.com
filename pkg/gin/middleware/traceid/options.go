package traceid

type Option func(*trace)

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
	traceKey string
	idFunc   IdFunc
	cookie   Cookie
}

func WithTraceKey(key string) Option {
	return func(r *trace) {
		r.options.traceKey = key
	}
}

func WithIdFunc(idFunc IdFunc) Option {
	return func(r *trace) {
		r.options.idFunc = idFunc
	}
}

func WithCookie(cookie Cookie) Option {
	return func(r *trace) {
		r.options.cookie = cookie
	}
}
