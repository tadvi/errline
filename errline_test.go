package errline

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {
	tests := []struct {
		in       error
		expected error
	}{
		{nil, nil},
		{
			errors.New("random error"),
			&withFileLine{errors.New("random error"),
				"errline_test.go", 24},
		},
	}

	for _, tt := range tests {
		actual := Wrap(tt.in)
		if !equal(actual, tt.expected) {
			t.Error("Failed")
			t.Fail()
		}
	}
}

func TestShortFilename(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"???", "???"},
		{"filename.go", "filename.go"},
		{"hello/filename.go", "filename.go"},
		{"main/hello/filename.go", "filename.go"},
	}

	for _, tt := range tests {
		actual := getShortFilename(tt.in)
		if strings.Compare(actual, tt.expected) != 0 {
			t.Fail()
		}
	}
}

func TestFormat(t *testing.T) {
	errRand := errors.New("random err")
	errWithFileLine := withFileLine{errRand, "file.go", 23}
	tests := []struct {
		e       withFileLine
		inState fmt.State
		inVerb  rune

		expected string
	}{
		{errWithFileLine, fakeFmtState{}, 'v', "file.go:23: random err"},
		{errWithFileLine, fakeFmtState{}, 's', "random err"},
		{errWithFileLine, fakeFmtState{}, 'q', "\"random err\""},
	}

	for _, tt := range tests {
		tt.e.Format(tt.inState, tt.inVerb)
		if strings.Compare(string(buf), tt.expected) != 0 {
			t.Fail()
		}
	}
}

// Main test entrypoint
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// Utilities
var buf []byte

type fakeFmtState struct{}

func (fakeFmtState) Write(b []byte) (n int, err error) {
	buf = b
	return len(string(b)), nil
}

func (fakeFmtState) Width() (wid int, ok bool) {
	return -1, false
}

func (fakeFmtState) Precision() (prec int, ok bool) {
	return -1, false
}

func (fakeFmtState) Flag(c int) bool {
	return c == '+'
}

func equal(e1, e2 error) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	err1, ok := e1.(*withFileLine)
	if !ok {
		return false
	}
	err2, ok := e2.(*withFileLine)
	if !ok {
		return false
	}

	return strings.Compare(err1.file, err2.file) == 0 &&
		strings.Compare(err1.Error(), err2.Error()) == 0 &&
		err1.line == err1.line
}
