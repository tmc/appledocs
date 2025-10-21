// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// FSVolumeReadWriteOperationsProtocol is the FSVolumeReadWriteOperations protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to FSVolumeReadWriteOperations.
var FSVolumeReadWriteOperationsProtocol *objc.Protocol

func init() {
	FSVolumeReadWriteOperationsProtocol = objc.GetProtocol("FSVolumeReadWriteOperations")
}
