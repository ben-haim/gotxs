package gotxs_test

import "testing"

import . "gotxs"

func TestCreateNym(t *testing.T) {
    retval, err := CreatePseudonym(1024, "", "")

    t.Logf("done", retval, err)
}
