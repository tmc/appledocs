// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReplicatorLayer] class.
var replicatorLayerClass = _ReplicatorLayerClass{objc.GetClass("CAReplicatorLayer")}

type _ReplicatorLayerClass struct {
	class objc.Class
}

// A layer that creates a specified number of sublayer copies with varying geometric, temporal, and color transformations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer

type ReplicatorLayer struct {
	Layer
}

// ReplicatorLayerFrom constructs a [ReplicatorLayer] from an unsafe.Pointer.
//
// A layer that creates a specified number of sublayer copies with varying geometric, temporal, and color transformations.
func ReplicatorLayerFrom(ptr unsafe.Pointer) ReplicatorLayer {
	return ReplicatorLayer{
		Layer: LayerFrom(ptr),
	}
}



