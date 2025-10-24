// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TransformLayer] class.
var (
	TransformLayerClass     _TransformLayerClass
	TransformLayerClassOnce sync.Once
)

func getTransformLayerClass() _TransformLayerClass {
	TransformLayerClassOnce.Do(func() {
		TransformLayerClass = _TransformLayerClass{objc.GetClass("CATransformLayer")}
	})
	return TransformLayerClass
}

type _TransformLayerClass struct {
	class objc.Class
}

// An interface definition for the [TransformLayer] class.
type ITransformLayer interface {
	ILayer
	// properties:
	ZPosition() float64
	SetZPosition(value float64)
	// methods:
}

// Objects used to create true 3D layer hierarchies, rather than the flattened hierarchy rendering model used by other layer types.
//
// Unlike normal layers, transform layers do not flatten their sublayers into the plane at . Due to this, they do not support many of the features of the class compositing model: Only the sublayers of a transform layer are rendered. The properties that are rendered by a layer are ignored, including: , , border style properties, stroke style properties, etc. The properties that assume 2D image processing are also ignored, including: , , , , , and shadow style properties. The property is applied to each sublayer individually, the transform layer does not form a compositing group. The method should never be called on a transform layer as they do not have a 2D coordinate space into which the point can be mapped.


// Objects used to create true 3D layer hierarchies, rather than the flattened hierarchy rendering model used by other layer types.
//
// [Full Topic]
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

// Alloc allocates a new instance without initialization.
func (tc _TransformLayerClass) Alloc() TransformLayer {
	rv := objc.Send[TransformLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TransformLayerClass) New() TransformLayer {
	rv := objc.Send[TransformLayer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TransformLayer) Init() TransformLayer {
	rv := objc.Send[TransformLayer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TransformLayer) Autorelease() TransformLayer {
	rv := objc.Send[TransformLayer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTransformLayer creates a new TransformLayer instance.
func NewTransformLayer() TransformLayer {
	return getTransformLayerClass().New()
}



// The layer’s position on the z axis. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/zposition
func (t_ TransformLayer) ZPosition() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("zPosition"))
	return rv
}


// The layer’s position on the z axis. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/zposition
func (t_ TransformLayer) SetZPosition(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setZPosition:"), value)
}



