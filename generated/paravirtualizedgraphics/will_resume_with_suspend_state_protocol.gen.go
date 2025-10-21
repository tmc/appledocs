// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// willResumeWithSuspendStateProtocol is the willResumeWithSuspendState: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to willResumeWithSuspendState:.
var willResumeWithSuspendStateProtocol *objc.Protocol

func init() {
	willResumeWithSuspendStateProtocol = objc.GetProtocol("willResumeWithSuspendState:")
}
