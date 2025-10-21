// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import "github.com/ebitengine/purego/objc"

// systemExtensionWillBecomeEnabledProtocol is the systemExtensionWillBecomeEnabled: protocol.
//
// Availability:
//   - macOS 15.1+
//
// Use this protocol when registering custom classes that conform to systemExtensionWillBecomeEnabled:.
var systemExtensionWillBecomeEnabledProtocol *objc.Protocol

func init() {
	systemExtensionWillBecomeEnabledProtocol = objc.GetProtocol("systemExtensionWillBecomeEnabled:")
}
