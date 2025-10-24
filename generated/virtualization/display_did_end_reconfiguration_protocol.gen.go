// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import "github.com/ebitengine/purego/objc"

// displayDidEndReconfigurationProtocol is the displayDidEndReconfiguration: protocol.
//
// Availability:
//   - macOS 14.0+
//
// Use this protocol when registering custom classes that conform to displayDidEndReconfiguration:.
var displayDidEndReconfigurationProtocol *objc.Protocol

func init() {
	displayDidEndReconfigurationProtocol = objc.GetProtocol("displayDidEndReconfiguration:")
}
