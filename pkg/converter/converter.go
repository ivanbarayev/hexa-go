package converter

import (
	"bytes"
	"github.com/goccy/go-json"
	"time"
)

// AnyToBytesBuffer Convert bytes to buffer helper
func AnyToBytesBuffer(i interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

func UnixToDate(unix int64) string {
	return time.UnixMilli(unix).UTC().String()
}
