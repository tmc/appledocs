// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// modeChangeHandlerProtocol is the modeChangeHandler protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to modeChangeHandler.
var modeChangeHandlerProtocol *objc.Protocol

func init() {
	modeChangeHandlerProtocol = objc.GetProtocol("modeChangeHandler")
}
