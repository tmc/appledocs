// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIRenderInfo */


/* debug [class_header]: Header for CIRenderInfo */
// The class instance for the [RenderInfo] class.
var (
	RenderInfoClass     _RenderInfoClass
	RenderInfoClassOnce sync.Once
)

func getRenderInfoClass() _RenderInfoClass {
	RenderInfoClassOnce.Do(func() {
		RenderInfoClass = _RenderInfoClass{objc.GetClass("CIRenderInfo")}
	})
	return RenderInfoClass
}

type _RenderInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RenderInfo */
// An interface definition for the [RenderInfo] class.
type IRenderInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RenderInfo */
	// properties:
	KernelCompileTime() float64
	KernelExecutionTime() float64
	PassCount() int
	PixelsProcessed() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RenderInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RenderInfo */
// Alloc allocates a new instance without initialization.
func (rc _RenderInfoClass) Alloc() RenderInfo {
	rv := objc.Send[RenderInfo](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderInfoClass) New() RenderInfo {
	rv := objc.Send[RenderInfo](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderInfo) Init() RenderInfo {
	rv := objc.Send[RenderInfo](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderInfo) Autorelease() RenderInfo {
	rv := objc.Send[RenderInfo](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderInfo creates a new RenderInfo instance.
func NewRenderInfo() RenderInfo {
	return getRenderInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RenderInfo */
// An encapsulation of a render task’s timing, passes, and pixels processed.
//
// A object allows Xcode Quick Look to visualize the render graph with detailed timing information.


// An encapsulation of a render task’s timing, passes, and pixels processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderInfo
type RenderInfo struct {
	objectivec.Object
}

// RenderInfoFrom constructs a [RenderInfo] from an unsafe.Pointer.
//
// An encapsulation of a render task’s timing, passes, and pixels processed.
func RenderInfoFrom(ptr unsafe.Pointer) RenderInfo {
	return RenderInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RenderInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RenderInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RenderInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RenderInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RenderInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderInfo/kernelCompileTime
func (r_ RenderInfo) KernelCompileTime() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("kernelCompileTime"))
	return rv
}/* debug [instance_properties/getter]: kernelCompileTime */


// The amount of time a render spent executing kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderInfo/kernelExecutionTime
func (r_ RenderInfo) KernelExecutionTime() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("kernelExecutionTime"))
	return rv
}/* debug [instance_properties/getter]: kernelExecutionTime */


// The number of passes the render took.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderInfo/passCount
func (r_ RenderInfo) PassCount() int {
	rv := objc.Send[int](r_.ID, objc.Sel("passCount"))
	return rv
}/* debug [instance_properties/getter]: passCount */


// The number of pixels the render produced executing kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderInfo/pixelsProcessed
func (r_ RenderInfo) PixelsProcessed() int {
	rv := objc.Send[int](r_.ID, objc.Sel("pixelsProcessed"))
	return rv
}/* debug [instance_properties/getter]: pixelsProcessed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIRenderInfo */



