// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIWarpKernel */


/* debug [class_header]: Header for CIWarpKernel */
// The class instance for the [WarpKernel] class.
var (
	WarpKernelClass     _WarpKernelClass
	WarpKernelClassOnce sync.Once
)

func getWarpKernelClass() _WarpKernelClass {
	WarpKernelClassOnce.Do(func() {
		WarpKernelClass = _WarpKernelClass{objc.GetClass("CIWarpKernel")}
	})
	return WarpKernelClass
}

type _WarpKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WarpKernel */
// An interface definition for the [WarpKernel] class.
type IWarpKernel interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for WarpKernel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WarpKernel */
	// methods:
	ApplyWithExtentRoiCallbackInputImageArguments(extent corefoundation.CGRect, callback KernelROICallback /* not a class type */, image ICIImage, args []objc.ID) IImage
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WarpKernel */
// Alloc allocates a new instance without initialization.
func (wc _WarpKernelClass) Alloc() WarpKernel {
	rv := objc.Send[WarpKernel](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WarpKernelClass) New() WarpKernel {
	rv := objc.Send[WarpKernel](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WarpKernel) Init() WarpKernel {
	rv := objc.Send[WarpKernel](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WarpKernel) Autorelease() WarpKernel {
	rv := objc.Send[WarpKernel](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWarpKernel creates a new WarpKernel instance.
func NewWarpKernel() WarpKernel {
	return getWarpKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WarpKernel */
// A GPU-based image-processing routine that processes only the geometry information in an image, used to create custom Core Image filters.
//
// The kernel language routine for a warp kernel has the following characteristics: It uses exactly one input image. Its return type is (Core Image Kernel Language) or (Metal Shading Language), specifying a position in source image coordinates. A warp kernel routine requires no input parameters (but can use additional custom parameters you declare). Typically, a warp kernel uses the destination coordinate function to look up the coordinates of the destination pixel currently being rendered, then computes a corresponding position in source image coordinates (output using the keyword). Core Image then samples from the source image at the returned coordinates to produce a pixel color for the output image. For example, the Metal Shading Language source below implements a filter that passes through its input image unchanged. The equivalent code in Core Image Kernel Language is: The Core Image Kernel Language is a dialect of the OpenGL Shading Language. See and for more details.


// A GPU-based image-processing routine that processes only the geometry information in an image, used to create custom Core Image filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIWarpKernel
type WarpKernel struct {
	Kernel
}

// WarpKernelFrom constructs a [WarpKernel] from an unsafe.Pointer.
//
// A GPU-based image-processing routine that processes only the geometry information in an image, used to create custom Core Image filters.
func WarpKernelFrom(ptr unsafe.Pointer) WarpKernel {
	return WarpKernel{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WarpKernel */

// Creates a warp kernel object from the specified kernel source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIWarpKernel/init(source:)
func NewWarpKernelWithString(string_ objc.IObject /* cross-framework: NSString */) WarpKernel {
	rv := objc.Send[WarpKernel](objc.ID(getWarpKernelClass().class), objc.Sel("kernelWithString:"), string_)
	return rv
}/* debug [class_init_methods/constructor]: NewWarpKernelWithString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WarpKernel */

// Creates a warp kernel object from the specified kernel source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIWarpKernel/init(source:)
func (wc _WarpKernelClass) KernelWithString(string_ objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("kernelWithString:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=KernelWithString) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WarpKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WarpKernel */

// Creates a new image using the kernel and the specified input image and arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIWarpKernel/apply(extent:roiCallback:image:arguments:)
func (w_ WarpKernel) ApplyWithExtentRoiCallbackInputImageArguments(extent corefoundation.CGRect, callback KernelROICallback /* not a class type */, image ICIImage, args []objc.ID) IImage {
	rv := objc.Send[Image](w_.ID, objc.Sel("applyWithExtent:roiCallback:inputImage:arguments:"), extent, callback, image, args)
	return rv
}/* debug [instance_methods/method]: ApplyWithExtentRoiCallbackInputImageArguments */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WarpKernel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIWarpKernel */


