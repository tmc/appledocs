// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// finishSuspendProtocol is the finishSuspend protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to finishSuspend.
var finishSuspendProtocol *objc.Protocol

func init() {
	finishSuspendProtocol = objc.GetProtocol("finishSuspend")
}
