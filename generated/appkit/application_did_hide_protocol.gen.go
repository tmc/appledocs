// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationDidHideProtocol is the applicationDidHide: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationDidHide:.
var applicationDidHideProtocol *objc.Protocol

func init() {
	applicationDidHideProtocol = objc.GetProtocol("applicationDidHide:")
}
