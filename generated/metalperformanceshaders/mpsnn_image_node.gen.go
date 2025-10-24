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
	

	// properties:
	Handle() Handle get set /* not a class type */
	SetHandle(value Handle get set /* not a class type */)
	ExportFromGraph() objectivec.IObject
	SetExportFromGraph(value objectivec.IObject)
	ImageAllocator() ImageAllocator get set /* not a class type */
	SetImageAllocator(value ImageAllocator get set /* not a class type */)
	Format() ImageFeatureChannelFormat get set /* not a class type */
	SetFormat(value ImageFeatureChannelFormat get set /* not a class type */)
	SynchronizeResource() objectivec.IObject
	SetSynchronizeResource(value objectivec.IObject)
	StopGradient() objectivec.IObject
	SetStopGradient(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageNodeClass) Alloc() ImageNode {
	rv := objc.Send[ImageNode](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A placeholder node denoting the position of a neural network image in a graph.


// A placeholder node denoting the position of a neural network image in a graph.
//
// [Full Topic]
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866483-initwithhandle
func NewImageNodeWithHandle(handle unsafe.Pointer) ImageNode {
	instance := getImageNodeClass().Alloc()
	rv := objc.Send[ImageNode](instance.ID, objc.Sel("initWithHandle:"), handle)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866440-exportednode
func (ic _ImageNodeClass) ExportedNode() {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("exportedNode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866440-exportednodewithhandle
func (ic _ImageNodeClass) ExportedNodeWithHandle(handle unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("exportedNodeWithHandle:"), handle)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866447-nodewithhandle
func (ic _ImageNodeClass) NodeWithHandle(handle unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("nodeWithHandle:"), handle)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866406-handle
func (i_ ImageNode) Handle() Handle get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("handle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866406-handle
func (i_ ImageNode) SetHandle(value Handle get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHandle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866478-exportfromgraph
func (i_ ImageNode) ExportFromGraph() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("exportFromGraph"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866478-exportfromgraph
func (i_ ImageNode) SetExportFromGraph(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExportFromGraph:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866490-imageallocator
func (i_ ImageNode) ImageAllocator() ImageAllocator get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("imageAllocator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866490-imageallocator
func (i_ ImageNode) SetImageAllocator(value ImageAllocator get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageAllocator:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866498-format
func (i_ ImageNode) Format() ImageFeatureChannelFormat get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("format"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2866498-format
func (i_ ImageNode) SetFormat(value ImageFeatureChannelFormat get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2942638-synchronizeresource
func (i_ ImageNode) SynchronizeResource() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("synchronizeResource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/2942638-synchronizeresource
func (i_ ImageNode) SetSynchronizeResource(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSynchronizeResource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/3020689-stopgradient
func (i_ ImageNode) StopGradient() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("stopGradient"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnimagenode/3020689-stopgradient
func (i_ ImageNode) SetStopGradient(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStopGradient:"), value)
}







