// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// connectionProtocol is the connection: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.13)
//
// Use this protocol when registering custom classes that conform to connection:.
var connectionProtocol *objc.Protocol

func init() {
	connectionProtocol = objc.GetProtocol("connection:")
}
