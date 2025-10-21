// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import "github.com/ebitengine/purego/objc"

// uuidProtocol is the uuid protocol.
//
// Availability:
//   - macOS 15.0+
//
// Use this protocol when registering custom classes that conform to uuid.
var uuidProtocol *objc.Protocol

func init() {
	uuidProtocol = objc.GetProtocol("uuid")
}
