// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// createFileNamedProtocol is the createFileNamed: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to createFileNamed:.
var createFileNamedProtocol *objc.Protocol

func init() {
	createFileNamedProtocol = objc.GetProtocol("createFileNamed:")
}
