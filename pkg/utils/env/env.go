package env

import (
	"errors"
	"os"
	"strconv"
)

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func EnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}

func EnvInt(key string) int {
	s, err := EnvStr(key)
	if err != nil {
		println(err.Error())
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

func EnvBool(key string) (bool, error) {
	s, err := EnvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}

func EnvFloat(key string) (float64, error) {
	s, err := EnvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}
