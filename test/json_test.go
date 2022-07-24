package test

import (
	"go-mock-server/src"
	"testing"
)

func Test_JSON(t *testing.T) {
	path, err := src.ParsePath()
	if err != nil {
		t.Log(err, "path error")
	}
	bytes, err := src.ReadJSON(path)
	if err != nil {
		t.Log(err, "read error")
	}
	m, err := src.DecodeJSONString(bytes)
	if err != nil {
		t.Log(err, "decode error")
	}
	t.Log(m)
}
