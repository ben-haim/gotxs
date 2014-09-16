package gotxs

import "errors"

import "gotxs/opentxs"

var otme opentxs.OT_ME

type apiCallError struct {
    method string
    message string
}

func (e *apiCallError) Error() string {
    return e.message
}

func CreatePseudonym(
    keybits int, nym_id_source, alt_location string) (string, error) {

    /*
    Create a new pseudonym in the local wallet.

    Crashes with OT_FAIL if keysize is invalid.

    Returns generated pseudonym id.
    */

    retval := otme.Create_pseudonym(keybits, nym_id_source, alt_location)

    if retval == "" {
        return "", errors.New("empty response")
    } else {
        return retval, nil
    }
}

func init() {
    opentxs.OTAPI_WrapAppInit()
    opentxs.OTAPI_WrapLoadWallet()

    otme = opentxs.NewOT_ME()
}
