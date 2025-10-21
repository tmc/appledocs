// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// newDisplayWithDescriptorProtocol is the newDisplayWithDescriptor: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to newDisplayWithDescriptor:.
var newDisplayWithDescriptorProtocol *objc.Protocol

func init() {
	newDisplayWithDescriptorProtocol = objc.GetProtocol("newDisplayWithDescriptor:")
}
