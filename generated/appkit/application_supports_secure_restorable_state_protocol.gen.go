// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationSupportsSecureRestorableStateProtocol is the applicationSupportsSecureRestorableState: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to applicationSupportsSecureRestorableState:.
var applicationSupportsSecureRestorableStateProtocol *objc.Protocol

func init() {
	applicationSupportsSecureRestorableStateProtocol = objc.GetProtocol("applicationSupportsSecureRestorableState:")
}
