// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import "github.com/ebitengine/purego/objc"

// clientConnectionInvalidatedProtocol is the clientConnectionInvalidated protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to clientConnectionInvalidated.
var clientConnectionInvalidatedProtocol *objc.Protocol

func init() {
	clientConnectionInvalidatedProtocol = objc.GetProtocol("clientConnectionInvalidated")
}
