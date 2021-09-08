package bizcode

var msg = map[int]string{
	Success: "ok",
	Fatal:   "error",
}

func Msg(code int) string {
	return msg[code]
}
