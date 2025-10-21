// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// popoverDidShowProtocol is the popoverDidShow: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to popoverDidShow:.
var popoverDidShowProtocol *objc.Protocol

func init() {
	popoverDidShowProtocol = objc.GetProtocol("popoverDidShow:")
}
