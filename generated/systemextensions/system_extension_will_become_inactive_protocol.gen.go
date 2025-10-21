// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import "github.com/ebitengine/purego/objc"

// systemExtensionWillBecomeInactiveProtocol is the systemExtensionWillBecomeInactive: protocol.
//
// Availability:
//   - macOS 15.1+
//
// Use this protocol when registering custom classes that conform to systemExtensionWillBecomeInactive:.
var systemExtensionWillBecomeInactiveProtocol *objc.Protocol

func init() {
	systemExtensionWillBecomeInactiveProtocol = objc.GetProtocol("systemExtensionWillBecomeInactive:")
}
