// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLSessionWebSocketMessage

// ExampleNewURLSessionWebSocketMessageWithData demonstrates how to create a URLSessionWebSocketMessage instance using NewURLSessionWebSocketMessageWithData.
func ExampleNewURLSessionWebSocketMessageWithData() {
	_ = foundation.NewURLSessionWebSocketMessageWithData(
		foundation.NSData{}, // data NSData
	)
	// Output:
}
// ExampleNewURLSessionWebSocketMessageWithString demonstrates how to create a URLSessionWebSocketMessage instance using NewURLSessionWebSocketMessageWithString.
func ExampleNewURLSessionWebSocketMessageWithString() {
	_ = foundation.NewURLSessionWebSocketMessageWithString(
		"string", // string string
	)
	// Output:
}
