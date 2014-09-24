// This package wraps all OTAPI_Wrap* methods.
// Wrappers for OT_ME methods can be found in gotxs/easy/easy.go
//
// The wrappers translate error conditions to the Go way of doing things, and
// provide some documentation.
package gotxs

import "errors"

// this are the lower-level, SWIG-generated bindings
import "github.com/monetas/gotxs/opentxs"

const (
	// these constants are common return values
	// defined in OTAPI_Exec.cpp
	otError = -1
	otFalse = 0
	otTrue  = 1
)

// MessageGetSuccess checks if a received message is valid.
// A message is valid if XML decoding succeeds and valid signatures
// are found.
// This is the first level of verification for incoming messages.
// Crashes the process if an empty message is passed.
// TODO: the actual verification algorithm can be found in
// OTContract::ParseRawFile(). Figure out what happens in detail.
func MessageGetSuccess(message string) (bool, error) {
	result := opentxs.OTAPI_WrapMessage_GetSuccess(message)

	if result == otError {
		return false, errors.New("error in MessageGetSuccess()")
	} else if result == otFalse {
		return false, nil
	} else if result == otTrue {
		return true, nil
	}

	return false, errors.New("unknown return value")
}

// The initialization can be done automatically. This function is called when
// the module is imported first. It calls the init functions of the opentxs api
// and creates a reference to a `OTME` object.
func init() {
	if opentxs.OTAPI_WrapAppInit() == false {
		panic("error in OTAPI_WrapAppInit()" +
			"- forgot gotxs.Cleanup()?")
	}

	if opentxs.OTAPI_WrapLoadWallet() == false {
		panic("error in OTAPI_WrapLoadWallet()" +
			"- forgot gotxs.Cleanup()?")
	}
}

// Should be called when the program exits. Unfortunately there is not `atexit`
// mechanism in Go that can call this function automatically
func Cleanup() {
	opentxs.OTAPI_WrapAppCleanup()
}
