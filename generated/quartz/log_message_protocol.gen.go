// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import "github.com/ebitengine/purego/objc"

// logMessageProtocol is the logMessage: protocol.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// Use this protocol when registering custom classes that conform to logMessage:.
var logMessageProtocol *objc.Protocol

func init() {
	logMessageProtocol = objc.GetProtocol("logMessage:")
}

