// This package wraps all OTAPI_Wrap* methods.
// Wrappers for OT_ME methods can be found in gotxs/easy/easy.go
//
// The wrappers translate error conditions to the Go way of doing things, and
// provide some documentation.
package gotxs

// this are the lower-level, SWIG-generated bindings
import "github.com/monetas/gotxs/opentxs"



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
