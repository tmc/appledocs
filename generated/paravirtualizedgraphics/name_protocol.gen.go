// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// nameProtocol is the name protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to name.
var nameProtocol *objc.Protocol

func init() {
	nameProtocol = objc.GetProtocol("name")
}
