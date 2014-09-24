// This is just a very basic test to see if the gotxs/opentxs package
// compiled correctly.
//
// Run 'go test -v'
package gotxs_test

import "testing"

import "github.com/monetas/gotxs" // for Cleanup()
import "github.com/monetas/gotxs/easy"

import check "gopkg.in/check.v1"

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type MySuite struct{}

var _ = check.Suite(&MySuite{})

// just make a basic api call and Cleanup() after that
func (s *MySuite) TestBasicApi(c *check.C) {
	keysize := 1024
	nymSource := ""
	altLocation := ""

	retval, err := easy.CreatePseudonym(keysize, nymSource, altLocation)

	if err == nil {
		c.Logf("created new pseudonym %s", retval)
	} else {
		c.Error("could not create new pseudoynm")
	}

}

func (s *MySuite) TearDownSuite(c *check.C) {
	gotxs.Cleanup()
}
