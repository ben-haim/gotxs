// An ideomatic go wrapper for the SWIG-generated go interface
//
// This package sets up all the required objects, translates error conditions to
// the go way of doing things, and provides some documentation.
package gotxs

import "errors"

// this are the lower-level, SWIG-generated bindings
import "gotxs/opentxs"

// Some API methods go through an instance of this object.
var otme opentxs.OT_ME

// Create a new pseudonym in the local wallet.
// Crashes with OT_FAIL if keysize is invalid.
// Returns generated pseudonym id.
func CreatePseudonym(
	keybits int, nym_id_source, alt_location string) (string, error) {

	retval := otme.Create_pseudonym(keybits, nym_id_source, alt_location)

	if retval == "" {
		return "", errors.New("empty response")
	} else {
		return retval, nil
	}
}

// The initialization can be done automatically. This function is called when
// the module is imported first. It calls the init functions of the opentxs api
// and creates a reference to a `OTME` object.
func init() {
        if opentxs.OTAPI_WrapAppInit() == false {
                panic("error in OTAPI_WrapAppInit()")
        }
	
        if opentxs.OTAPI_WrapLoadWallet() == false {
                panic("error in OTAPI_WrapLoadWallet()")
        }

	otme = opentxs.NewOT_ME()
}
