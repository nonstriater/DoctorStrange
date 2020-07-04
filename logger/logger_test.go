package logger

import (
	"fmt"
	"testing"
	"time"
)

func Test_suffixName(t *testing.T) {
	now := time.Now()
	fmt.Print(now)
	r := RotateHour
	s := r.suffixName(now)
	if s == "20171010" {
		t.Errorf("suffixName error %v", s)
	}
}
