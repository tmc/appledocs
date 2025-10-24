// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// containerProtocol is the container protocol.
//
// Availability:
//   - macOS 10.15+
//
// Use this protocol when registering custom classes that conform to container.
var containerProtocol *objc.Protocol

func init() {
	containerProtocol = objc.GetProtocol("container")
}

