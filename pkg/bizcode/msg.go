package bizcode

var msg = map[int]string{
	Success: "ok",
	Error:   "error",
}

func Msg(code int) string {
	return msg[code]
}
