// This package contains all wrappers for the methods exposed by the OT_ME class
// Like in the gotxs package, we translate errors and provide some documentation
package easy

import "errors"

import _ "github.com/monetas/gotxs" // import for init() side effect
import "github.com/monetas/gotxs/opentxs"

var otme opentxs.OT_ME

// CreatePseudonym creates a new pseudonym in the local wallet.
// Keybits must be one of 1024, 2048, 4096 or 8192
// The arguments ny_id_source and alt_location are optional.
// Returns generated pseudonym id.
func CreatePseudonym(keybits int, nymIdSource, altLocation string) (string, error) {

	retval := otme.Create_pseudonym(keybits, nymIdSource, altLocation)

	if retval == "" {
		return "", errors.New("empty return value")
	}

	return retval, nil
}

func RegisterNym(serverID, nymID string) (string, error) {
	message := otme.Register_nym(serverID, nymID)
	if message == "" {
		return "", errors.New("unable to register pseudonym, empty response")
	}
	success := opentxs.OTAPI_WrapMessage_GetSuccess(message)
	if success == -1 {
		return message, errors.New("unable to register pseudonym, error response")
	}
	if success != 1 {
		return message, errors.New("unable to register pseudonym, failed response")
	}
	return message, nil
}

// create a reference to a OT_ME() instance
// needs gotxs import which calls init()
func init() {
	otme = opentxs.NewOT_ME()
}
