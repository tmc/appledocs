// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	BitmapData() unsafe.Pointer
	BitmapFormat() BitmapFormat
	BitsPerPixel() int
	BytesPerPlane() int
	BytesPerRow() int
	CGImage() ImageRef /* not a class type */
	ColorSpace() IColorSpace
	Planar() bool
	NumberOfPlanes() int
	SamplesPerPixel() int
	TIFFRepresentation() objc.IObject /* cross-framework: NSData */
	IsPlanar() bool
	SetIsPlanar(value bool)
	// methods:
	CanBeCompressedUsing(compression TIFFCompression) bool
	ColorAtXY(x int, y int) IColor
	ColorizeByMappingGrayToColorBlackMappingWhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor)
	BitmapImageRepByConvertingToColorSpaceRenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent /* not a class type */) IBitmapImageRep
	GetBitmapDataPlanes(data unsafe.Pointer)
	GetCompressionFactor(compression TIFFCompression, factor unsafe.Pointer)
	GetPixelAtXY(p []uint, x int, y int)
	IncrementalLoadFromDataComplete(data objc.IObject /* cross-framework: NSData */, complete bool) int
	RepresentationUsingTypeProperties(storageType BitmapImageFileType, properties foundation.IDictionary) foundation.Data
	BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) IBitmapImageRep
	SetColorAtXY(color IColor, x int, y int)
	SetCompressionFactor(compression TIFFCompression, factor float32)
	SetPixelAtXY(p []uint, x int, y int)
	SetPropertyWithValue(property objc.IObject /* cross-framework: BitmapImageRepPropertyKey */, value objc.IObject)
	TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) foundation.Data
	ValueForProperty(property objc.IObject /* cross-framework: BitmapImageRepPropertyKey */) objc.ID
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



// Initializes a newly allocated bitmap image representation for incremental loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(forIncrementalLoad:)
func NewBitmapImageRepForIncrementalLoad() BitmapImageRep {
	instance := getBitmapImageRepClass().Alloc()
	rv := objc.Send[BitmapImageRep](instance.ID, objc.Sel("initForIncrementalLoad"))
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated bitmap image representation so it can render the specified image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(bitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bitmapFormat:bytesPerRow:bitsPerPixel:)
func NewBitmapImageRepWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBitmapFormatBytesPerRowBitsPerPixel(planes unsafe.Pointer, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName objc.IObject /* cross-framework: ColorSpaceName */, bitmapFormat BitmapFormat, rBytes int, pBits int) BitmapImageRep {
	instance := getBitmapImageRepClass().Alloc()
	rv := objc.Send[BitmapImageRep](instance.ID, objc.Sel("initWithBitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bitmapFormat:bytesPerRow:bitsPerPixel:"), planes, width, height, bps, spp, alpha, isPlanar, colorSpaceName, bitmapFormat, rBytes, pBits)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated bitmap image representation so it can render the specified image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(bitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bytesPerRow:bitsPerPixel:)
func NewBitmapImageRepWithBitmapDataPlanesPixelsWidePixelsHighBitsPerSampleSamplesPerPixelHasAlphaIsPlanarColorSpaceNameBytesPerRowBitsPerPixel(planes unsafe.Pointer, width int, height int, bps int, spp int, alpha bool, isPlanar bool, colorSpaceName objc.IObject /* cross-framework: ColorSpaceName */, rBytes int, pBits int) BitmapImageRep {
	instance := getBitmapImageRepClass().Alloc()
	rv := objc.Send[BitmapImageRep](instance.ID, objc.Sel("initWithBitmapDataPlanes:pixelsWide:pixelsHigh:bitsPerSample:samplesPerPixel:hasAlpha:isPlanar:colorSpaceName:bytesPerRow:bitsPerPixel:"), planes, width, height, bps, spp, alpha, isPlanar, colorSpaceName, rBytes, pBits)
	rv.Autorelease()
	return rv
}


// Returns a bitmap image representation from a Core Graphics image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(cgImage:)
func NewBitmapImageRepWithCGImage(cgImage ImageRef /* not a class type */) BitmapImageRep {
	instance := getBitmapImageRepClass().Alloc()
	rv := objc.Send[BitmapImageRep](instance.ID, objc.Sel("initWithCGImage:"), cgImage)
	rv.Autorelease()
	return rv
}


// Returns a bitmap image representation from a Core Image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(ciImage:)
func NewBitmapImageRepWithCIImage(ciImage IImage) BitmapImageRep {
	instance := getBitmapImageRepClass().Alloc()
	rv := objc.Send[BitmapImageRep](instance.ID, objc.Sel("initWithCIImage:"), ciImage)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated bitmap image representation from the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(data:)
func NewBitmapImageRepWithData(data objc.IObject /* cross-framework: NSData */) BitmapImageRep {
	instance := getBitmapImageRepClass().Alloc()
	rv := objc.Send[BitmapImageRep](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated bitmap image representation with bitmap data from a rendered image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/init(focusedViewRect:)
func NewBitmapImageRepWithFocusedViewRect(rect objc.IObject /* cross-framework: Rect */) BitmapImageRep {
	instance := getBitmapImageRepClass().Alloc()
	rv := objc.Send[BitmapImageRep](instance.ID, objc.Sel("initWithFocusedViewRect:"), rect)
	rv.Autorelease()
	return rv
}



// Returns by indirection an array of all available compression types that can be used when writing a TIFF image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/getTIFFCompressionTypes(_:count:)
func (bc _BitmapImageRepClass) GetTIFFCompressionTypesCount(list TIFFCompression, numTypes int) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("getTIFFCompressionTypes:count:"), list, numTypes)
}


// Creates and returns a bitmap image representation with the first image in the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/imageRepWithData:
func (bc _BitmapImageRepClass) ImageRepWithData(data objc.IObject /* cross-framework: NSData */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("imageRepWithData:"), data)
	return rv
}


// Creates and returns an array of bitmap image representation objects that correspond to the images in the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/imageReps(with:)
func (bc _BitmapImageRepClass) ImageRepsWithData(data objc.IObject /* cross-framework: NSData */) []ImageRep {
	rv := objc.Send[[]ImageRep](objc.ID(bc.class), objc.Sel("imageRepsWithData:"), data)
	return rv
}


// Returns an autoreleased string containing the localized name for the specified compression type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/localizedName(forTIFFCompressionType:)
func (bc _BitmapImageRepClass) LocalizedNameForTIFFCompressionType(compression TIFFCompression) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(bc.class), objc.Sel("localizedNameForTIFFCompressionType:"), compression)
	return rv
}


// Formats the specified bitmap images using the specified storage type and properties and returns them in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/representationOfImageReps(in:using:properties:)
func (bc _BitmapImageRepClass) RepresentationOfImageRepsInArrayUsingTypeProperties(imageReps []ImageRep, storageType BitmapImageFileType, properties foundation.IDictionary) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(bc.class), objc.Sel("representationOfImageRepsInArray:usingType:properties:"), imageReps, storageType, properties)
	return rv
}


// Returns a TIFF representation of the specified images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentationOfImageReps(in:)
func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray(array []ImageRep) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(bc.class), objc.Sel("TIFFRepresentationOfImageRepsInArray:"), array)
	return rv
}


// Returns a TIFF representation of the specified images using the specified compression scheme and factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentationOfImageReps(in:using:factor:)
func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArrayUsingCompressionFactor(array []ImageRep, comp TIFFCompression, factor float32) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(bc.class), objc.Sel("TIFFRepresentationOfImageRepsInArray:usingCompression:factor:"), array, comp, factor)
	return rv
}


// Tests whether the bitmap image representation can be compressed by the specified compression scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/canBeCompressed(using:)
func (b_ BitmapImageRep) CanBeCompressedUsing(compression TIFFCompression) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("canBeCompressedUsing:"), compression)
	return rv
}


// Returns the color of the pixel at the specified coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/colorAt(x:y:)
func (b_ BitmapImageRep) ColorAtXY(x int, y int) IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("colorAtX:y:"), x, y)
	return rv
}


// Colorizes a grayscale image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/colorize(byMappingGray:to:blackMapping:whiteMapping:)
func (b_ BitmapImageRep) ColorizeByMappingGrayToColorBlackMappingWhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("colorizeByMappingGray:toColor:blackMapping:whiteMapping:"), midPoint, midPointColor, shadowColor, lightColor)
}


// Converts the bitmap image representation to the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/converting(to:renderingIntent:)
func (b_ BitmapImageRep) BitmapImageRepByConvertingToColorSpaceRenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent /* not a class type */) IBitmapImageRep {
	rv := objc.Send[BitmapImageRep](b_.ID, objc.Sel("bitmapImageRepByConvertingToColorSpace:renderingIntent:"), targetSpace, renderingIntent)
	return rv
}


// Returns by indirection bitmap data of the bitmap image representation separated into planes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/getBitmapDataPlanes(_:)
func (b_ BitmapImageRep) GetBitmapDataPlanes(data unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getBitmapDataPlanes:"), data)
}


// Returns by indirection the bitmap image representation’s compression type and compression factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/getCompression(_:factor:)
func (b_ BitmapImageRep) GetCompressionFactor(compression TIFFCompression, factor unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getCompression:factor:"), compression, factor)
}


// Returns by indirection the pixel data for the specified location in the bitmap image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/getPixel(_:atX:y:)
func (b_ BitmapImageRep) GetPixelAtXY(p []uint, x int, y int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getPixel:atX:y:"), p, x, y)
}


// Loads the current image data into an incrementally-loaded image representation and returns the current status of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/incrementalLoad(from:complete:)
func (b_ BitmapImageRep) IncrementalLoadFromDataComplete(data objc.IObject /* cross-framework: NSData */, complete bool) int {
	rv := objc.Send[int](b_.ID, objc.Sel("incrementalLoadFromData:complete:"), data, complete)
	return rv
}


// Formats the bitmap representation’s image data using the specified storage type and properties and returns it in a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/representation(using:properties:)
func (b_ BitmapImageRep) RepresentationUsingTypeProperties(storageType BitmapImageFileType, properties foundation.IDictionary) foundation.Data {
	rv := objc.Send[foundation.Data](b_.ID, objc.Sel("representationUsingType:properties:"), storageType, properties)
	return rv
}


// Changes the color space tag of the bitmap image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/retagging(with:)
func (b_ BitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) IBitmapImageRep {
	rv := objc.Send[BitmapImageRep](b_.ID, objc.Sel("bitmapImageRepByRetaggingWithColorSpace:"), newSpace)
	return rv
}


// Changes the color of the pixel at the specified coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/setColor(_:atX:y:)
func (b_ BitmapImageRep) SetColorAtXY(color IColor, x int, y int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setColor:atX:y:"), color, x, y)
}


// Sets the bitmap image representation’s compression type and compression factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/setCompression(_:factor:)
func (b_ BitmapImageRep) SetCompressionFactor(compression TIFFCompression, factor float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCompression:factor:"), compression, factor)
}


// Sets the bitmap image representation’s pixel at the specified coordinates to the specified raw pixel values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/setPixel(_:atX:y:)
func (b_ BitmapImageRep) SetPixelAtXY(p []uint, x int, y int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPixel:atX:y:"), p, x, y)
}


// Sets the specified property of the bitmap image representation to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/setProperty(_:withValue:)
func (b_ BitmapImageRep) SetPropertyWithValue(property objc.IObject /* cross-framework: BitmapImageRepPropertyKey */, value objc.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setProperty:withValue:"), property, value)
}


// Returns a TIFF representation of the image using the specified compression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentation(using:factor:)
func (b_ BitmapImageRep) TIFFRepresentationUsingCompressionFactor(comp TIFFCompression, factor float32) foundation.Data {
	rv := objc.Send[foundation.Data](b_.ID, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}


// Returns the value for the specified property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/value(forProperty:)
func (b_ BitmapImageRep) ValueForProperty(property objc.IObject /* cross-framework: BitmapImageRepPropertyKey */) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("valueForProperty:"), property)
	return rv
}


// A pointer to the bitmap data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bitmapData
func (b_ BitmapImageRep) BitmapData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bitmapData"))
	return rv
}


// The format of the bitmap image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bitmapFormat
func (b_ BitmapImageRep) BitmapFormat() BitmapFormat {
	rv := objc.Send[BitmapFormat](b_.ID, objc.Sel("bitmapFormat"))
	return rv
}


// The number of bits allocated for each pixel in each plane of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bitsPerPixel
func (b_ BitmapImageRep) BitsPerPixel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("bitsPerPixel"))
	return rv
}


// The number of bytes in each plane or channel of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bytesPerPlane
func (b_ BitmapImageRep) BytesPerPlane() int {
	rv := objc.Send[int](b_.ID, objc.Sel("bytesPerPlane"))
	return rv
}


// The minimum number of bytes required to specify a scan line in each data plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/bytesPerRow
func (b_ BitmapImageRep) BytesPerRow() int {
	rv := objc.Send[int](b_.ID, objc.Sel("bytesPerRow"))
	return rv
}


// A Core Graphics image object based on the bitmap image representation’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/cgImage
func (b_ BitmapImageRep) CGImage() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](b_.ID, objc.Sel("CGImage"))
	return rv
}


// The color space of the bitmap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/colorSpace
func (b_ BitmapImageRep) ColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](b_.ID, objc.Sel("colorSpace"))
	return rv
}


// A Boolean value that indicates whether the image data is in a planar configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/isPlanar
func (b_ BitmapImageRep) Planar() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("planar"))
	return rv
}


// The number of separate planes into which the image data is organized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/numberOfPlanes
func (b_ BitmapImageRep) NumberOfPlanes() int {
	rv := objc.Send[int](b_.ID, objc.Sel("numberOfPlanes"))
	return rv
}


// The number of components for each pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/samplesPerPixel
func (b_ BitmapImageRep) SamplesPerPixel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("samplesPerPixel"))
	return rv
}


// A TIFF representation of the bitmap image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/tiffRepresentation
func (b_ BitmapImageRep) TIFFRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](b_.ID, objc.Sel("TIFFRepresentation"))
	return rv
}


// A Boolean value that indicates whether the image data is in a planar configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/isplanar
func (b_ BitmapImageRep) IsPlanar() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isPlanar"))
	return rv
}


// A Boolean value that indicates whether the image data is in a planar configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbitmapimagerep/isplanar
func (b_ BitmapImageRep) SetIsPlanar(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsPlanar:"), value)
}


