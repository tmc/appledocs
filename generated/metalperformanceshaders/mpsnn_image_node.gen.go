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


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/imageallocator
func (i_ ImageNode) ImageAllocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageAllocator"))
	return rv
}


// SetImageAllocator sets the value of the imageAllocator property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/imageallocator
func (i_ ImageNode) SetImageAllocator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageAllocator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/handle
func (i_ ImageNode) Handle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("handle"))
	return rv
}


// SetHandle sets the value of the handle property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/handle
func (i_ ImageNode) SetHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHandle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/exportfromgraph
func (i_ ImageNode) ExportFromGraph() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("exportFromGraph"))
	return rv
}


// SetExportFromGraph sets the value of the exportFromGraph property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/exportfromgraph
func (i_ ImageNode) SetExportFromGraph(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExportFromGraph:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/stopgradient
func (i_ ImageNode) StopGradient() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("stopGradient"))
	return rv
}


// SetStopGradient sets the value of the stopGradient property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/stopgradient
func (i_ ImageNode) SetStopGradient(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStopGradient:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/synchronizeresource
func (i_ ImageNode) SynchronizeResource() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("synchronizeResource"))
	return rv
}


// SetSynchronizeResource sets the value of the synchronizeResource property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/synchronizeresource
func (i_ ImageNode) SetSynchronizeResource(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSynchronizeResource:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/format
func (i_ ImageNode) Format() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("format"))
	return rv
}


// SetFormat sets the value of the format property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/format
func (i_ ImageNode) SetFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFormat:"), value)
}



