// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// stopProtocol is the stop protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to stop.
var stopProtocol *objc.Protocol

func init() {
	stopProtocol = objc.GetProtocol("stop")
}

