// Code generated from Apple documentation for HTTP-Live-Streaming. DO NOT EDIT.

// Package http-live-streaming provides Go bindings for the HTTP-Live-Streaming framework.
//
// Send audio and video to iOS, tvOS, and macOS devices. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to HTTP-Live-Streaming without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/HTTP-Live-Streaming
package http-live-streaming

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/HTTP-Live-Streaming.framework/HTTP-Live-Streaming"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


