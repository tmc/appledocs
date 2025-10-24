// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// animationShouldStartProtocol is the animationShouldStart: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to animationShouldStart:.
var animationShouldStartProtocol *objc.Protocol

func init() {
	animationShouldStartProtocol = objc.GetProtocol("animationShouldStart:")
}
