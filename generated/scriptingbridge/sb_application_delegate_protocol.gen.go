// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import "github.com/ebitengine/purego/objc"

// SBApplicationDelegateProtocol is the SBApplicationDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to SBApplicationDelegate.
var SBApplicationDelegateProtocol *objc.Protocol

func init() {
	SBApplicationDelegateProtocol = objc.GetProtocol("SBApplicationDelegate")
}
