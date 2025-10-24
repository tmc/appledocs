// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// pauseProtocol is the pause protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to pause.
var pauseProtocol *objc.Protocol

func init() {
	pauseProtocol = objc.GetProtocol("pause")
}

