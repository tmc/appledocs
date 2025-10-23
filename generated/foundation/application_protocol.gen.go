// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// applicationProtocol is the application: protocol.
//
// Availability:
//   - macOS 10.13+
//
// Use this protocol when registering custom classes that conform to application:.
var applicationProtocol *objc.Protocol

func init() {
	applicationProtocol = objc.GetProtocol("application:")
}
