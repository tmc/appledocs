// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// queueProtocol is the queue protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to queue.
var queueProtocol *objc.Protocol

func init() {
	queueProtocol = objc.GetProtocol("queue")
}


