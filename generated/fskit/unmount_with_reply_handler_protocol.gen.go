// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// unmountWithReplyHandlerProtocol is the unmountWithReplyHandler: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to unmountWithReplyHandler:.
var unmountWithReplyHandlerProtocol *objc.Protocol

func init() {
	unmountWithReplyHandlerProtocol = objc.GetProtocol("unmountWithReplyHandler:")
}
