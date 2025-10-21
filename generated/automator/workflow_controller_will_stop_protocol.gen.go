// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import "github.com/ebitengine/purego/objc"

// workflowControllerWillStopProtocol is the workflowControllerWillStop: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 10.4+
//
// Use this protocol when registering custom classes that conform to workflowControllerWillStop:.
var workflowControllerWillStopProtocol *objc.Protocol

func init() {
	workflowControllerWillStopProtocol = objc.GetProtocol("workflowControllerWillStop:")
}
