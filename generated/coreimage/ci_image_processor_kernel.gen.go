// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageProcessorKernel] class.
var (
	ImageProcessorKernelClass     _ImageProcessorKernelClass
	ImageProcessorKernelClassOnce sync.Once
)

func getImageProcessorKernelClass() _ImageProcessorKernelClass {
	ImageProcessorKernelClassOnce.Do(func() {
		ImageProcessorKernelClass = _ImageProcessorKernelClass{objc.GetClass("CIImageProcessorKernel")}
	})
	return ImageProcessorKernelClass
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
func (ic _ImageProcessorKernelClass) ApplyWithExtentInputsArgumentsError(extent coregraphics.CGRect, inputs []Image, arguments unsafe.Pointer, error_ unsafe.Pointer) Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("applyWithExtent:inputs:arguments:error:"), extent, inputs, arguments, error_)
	return rv
}

// Call this method on your multiple-output Core Image Processor Kernel subclass to create an array of new image objects given the specified array of extents.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/apply(withExtents:inputs:arguments:)
func (ic _ImageProcessorKernelClass) ApplyWithExtentsInputsArgumentsError(extents []Vector, inputs []Image, arguments unsafe.Pointer, error_ unsafe.Pointer) []Image {
	rv := objc.Send[[]Image](objc.ID(ic.class), objc.Sel("applyWithExtents:inputs:arguments:error:"), extents, inputs, arguments, error_)
	return rv
}

// Override this class method if you want your any of the inputs to be in a specific pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/formatForInput(at:)
func (ic _ImageProcessorKernelClass) FormatForInputAtIndex(inputIndex int) Format {
	rv := objc.Send[Format](objc.ID(ic.class), objc.Sel("formatForInputAtIndex:"), inputIndex)
	return rv
}

// Override this class method if your processor has more than one output and you want your processor’s output to be in a specific supported .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputFormat(at:arguments:)
func (ic _ImageProcessorKernelClass) OutputFormatAtIndexArguments(outputIndex int, arguments unsafe.Pointer) Format {
	rv := objc.Send[Format](objc.ID(ic.class), objc.Sel("outputFormatAtIndex:arguments:"), outputIndex, arguments)
	return rv
}

// Override this class method to implement your Core Image Processor Kernel subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/process(with:arguments:output:)
func (ic _ImageProcessorKernelClass) ProcessWithInputsArgumentsOutputError(inputs []objc.ID, arguments unsafe.Pointer, output objectivec.IObject, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("processWithInputs:arguments:output:error:"), inputs, arguments, output, error_)
	return rv
}

// Override this class method of your Core Image Processor Kernel subclass if it needs to produce multiple outputs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/process(with:arguments:outputs:)
func (ic _ImageProcessorKernelClass) ProcessWithInputsArgumentsOutputsError(inputs []objc.ID, arguments unsafe.Pointer, outputs []objc.ID, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("processWithInputs:arguments:outputs:error:"), inputs, arguments, outputs, error_)
	return rv
}

// Override this class method to implement your processor’s ROI callback.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/roi(forInput:arguments:outputRect:)
func (ic _ImageProcessorKernelClass) RoiForInputArgumentsOutputRect(inputIndex int, arguments unsafe.Pointer, outputRect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](objc.ID(ic.class), objc.Sel("roiForInput:arguments:outputRect:"), inputIndex, arguments, outputRect)
	return rv
}

// Override this class method to implement your processor’s tiled ROI callback.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/roiTileArray(forInput:arguments:outputRect:)
func (ic _ImageProcessorKernelClass) RoiTileArrayForInputArgumentsOutputRect(inputIndex int, arguments unsafe.Pointer, outputRect coregraphics.CGRect) []Vector {
	rv := objc.Send[[]Vector](objc.ID(ic.class), objc.Sel("roiTileArrayForInput:arguments:outputRect:"), inputIndex, arguments, outputRect)
	return rv
}

// Override this class property if you want your processor’s output to be in a specific pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputFormat
func (ic _ImageProcessorKernelClass) OutputFormat() Format {
	rv := objc.Send[Format](objc.ID(ic.class), objc.Sel("outputFormat"))
	return rv
}
// Override this class property if your processor’s output stores 1.0 into the alpha channel of all pixels within the output extent.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputIsOpaque
func (ic _ImageProcessorKernelClass) OutputIsOpaque() bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("outputIsOpaque"))
	return rv
}
// Override this class property to return false if you want your processor to be given input objects that have not been synchronized for CPU access.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/synchronizeInputs
func (ic _ImageProcessorKernelClass) SynchronizeInputs() bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("synchronizeInputs"))
	return rv
}
// Override this class property if you want your processor’s output to be in a specific pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputFormat
func (i_ ImageProcessorKernel) OutputFormat() Format {
	rv := objc.Send[Format](i_.ID, objc.Sel("outputFormat"))
	return rv
}

// Override this class property if your processor’s output stores 1.0 into the alpha channel of all pixels within the output extent.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputIsOpaque
func (i_ ImageProcessorKernel) OutputIsOpaque() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("outputIsOpaque"))
	return rv
}

// Override this class property to return false if you want your processor to be given input objects that have not been synchronized for CPU access.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/synchronizeInputs
func (i_ ImageProcessorKernel) SynchronizeInputs() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("synchronizeInputs"))
	return rv
}



