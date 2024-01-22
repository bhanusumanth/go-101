package api_test

import (
	"testing"

	"github.com/bhanusumanth/go/cryptograsp/api"
)

func TestCEXGetRate(t *testing.T) {
	res, err := api.GetRate("")
	if err == nil {
		t.Error("Test Error : response received for no input : ", res)
	}
}
