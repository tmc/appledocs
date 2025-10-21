// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import "github.com/ebitengine/purego/objc"

// didResumeProtocol is the didResume protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to didResume.
var didResumeProtocol *objc.Protocol

func init() {
	didResumeProtocol = objc.GetProtocol("didResume")
}
