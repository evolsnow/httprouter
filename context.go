package httprouter

//add functions to share values between middlewares
func (ps *Params) Set(key, value string) {
	add := Param{Key: key, Value: value}
	*ps = append(*ps, add)
}

func (ps Params) Get(key string) string {
	return ps.ByName(key)
}
