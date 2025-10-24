// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// didFinishLoadingProtocol is the didFinishLoading protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to didFinishLoading.
var didFinishLoadingProtocol *objc.Protocol

func init() {
	didFinishLoadingProtocol = objc.GetProtocol("didFinishLoading")
}

