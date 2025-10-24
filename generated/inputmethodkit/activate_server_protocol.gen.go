// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import "github.com/ebitengine/purego/objc"

// activateServerProtocol is the activateServer: protocol.
//
// Availability:
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to activateServer:.
var activateServerProtocol *objc.Protocol

func init() {
	activateServerProtocol = objc.GetProtocol("activateServer:")
}

