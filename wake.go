package wake

import (
	"time"

	"golang.org/x/sys/windows"
)

const (
	esSystemRequired   uintptr = 0x00000001
	esDisplayRequired  uintptr = 0x00000002
	// esAwayModeRequired uintptr = 0x00000040
	esContinuous       uintptr = 0x80000000
)

var (
	kernel32 *windows.LazyDLL
	setState *windows.LazyProc
	sleepChan chan bool
	wideAwake bool
)

func init() {
	kernel32 = windows.NewLazyDLL("kernel32.dll")
	setState = kernel32.NewProc("SetThreadExecutionState")
	sleepChan = make(chan bool)
	wideAwake = false
}

func StayAwake() {
	if wideAwake { return }
	pulse := time.NewTicker(5 * time.Minute)

	go func() {
		wideAwake = true
		defer backToSleep()
		for {
			select {
			case <-sleepChan:
				setState.Call(esContinuous)
				return
			case <-pulse.C:
				setState.Call(esSystemRequired)
			}
		}
	}()
}

func KeepScreenOn() {
	if wideAwake { return }
	pulse := time.NewTicker(5 * time.Minute)

	go func() {
		wideAwake = true
		defer backToSleep()
		for {
			select {
			case <-sleepChan:
				setState.Call(esContinuous)
				return
			case <-pulse.C:
				setState.Call(esDisplayRequired)
			}
		}
	}()
}

func IsWideAwake() bool {
	return wideAwake
}

func AllowSleep() {
	if !wideAwake { return }

	sleepChan <- true
}

func backToSleep() {
	wideAwake = false
}
