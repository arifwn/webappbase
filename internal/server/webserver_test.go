package server

import (
	"os"
	"testing"
)

func TestFakeTestCase(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log(wd)
}
