// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// mountWithOptionsProtocol is the mountWithOptions: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to mountWithOptions:.
var mountWithOptionsProtocol *objc.Protocol

func init() {
	mountWithOptionsProtocol = objc.GetProtocol("mountWithOptions:")
}
