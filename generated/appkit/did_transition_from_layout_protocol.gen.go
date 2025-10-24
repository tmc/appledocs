// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// didTransitionFromLayoutProtocol is the didTransitionFromLayout: protocol.
//
// Availability:
//   - macOS 10.11+
//
// Use this protocol when registering custom classes that conform to didTransitionFromLayout:.
var didTransitionFromLayoutProtocol *objc.Protocol

func init() {
	didTransitionFromLayoutProtocol = objc.GetProtocol("didTransitionFromLayout:")
}
