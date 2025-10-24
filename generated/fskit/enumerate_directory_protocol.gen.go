// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// enumerateDirectoryProtocol is the enumerateDirectory: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to enumerateDirectory:.
var enumerateDirectoryProtocol *objc.Protocol

func init() {
	enumerateDirectoryProtocol = objc.GetProtocol("enumerateDirectory:")
}

