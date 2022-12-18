package common

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func HTTPResponser(data interface{}, status_code int, status bool, message string) fiber.Map {
	return fiber.Map{
		"error":   status,
		"message": message,
		"data":    data,
	}
}

func GenNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(999999-100000) + 100000
}

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

func ValueTrim(value string) string {
	result := strings.TrimSpace(value)
	return result
}

func Placeholder(data []string) string {
	var columns []string
	for i := 1; i <= len(data); i++ {
		columns = append(columns, fmt.Sprintf("$%d", i))
	}
	return strings.Join(columns, ",")
}

func Column(data []string) string {
	var columns []string
	for i := 0; i < len(data); i++ {
		columns = append(columns, data[i])
	}
	return strings.Join(columns, ",")
}

func ThrowError(defination string) error {
	return errors.New(defination)
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

// CheckStringIfContains check a string if contains given param
func CheckStringIfContains(input_text string, search_text string) bool {
	CheckContains := strings.Contains(input_text, search_text)
	return CheckContains
}
