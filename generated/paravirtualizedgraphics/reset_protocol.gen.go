// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// resetProtocol is the reset protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to reset.
var resetProtocol *objc.Protocol

func init() {
	resetProtocol = objc.GetProtocol("reset")
}

