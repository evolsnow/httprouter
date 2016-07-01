package httprouter

import "strconv"

//add functions to share values between middlewares
func (ps *Params) Set(key, value string) {
	if ps.ByName(key) == "" {
		add := Param{Key: key, Value: value}
		*ps = append(*ps, add)
	} else {
		for i := range *ps {
			if (*ps)[i].Key == key {
				(*ps)[i].Value = value
			}
		}
	}
}

func (ps Params) Get(key string) string {
	return ps.ByName(key)
}

func (ps Params) GetInt(key string) (value int) {
	value, err := strconv.Atoi(ps.ByName(key))
	if err != nil {
		return
	}
	return
}

func (ps Params) GetUint(key string) (value uint) {
	intValue, err := strconv.Atoi(ps.ByName(key))
	if err != nil {
		return
	}
	return uint(intValue)
}
