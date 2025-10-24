// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import "github.com/ebitengine/purego/objc"

// ssidDidChangeForWiFiInterfaceWithNameProtocol is the ssidDidChangeForWiFiInterfaceWithName: protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to ssidDidChangeForWiFiInterfaceWithName:.
var ssidDidChangeForWiFiInterfaceWithNameProtocol *objc.Protocol

func init() {
	ssidDidChangeForWiFiInterfaceWithNameProtocol = objc.GetProtocol("ssidDidChangeForWiFiInterfaceWithName:")
}

