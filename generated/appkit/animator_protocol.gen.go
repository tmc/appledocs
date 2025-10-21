// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// animatorProtocol is the animator protocol.
//
// Availability:
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to animator.
var animatorProtocol *objc.Protocol

func init() {
	animatorProtocol = objc.GetProtocol("animator")
}
