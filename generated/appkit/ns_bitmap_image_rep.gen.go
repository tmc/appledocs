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
	// properties:
	TIFFRepresentation() foundation.objc.IObject /* cross-framework: NSData */
	BitmapData() unsafe.Pointer
	SetBitmapData(value unsafe.Pointer)
	BitmapFormat() unsafe.Pointer
	SetBitmapFormat(value unsafe.Pointer)
	BitsPerPixel() int /* primitive/slice/pointer. */
	SetBitsPerPixel(value int /* primitive/slice/pointer. */)
	BytesPerPlane() int /* primitive/slice/pointer. */
	SetBytesPerPlane(value int /* primitive/slice/pointer. */)
	BytesPerRow() int /* primitive/slice/pointer. */
	SetBytesPerRow(value int /* primitive/slice/pointer. */)
	CgImage() IImage
	SetCgImage(value IImage)
	ColorSpace() IColorSpace
	SetColorSpace(value IColorSpace)
	IsPlanar() bool /* primitive/slice/pointer. */
	SetIsPlanar(value bool /* primitive/slice/pointer. */)
	NumberOfPlanes() int /* primitive/slice/pointer. */
	SetNumberOfPlanes(value int /* primitive/slice/pointer. */)
	SamplesPerPixel() int /* primitive/slice/pointer. */
	SetSamplesPerPixel(value int /* primitive/slice/pointer. */)
	// methods:
	RepresentationUsingTypeProperties(storageType BitmapImageFileType /* not a class type */, properties foundation.IDictionary /* already interface */) objc.IObject /* cross-framework: Data */
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32 /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Data */
}

// An object that renders an image from bitmap data.
//
// Supported bitmap data formats include GIF, JPEG, TIFF, PNG, and various permutations of raw bitmap data.


// An object that renders an image from bitmap data.
//
// [Full Topic]
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



// Formats the bitmap representation’s image data using the specified storage type and properties and returns it in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/representation(using:properties:)
func (b_ BitmapImageRep) RepresentationUsingTypeProperties(storageType BitmapImageFileType /* not a class type */, properties foundation.IDictionary /* already interface */) objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[Data](b_.ID, objc.Sel("representationUsingType:properties:"), storageType, properties)
	return rv
}


// Returns a TIFF representation of the image using the specified compression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentation(using:factor:)
func (b_ BitmapImageRep) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32 /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[Data](b_.ID, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}


// A TIFF representation of the bitmap image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentation
func (b_ BitmapImageRep) TIFFRepresentation() foundation.objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](b_.ID, objc.Sel("TIFFRepresentation"))
	return rv
}


// A pointer to the bitmap data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitmapdata
func (b_ BitmapImageRep) BitmapData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bitmapData"))
	return rv
}


// A pointer to the bitmap data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitmapdata
func (b_ BitmapImageRep) SetBitmapData(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBitmapData:"), value)
}


// The format of the bitmap image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitmapformat
func (b_ BitmapImageRep) BitmapFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bitmapFormat"))
	return rv
}


// The format of the bitmap image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitmapformat
func (b_ BitmapImageRep) SetBitmapFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBitmapFormat:"), value)
}


// The number of bits allocated for each pixel in each plane of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitsperpixel
func (b_ BitmapImageRep) BitsPerPixel() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("bitsPerPixel"))
	return rv
}


// The number of bits allocated for each pixel in each plane of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bitsperpixel
func (b_ BitmapImageRep) SetBitsPerPixel(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBitsPerPixel:"), value)
}


// The number of bytes in each plane or channel of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bytesperplane
func (b_ BitmapImageRep) BytesPerPlane() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("bytesPerPlane"))
	return rv
}


// The number of bytes in each plane or channel of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bytesperplane
func (b_ BitmapImageRep) SetBytesPerPlane(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBytesPerPlane:"), value)
}


// The minimum number of bytes required to specify a scan line in each data plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bytesperrow
func (b_ BitmapImageRep) BytesPerRow() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("bytesPerRow"))
	return rv
}


// The minimum number of bytes required to specify a scan line in each data plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/bytesperrow
func (b_ BitmapImageRep) SetBytesPerRow(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBytesPerRow:"), value)
}


// A Core Graphics image object based on the bitmap image representation’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/cgimage
func (b_ BitmapImageRep) CgImage() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("cgImage"))
	return rv
}


// A Core Graphics image object based on the bitmap image representation’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/cgimage
func (b_ BitmapImageRep) SetCgImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCgImage:"), value)
}


// The color space of the bitmap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/colorspace
func (b_ BitmapImageRep) ColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](b_.ID, objc.Sel("colorSpace"))
	return rv
}


// The color space of the bitmap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/colorspace
func (b_ BitmapImageRep) SetColorSpace(value IColorSpace) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setColorSpace:"), value)
}


// A Boolean value that indicates whether the image data is in a planar configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/isplanar
func (b_ BitmapImageRep) IsPlanar() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isPlanar"))
	return rv
}


// A Boolean value that indicates whether the image data is in a planar configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/isplanar
func (b_ BitmapImageRep) SetIsPlanar(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsPlanar:"), value)
}


// The number of separate planes into which the image data is organized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/numberofplanes
func (b_ BitmapImageRep) NumberOfPlanes() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("numberOfPlanes"))
	return rv
}


// The number of separate planes into which the image data is organized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/numberofplanes
func (b_ BitmapImageRep) SetNumberOfPlanes(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNumberOfPlanes:"), value)
}


// The number of components for each pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/samplesperpixel
func (b_ BitmapImageRep) SamplesPerPixel() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("samplesPerPixel"))
	return rv
}


// The number of components for each pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/samplesperpixel
func (b_ BitmapImageRep) SetSamplesPerPixel(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSamplesPerPixel:"), value)
}



