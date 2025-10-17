// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RemoteLayerServer] class.
var remoteLayerServerClass = _RemoteLayerServerClass{objc.GetClass("CARemoteLayerServer")}

type _RemoteLayerServerClass struct {
	class objc.Class
}

// A legacy class for cross-process rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARemoteLayerServer

type RemoteLayerServer struct {
	objectivec.Object
}

// RemoteLayerServerFrom constructs a [RemoteLayerServer] from an unsafe.Pointer.
//
// A legacy class for cross-process rendering.
func RemoteLayerServerFrom(ptr unsafe.Pointer) RemoteLayerServer {
	return RemoteLayerServer{objectivec.Object{objc.ID(ptr)}}
}



