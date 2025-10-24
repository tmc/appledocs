// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ImageIO Functions (36 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CGAnimateImageAtURLWithBlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGAnimateImageDataWithBlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationAddAuxiliaryDataInfo func(ImageDestinationRef, unsafe.Pointer, unsafe.Pointer)
	_CGImageDestinationAddImage func(ImageDestinationRef, ImageRef, unsafe.Pointer)
	_CGImageDestinationAddImageFromSource func(ImageDestinationRef, ImageSourceRef, uintptr, unsafe.Pointer)
	_CGImageDestinationCopyImageSource func(ImageDestinationRef, ImageSourceRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGImageDestinationCopyTypeIdentifiers func() unsafe.Pointer
	_CGImageDestinationCreateWithData func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) ImageDestinationRef
	_CGImageDestinationCreateWithDataConsumer func(DataConsumerRef, unsafe.Pointer, uintptr, unsafe.Pointer) ImageDestinationRef
	_CGImageDestinationCreateWithURL func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) ImageDestinationRef
	_CGImageDestinationFinalize func(ImageDestinationRef) bool
	_CGImageDestinationGetTypeID func() unsafe.Pointer
	_CGImageDestinationSetProperties func(ImageDestinationRef, unsafe.Pointer)
	_CGImageMetadataCopyTags func(ImageMetadataRef) unsafe.Pointer
	_CGImageMetadataCreateMutableCopy func(ImageMetadataRef) MutableImageMetadataRef
	_CGImageMetadataTagGetTypeID func() unsafe.Pointer
	_CGImageSourceCopyAuxiliaryDataInfoAtIndex func(ImageSourceRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCopyProperties func(ImageSourceRef, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCopyPropertiesAtIndex func(ImageSourceRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCopyTypeIdentifiers func() unsafe.Pointer
	_CGImageSourceCreateImageAtIndex func(ImageSourceRef, uintptr, unsafe.Pointer) ImageRef
	_CGImageSourceCreateIncremental func(unsafe.Pointer) ImageSourceRef
	_CGImageSourceCreateThumbnailAtIndex func(ImageSourceRef, uintptr, unsafe.Pointer) ImageRef
	_CGImageSourceCreateWithData func(unsafe.Pointer, unsafe.Pointer) ImageSourceRef
	_CGImageSourceCreateWithDataProvider func(DataProviderRef, unsafe.Pointer) ImageSourceRef
	_CGImageSourceCreateWithURL func(unsafe.Pointer, unsafe.Pointer) ImageSourceRef
	_CGImageSourceGetCount func(ImageSourceRef) uintptr
	_CGImageSourceGetPrimaryImageIndex func(ImageSourceRef) uintptr
	_CGImageSourceGetStatus func(ImageSourceRef) unsafe.Pointer
	_CGImageSourceGetStatusAtIndex func(ImageSourceRef, uintptr) unsafe.Pointer
	_CGImageSourceGetType func(ImageSourceRef) unsafe.Pointer
	_CGImageSourceGetTypeID func() unsafe.Pointer
	_CGImageSourceRemoveCacheAtIndex func(ImageSourceRef, uintptr)
	_CGImageSourceSetAllowableTypes func(unsafe.Pointer) unsafe.Pointer
	_CGImageSourceUpdateData func(ImageSourceRef, unsafe.Pointer, bool)
	_CGImageSourceUpdateDataProvider func(ImageSourceRef, DataProviderRef, bool)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CGAnimateImageAtURLWithBlock, lib, "CGAnimateImageAtURLWithBlock")
	tryRegister(&_CGAnimateImageDataWithBlock, lib, "CGAnimateImageDataWithBlock")
	tryRegister(&_CGImageDestinationAddAuxiliaryDataInfo, lib, "CGImageDestinationAddAuxiliaryDataInfo")
	tryRegister(&_CGImageDestinationAddImage, lib, "CGImageDestinationAddImage")
	tryRegister(&_CGImageDestinationAddImageFromSource, lib, "CGImageDestinationAddImageFromSource")
	tryRegister(&_CGImageDestinationCopyImageSource, lib, "CGImageDestinationCopyImageSource")
	tryRegister(&_CGImageDestinationCopyTypeIdentifiers, lib, "CGImageDestinationCopyTypeIdentifiers")
	tryRegister(&_CGImageDestinationCreateWithData, lib, "CGImageDestinationCreateWithData")
	tryRegister(&_CGImageDestinationCreateWithDataConsumer, lib, "CGImageDestinationCreateWithDataConsumer")
	tryRegister(&_CGImageDestinationCreateWithURL, lib, "CGImageDestinationCreateWithURL")
	tryRegister(&_CGImageDestinationFinalize, lib, "CGImageDestinationFinalize")
	tryRegister(&_CGImageDestinationGetTypeID, lib, "CGImageDestinationGetTypeID")
	tryRegister(&_CGImageDestinationSetProperties, lib, "CGImageDestinationSetProperties")
	tryRegister(&_CGImageMetadataCopyTags, lib, "CGImageMetadataCopyTags")
	tryRegister(&_CGImageMetadataCreateMutableCopy, lib, "CGImageMetadataCreateMutableCopy")
	tryRegister(&_CGImageMetadataTagGetTypeID, lib, "CGImageMetadataTagGetTypeID")
	tryRegister(&_CGImageSourceCopyAuxiliaryDataInfoAtIndex, lib, "CGImageSourceCopyAuxiliaryDataInfoAtIndex")
	tryRegister(&_CGImageSourceCopyProperties, lib, "CGImageSourceCopyProperties")
	tryRegister(&_CGImageSourceCopyPropertiesAtIndex, lib, "CGImageSourceCopyPropertiesAtIndex")
	tryRegister(&_CGImageSourceCopyTypeIdentifiers, lib, "CGImageSourceCopyTypeIdentifiers")
	tryRegister(&_CGImageSourceCreateImageAtIndex, lib, "CGImageSourceCreateImageAtIndex")
	tryRegister(&_CGImageSourceCreateIncremental, lib, "CGImageSourceCreateIncremental")
	tryRegister(&_CGImageSourceCreateThumbnailAtIndex, lib, "CGImageSourceCreateThumbnailAtIndex")
	tryRegister(&_CGImageSourceCreateWithData, lib, "CGImageSourceCreateWithData")
	tryRegister(&_CGImageSourceCreateWithDataProvider, lib, "CGImageSourceCreateWithDataProvider")
	tryRegister(&_CGImageSourceCreateWithURL, lib, "CGImageSourceCreateWithURL")
	tryRegister(&_CGImageSourceGetCount, lib, "CGImageSourceGetCount")
	tryRegister(&_CGImageSourceGetPrimaryImageIndex, lib, "CGImageSourceGetPrimaryImageIndex")
	tryRegister(&_CGImageSourceGetStatus, lib, "CGImageSourceGetStatus")
	tryRegister(&_CGImageSourceGetStatusAtIndex, lib, "CGImageSourceGetStatusAtIndex")
	tryRegister(&_CGImageSourceGetType, lib, "CGImageSourceGetType")
	tryRegister(&_CGImageSourceGetTypeID, lib, "CGImageSourceGetTypeID")
	tryRegister(&_CGImageSourceRemoveCacheAtIndex, lib, "CGImageSourceRemoveCacheAtIndex")
	tryRegister(&_CGImageSourceSetAllowableTypes, lib, "CGImageSourceSetAllowableTypes")
	tryRegister(&_CGImageSourceUpdateData, lib, "CGImageSourceUpdateData")
	tryRegister(&_CGImageSourceUpdateDataProvider, lib, "CGImageSourceUpdateDataProvider")
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



// Animate the sequence of images in the Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file at the specified URL.
//
// Added in macOS 10.15.
// Animate the sequence of images in the Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGAnimateImageAtURLWithBlock(_:_:_:)
func CGAnimateImageAtURLWithBlock(url unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	return _CGAnimateImageAtURLWithBlock(url, options, block)
}

// Animate the sequence of images using data from a Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file file.
//
// Added in macOS 10.15.
// Animate the sequence of images using data from a Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGAnimateImageDataWithBlock(_:_:_:)
func CGAnimateImageDataWithBlock(data unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	return _CGAnimateImageDataWithBlock(data, options, block)
}

// Sets the auxiliary data, such as mattes and depth information, that accompany the image.
//
// Added in macOS 10.13.
// Sets the auxiliary data, such as mattes and depth information, that accompany the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)
func CGImageDestinationAddAuxiliaryDataInfo(idst ImageDestinationRef, auxiliaryImageDataType unsafe.Pointer, auxiliaryDataInfoDictionary unsafe.Pointer) {
	_CGImageDestinationAddAuxiliaryDataInfo(idst, auxiliaryImageDataType, auxiliaryDataInfoDictionary)
}

// Adds an image to an image destination.
//
// Added in macOS 10.4.
// Adds an image to an image destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImage(_:_:_:)
func CGImageDestinationAddImage(idst ImageDestinationRef, image ImageRef, properties unsafe.Pointer) {
	_CGImageDestinationAddImage(idst, image, properties)
}

// Adds an image from an image source to an image destination.
//
// Added in macOS 10.4.
// Adds an image from an image source to an image destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImageFromSource(_:_:_:_:)
func CGImageDestinationAddImageFromSource(idst ImageDestinationRef, isrc ImageSourceRef, index uintptr, properties unsafe.Pointer) {
	_CGImageDestinationAddImageFromSource(idst, isrc, index, properties)
}

// CGImageDestinationCopyImageSource is a ImageIO function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyImageSource(_:_:_:_:)
func CGImageDestinationCopyImageSource(idst ImageDestinationRef, isrc ImageSourceRef, options unsafe.Pointer, err unsafe.Pointer) bool {
	return _CGImageDestinationCopyImageSource(idst, isrc, options, err)
}

// Returns an array of the uniform type identifiers that are supported for image destinations.
//
// Added in macOS 10.4.
// Returns an array of the uniform type identifiers that are supported for image destinations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyTypeIdentifiers()
func CGImageDestinationCopyTypeIdentifiers() unsafe.Pointer {
	return _CGImageDestinationCopyTypeIdentifiers()
}

// Creates an image destination that writes to a Core Foundation mutable data object.
//
// Added in macOS 10.4.
// Creates an image destination that writes to a Core Foundation mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithData(_:_:_:_:)
func CGImageDestinationCreateWithData(data unsafe.Pointer, type_ unsafe.Pointer, count uintptr, options unsafe.Pointer) ImageDestinationRef {
	return _CGImageDestinationCreateWithData(data, type_, count, options)
}

// Creates an image destination that writes to the specified data consumer.
//
// Added in macOS 10.4.
// Creates an image destination that writes to the specified data consumer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithDataConsumer(_:_:_:_:)
func CGImageDestinationCreateWithDataConsumer(consumer DataConsumerRef, type_ unsafe.Pointer, count uintptr, options unsafe.Pointer) ImageDestinationRef {
	return _CGImageDestinationCreateWithDataConsumer(consumer, type_, count, options)
}

// Creates an image destination that writes image data to the specified URL.
//
// Added in macOS 10.4.
// Creates an image destination that writes image data to the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithURL(_:_:_:_:)
func CGImageDestinationCreateWithURL(url unsafe.Pointer, type_ unsafe.Pointer, count uintptr, options unsafe.Pointer) ImageDestinationRef {
	return _CGImageDestinationCreateWithURL(url, type_, count, options)
}

// Writes image data and properties to the data, URL, or data consumer associated with the image destination.
//
// Added in macOS 10.4.
// Writes image data and properties to the data, URL, or data consumer associated with the image destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationFinalize(_:)
func CGImageDestinationFinalize(idst ImageDestinationRef) bool {
	return _CGImageDestinationFinalize(idst)
}

// Returns the unique type identifier of an image destination opaque type.
//
// Added in macOS 10.4.
// Returns the unique type identifier of an image destination opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationGetTypeID()
func CGImageDestinationGetTypeID() unsafe.Pointer {
	return _CGImageDestinationGetTypeID()
}

// Applies one or more properties to all images in an image destination.
//
// Added in macOS 10.4.
// Applies one or more properties to all images in an image destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationSetProperties(_:_:)
func CGImageDestinationSetProperties(idst ImageDestinationRef, properties unsafe.Pointer) {
	_CGImageDestinationSetProperties(idst, properties)
}

// Returns an array of root-level metadata tags from the specified metadata object.
//
// Added in macOS 10.8.
// Returns an array of root-level metadata tags from the specified metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTags(_:)
func CGImageMetadataCopyTags(metadata ImageMetadataRef) unsafe.Pointer {
	return _CGImageMetadataCopyTags(metadata)
}

// Creates a deep, mutable copy of the specified metadata information.
//
// Added in macOS 10.8.
// Creates a deep, mutable copy of the specified metadata information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateMutableCopy(_:)
func CGImageMetadataCreateMutableCopy(metadata ImageMetadataRef) MutableImageMetadataRef {
	return _CGImageMetadataCreateMutableCopy(metadata)
}

// Returns the type identifier for the image metadata tag opaque type
//
// Added in macOS 10.8.
// Returns the type identifier for the image metadata tag opaque type
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagGetTypeID()
func CGImageMetadataTagGetTypeID() unsafe.Pointer {
	return _CGImageMetadataTagGetTypeID()
}

// Returns auxiliary data, such as mattes and depth information, that accompany the image.
//
// Added in macOS 10.13.
// Returns auxiliary data, such as mattes and depth information, that accompany the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
func CGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc ImageSourceRef, index uintptr, auxiliaryImageDataType unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc, index, auxiliaryImageDataType)
}

// Returns the properties of the image source.
//
// Added in macOS 10.4.
// Returns the properties of the image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyProperties(_:_:)
func CGImageSourceCopyProperties(isrc ImageSourceRef, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCopyProperties(isrc, options)
}

// Returns the properties of the image at a specified location in an image source.
//
// Added in macOS 10.4.
// Returns the properties of the image at a specified location in an image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyPropertiesAtIndex(_:_:_:)
func CGImageSourceCopyPropertiesAtIndex(isrc ImageSourceRef, index uintptr, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCopyPropertiesAtIndex(isrc, index, options)
}

// Returns an array of uniform type identifiers that are supported for image sources.
//
// Added in macOS 10.4.
// Returns an array of uniform type identifiers that are supported for image sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyTypeIdentifiers()
func CGImageSourceCopyTypeIdentifiers() unsafe.Pointer {
	return _CGImageSourceCopyTypeIdentifiers()
}

// Creates an image object from the data at the specified index in an image source.
//
// Added in macOS 10.4.
// Creates an image object from the data at the specified index in an image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateImageAtIndex(_:_:_:)
func CGImageSourceCreateImageAtIndex(isrc ImageSourceRef, index uintptr, options unsafe.Pointer) ImageRef {
	return _CGImageSourceCreateImageAtIndex(isrc, index, options)
}

// Creates an empty image source that you can use to accumulate incremental image data.
//
// Added in macOS 10.4.
// Creates an empty image source that you can use to accumulate incremental image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateIncremental(_:)
func CGImageSourceCreateIncremental(options unsafe.Pointer) ImageSourceRef {
	return _CGImageSourceCreateIncremental(options)
}

// Creates a thumbnail version of the image at the specified index in an image source.
//
// Added in macOS 10.4.
// Creates a thumbnail version of the image at the specified index in an image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateThumbnailAtIndex(_:_:_:)
func CGImageSourceCreateThumbnailAtIndex(isrc ImageSourceRef, index uintptr, options unsafe.Pointer) ImageRef {
	return _CGImageSourceCreateThumbnailAtIndex(isrc, index, options)
}

// Creates an image source that reads from a Core Foundation data object.
//
// Added in macOS 10.4.
// Creates an image source that reads from a Core Foundation data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithData(_:_:)
func CGImageSourceCreateWithData(data unsafe.Pointer, options unsafe.Pointer) ImageSourceRef {
	return _CGImageSourceCreateWithData(data, options)
}

// Creates an image source that reads data from the specified data provider.
//
// Added in macOS 10.4.
// Creates an image source that reads data from the specified data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithDataProvider(_:_:)
func CGImageSourceCreateWithDataProvider(provider DataProviderRef, options unsafe.Pointer) ImageSourceRef {
	return _CGImageSourceCreateWithDataProvider(provider, options)
}

// Creates an image source that reads from a location specified by a URL.
//
// Added in macOS 10.4.
// Creates an image source that reads from a location specified by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithURL(_:_:)
func CGImageSourceCreateWithURL(url unsafe.Pointer, options unsafe.Pointer) ImageSourceRef {
	return _CGImageSourceCreateWithURL(url, options)
}

// Returns the number of images (not including thumbnails) in the image source.
//
// Added in macOS 10.4.
// Returns the number of images (not including thumbnails) in the image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetCount(_:)
func CGImageSourceGetCount(isrc ImageSourceRef) uintptr {
	return _CGImageSourceGetCount(isrc)
}

// Returns the index of the primary image for an High Efficiency Image File Format (HEIF) image.
//
// Added in macOS 10.14.
// Returns the index of the primary image for an High Efficiency Image File Format (HEIF) image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetPrimaryImageIndex(_:)
func CGImageSourceGetPrimaryImageIndex(isrc ImageSourceRef) uintptr {
	return _CGImageSourceGetPrimaryImageIndex(isrc)
}

// Return the status of an image source.
//
// Added in macOS 10.4.
// Return the status of an image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetStatus(_:)
func CGImageSourceGetStatus(isrc ImageSourceRef) unsafe.Pointer {
	return _CGImageSourceGetStatus(isrc)
}

// Returns the current status of an image at the specified location in the image source.
//
// Added in macOS 10.4.
// Returns the current status of an image at the specified location in the image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetStatusAtIndex(_:_:)
func CGImageSourceGetStatusAtIndex(isrc ImageSourceRef, index uintptr) unsafe.Pointer {
	return _CGImageSourceGetStatusAtIndex(isrc, index)
}

// Returns the uniform type identifier of the source container.
//
// Added in macOS 10.4.
// Returns the uniform type identifier of the source container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetType(_:)
func CGImageSourceGetType(isrc ImageSourceRef) unsafe.Pointer {
	return _CGImageSourceGetType(isrc)
}

// Returns the unique type identifier of an image source opaque type.
//
// Added in macOS 10.4.
// Returns the unique type identifier of an image source opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetTypeID()
func CGImageSourceGetTypeID() unsafe.Pointer {
	return _CGImageSourceGetTypeID()
}

// CGImageSourceRemoveCacheAtIndex is a ImageIO function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceRemoveCacheAtIndex(_:_:)
func CGImageSourceRemoveCacheAtIndex(isrc ImageSourceRef, index uintptr) {
	_CGImageSourceRemoveCacheAtIndex(isrc, index)
}

// CGImageSourceSetAllowableTypes is a ImageIO function.
//
// Added in macOS 14.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceSetAllowableTypes(_:)
func CGImageSourceSetAllowableTypes(allowableTypes unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceSetAllowableTypes(allowableTypes)
}

// Updates the data in an incremental image source.
//
// Added in macOS 10.4.
// Updates the data in an incremental image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceUpdateData(_:_:_:)
func CGImageSourceUpdateData(isrc ImageSourceRef, data unsafe.Pointer, final bool) {
	_CGImageSourceUpdateData(isrc, data, final)
}

// Updates an incremental image source with a new data provider.
//
// Added in macOS 10.4.
// Updates an incremental image source with a new data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceUpdateDataProvider(_:_:_:)
func CGImageSourceUpdateDataProvider(isrc ImageSourceRef, provider DataProviderRef, final bool) {
	_CGImageSourceUpdateDataProvider(isrc, provider, final)
}



