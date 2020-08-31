package utils

import "os"

func EnvDefaultVal(k1, defVal string) string {
	val, ok := os.LookupEnv(k1)
	if !ok {
		return defVal
	}
	return val
}
