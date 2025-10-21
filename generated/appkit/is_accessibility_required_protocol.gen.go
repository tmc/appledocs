// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// isAccessibilityRequiredProtocol is the isAccessibilityRequired protocol.
//
// Availability:
//   - macOS 10.12+
//
// Use this protocol when registering custom classes that conform to isAccessibilityRequired.
var isAccessibilityRequiredProtocol *objc.Protocol

func init() {
	isAccessibilityRequiredProtocol = objc.GetProtocol("isAccessibilityRequired")
}
