package easy_test

import (
	"testing"

	"github.com/go-check/check"
	"github.com/monetas/gotxs" // for Cleanup()
	"github.com/monetas/gotxs/easy"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type MySuite struct{}

var _ = check.Suite(&MySuite{})

// Make an API call and Cleanup() after that.
func (s *MySuite) TestCreateNym(c *check.C) {
	keysize := 1024
	nymSource := ""
	altLocation := ""

	nymId, err := easy.CreateNym(keysize, nymSource, altLocation)

	if err == nil {
		c.Logf("created new nym %s", nymId)
	} else {
		c.Errorf("could not create new pseudoynm: %s", err)
	}

}

func (s *MySuite) TearDownSuite(c *check.C) {
	gotxs.Cleanup()
}
