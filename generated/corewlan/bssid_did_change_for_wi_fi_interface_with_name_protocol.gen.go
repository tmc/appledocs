// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import "github.com/ebitengine/purego/objc"

// bssidDidChangeForWiFiInterfaceWithNameProtocol is the bssidDidChangeForWiFiInterfaceWithName: protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to bssidDidChangeForWiFiInterfaceWithName:.
var bssidDidChangeForWiFiInterfaceWithNameProtocol *objc.Protocol

func init() {
	bssidDidChangeForWiFiInterfaceWithNameProtocol = objc.GetProtocol("bssidDidChangeForWiFiInterfaceWithName:")
}
