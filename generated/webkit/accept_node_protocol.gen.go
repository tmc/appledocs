// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// acceptNodeProtocol is the acceptNode: protocol.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to acceptNode:.
var acceptNodeProtocol *objc.Protocol

func init() {
	acceptNodeProtocol = objc.GetProtocol("acceptNode:")
}

