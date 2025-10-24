// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNResizeBilinear */


/* debug [class_header]: Header for MPSNNResizeBilinear */
// The class instance for the [ResizeBilinear] class.
var (
	ResizeBilinearClass     _ResizeBilinearClass
	ResizeBilinearClassOnce sync.Once
)

func getResizeBilinearClass() _ResizeBilinearClass {
	ResizeBilinearClassOnce.Do(func() {
		ResizeBilinearClass = _ResizeBilinearClass{objc.GetClass("MPSNNResizeBilinear")}
	})
	return ResizeBilinearClass
}

type _ResizeBilinearClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ResizeBilinear */
// An interface definition for the [ResizeBilinear] class.
type IResizeBilinear interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for ResizeBilinear */
	// properties:
	AlignCorners() objectivec.IObject
	SetAlignCorners(value objectivec.IObject)
	ResizeHeight() objectivec.IObject
	SetResizeHeight(value objectivec.IObject)
	ResizeWidth() objectivec.IObject
	SetResizeWidth(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ResizeBilinear */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ResizeBilinear */
// Alloc allocates a new instance without initialization.
func (rc _ResizeBilinearClass) Alloc() ResizeBilinear {
	rv := objc.Send[ResizeBilinear](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ResizeBilinearClass) New() ResizeBilinear {
	rv := objc.Send[ResizeBilinear](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResizeBilinear) Init() ResizeBilinear {
	rv := objc.Send[ResizeBilinear](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResizeBilinear) Autorelease() ResizeBilinear {
	rv := objc.Send[ResizeBilinear](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResizeBilinear creates a new ResizeBilinear instance.
func NewResizeBilinear() ResizeBilinear {
	return getResizeBilinearClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ResizeBilinear */
// A bilinear resizing filter.


// A bilinear resizing filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear
type ResizeBilinear struct {
	CNNKernel
}

// ResizeBilinearFrom constructs a [ResizeBilinear] from an unsafe.Pointer.
//
// A bilinear resizing filter.
func ResizeBilinearFrom(ptr unsafe.Pointer) ResizeBilinear {
	return ResizeBilinear{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ResizeBilinear */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3013794-initwithcoder
func NewResizeBilinearWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ResizeBilinear {
	instance := getResizeBilinearClass().Alloc()
	rv := objc.Send[ResizeBilinear](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewResizeBilinearWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012966-initwithdevice
func NewResizeBilinearWithDeviceResizeWidthResizeHeightAlignCorners(device unsafe.Pointer, resizeWidth uint, resizeHeight uint, alignCorners bool) ResizeBilinear {
	instance := getResizeBilinearClass().Alloc()
	rv := objc.Send[ResizeBilinear](instance.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:alignCorners:"), device, resizeWidth, resizeHeight, alignCorners)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewResizeBilinearWithDeviceResizeWidthResizeHeightAlignCorners */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ResizeBilinear */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ResizeBilinear */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ResizeBilinear */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ResizeBilinear */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012965-aligncorners
func (r_ ResizeBilinear) AlignCorners() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("alignCorners"))
	return rv
}/* debug [instance_properties/getter]: alignCorners */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012965-aligncorners
func (r_ ResizeBilinear) SetAlignCorners(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAlignCorners:"), value)
}/* debug [instance_properties/setter]: alignCorners */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012968-resizeheight
func (r_ ResizeBilinear) ResizeHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("resizeHeight"))
	return rv
}/* debug [instance_properties/getter]: resizeHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012968-resizeheight
func (r_ ResizeBilinear) SetResizeHeight(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResizeHeight:"), value)
}/* debug [instance_properties/setter]: resizeHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012969-resizewidth
func (r_ ResizeBilinear) ResizeWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("resizeWidth"))
	return rv
}/* debug [instance_properties/getter]: resizeWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012969-resizewidth
func (r_ ResizeBilinear) SetResizeWidth(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResizeWidth:"), value)
}/* debug [instance_properties/setter]: resizeWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNResizeBilinear */


