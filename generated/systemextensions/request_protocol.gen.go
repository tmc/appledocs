// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import "github.com/ebitengine/purego/objc"

// requestProtocol is the request: protocol.
//
// Availability:
//   - macOS 10.15+
//
// Use this protocol when registering custom classes that conform to request:.
var requestProtocol *objc.Protocol

func init() {
	requestProtocol = objc.GetProtocol("request:")
}
