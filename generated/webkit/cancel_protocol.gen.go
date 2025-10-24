// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// cancelProtocol is the cancel protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to cancel.
var cancelProtocol *objc.Protocol

func init() {
	cancelProtocol = objc.GetProtocol("cancel")
}

