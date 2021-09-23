package bizcode

var msg = map[int]string{
	Success:        "ok",
	Fatal:          "error",
	RequestTimeout: "request timeout",
}

func Msg(code int) string {
	return msg[code]
}
