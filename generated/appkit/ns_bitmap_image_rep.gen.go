// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BitmapImageRep] class.
var (
	BitmapImageRepClass     _BitmapImageRepClass
	BitmapImageRepClassOnce sync.Once
)

func getBitmapImageRepClass() _BitmapImageRepClass {
	BitmapImageRepClassOnce.Do(func() {
		BitmapImageRepClass = _BitmapImageRepClass{objc.GetClass("NSBitmapImageRep")}
	})
	return BitmapImageRepClass
}

type _BitmapImageRepClass struct {
	class objc.Class
}

// An interface definition for the [BitmapImageRep] class.
type IBitmapImageRep interface {
	IImageRep
}

// An object that renders an image from bitmap data.
//
// Supported bitmap data formats include GIF, JPEG, TIFF, PNG, and various permutations of raw bitmap data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep
type BitmapImageRep struct {
	ImageRep
}

// BitmapImageRepFrom constructs a [BitmapImageRep] from an unsafe.Pointer.
//
// An object that renders an image from bitmap data.
func BitmapImageRepFrom(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BitmapImageRepClass) Alloc() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BitmapImageRepClass) New() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BitmapImageRep) Init() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BitmapImageRep) Autorelease() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBitmapImageRep creates a new BitmapImageRep instance.
func NewBitmapImageRep() BitmapImageRep {
	return getBitmapImageRepClass().New()
}


// A pointer to the bitmap data.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitmapdata
func (b_ BitmapImageRep) BitmapData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bitmapData"))
	return rv
}


// SetBitmapData sets the value of the bitmapData property.
// A pointer to the bitmap data.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitmapdata
func (b_ BitmapImageRep) SetBitmapData(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBitmapData:"), value)
}

// The format of the bitmap image representation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitmapformat
func (b_ BitmapImageRep) BitmapFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bitmapFormat"))
	return rv
}


// SetBitmapFormat sets the value of the bitmapFormat property.
// The format of the bitmap image representation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitmapformat
func (b_ BitmapImageRep) SetBitmapFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBitmapFormat:"), value)
}

// The number of bits allocated for each pixel in each plane of data.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitsperpixel
func (b_ BitmapImageRep) BitsPerPixel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("bitsPerPixel"))
	return rv
}


// SetBitsPerPixel sets the value of the bitsPerPixel property.
// The number of bits allocated for each pixel in each plane of data.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitsperpixel
func (b_ BitmapImageRep) SetBitsPerPixel(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBitsPerPixel:"), value)
}

// The number of bytes in each plane or channel of data.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bytesperplane
func (b_ BitmapImageRep) BytesPerPlane() int {
	rv := objc.Send[int](b_.ID, objc.Sel("bytesPerPlane"))
	return rv
}


// SetBytesPerPlane sets the value of the bytesPerPlane property.
// The number of bytes in each plane or channel of data.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bytesperplane
func (b_ BitmapImageRep) SetBytesPerPlane(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBytesPerPlane:"), value)
}

// The minimum number of bytes required to specify a scan line in each data plane.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bytesperrow
func (b_ BitmapImageRep) BytesPerRow() int {
	rv := objc.Send[int](b_.ID, objc.Sel("bytesPerRow"))
	return rv
}


// SetBytesPerRow sets the value of the bytesPerRow property.
// The minimum number of bytes required to specify a scan line in each data plane.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bytesperrow
func (b_ BitmapImageRep) SetBytesPerRow(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBytesPerRow:"), value)
}

// A Core Graphics image object based on the bitmap image representation’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/cgimage
func (b_ BitmapImageRep) CgImage() Image {
	rv := objc.Send[Image](b_.ID, objc.Sel("cgImage"))
	return rv
}


// SetCgImage sets the value of the cgImage property.
// A Core Graphics image object based on the bitmap image representation’s data.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/cgimage
func (b_ BitmapImageRep) SetCgImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCgImage:"), value)
}

// The color space of the bitmap.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/colorspace
func (b_ BitmapImageRep) ColorSpace() NSColorSpace {
	rv := objc.Send[NSColorSpace](b_.ID, objc.Sel("colorSpace"))
	return rv
}


// SetColorSpace sets the value of the colorSpace property.
// The color space of the bitmap.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/colorspace
func (b_ BitmapImageRep) SetColorSpace(value IColorSpace) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setColorSpace:"), value)
}

// A Boolean value that indicates whether the image data is in a planar configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/isplanar
func (b_ BitmapImageRep) IsPlanar() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isPlanar"))
	return rv
}


// SetIsPlanar sets the value of the isPlanar property.
// A Boolean value that indicates whether the image data is in a planar configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/isplanar
func (b_ BitmapImageRep) SetIsPlanar(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsPlanar:"), value)
}

// The number of separate planes into which the image data is organized.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/numberofplanes
func (b_ BitmapImageRep) NumberOfPlanes() int {
	rv := objc.Send[int](b_.ID, objc.Sel("numberOfPlanes"))
	return rv
}


// SetNumberOfPlanes sets the value of the numberOfPlanes property.
// The number of separate planes into which the image data is organized.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/numberofplanes
func (b_ BitmapImageRep) SetNumberOfPlanes(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNumberOfPlanes:"), value)
}

// The number of components for each pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/samplesperpixel
func (b_ BitmapImageRep) SamplesPerPixel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("samplesPerPixel"))
	return rv
}


// SetSamplesPerPixel sets the value of the samplesPerPixel property.
// The number of components for each pixel.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/samplesperpixel
func (b_ BitmapImageRep) SetSamplesPerPixel(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSamplesPerPixel:"), value)
}

// A TIFF representation of the bitmap image data.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/tiffrepresentation
func (b_ BitmapImageRep) TiffRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](b_.ID, objc.Sel("tiffRepresentation"))
	return rv
}


// SetTiffRepresentation sets the value of the tiffRepresentation property.
// A TIFF representation of the bitmap image data.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/tiffrepresentation
func (b_ BitmapImageRep) SetTiffRepresentation(value foundation.IData) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTiffRepresentation:"), value)
}



