package gotxs_test

import (
	"testing"

	"github.com/go-check/check"
	"github.com/monetas/gotxs" // for Cleanup()
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type MySuite struct{}

var _ = check.Suite(&MySuite{})

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
