// This package contains all wrappers for the methods exposed by the OT_ME class.
// Like in the gotxs package, we translate errors and provide some documentation.
package easy

import "errors"

import _ "github.com/monetas/gotxs" // import for init() side effect
import "github.com/monetas/gotxs/opentxs"

var otme opentxs.OT_ME

// CreateNym creates a new nym in the local wallet.
// keybits must be one of 1024, 2048, 4096 or 8192
// The arguments nymIdSource and altLocation can be empty strings.
// Returns the id of the created nym.
func CreateNym(keybits int, nymIdSource, altLocation string) (string, error) {

	result := otme.Create_nym(keybits, nymIdSource, altLocation)

	if result == "" {
		return "", errors.New("empty return value")
	}

	return result, nil
}

// RegisterNym takes a locally created Nym and registers it at the provided server.
// Returns the message from the server. This message is a SAMY hashed document
// containing an XML snippet and a signature. The snippet is an OTMessage with
// a '@createUserAccount' node.
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

// Create a reference to a OT_ME() instance.
// Needs gotxs import which calls init().
func init() {
	otme = opentxs.NewOT_ME()
}
