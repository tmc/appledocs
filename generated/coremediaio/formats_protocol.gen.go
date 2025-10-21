// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import "github.com/ebitengine/purego/objc"

// formatsProtocol is the formats protocol.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - macOS 12.3+
//
// Use this protocol when registering custom classes that conform to formats.
var formatsProtocol *objc.Protocol

func init() {
	formatsProtocol = objc.GetProtocol("formats")
}
