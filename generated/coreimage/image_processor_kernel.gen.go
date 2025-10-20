// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ImageProcessorKernel] class.
var (
	imageProcessorKernelClass     _ImageProcessorKernelClass
	imageProcessorKernelClassOnce sync.Once
)

func getImageProcessorKernelClass() _ImageProcessorKernelClass {
	imageProcessorKernelClassOnce.Do(func() {
		imageProcessorKernelClass = _ImageProcessorKernelClass{objc.GetClass("CIImageProcessorKernel")}
	})
	return imageProcessorKernelClass
}

type _ImageProcessorKernelClass struct {
	class objc.Class
}

// An interface definition for the [ImageProcessorKernel] class.
type IImageProcessorKernel interface {
	objectivec.IObject
}

// The abstract class you extend to create custom image processors that can integrate with Core Image workflows.
//
// Unlike the class and its other subclasses that allow you to create new image-processing effects with the Core Image Kernel Language, the class provides direct access to the underlying bitmap image data for a step in the Core Image processing pipeline. As such, you can create subclasses of this class to integrate other image-processing technologies—such as Metal compute shaders, , operations, or your own CPU-based image-processing routines—with a Core Image filter chain. Your custom image processing operation is invoked by your subclassed image processor kernel’s method. The method can accept zero, one or more inputs: kernels that generate imagery (such as a noise or pattern generator) need no inputs, while kernels that composite source images together require multiple inputs. The dictionary allows the caller to pass in additional parameter values (such as the radius of a blur) and the contains the destination for your image processing code to write to. The following code shows how you can subclass to apply the Metal Performance Shader kernel to a : To apply to kernel to an image, the calling side invokes the image processor’s method. The following code generates a new object named which contains a thresholded version of the source image, .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel
type ImageProcessorKernel struct {
	objectivec.Object
}

// ImageProcessorKernelFrom constructs a [ImageProcessorKernel] from an unsafe.Pointer.
//
// The abstract class you extend to create custom image processors that can integrate with Core Image workflows.
func ImageProcessorKernelFrom(ptr unsafe.Pointer) ImageProcessorKernel {
	return ImageProcessorKernel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageProcessorKernelClass) Alloc() ImageProcessorKernel {
	rv := objc.Send[ImageProcessorKernel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageProcessorKernelClass) New() ImageProcessorKernel {
	rv := objc.Send[ImageProcessorKernel](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageProcessorKernel) Init() ImageProcessorKernel {
	rv := objc.Send[ImageProcessorKernel](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageProcessorKernel) Autorelease() ImageProcessorKernel {
	rv := objc.Send[ImageProcessorKernel](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageProcessorKernel creates a new ImageProcessorKernel instance.
func NewImageProcessorKernel() ImageProcessorKernel {
	return getImageProcessorKernelClass().New()
}


// Call this method on your Core Image Processor Kernel subclass to create a new image of the specified extent.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/apply(withExtent:inputs:arguments:)
func (ic _ImageProcessorKernelClass) ApplyWithExtentInputsArgumentsError(extent coregraphics.CGRect, inputs unsafe.Pointer, arguments unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("applyWithExtent:inputs:arguments:error:"), extent, inputs, arguments, error)
	return rv
}

// Call this method on your multiple-output Core Image Processor Kernel subclass to create an array of new image objects given the specified array of extents.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/apply(withExtents:inputs:arguments:)
func (ic _ImageProcessorKernelClass) ApplyWithExtentsInputsArgumentsError(extents unsafe.Pointer, inputs unsafe.Pointer, arguments unsafe.Pointer, error unsafe.Pointer) []Image {
	rv := objc.Send[[]Image](objc.ID(ic.class), objc.Sel("applyWithExtents:inputs:arguments:error:"), extents, inputs, arguments, error)
	return rv
}

// Override this class method if you want your any of the inputs to be in a specific pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/formatForInput(at:)
func (ic _ImageProcessorKernelClass) FormatForInputAtIndex(inputIndex unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("formatForInputAtIndex:"), inputIndex)
	return rv
}

// Override this class method if your processor has more than one output and you want your processor’s output to be in a specific supported .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputFormat(at:arguments:)
func (ic _ImageProcessorKernelClass) OutputFormatAtIndexArguments(outputIndex unsafe.Pointer, arguments unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("outputFormatAtIndex:arguments:"), outputIndex, arguments)
	return rv
}

// Override this class method to implement your Core Image Processor Kernel subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/process(with:arguments:output:)
func (ic _ImageProcessorKernelClass) ProcessWithInputsArgumentsOutputError(inputs unsafe.Pointer, arguments unsafe.Pointer, output unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("processWithInputs:arguments:output:error:"), inputs, arguments, output, error)
	return rv
}

// Override this class method of your Core Image Processor Kernel subclass if it needs to produce multiple outputs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/process(with:arguments:outputs:)
func (ic _ImageProcessorKernelClass) ProcessWithInputsArgumentsOutputsError(inputs unsafe.Pointer, arguments unsafe.Pointer, outputs unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("processWithInputs:arguments:outputs:error:"), inputs, arguments, outputs, error)
	return rv
}

// Override this class method to implement your processor’s ROI callback.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/roi(forInput:arguments:outputRect:)
func (ic _ImageProcessorKernelClass) RoiForInputArgumentsOutputRect(inputIndex unsafe.Pointer, arguments unsafe.Pointer, outputRect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](objc.ID(ic.class), objc.Sel("roiForInput:arguments:outputRect:"), inputIndex, arguments, outputRect)
	return rv
}

// Override this class method to implement your processor’s tiled ROI callback.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/roiTileArray(forInput:arguments:outputRect:)
func (ic _ImageProcessorKernelClass) RoiTileArrayForInputArgumentsOutputRect(inputIndex unsafe.Pointer, arguments unsafe.Pointer, outputRect coregraphics.CGRect) []Vector {
	rv := objc.Send[[]Vector](objc.ID(ic.class), objc.Sel("roiTileArrayForInput:arguments:outputRect:"), inputIndex, arguments, outputRect)
	return rv
}



