// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import "github.com/ebitengine/purego/objc"

// systemExtensionWillBecomeDisabledProtocol is the systemExtensionWillBecomeDisabled: protocol.
//
// Availability:
//   - macOS 15.1+
//
// Use this protocol when registering custom classes that conform to systemExtensionWillBecomeDisabled:.
var systemExtensionWillBecomeDisabledProtocol *objc.Protocol

func init() {
	systemExtensionWillBecomeDisabledProtocol = objc.GetProtocol("systemExtensionWillBecomeDisabled:")
}

