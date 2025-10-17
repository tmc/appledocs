// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TransformLayer] class.
var transformLayerClass = _TransformLayerClass{objc.GetClass("CATransformLayer")}

type _TransformLayerClass struct {
	class objc.Class
}

// Objects used to create true 3D layer hierarchies, rather than the flattened hierarchy rendering model used by other layer types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransformLayer

type TransformLayer struct {
	Layer
}

// TransformLayerFrom constructs a [TransformLayer] from an unsafe.Pointer.
//
// Objects used to create true 3D layer hierarchies, rather than the flattened hierarchy rendering model used by other layer types.
func TransformLayerFrom(ptr unsafe.Pointer) TransformLayer {
	return TransformLayer{
		Layer: LayerFrom(ptr),
	}
}



