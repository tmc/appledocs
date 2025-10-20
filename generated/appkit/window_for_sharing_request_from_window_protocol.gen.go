// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowForSharingRequestFromWindowProtocol is the windowForSharingRequestFromWindow: protocol.
//
// Availability:
//   - macOS 15.0+
//
// Use this protocol when registering custom classes that conform to windowForSharingRequestFromWindow:.
var windowForSharingRequestFromWindowProtocol *objc.Protocol

func init() {
	windowForSharingRequestFromWindowProtocol = objc.GetProtocol("windowForSharingRequestFromWindow:")
}
