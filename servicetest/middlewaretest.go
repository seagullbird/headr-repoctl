package servicetest

import (
	"bytes"
	"github.com/seagullbird/headr-repoctl/config"
	"os"
	"testing"
)

// MiddlewareTest describes a ServiceTest middleware.
type MiddlewareTest func(ServiceTest) ServiceTest

// LoggingMiddlewareTest takes a buffer as a dependency
// and returns  a ServiceTest middleware.
func LoggingMiddlewareTest(buffer *bytes.Buffer) MiddlewareTest {
	return func(nextTest ServiceTest) ServiceTest {
		return loggingMiddlewareTest{
			buffer: buffer,
			next:   nextTest,
		}
	}
}

type loggingMiddlewareTest struct {
	buffer *bytes.Buffer
	next   ServiceTest
}

func (mwt loggingMiddlewareTest) TestNewSite(t *testing.T) {
	mwt.next.TestNewSite(t)
	want := "method=NewSite siteID=0 err=\"invalid siteID\"\n" +
		"method=NewSite siteID=1 err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("NewSite log mismatches")
	}
}

func (mwt loggingMiddlewareTest) TestDeleteSite(t *testing.T) {
	mwt.next.TestDeleteSite(t)
	want := "method=DeleteSite siteID=0 err=\"invalid siteID\"\n" +
		"method=DeleteSite siteID=2 err=\"path does not exist\"\n" +
		"method=DeleteSite siteID=1 err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("DeleteSite log mismatches")
	}
}

func (mwt loggingMiddlewareTest) TestWritePost(t *testing.T) {
	mwt.next.TestWritePost(t)
	want := "method=WritePost siteID=0 filename= err=\"invalid siteID\"\n" +
		"method=WritePost siteID=1 filename=test-write-post.md err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("WritePost log mismatches")
	}
}

func (mwt loggingMiddlewareTest) TestRemovePost(t *testing.T) {
	mwt.next.TestRemovePost(t)
	want := "method=RemovePost siteID=0 filename= err=\"invalid siteID\"\n" +
		"method=RemovePost siteID=1 filename=test-write-post.md err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("RemovePost log mismatches")
	}
}

func (mwt loggingMiddlewareTest) TestReadPost(t *testing.T) {
	mwt.next.TestReadPost(t)
	want := "method=ReadPost siteID=0 filename= err=\"invalid siteID\"\n" +
		"method=ReadPost siteID=1 filename=test-write-post.md err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("ReadPost log mismatches")
	}
}

// EnvClearTestMiddleware is a dedicated middleware for testing
// It is the outermost middleware which does only environment clearing
func EnvClearTestMiddleware() MiddlewareTest {
	return func(next ServiceTest) ServiceTest {
		return envClearTestMiddleware{
			next: next,
		}
	}
}

type envClearTestMiddleware struct {
	next ServiceTest
}

func (tmw envClearTestMiddleware) TestNewSite(t *testing.T) {
	clearEnv(t)
	tmw.next.TestNewSite(t)
}

func (tmw envClearTestMiddleware) TestDeleteSite(t *testing.T) {
	clearEnv(t)
	tmw.next.TestDeleteSite(t)
}

func (tmw envClearTestMiddleware) TestWritePost(t *testing.T) {
	clearEnv(t)
	tmw.next.TestWritePost(t)
}

func (tmw envClearTestMiddleware) TestRemovePost(t *testing.T) {
	clearEnv(t)
	tmw.next.TestRemovePost(t)
}

func (tmw envClearTestMiddleware) TestReadPost(t *testing.T) {
	clearEnv(t)
	tmw.next.TestReadPost(t)
}

func clearEnv(t *testing.T) {
	if err := os.RemoveAll(config.SITESDIR); !(err == nil || os.IsNotExist(err)) {
		t.Fatalf("Removing SITESDIR failed: %v", err)
	}

	if err := os.MkdirAll(config.SITESDIR, 0644); err != nil {
		t.Fatalf("Creating SITESDIR failed: %v", err)
	}
}