// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationProtectedDataDidBecomeAvailableProtocol is the applicationProtectedDataDidBecomeAvailable: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to applicationProtectedDataDidBecomeAvailable:.
var applicationProtectedDataDidBecomeAvailableProtocol *objc.Protocol

func init() {
	applicationProtectedDataDidBecomeAvailableProtocol = objc.GetProtocol("applicationProtectedDataDidBecomeAvailable:")
}
