// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import "github.com/ebitengine/purego/objc"

// IOExtensionStreamSourceProtocol is the CMIOExtensionStreamSource protocol.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - macOS 12.3+
//
// Use this protocol when registering custom classes that conform to CMIOExtensionStreamSource.
var IOExtensionStreamSourceProtocol *objc.Protocol

func init() {
	IOExtensionStreamSourceProtocol = objc.GetProtocol("CMIOExtensionStreamSource")
}
