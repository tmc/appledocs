// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageNode] class.
var (
	ImageNodeClass     _ImageNodeClass
	ImageNodeClassOnce sync.Once
)

func getImageNodeClass() _ImageNodeClass {
	ImageNodeClassOnce.Do(func() {
		ImageNodeClass = _ImageNodeClass{objc.GetClass("MPSNNImageNode")}
	})
	return ImageNodeClass
}

type _ImageNodeClass struct {
	class objc.Class
}

// An interface definition for the [ImageNode] class.
type IImageNode interface {
	objectivec.IObject
}

// A placeholder node denoting the position of a neural network image in a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode
type ImageNode struct {
	objectivec.Object
}

// ImageNodeFrom constructs a [ImageNode] from an unsafe.Pointer.
//
// A placeholder node denoting the position of a neural network image in a graph.
func ImageNodeFrom(ptr unsafe.Pointer) ImageNode {
	return ImageNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageNodeClass) Alloc() ImageNode {
	rv := objc.Send[ImageNode](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageNodeClass) New() ImageNode {
	rv := objc.Send[ImageNode](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageNode) Init() ImageNode {
	rv := objc.Send[ImageNode](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageNode) Autorelease() ImageNode {
	rv := objc.Send[ImageNode](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageNode creates a new ImageNode instance.
func NewImageNode() ImageNode {
	return getImageNodeClass().New()
}




