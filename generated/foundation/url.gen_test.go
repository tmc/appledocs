// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError demonstrates how to create a URL instance using NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError.
// Initializes a newly created NSURL that points to a location specified by resolving bookmark data.
func ExampleNewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError() {
	_ = foundation.NewURLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError(
		nil, // bookmarkData unsafe.Pointer
		nil, // options unsafe.Pointer
		nil, // relativeURL unsafe.Pointer
		nil, // isStale unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewURLWithString demonstrates how to create a URL instance using NewURLWithString.
// Initializes an NSURL object with a provided URL string.
func ExampleNewURLWithString() {
	_ = foundation.NewURLWithString(
		"URLString", // URLString string
	)
	// Output:
}

// ExampleNewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL demonstrates how to create a URL instance using NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL.
// Initializes a URL object with a C string representing a local file system path.
func ExampleNewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL() {
	_ = foundation.NewURLFileURLWithFileSystemRepresentationIsDirectoryRelativeToURL(
		nil, // path unsafe.Pointer
		false, // isDir bool
		nil, // baseURL unsafe.Pointer
	)
	// Output:
}

// ExampleNewURLFileURLWithPathIsDirectory demonstrates how to create a URL instance using NewURLFileURLWithPathIsDirectory.
// Initializes a newly created NSURL referencing the local file or directory at  .
func ExampleNewURLFileURLWithPathIsDirectory() {
	_ = foundation.NewURLFileURLWithPathIsDirectory(
		"path", // path string
		false, // isDir bool
	)
	// Output:
}

// ExampleNewURLFromPasteboard demonstrates how to create a URL instance using NewURLFromPasteboard.
// Reads an NSURL object off of the specified pasteboard.
func ExampleNewURLFromPasteboard() {
	_ = foundation.NewURLFromPasteboard(
		nil, // pasteBoard unsafe.Pointer
	)
	// Output:
}

// ExampleNewURLWithSchemeHostPath demonstrates how to create a URL instance using NewURLWithSchemeHostPath.
// Initializes a newly created NSURL with a specified scheme, host, and path.
func ExampleNewURLWithSchemeHostPath() {
	_ = foundation.NewURLWithSchemeHostPath(
		"scheme", // scheme string
		"host", // host string
		"path", // path string
	)
	// Output:
}

// ExampleNewURLWithStringEncodingInvalidCharacters demonstrates how to create a URL instance using NewURLWithStringEncodingInvalidCharacters.
// Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
func ExampleNewURLWithStringEncodingInvalidCharacters() {
	_ = foundation.NewURLWithStringEncodingInvalidCharacters(
		"URLString", // URLString string
		false, // encodingInvalidCharacters bool
	)
	// Output:
}

// ExampleNewURLWithStringRelativeToURL demonstrates how to create a URL instance using NewURLWithStringRelativeToURL.
// Initializes an NSURL object with a base URL and a relative string.
func ExampleNewURLWithStringRelativeToURL() {
	_ = foundation.NewURLWithStringRelativeToURL(
		"URLString", // URLString string
		nil, // baseURL unsafe.Pointer
	)
	// Output:
}

// ExampleNewURLAbsoluteURLWithDataRepresentationRelativeToURL demonstrates how to create a URL instance using NewURLAbsoluteURLWithDataRepresentationRelativeToURL.
func ExampleNewURLAbsoluteURLWithDataRepresentationRelativeToURL() {
	_ = foundation.NewURLAbsoluteURLWithDataRepresentationRelativeToURL(
		nil, // data unsafe.Pointer
		nil, // baseURL unsafe.Pointer
	)
	// Output:
}

// ExampleNewURLWithDataRepresentationRelativeToURL demonstrates how to create a URL instance using NewURLWithDataRepresentationRelativeToURL.
func ExampleNewURLWithDataRepresentationRelativeToURL() {
	_ = foundation.NewURLWithDataRepresentationRelativeToURL(
		nil, // data unsafe.Pointer
		nil, // baseURL unsafe.Pointer
	)
	// Output:
}

// ExampleNewURLFileURLWithPath demonstrates how to create a URL instance using NewURLFileURLWithPath.
// Initializes a newly created NSURL referencing the local file or directory at  .
func ExampleNewURLFileURLWithPath() {
	_ = foundation.NewURLFileURLWithPath(
		"path", // path string
	)
	// Output:
}

// ExampleNewURLFileURLWithPathIsDirectoryRelativeToURL demonstrates how to create a URL instance using NewURLFileURLWithPathIsDirectoryRelativeToURL.
func ExampleNewURLFileURLWithPathIsDirectoryRelativeToURL() {
	_ = foundation.NewURLFileURLWithPathIsDirectoryRelativeToURL(
		"path", // path string
		false, // isDir bool
		nil, // baseURL unsafe.Pointer
	)
	// Output:
}

// ExampleNewURLFileURLWithPathRelativeToURL demonstrates how to create a URL instance using NewURLFileURLWithPathRelativeToURL.
func ExampleNewURLFileURLWithPathRelativeToURL() {
	_ = foundation.NewURLFileURLWithPathRelativeToURL(
		"path", // path string
		nil, // baseURL unsafe.Pointer
	)
	// Output:
}

// ExampleNewURLByResolvingAliasFileAtURLOptionsError demonstrates how to create a URL instance using NewURLByResolvingAliasFileAtURLOptionsError.
// Returns a new URL made by resolving the alias file at  .
func ExampleNewURLByResolvingAliasFileAtURLOptionsError() {
	_ = foundation.NewURLByResolvingAliasFileAtURLOptionsError(
		nil, // url unsafe.Pointer
		nil, // options unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}


