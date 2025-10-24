// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationProtectedDataWillBecomeUnavailableProtocol is the applicationProtectedDataWillBecomeUnavailable: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to applicationProtectedDataWillBecomeUnavailable:.
var applicationProtectedDataWillBecomeUnavailableProtocol *objc.Protocol

func init() {
	applicationProtectedDataWillBecomeUnavailableProtocol = objc.GetProtocol("applicationProtectedDataWillBecomeUnavailable:")
}
