package main

import (
	"errors"
	"github.com/realHoangHai/logh"
)

func main() {
	err := errors.New("wrong")
	lo := logh.NewLogh(logh.DebugLevel, "hailth")
	lo.Infof("%s", err)
	// --------------------------------------------------
	logh.Errorf("error: %v", err)
	logh.Infof("info: success")
	logh.Warnf("warn: %v", err)
	logh.Debugf("debug: %v", err)
	logh.Fatalf("fatal: %v", err)
}
