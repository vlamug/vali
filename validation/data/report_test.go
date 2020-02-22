package data

import (
	"bytes"
	"testing"

	. "gopkg.in/check.v1"
)

func TestReportTest(t *testing.T) { TestingT(t) }

type ReportSuite struct{}

var _ = Suite(&ReportSuite{})

func (s ReportSuite) TestAdd(c *C) {
	buf := &bytes.Buffer{}
	report := NewWriterReport(buf)
	report.Add("first_message")
	report.Add("second_message")
	report.Add("third_message")

	_ = report.Publish()

	c.Assert(buf.String(), Equals, "first_message\nsecond_message\nthird_message\n")
}

func (s ReportSuite) TestPublish(c *C) {
	buf := &bytes.Buffer{}
	report := NewWriterReport(buf)

	report.Add("first_message")
	report.Add("second_message")

	err := report.Publish()

	c.Assert(err, Equals, nil)
	c.Assert(buf.String(), Equals, "first_message\nsecond_message\n")
}
