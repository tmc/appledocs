// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// isRemovableProtocol is the isRemovable protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.13+
//
// Use this protocol when registering custom classes that conform to isRemovable.
var isRemovableProtocol *objc.Protocol

func init() {
	isRemovableProtocol = objc.GetProtocol("isRemovable")
}
