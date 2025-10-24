// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Layer] class.
var (
	LayerClass     _LayerClass
	LayerClassOnce sync.Once
)

func getLayerClass() _LayerClass {
	LayerClassOnce.Do(func() {
		LayerClass = _LayerClass{objc.GetClass("CALayer")}
	})
	return LayerClass
}

type _LayerClass struct {
	class objc.Class
}

// An interface definition for the [Layer] class.
type ILayer interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type Layer struct {
	objectivec.Object
}

// LayerFrom constructs a [Layer] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func LayerFrom(ptr unsafe.Pointer) Layer {
	return Layer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayerClass) Alloc() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayerClass) New() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Layer) Init() Layer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Layer) Autorelease() Layer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayer creates a new Layer instance.
func NewLayer() Layer {
	return getLayerClass().New()
}




