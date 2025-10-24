// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// addEventListenerProtocol is the addEventListener: protocol.
//
// Availability:
//   - macOS 10.5+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to addEventListener:.
var addEventListenerProtocol *objc.Protocol

func init() {
	addEventListenerProtocol = objc.GetProtocol("addEventListener:")
}
