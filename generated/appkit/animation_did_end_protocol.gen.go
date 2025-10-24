// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// animationDidEndProtocol is the animationDidEnd: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to animationDidEnd:.
var animationDidEndProtocol *objc.Protocol

func init() {
	animationDidEndProtocol = objc.GetProtocol("animationDidEnd:")
}
