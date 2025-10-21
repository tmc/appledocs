// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// closeItemProtocol is the closeItem: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to closeItem:.
var closeItemProtocol *objc.Protocol

func init() {
	closeItemProtocol = objc.GetProtocol("closeItem:")
}
