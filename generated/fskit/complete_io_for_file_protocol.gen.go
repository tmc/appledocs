// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// completeIOForFileProtocol is the completeIOForFile: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to completeIOForFile:.
var completeIOForFileProtocol *objc.Protocol

func init() {
	completeIOForFileProtocol = objc.GetProtocol("completeIOForFile:")
}
