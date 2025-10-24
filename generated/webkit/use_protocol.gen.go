// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// useProtocol is the use protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to use.
var useProtocol *objc.Protocol

func init() {
	useProtocol = objc.GetProtocol("use")
}
