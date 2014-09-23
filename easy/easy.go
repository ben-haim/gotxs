// This package contains all wrappers for the methods exposed by the OT_ME class
// Like in the gotxs package, we translate errors and provide some documentation
package easy

import "errors"

import _ "gotxs" // import for init() side effect
import "gotxs/opentxs"

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

// create a reference to a OT_ME() instance
// needs gotxs import which calls init()
func init() {
	otme = opentxs.NewOT_ME()
}
