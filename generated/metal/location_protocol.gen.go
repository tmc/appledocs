// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// locationProtocol is the location protocol.
//
// Availability:
//   - macOS 10.15+
//
// Use this protocol when registering custom classes that conform to location.
var locationProtocol *objc.Protocol

func init() {
	locationProtocol = objc.GetProtocol("location")
}
