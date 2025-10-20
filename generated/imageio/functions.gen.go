// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ImageIO Functions (14 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CGAnimateImageAtURLWithBlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGAnimateImageDataWithBlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationAddImage func(unsafe.Pointer, CGImageRef, unsafe.Pointer)
	_CGImageDestinationCopyImageSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CGImageDestinationCreateWithURL func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationFinalize func(unsafe.Pointer) bool
	_CGImageMetadataCopyTags func(CGImageMetadataRef) unsafe.Pointer
	_CGImageMetadataCreateMutableCopy func(CGImageMetadataRef) CGMutableImageMetadataRef
	_CGImageMetadataTagGetTypeID func() unsafe.Pointer
	_CGImageSourceCopyTypeIdentifiers func() unsafe.Pointer
	_CGImageSourceCreateIncremental func(unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCreateWithURL func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceRemoveCacheAtIndex func(unsafe.Pointer, uintptr)
	_CGImageSourceSetAllowableTypes func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CGAnimateImageAtURLWithBlock, lib, "CGAnimateImageAtURLWithBlock")
	tryRegister(&_CGAnimateImageDataWithBlock, lib, "CGAnimateImageDataWithBlock")
	tryRegister(&_CGImageDestinationAddImage, lib, "CGImageDestinationAddImage")
	tryRegister(&_CGImageDestinationCopyImageSource, lib, "CGImageDestinationCopyImageSource")
	tryRegister(&_CGImageDestinationCreateWithURL, lib, "CGImageDestinationCreateWithURL")
	tryRegister(&_CGImageDestinationFinalize, lib, "CGImageDestinationFinalize")
	tryRegister(&_CGImageMetadataCopyTags, lib, "CGImageMetadataCopyTags")
	tryRegister(&_CGImageMetadataCreateMutableCopy, lib, "CGImageMetadataCreateMutableCopy")
	tryRegister(&_CGImageMetadataTagGetTypeID, lib, "CGImageMetadataTagGetTypeID")
	tryRegister(&_CGImageSourceCopyTypeIdentifiers, lib, "CGImageSourceCopyTypeIdentifiers")
	tryRegister(&_CGImageSourceCreateIncremental, lib, "CGImageSourceCreateIncremental")
	tryRegister(&_CGImageSourceCreateWithURL, lib, "CGImageSourceCreateWithURL")
	tryRegister(&_CGImageSourceRemoveCacheAtIndex, lib, "CGImageSourceRemoveCacheAtIndex")
	tryRegister(&_CGImageSourceSetAllowableTypes, lib, "CGImageSourceSetAllowableTypes")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Animate the sequence of images in the Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file at the specified URL. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGAnimateImageAtURLWithBlock(_:_:_:)
func CGAnimateImageAtURLWithBlock(url unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	return _CGAnimateImageAtURLWithBlock(url, options, block)
	}


// Animate the sequence of images using data from a Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file file. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGAnimateImageDataWithBlock(_:_:_:)
func CGAnimateImageDataWithBlock(data unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	return _CGAnimateImageDataWithBlock(data, options, block)
	}


// Adds an image to an image destination. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImage(_:_:_:)
func CGImageDestinationAddImage(idst unsafe.Pointer, image CGImageRef, properties unsafe.Pointer) {
	_CGImageDestinationAddImage(idst, image, properties)
	}


// CGImageDestinationCopyImageSource is a ImageIO function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyImageSource(_:_:_:_:)
func CGImageDestinationCopyImageSource(idst unsafe.Pointer, isrc unsafe.Pointer, options unsafe.Pointer, err unsafe.Pointer) bool {
	return _CGImageDestinationCopyImageSource(idst, isrc, options, err)
	}


// Creates an image destination that writes image data to the specified URL. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithURL(_:_:_:_:)
func CGImageDestinationCreateWithURL(url unsafe.Pointer, type_ unsafe.Pointer, count uintptr, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageDestinationCreateWithURL(url, type_, count, options)
	}


// Writes image data and properties to the data, URL, or data consumer associated with the image destination. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationFinalize(_:)
func CGImageDestinationFinalize(idst unsafe.Pointer) bool {
	return _CGImageDestinationFinalize(idst)
	}


// Returns an array of root-level metadata tags from the specified metadata object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTags(_:)
func CGImageMetadataCopyTags(metadata CGImageMetadataRef) unsafe.Pointer {
	return _CGImageMetadataCopyTags(metadata)
	}


// Creates a deep, mutable copy of the specified metadata information. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateMutableCopy(_:)
func CGImageMetadataCreateMutableCopy(metadata CGImageMetadataRef) CGMutableImageMetadataRef {
	return _CGImageMetadataCreateMutableCopy(metadata)
	}


// Returns the type identifier for the image metadata tag opaque type [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagGetTypeID()
func CGImageMetadataTagGetTypeID() unsafe.Pointer {
	return _CGImageMetadataTagGetTypeID()
	}


// Returns an array of uniform type identifiers that are supported for image sources. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyTypeIdentifiers()
func CGImageSourceCopyTypeIdentifiers() unsafe.Pointer {
	return _CGImageSourceCopyTypeIdentifiers()
	}


// Creates an empty image source that you can use to accumulate incremental image data. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateIncremental(_:)
func CGImageSourceCreateIncremental(options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCreateIncremental(options)
	}


// Creates an image source that reads from a location specified by a URL. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithURL(_:_:)
func CGImageSourceCreateWithURL(url unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCreateWithURL(url, options)
	}


// CGImageSourceRemoveCacheAtIndex is a ImageIO function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceRemoveCacheAtIndex(_:_:)
func CGImageSourceRemoveCacheAtIndex(isrc unsafe.Pointer, index uintptr) {
	_CGImageSourceRemoveCacheAtIndex(isrc, index)
	}


// CGImageSourceSetAllowableTypes is a ImageIO function. [Full Topic]
//
// Added in macOS 14.2.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceSetAllowableTypes(_:)
func CGImageSourceSetAllowableTypes(allowableTypes unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceSetAllowableTypes(allowableTypes)
	}




