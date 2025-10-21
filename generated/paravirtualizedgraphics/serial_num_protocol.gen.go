// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// serialNumProtocol is the serialNum protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to serialNum.
var serialNumProtocol *objc.Protocol

func init() {
	serialNumProtocol = objc.GetProtocol("serialNum")
}
