// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// dispatchEventProtocol is the dispatchEvent: protocol.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to dispatchEvent:.
var dispatchEventProtocol *objc.Protocol

func init() {
	dispatchEventProtocol = objc.GetProtocol("dispatchEvent:")
}
