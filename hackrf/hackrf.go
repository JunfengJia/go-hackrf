/*
Package hackrf provides an interface to the HackRF SDR hardware.

This package wraps libhackrf using cgo.
*/
package hackrf

// #cgo darwin CFLAGS: -I/usr/local/include
// #cgo darwin LDFLAGS: -L/usr/local/lib
// #cgo LDFLAGS: -lhackrf
// #include <libhackrf/hackrf.h>
import "C"
import (
	"errors"
	"fmt"
)

var (
	ErrInvalidParam        = errors.New("hackrf: invalid param")
	ErrNotFound            = errors.New("hackrf: not found")
	ErrBusy                = errors.New("hackrf: busy")
	ErrNoMem               = errors.New("hackrf: no mem")
	ErrLibUSB              = errors.New("hackrf: libusb error")
	ErrThread              = errors.New("hackrf: thread error")
	ErrStreamingThreadErr  = errors.New("hackrf: streaming thread error")
	ErrStreamingStopped    = errors.New("hackrf: streaming stopped")
	ErrStreamingExitCalled = errors.New("hackrf: streaming exit called")
	ErrOther               = errors.New("hackrf: other error")
)

type ErrUnknown int

func (e ErrUnknown) Error() string {
	return fmt.Sprintf("hackrf: unknown error %d", int(e))
}

// Init must be called once at the start of the program.
func Init() error {
	return toError(nil, C.hackrf_init())
}

// Exit should be called once at the end of the program.
func Exit() error {
	return toError(nil, C.hackrf_exit())
}

func toError(d *Device, r C.int) error {
	if r == C.HACKRF_SUCCESS {
		return nil
	}
	var err error
	switch r {
	case C.HACKRF_ERROR_INVALID_PARAM:
		err = ErrInvalidParam
	case C.HACKRF_ERROR_NOT_FOUND:
		err = ErrNotFound
	case C.HACKRF_ERROR_BUSY:
		err = ErrBusy
	case C.HACKRF_ERROR_NO_MEM:
		err = ErrNoMem
	case C.HACKRF_ERROR_LIBUSB:
		err = ErrLibUSB
	case C.HACKRF_ERROR_THREAD:
		err = ErrThread
	case C.HACKRF_ERROR_STREAMING_THREAD_ERR:
		err = ErrStreamingThreadErr
	case C.HACKRF_ERROR_STREAMING_STOPPED:
		err = ErrStreamingStopped
	case C.HACKRF_ERROR_STREAMING_EXIT_CALLED:
		err = ErrStreamingExitCalled
	case C.HACKRF_ERROR_OTHER:
		err = ErrOther
	default:
		err = ErrUnknown(int(r))
	}
	if d != nil {
		d.err = err
	}

	return err
}
