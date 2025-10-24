// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [BlendKernel] class.
var (
	BlendKernelClass     _BlendKernelClass
	BlendKernelClassOnce sync.Once
)

func getBlendKernelClass() _BlendKernelClass {
	BlendKernelClassOnce.Do(func() {
		BlendKernelClass = _BlendKernelClass{objc.GetClass("CIBlendKernel")}
	})
	return BlendKernelClass
}

type _BlendKernelClass struct {
	class objc.Class
}





// An interface definition for the [BlendKernel] class.
type IBlendKernel interface {
	IColorKernel
	

	// properties:


	

	// methods:
	ApplyWithForegroundBackground(foreground ICIImage, background ICIImage) IImage
	ApplyWithForegroundBackgroundColorSpace(foreground ICIImage, background ICIImage, colorSpace ColorSpaceRef /* not a class type */) IImage


}





// Alloc allocates a new instance without initialization.
func (bc _BlendKernelClass) Alloc() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BlendKernelClass) New() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BlendKernel) Init() BlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BlendKernel) Autorelease() BlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlendKernel creates a new BlendKernel instance.
func NewBlendKernel() BlendKernel {
	return getBlendKernelClass().New()
}





// A GPU-based image-processing routine that is optimized for blending two images.
//
// The blend kernel function has the following characteristics: It has two arguments of type (Core Image Kernel Language) or (Metal Shading Language), representing the foreground and background images. Its return type is (Core Image Kernel Language) or (Metal Shading Language); that is, it returns a pixel color for the output image. A blend kernel routine receives as input single-pixel colors (one sampled from each input image) and computes a final pixel color (output using the return keyword). For example, the Metal Shading Language source below implements a filter that returns the average of its two input images. Generally, the extent of the output image is the union of the extents of the foreground and background images.


// A GPU-based image-processing routine that is optimized for blending two images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel
type BlendKernel struct {
	ColorKernel
}

// BlendKernelFrom constructs a [BlendKernel] from an unsafe.Pointer.
//
// A GPU-based image-processing routine that is optimized for blending two images.
func BlendKernelFrom(ptr unsafe.Pointer) BlendKernel {
	return BlendKernel{
		ColorKernel: ColorKernelFrom(ptr),
	}
}






// Creates a custom blend kernel from a program string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/init(source:)
func NewBlendKernelWithString(string_ objc.IObject /* cross-framework: NSString */) BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(getBlendKernelClass().class), objc.Sel("kernelWithString:"), string_)
	return rv
}







// Creates a custom blend kernel from a program string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/init(source:)
func (bc _BlendKernelClass) KernelWithString(string_ objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("kernelWithString:"), string_)
	return rv
}







// A blend kernel that returns a clear color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/clear
func (bc _BlendKernelClass) Clear() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("clear"))
	return rv
}

// A blend kernel that uses the luminance values of the background with the hue and saturation values of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/color
func (bc _BlendKernelClass) Color() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("color"))
	return rv
}

// A blend kernel that darkens the background image samples to reflect the foreground image samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/colorBurn
func (bc _BlendKernelClass) ColorBurn() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("colorBurn"))
	return rv
}

// A blend kernel that brightens the background image samples to reflect the foreground image samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/colorDodge
func (bc _BlendKernelClass) ColorDodge() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("colorDodge"))
	return rv
}

// A blend kernel that adds color components to achieve a brightening effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentAdd
func (bc _BlendKernelClass) ComponentAdd() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("componentAdd"))
	return rv
}

// A blend kernel that creates an image using the maximum values of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMax
func (bc _BlendKernelClass) ComponentMax() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("componentMax"))
	return rv
}

// A blend kernel that creates an image using the minimum values of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMin
func (bc _BlendKernelClass) ComponentMin() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("componentMin"))
	return rv
}

// A blend kernel that multiplies the color components of its input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMultiply
func (bc _BlendKernelClass) ComponentMultiply() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("componentMultiply"))
	return rv
}

// A blend kernel that creates an image using the darker values of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/darken
func (bc _BlendKernelClass) Darken() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("darken"))
	return rv
}

// A blend kernel that creates an image using the darker color of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/darkerColor
func (bc _BlendKernelClass) DarkerColor() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("darkerColor"))
	return rv
}

// A blend kernel that returns the background input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destination
func (bc _BlendKernelClass) Destination() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("destination"))
	return rv
}

// A blend kernel that places the background over the foreground and crops based on the visibility of the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationAtop
func (bc _BlendKernelClass) DestinationAtop() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("destinationAtop"))
	return rv
}

// A blend kernel that places the background over the foreground and crops based on the visibility of both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationIn
func (bc _BlendKernelClass) DestinationIn() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("destinationIn"))
	return rv
}

// A blend kernel that uses the background image to define what to take out of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationOut
func (bc _BlendKernelClass) DestinationOut() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("destinationOut"))
	return rv
}

// A blend kernel that places the background image over the input foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationOver
func (bc _BlendKernelClass) DestinationOver() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("destinationOver"))
	return rv
}

// A blend kernel that creates an image using the difference between the background and foreground images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/difference
func (bc _BlendKernelClass) Difference() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("difference"))
	return rv
}

// A blend kernel that divides the background image sample color with the foreground image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/divide
func (bc _BlendKernelClass) Divide() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("divide"))
	return rv
}

// A blend kernel that produces an effect similar to difference blending but with lower contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/exclusion
func (bc _BlendKernelClass) Exclusion() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("exclusion"))
	return rv
}

// A blend kernel that returns either the foreground or background image if the other contains a clear color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/exclusiveOr
func (bc _BlendKernelClass) ExclusiveOr() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("exclusiveOr"))
	return rv
}

// A blend kernel that either multiplies or screens colors, depending on the source image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hardLight
func (bc _BlendKernelClass) HardLight() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("hardLight"))
	return rv
}

// A blend kernel that adds two images together, setting each color channel value to either 0 or 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hardMix
func (bc _BlendKernelClass) HardMix() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("hardMix"))
	return rv
}

// A blend kernel that uses the luminance and saturation values of the background image with the hue of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hue
func (bc _BlendKernelClass) Hue() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("hue"))
	return rv
}

// A blend kernel that creates an image using the lighter values of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/lighten
func (bc _BlendKernelClass) Lighten() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("lighten"))
	return rv
}

// A blend kernel that creates an image using the lighter color of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/lighterColor
func (bc _BlendKernelClass) LighterColor() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("lighterColor"))
	return rv
}

// A blend kernel that darkens the background image samples to reflect the foreground image samples while also increasing contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearBurn
func (bc _BlendKernelClass) LinearBurn() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("linearBurn"))
	return rv
}

// A blend kernel that lightens the background image samples to reflect the foreground image samples while also increasing contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearDodge
func (bc _BlendKernelClass) LinearDodge() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("linearDodge"))
	return rv
}

// A blend kernel that burns or dodges colors by changing brightness, depending on the blend color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearLight
func (bc _BlendKernelClass) LinearLight() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("linearLight"))
	return rv
}

// A blend kernel that uses the hue and saturation of the background image with the luminance of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/luminosity
func (bc _BlendKernelClass) Luminosity() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("luminosity"))
	return rv
}

// A blend kernel that multiplies the background image sample color with the foreground image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/multiply
func (bc _BlendKernelClass) Multiply() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("multiply"))
	return rv
}

// A blend kernel that either multiplies or screens the foreground image samples with the background image samples, depending on the background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/overlay
func (bc _BlendKernelClass) Overlay() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("overlay"))
	return rv
}

// A blend kernel that conditionally replaces background image samples with source image samples depending on the brightness of the source image samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/pinLight
func (bc _BlendKernelClass) PinLight() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("pinLight"))
	return rv
}

// A blend kernel that uses the luminance and hue values of the background image with the saturation of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/saturation
func (bc _BlendKernelClass) Saturation() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("saturation"))
	return rv
}

// A blend kernel that multiplies the inverse of the foreground image samples with the inverse of the background image samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/screen
func (bc _BlendKernelClass) Screen() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("screen"))
	return rv
}

// A blend kernel that either darkens or lightens colors, depending on the foreground image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/softLight
func (bc _BlendKernelClass) SoftLight() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("softLight"))
	return rv
}

// A blend kernel that returns the foreground input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/source
func (bc _BlendKernelClass) Source() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("source"))
	return rv
}

// A blend kernel that places the foreground over the background and crops based on the visibility of the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceAtop
func (bc _BlendKernelClass) SourceAtop() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("sourceAtop"))
	return rv
}

// A blend kernel that places the foreground over the background and crops based on the visibility of both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceIn
func (bc _BlendKernelClass) SourceIn() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("sourceIn"))
	return rv
}

// A blend kernel that uses the foreground image to define what to take out of the background image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceOut
func (bc _BlendKernelClass) SourceOut() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("sourceOut"))
	return rv
}

// A blend kernel that places the foreground image over the input background image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceOver
func (bc _BlendKernelClass) SourceOver() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("sourceOver"))
	return rv
}

// A blend kernel that subtracts the background image sample color from the foreground image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/subtract
func (bc _BlendKernelClass) Subtract() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("subtract"))
	return rv
}

// A blend kernel that burns or dodges colors by changing contrast, depending on the blend color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/vividLight
func (bc _BlendKernelClass) VividLight() BlendKernel {
	rv := objc.Send[BlendKernel](objc.ID(bc.class), objc.Sel("vividLight"))
	return rv
}






// Creates a new image using the blend kernel and specified foreground and background images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/apply(foreground:background:)
func (b_ BlendKernel) ApplyWithForegroundBackground(foreground ICIImage, background ICIImage) IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("applyWithForeground:background:"), foreground, background)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/apply(foreground:background:colorSpace:)
func (b_ BlendKernel) ApplyWithForegroundBackgroundColorSpace(foreground ICIImage, background ICIImage, colorSpace ColorSpaceRef /* not a class type */) IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("applyWithForeground:background:colorSpace:"), foreground, background, colorSpace)
	return rv
}







// A blend kernel that returns a clear color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/clear
func (b_ BlendKernel) Clear() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("clear"))
	return rv
}


// A blend kernel that uses the luminance values of the background with the hue and saturation values of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/color
func (b_ BlendKernel) Color() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("color"))
	return rv
}


// A blend kernel that darkens the background image samples to reflect the foreground image samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/colorBurn
func (b_ BlendKernel) ColorBurn() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("colorBurn"))
	return rv
}


// A blend kernel that brightens the background image samples to reflect the foreground image samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/colorDodge
func (b_ BlendKernel) ColorDodge() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("colorDodge"))
	return rv
}


// A blend kernel that adds color components to achieve a brightening effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentAdd
func (b_ BlendKernel) ComponentAdd() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("componentAdd"))
	return rv
}


// A blend kernel that creates an image using the maximum values of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMax
func (b_ BlendKernel) ComponentMax() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("componentMax"))
	return rv
}


// A blend kernel that creates an image using the minimum values of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMin
func (b_ BlendKernel) ComponentMin() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("componentMin"))
	return rv
}


// A blend kernel that multiplies the color components of its input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMultiply
func (b_ BlendKernel) ComponentMultiply() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("componentMultiply"))
	return rv
}


// A blend kernel that creates an image using the darker values of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/darken
func (b_ BlendKernel) Darken() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("darken"))
	return rv
}


// A blend kernel that creates an image using the darker color of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/darkerColor
func (b_ BlendKernel) DarkerColor() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("darkerColor"))
	return rv
}


// A blend kernel that returns the background input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destination
func (b_ BlendKernel) Destination() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("destination"))
	return rv
}


// A blend kernel that places the background over the foreground and crops based on the visibility of the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationAtop
func (b_ BlendKernel) DestinationAtop() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("destinationAtop"))
	return rv
}


// A blend kernel that places the background over the foreground and crops based on the visibility of both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationIn
func (b_ BlendKernel) DestinationIn() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("destinationIn"))
	return rv
}


// A blend kernel that uses the background image to define what to take out of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationOut
func (b_ BlendKernel) DestinationOut() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("destinationOut"))
	return rv
}


// A blend kernel that places the background image over the input foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationOver
func (b_ BlendKernel) DestinationOver() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("destinationOver"))
	return rv
}


// A blend kernel that creates an image using the difference between the background and foreground images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/difference
func (b_ BlendKernel) Difference() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("difference"))
	return rv
}


// A blend kernel that divides the background image sample color with the foreground image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/divide
func (b_ BlendKernel) Divide() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("divide"))
	return rv
}


// A blend kernel that produces an effect similar to difference blending but with lower contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/exclusion
func (b_ BlendKernel) Exclusion() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("exclusion"))
	return rv
}


// A blend kernel that returns either the foreground or background image if the other contains a clear color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/exclusiveOr
func (b_ BlendKernel) ExclusiveOr() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("exclusiveOr"))
	return rv
}


// A blend kernel that either multiplies or screens colors, depending on the source image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hardLight
func (b_ BlendKernel) HardLight() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("hardLight"))
	return rv
}


// A blend kernel that adds two images together, setting each color channel value to either 0 or 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hardMix
func (b_ BlendKernel) HardMix() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("hardMix"))
	return rv
}


// A blend kernel that uses the luminance and saturation values of the background image with the hue of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hue
func (b_ BlendKernel) Hue() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("hue"))
	return rv
}


// A blend kernel that creates an image using the lighter values of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/lighten
func (b_ BlendKernel) Lighten() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("lighten"))
	return rv
}


// A blend kernel that creates an image using the lighter color of two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/lighterColor
func (b_ BlendKernel) LighterColor() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("lighterColor"))
	return rv
}


// A blend kernel that darkens the background image samples to reflect the foreground image samples while also increasing contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearBurn
func (b_ BlendKernel) LinearBurn() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("linearBurn"))
	return rv
}


// A blend kernel that lightens the background image samples to reflect the foreground image samples while also increasing contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearDodge
func (b_ BlendKernel) LinearDodge() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("linearDodge"))
	return rv
}


// A blend kernel that burns or dodges colors by changing brightness, depending on the blend color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearLight
func (b_ BlendKernel) LinearLight() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("linearLight"))
	return rv
}


// A blend kernel that uses the hue and saturation of the background image with the luminance of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/luminosity
func (b_ BlendKernel) Luminosity() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("luminosity"))
	return rv
}


// A blend kernel that multiplies the background image sample color with the foreground image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/multiply
func (b_ BlendKernel) Multiply() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("multiply"))
	return rv
}


// A blend kernel that either multiplies or screens the foreground image samples with the background image samples, depending on the background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/overlay
func (b_ BlendKernel) Overlay() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("overlay"))
	return rv
}


// A blend kernel that conditionally replaces background image samples with source image samples depending on the brightness of the source image samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/pinLight
func (b_ BlendKernel) PinLight() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("pinLight"))
	return rv
}


// A blend kernel that uses the luminance and hue values of the background image with the saturation of the foreground image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/saturation
func (b_ BlendKernel) Saturation() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("saturation"))
	return rv
}


// A blend kernel that multiplies the inverse of the foreground image samples with the inverse of the background image samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/screen
func (b_ BlendKernel) Screen() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("screen"))
	return rv
}


// A blend kernel that either darkens or lightens colors, depending on the foreground image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/softLight
func (b_ BlendKernel) SoftLight() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("softLight"))
	return rv
}


// A blend kernel that returns the foreground input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/source
func (b_ BlendKernel) Source() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("source"))
	return rv
}


// A blend kernel that places the foreground over the background and crops based on the visibility of the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceAtop
func (b_ BlendKernel) SourceAtop() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("sourceAtop"))
	return rv
}


// A blend kernel that places the foreground over the background and crops based on the visibility of both.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceIn
func (b_ BlendKernel) SourceIn() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("sourceIn"))
	return rv
}


// A blend kernel that uses the foreground image to define what to take out of the background image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceOut
func (b_ BlendKernel) SourceOut() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("sourceOut"))
	return rv
}


// A blend kernel that places the foreground image over the input background image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceOver
func (b_ BlendKernel) SourceOver() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("sourceOver"))
	return rv
}


// A blend kernel that subtracts the background image sample color from the foreground image sample color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/subtract
func (b_ BlendKernel) Subtract() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("subtract"))
	return rv
}


// A blend kernel that burns or dodges colors by changing contrast, depending on the blend color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/vividLight
func (b_ BlendKernel) VividLight() ICIBlendKernel {
	rv := objc.Send[BlendKernel](b_.ID, objc.Sel("vividLight"))
	return rv
}







