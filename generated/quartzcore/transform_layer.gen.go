// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TransformLayer] class.
var (
	transformLayerClass     _TransformLayerClass
	transformLayerClassOnce sync.Once
)

func getTransformLayerClass() _TransformLayerClass {
	transformLayerClassOnce.Do(func() {
		transformLayerClass = _TransformLayerClass{objc.GetClass("CATransformLayer")}
	})
	return transformLayerClass
}

type _TransformLayerClass struct {
	class objc.Class
}

// An interface definition for the [TransformLayer] class.
type ITransformLayer interface {
	ILayer
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




