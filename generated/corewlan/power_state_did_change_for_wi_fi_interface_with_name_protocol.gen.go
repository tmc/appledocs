// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import "github.com/ebitengine/purego/objc"

// powerStateDidChangeForWiFiInterfaceWithNameProtocol is the powerStateDidChangeForWiFiInterfaceWithName: protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to powerStateDidChangeForWiFiInterfaceWithName:.
var powerStateDidChangeForWiFiInterfaceWithNameProtocol *objc.Protocol

func init() {
	powerStateDidChangeForWiFiInterfaceWithNameProtocol = objc.GetProtocol("powerStateDidChangeForWiFiInterfaceWithName:")
}

