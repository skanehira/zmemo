package logger

import (
	"errors"
	"testing"
	"zmemo/api/common"
)

func TestErrorf(t *testing.T) {
	Init()

	Errorf("faild test cause ", common.Wrap(errors.New("cannot open file")))
	Errorf("faild test %d", 10222103021093)
}
