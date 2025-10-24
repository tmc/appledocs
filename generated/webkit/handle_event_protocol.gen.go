// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// handleEventProtocol is the handleEvent: protocol.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to handleEvent:.
var handleEventProtocol *objc.Protocol

func init() {
	handleEventProtocol = objc.GetProtocol("handleEvent:")
}
