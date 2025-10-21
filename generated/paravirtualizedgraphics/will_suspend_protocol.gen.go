// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// willSuspendProtocol is the willSuspend protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to willSuspend.
var willSuspendProtocol *objc.Protocol

func init() {
	willSuspendProtocol = objc.GetProtocol("willSuspend")
}
