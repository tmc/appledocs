// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import "github.com/ebitengine/purego/objc"

// eventDidFailProtocol is the eventDidFail: protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to eventDidFail:.
var eventDidFailProtocol *objc.Protocol

func init() {
	eventDidFailProtocol = objc.GetProtocol("eventDidFail:")
}
