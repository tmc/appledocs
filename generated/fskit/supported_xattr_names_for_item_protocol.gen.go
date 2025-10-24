// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// supportedXattrNamesForItemProtocol is the supportedXattrNamesForItem: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to supportedXattrNamesForItem:.
var supportedXattrNamesForItemProtocol *objc.Protocol

func init() {
	supportedXattrNamesForItemProtocol = objc.GetProtocol("supportedXattrNamesForItem:")
}

