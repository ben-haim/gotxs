// This is just a very basic test to see if the gotxs/opentxs package
// compiled correctly. 
//
// Run 'go test -v'
package gotxs_test

import "testing"

import "gotxs"

// just make a basic api call and Cleanup() after that
func TestBasicApi(t *testing.T) {
	keysize := 1024
	nym_source := ""
	alt_location := ""

	retval, err := gotxs.CreatePseudonym(keysize, nym_source, alt_location)

	if err == nil {
		t.Logf("created new pseudonym %s", retval)
	} else {
		t.Error("could not create new pseudoynm")
	}

	gotxs.Cleanup()
}
