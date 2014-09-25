// Check that the gotxs/opentxs package compiled correctly.
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

// Make an API call and Cleanup() after that.
func (s *MySuite) TestCreatePseudonym(c *check.C) {
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

// Test basic message verification (only that it throws an error).
func (s *MySuite) TestMessageGetSuccess(c *check.C) {
	// It's non-trival to construct a valid message with signatures and
	// all. This just checks if we can call the function, that it
	// doesn't crash and returns an error (no signatures can be found).
	const invalidMessage = "invalid message"
	_, err := gotxs.MessageGetSuccess(invalidMessage)
	c.Assert(err, check.NotNil)
}

func (s *MySuite) TearDownSuite(c *check.C) {
	gotxs.Cleanup()
}
