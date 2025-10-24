// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// containerStatusProtocol is the containerStatus protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to containerStatus.
var containerStatusProtocol *objc.Protocol

func init() {
	containerStatusProtocol = objc.GetProtocol("containerStatus")
}

