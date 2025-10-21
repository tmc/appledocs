// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ImageIO Functions (58 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CGAnimateImageAtURLWithBlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGAnimateImageDataWithBlock func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationAddAuxiliaryDataInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationAddImage func(unsafe.Pointer, CGImageRef, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationAddImageAndMetadata func(unsafe.Pointer, CGImageRef, imageio.CGImageMetadataRef, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationAddImageFromSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationCopyImageSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationCopyTypeIdentifiers func() unsafe.Pointer
	_CGImageDestinationCreateWithData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationCreateWithDataConsumer func(CGDataConsumerRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationCreateWithURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationFinalize func(unsafe.Pointer) unsafe.Pointer
	_CGImageDestinationGetTypeID func() unsafe.Pointer
	_CGImageDestinationSetProperties func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageMetadataCopyStringValueWithPath func(imageio.CGImageMetadataRef, imageio.CGImageMetadataTagRef, unsafe.Pointer) unsafe.Pointer
	_CGImageMetadataCopyTagMatchingImageProperty func(imageio.CGImageMetadataRef, unsafe.Pointer, unsafe.Pointer) imageio.CGImageMetadataTagRef
	_CGImageMetadataCopyTagWithPath func(imageio.CGImageMetadataRef, imageio.CGImageMetadataTagRef, unsafe.Pointer) imageio.CGImageMetadataTagRef
	_CGImageMetadataCopyTags func(imageio.CGImageMetadataRef) unsafe.Pointer
	_CGImageMetadataCreateFromXMPData func(unsafe.Pointer) imageio.CGImageMetadataRef
	_CGImageMetadataCreateMutable func() imageio.CGMutableImageMetadataRef
	_CGImageMetadataCreateMutableCopy func(imageio.CGImageMetadataRef) imageio.CGMutableImageMetadataRef
	_CGImageMetadataCreateXMPData func(imageio.CGImageMetadataRef, unsafe.Pointer) unsafe.Pointer
	_CGImageMetadataEnumerateTagsUsingBlock func(imageio.CGImageMetadataRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageMetadataGetTypeID func() unsafe.Pointer
	_CGImageMetadataRegisterNamespaceForPrefix func(imageio.CGMutableImageMetadataRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageMetadataRemoveTagWithPath func(imageio.CGMutableImageMetadataRef, imageio.CGImageMetadataTagRef, unsafe.Pointer) unsafe.Pointer
	_CGImageMetadataSetTagWithPath func(imageio.CGMutableImageMetadataRef, imageio.CGImageMetadataTagRef, unsafe.Pointer, imageio.CGImageMetadataTagRef) unsafe.Pointer
	_CGImageMetadataSetValueMatchingImageProperty func(imageio.CGMutableImageMetadataRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageMetadataSetValueWithPath func(imageio.CGMutableImageMetadataRef, imageio.CGImageMetadataTagRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageMetadataTagCopyName func(imageio.CGImageMetadataTagRef) unsafe.Pointer
	_CGImageMetadataTagCopyNamespace func(imageio.CGImageMetadataTagRef) unsafe.Pointer
	_CGImageMetadataTagCopyPrefix func(imageio.CGImageMetadataTagRef) unsafe.Pointer
	_CGImageMetadataTagCopyQualifiers func(imageio.CGImageMetadataTagRef) unsafe.Pointer
	_CGImageMetadataTagCopyValue func(imageio.CGImageMetadataTagRef) unsafe.Pointer
	_CGImageMetadataTagCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) imageio.CGImageMetadataTagRef
	_CGImageMetadataTagGetType func(imageio.CGImageMetadataTagRef) unsafe.Pointer
	_CGImageMetadataTagGetTypeID func() unsafe.Pointer
	_CGImageSourceCopyAuxiliaryDataInfoAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCopyMetadataAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) imageio.CGImageMetadataRef
	_CGImageSourceCopyProperties func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCopyPropertiesAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCopyTypeIdentifiers func() unsafe.Pointer
	_CGImageSourceCreateImageAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGImageRef
	_CGImageSourceCreateIncremental func(unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCreateThumbnailAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGImageRef
	_CGImageSourceCreateWithData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCreateWithDataProvider func(CGDataProviderRef, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceCreateWithURL func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceGetCount func(unsafe.Pointer) unsafe.Pointer
	_CGImageSourceGetPrimaryImageIndex func(unsafe.Pointer) unsafe.Pointer
	_CGImageSourceGetStatus func(unsafe.Pointer) unsafe.Pointer
	_CGImageSourceGetStatusAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceGetType func(unsafe.Pointer) unsafe.Pointer
	_CGImageSourceGetTypeID func() unsafe.Pointer
	_CGImageSourceRemoveCacheAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceSetAllowableTypes func(unsafe.Pointer) unsafe.Pointer
	_CGImageSourceUpdateData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGImageSourceUpdateDataProvider func(unsafe.Pointer, CGDataProviderRef, unsafe.Pointer) unsafe.Pointer
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
	tryRegister(&_CGImageDestinationAddImageAndMetadata, lib, "CGImageDestinationAddImageAndMetadata")
	tryRegister(&_CGImageDestinationAddImageFromSource, lib, "CGImageDestinationAddImageFromSource")
	tryRegister(&_CGImageDestinationCopyImageSource, lib, "CGImageDestinationCopyImageSource")
	tryRegister(&_CGImageDestinationCopyTypeIdentifiers, lib, "CGImageDestinationCopyTypeIdentifiers")
	tryRegister(&_CGImageDestinationCreateWithData, lib, "CGImageDestinationCreateWithData")
	tryRegister(&_CGImageDestinationCreateWithDataConsumer, lib, "CGImageDestinationCreateWithDataConsumer")
	tryRegister(&_CGImageDestinationCreateWithURL, lib, "CGImageDestinationCreateWithURL")
	tryRegister(&_CGImageDestinationFinalize, lib, "CGImageDestinationFinalize")
	tryRegister(&_CGImageDestinationGetTypeID, lib, "CGImageDestinationGetTypeID")
	tryRegister(&_CGImageDestinationSetProperties, lib, "CGImageDestinationSetProperties")
	tryRegister(&_CGImageMetadataCopyStringValueWithPath, lib, "CGImageMetadataCopyStringValueWithPath")
	tryRegister(&_CGImageMetadataCopyTagMatchingImageProperty, lib, "CGImageMetadataCopyTagMatchingImageProperty")
	tryRegister(&_CGImageMetadataCopyTagWithPath, lib, "CGImageMetadataCopyTagWithPath")
	tryRegister(&_CGImageMetadataCopyTags, lib, "CGImageMetadataCopyTags")
	tryRegister(&_CGImageMetadataCreateFromXMPData, lib, "CGImageMetadataCreateFromXMPData")
	tryRegister(&_CGImageMetadataCreateMutable, lib, "CGImageMetadataCreateMutable")
	tryRegister(&_CGImageMetadataCreateMutableCopy, lib, "CGImageMetadataCreateMutableCopy")
	tryRegister(&_CGImageMetadataCreateXMPData, lib, "CGImageMetadataCreateXMPData")
	tryRegister(&_CGImageMetadataEnumerateTagsUsingBlock, lib, "CGImageMetadataEnumerateTagsUsingBlock")
	tryRegister(&_CGImageMetadataGetTypeID, lib, "CGImageMetadataGetTypeID")
	tryRegister(&_CGImageMetadataRegisterNamespaceForPrefix, lib, "CGImageMetadataRegisterNamespaceForPrefix")
	tryRegister(&_CGImageMetadataRemoveTagWithPath, lib, "CGImageMetadataRemoveTagWithPath")
	tryRegister(&_CGImageMetadataSetTagWithPath, lib, "CGImageMetadataSetTagWithPath")
	tryRegister(&_CGImageMetadataSetValueMatchingImageProperty, lib, "CGImageMetadataSetValueMatchingImageProperty")
	tryRegister(&_CGImageMetadataSetValueWithPath, lib, "CGImageMetadataSetValueWithPath")
	tryRegister(&_CGImageMetadataTagCopyName, lib, "CGImageMetadataTagCopyName")
	tryRegister(&_CGImageMetadataTagCopyNamespace, lib, "CGImageMetadataTagCopyNamespace")
	tryRegister(&_CGImageMetadataTagCopyPrefix, lib, "CGImageMetadataTagCopyPrefix")
	tryRegister(&_CGImageMetadataTagCopyQualifiers, lib, "CGImageMetadataTagCopyQualifiers")
	tryRegister(&_CGImageMetadataTagCopyValue, lib, "CGImageMetadataTagCopyValue")
	tryRegister(&_CGImageMetadataTagCreate, lib, "CGImageMetadataTagCreate")
	tryRegister(&_CGImageMetadataTagGetType, lib, "CGImageMetadataTagGetType")
	tryRegister(&_CGImageMetadataTagGetTypeID, lib, "CGImageMetadataTagGetTypeID")
	tryRegister(&_CGImageSourceCopyAuxiliaryDataInfoAtIndex, lib, "CGImageSourceCopyAuxiliaryDataInfoAtIndex")
	tryRegister(&_CGImageSourceCopyMetadataAtIndex, lib, "CGImageSourceCopyMetadataAtIndex")
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


// Sets the auxiliary data, such as mattes and depth information, that accompany the image. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)
func CGImageDestinationAddAuxiliaryDataInfo(idst unsafe.Pointer, auxiliaryImageDataType unsafe.Pointer, auxiliaryDataInfoDictionary unsafe.Pointer) {
	_CGImageDestinationAddAuxiliaryDataInfo(idst, auxiliaryImageDataType, auxiliaryDataInfoDictionary)
	}


// Adds an image to an image destination. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImage(_:_:_:)
func CGImageDestinationAddImage(idst unsafe.Pointer, image CGImageRef, properties unsafe.Pointer) {
	_CGImageDestinationAddImage(idst, image, properties)
	}


// CGImageDestinationAddImageAndMetadata is a ImageIO function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImageAndMetadata(_:_:_:_:)
func CGImageDestinationAddImageAndMetadata(idst unsafe.Pointer, image CGImageRef, metadata imageio.CGImageMetadataRef, options unsafe.Pointer) {
	_CGImageDestinationAddImageAndMetadata(idst, image, metadata, options)
	}


// Adds an image from an image source to an image destination. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImageFromSource(_:_:_:_:)
func CGImageDestinationAddImageFromSource(idst unsafe.Pointer, isrc unsafe.Pointer, index unsafe.Pointer, properties unsafe.Pointer) {
	_CGImageDestinationAddImageFromSource(idst, isrc, index, properties)
	}


// CGImageDestinationCopyImageSource is a ImageIO function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyImageSource(_:_:_:_:)
func CGImageDestinationCopyImageSource(idst unsafe.Pointer, isrc unsafe.Pointer, options unsafe.Pointer, err unsafe.Pointer) unsafe.Pointer {
	return _CGImageDestinationCopyImageSource(idst, isrc, options, err)
	}


// Returns an array of the uniform type identifiers that are supported for image destinations. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyTypeIdentifiers()
func CGImageDestinationCopyTypeIdentifiers() unsafe.Pointer {
	return _CGImageDestinationCopyTypeIdentifiers()
	}


// Creates an image destination that writes to a Core Foundation mutable data object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithData(_:_:_:_:)
func CGImageDestinationCreateWithData(data unsafe.Pointer, type_ unsafe.Pointer, count unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageDestinationCreateWithData(data, type_, count, options)
	}


// Creates an image destination that writes to the specified data consumer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithDataConsumer(_:_:_:_:)
func CGImageDestinationCreateWithDataConsumer(consumer CGDataConsumerRef, type_ unsafe.Pointer, count unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageDestinationCreateWithDataConsumer(consumer, type_, count, options)
	}


// Creates an image destination that writes image data to the specified URL. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithURL(_:_:_:_:)
func CGImageDestinationCreateWithURL(url unsafe.Pointer, type_ unsafe.Pointer, count unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageDestinationCreateWithURL(url, type_, count, options)
	}


// Writes image data and properties to the data, URL, or data consumer associated with the image destination. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationFinalize(_:)
func CGImageDestinationFinalize(idst unsafe.Pointer) unsafe.Pointer {
	return _CGImageDestinationFinalize(idst)
	}


// Returns the unique type identifier of an image destination opaque type. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationGetTypeID()
func CGImageDestinationGetTypeID() unsafe.Pointer {
	return _CGImageDestinationGetTypeID()
	}


// Applies one or more properties to all images in an image destination. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationSetProperties(_:_:)
func CGImageDestinationSetProperties(idst unsafe.Pointer, properties unsafe.Pointer) {
	_CGImageDestinationSetProperties(idst, properties)
	}


// Searches the metadata for the specified tag, and returns its string value if it exists. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyStringValueWithPath(_:_:_:)
func CGImageMetadataCopyStringValueWithPath(metadata imageio.CGImageMetadataRef, parent imageio.CGImageMetadataTagRef, path unsafe.Pointer) unsafe.Pointer {
	return _CGImageMetadataCopyStringValueWithPath(metadata, parent, path)
	}


// Searches for the specified image property and, if found, returns the corresponding tag object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTagMatchingImageProperty(_:_:_:)
func CGImageMetadataCopyTagMatchingImageProperty(metadata imageio.CGImageMetadataRef, dictionaryName unsafe.Pointer, propertyName unsafe.Pointer) imageio.CGImageMetadataTagRef {
	return _CGImageMetadataCopyTagMatchingImageProperty(metadata, dictionaryName, propertyName)
	}


// Searches for a specific metadata tag within a metadata collection. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTagWithPath(_:_:_:)
func CGImageMetadataCopyTagWithPath(metadata imageio.CGImageMetadataRef, parent imageio.CGImageMetadataTagRef, path unsafe.Pointer) imageio.CGImageMetadataTagRef {
	return _CGImageMetadataCopyTagWithPath(metadata, parent, path)
	}


// Returns an array of root-level metadata tags from the specified metadata object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTags(_:)
func CGImageMetadataCopyTags(metadata imageio.CGImageMetadataRef) unsafe.Pointer {
	return _CGImageMetadataCopyTags(metadata)
	}


// Creates a collection of metadata tags from the specified XMP data. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateFromXMPData(_:)
func CGImageMetadataCreateFromXMPData(data unsafe.Pointer) imageio.CGImageMetadataRef {
	return _CGImageMetadataCreateFromXMPData(data)
	}


// Creates an empty, mutable image metdata opaque type. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateMutable()
func CGImageMetadataCreateMutable() imageio.CGMutableImageMetadataRef {
	return _CGImageMetadataCreateMutable()
	}


// Creates a deep, mutable copy of the specified metadata information. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateMutableCopy(_:)
func CGImageMetadataCreateMutableCopy(metadata imageio.CGImageMetadataRef) imageio.CGMutableImageMetadataRef {
	return _CGImageMetadataCreateMutableCopy(metadata)
	}


// Returns a data object that contains the metadata object’s contents serialized into the XMP format. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateXMPData(_:_:)
func CGImageMetadataCreateXMPData(metadata imageio.CGImageMetadataRef, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageMetadataCreateXMPData(metadata, options)
	}


// Enumerates the tags of a metadata object and executes the specified block on each tag. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataEnumerateTagsUsingBlock(_:_:_:_:)
func CGImageMetadataEnumerateTagsUsingBlock(metadata imageio.CGImageMetadataRef, rootPath unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) {
	_CGImageMetadataEnumerateTagsUsingBlock(metadata, rootPath, options, block)
	}


// Returns the type identifier for metadata objects. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataGetTypeID()
func CGImageMetadataGetTypeID() unsafe.Pointer {
	return _CGImageMetadataGetTypeID()
	}


// Registers the specified namespace and prefix with the metadata object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataRegisterNamespaceForPrefix(_:_:_:_:)
func CGImageMetadataRegisterNamespaceForPrefix(metadata imageio.CGMutableImageMetadataRef, xmlns unsafe.Pointer, prefix unsafe.Pointer, err unsafe.Pointer) unsafe.Pointer {
	return _CGImageMetadataRegisterNamespaceForPrefix(metadata, xmlns, prefix, err)
	}


// Removes the tag at the specified path from the metadata object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataRemoveTagWithPath(_:_:_:)
func CGImageMetadataRemoveTagWithPath(metadata imageio.CGMutableImageMetadataRef, parent imageio.CGImageMetadataTagRef, path unsafe.Pointer) unsafe.Pointer {
	return _CGImageMetadataRemoveTagWithPath(metadata, parent, path)
	}


// Sets the tag at the specified path in the metadata object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetTagWithPath(_:_:_:_:)
func CGImageMetadataSetTagWithPath(metadata imageio.CGMutableImageMetadataRef, parent imageio.CGImageMetadataTagRef, path unsafe.Pointer, tag imageio.CGImageMetadataTagRef) unsafe.Pointer {
	return _CGImageMetadataSetTagWithPath(metadata, parent, path, tag)
	}


// Updates the value of the metadata tag assigned to the specified image property. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetValueMatchingImageProperty(_:_:_:_:)
func CGImageMetadataSetValueMatchingImageProperty(metadata imageio.CGMutableImageMetadataRef, dictionaryName unsafe.Pointer, propertyName unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CGImageMetadataSetValueMatchingImageProperty(metadata, dictionaryName, propertyName, value)
	}


// Update the value of an existing metadata tag, or create a new tag using the specified information. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetValueWithPath(_:_:_:_:)
func CGImageMetadataSetValueWithPath(metadata imageio.CGMutableImageMetadataRef, parent imageio.CGImageMetadataTagRef, path unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CGImageMetadataSetValueWithPath(metadata, parent, path, value)
	}


// Returns an immutable copy of the tag’s name. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyName(_:)
func CGImageMetadataTagCopyName(tag imageio.CGImageMetadataTagRef) unsafe.Pointer {
	return _CGImageMetadataTagCopyName(tag)
	}


// Returns an immutable copy of the tag’s XMP namespace. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyNamespace(_:)
func CGImageMetadataTagCopyNamespace(tag imageio.CGImageMetadataTagRef) unsafe.Pointer {
	return _CGImageMetadataTagCopyNamespace(tag)
	}


// Returns an immutable copy of the tag’s prefix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyPrefix(_:)
func CGImageMetadataTagCopyPrefix(tag imageio.CGImageMetadataTagRef) unsafe.Pointer {
	return _CGImageMetadataTagCopyPrefix(tag)
	}


// Returns a shallow copy of the metadata tags that act as qualifiers for the current tag. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyQualifiers(_:)
func CGImageMetadataTagCopyQualifiers(tag imageio.CGImageMetadataTagRef) unsafe.Pointer {
	return _CGImageMetadataTagCopyQualifiers(tag)
	}


// Returns a shallow copy of the tag’s value, which is suitable only for reading. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyValue(_:)
func CGImageMetadataTagCopyValue(tag imageio.CGImageMetadataTagRef) unsafe.Pointer {
	return _CGImageMetadataTagCopyValue(tag)
	}


// Creates a new image metadata tag, and fills it with the specified information. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCreate(_:_:_:_:_:)
func CGImageMetadataTagCreate(xmlns unsafe.Pointer, prefix unsafe.Pointer, name unsafe.Pointer, type_ unsafe.Pointer, value unsafe.Pointer) imageio.CGImageMetadataTagRef {
	return _CGImageMetadataTagCreate(xmlns, prefix, name, type_, value)
	}


// Returns the type of the metadata tag’s value. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagGetType(_:)
func CGImageMetadataTagGetType(tag imageio.CGImageMetadataTagRef) unsafe.Pointer {
	return _CGImageMetadataTagGetType(tag)
	}


// Returns the type identifier for the image metadata tag opaque type [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagGetTypeID()
func CGImageMetadataTagGetTypeID() unsafe.Pointer {
	return _CGImageMetadataTagGetTypeID()
	}


// Returns auxiliary data, such as mattes and depth information, that accompany the image. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
func CGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc unsafe.Pointer, index unsafe.Pointer, auxiliaryImageDataType unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc, index, auxiliaryImageDataType)
	}


// CGImageSourceCopyMetadataAtIndex is a ImageIO function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyMetadataAtIndex(_:_:_:)
func CGImageSourceCopyMetadataAtIndex(isrc unsafe.Pointer, index unsafe.Pointer, options unsafe.Pointer) imageio.CGImageMetadataRef {
	return _CGImageSourceCopyMetadataAtIndex(isrc, index, options)
	}


// Returns the properties of the image source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyProperties(_:_:)
func CGImageSourceCopyProperties(isrc unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCopyProperties(isrc, options)
	}


// Returns the properties of the image at a specified location in an image source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyPropertiesAtIndex(_:_:_:)
func CGImageSourceCopyPropertiesAtIndex(isrc unsafe.Pointer, index unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCopyPropertiesAtIndex(isrc, index, options)
	}


// Returns an array of uniform type identifiers that are supported for image sources. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyTypeIdentifiers()
func CGImageSourceCopyTypeIdentifiers() unsafe.Pointer {
	return _CGImageSourceCopyTypeIdentifiers()
	}


// Creates an image object from the data at the specified index in an image source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateImageAtIndex(_:_:_:)
func CGImageSourceCreateImageAtIndex(isrc unsafe.Pointer, index unsafe.Pointer, options unsafe.Pointer) CGImageRef {
	return _CGImageSourceCreateImageAtIndex(isrc, index, options)
	}


// Creates an empty image source that you can use to accumulate incremental image data. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateIncremental(_:)
func CGImageSourceCreateIncremental(options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCreateIncremental(options)
	}


// Creates a thumbnail version of the image at the specified index in an image source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateThumbnailAtIndex(_:_:_:)
func CGImageSourceCreateThumbnailAtIndex(isrc unsafe.Pointer, index unsafe.Pointer, options unsafe.Pointer) CGImageRef {
	return _CGImageSourceCreateThumbnailAtIndex(isrc, index, options)
	}


// Creates an image source that reads from a Core Foundation data object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithData(_:_:)
func CGImageSourceCreateWithData(data unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCreateWithData(data, options)
	}


// Creates an image source that reads data from the specified data provider. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithDataProvider(_:_:)
func CGImageSourceCreateWithDataProvider(provider CGDataProviderRef, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCreateWithDataProvider(provider, options)
	}


// Creates an image source that reads from a location specified by a URL. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithURL(_:_:)
func CGImageSourceCreateWithURL(url unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceCreateWithURL(url, options)
	}


// Returns the number of images (not including thumbnails) in the image source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetCount(_:)
func CGImageSourceGetCount(isrc unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceGetCount(isrc)
	}


// Returns the index of the primary image for an High Efficiency Image File Format (HEIF) image. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetPrimaryImageIndex(_:)
func CGImageSourceGetPrimaryImageIndex(isrc unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceGetPrimaryImageIndex(isrc)
	}


// Return the status of an image source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetStatus(_:)
func CGImageSourceGetStatus(isrc unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceGetStatus(isrc)
	}


// Returns the current status of an image at the specified location in the image source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetStatusAtIndex(_:_:)
func CGImageSourceGetStatusAtIndex(isrc unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceGetStatusAtIndex(isrc, index)
	}


// Returns the uniform type identifier of the source container. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetType(_:)
func CGImageSourceGetType(isrc unsafe.Pointer) unsafe.Pointer {
	return _CGImageSourceGetType(isrc)
	}


// Returns the unique type identifier of an image source opaque type. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetTypeID()
func CGImageSourceGetTypeID() unsafe.Pointer {
	return _CGImageSourceGetTypeID()
	}


// CGImageSourceRemoveCacheAtIndex is a ImageIO function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceRemoveCacheAtIndex(_:_:)
func CGImageSourceRemoveCacheAtIndex(isrc unsafe.Pointer, index unsafe.Pointer) {
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


// Updates the data in an incremental image source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceUpdateData(_:_:_:)
func CGImageSourceUpdateData(isrc unsafe.Pointer, data unsafe.Pointer, final unsafe.Pointer) {
	_CGImageSourceUpdateData(isrc, data, final)
	}


// Updates an incremental image source with a new data provider. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceUpdateDataProvider(_:_:_:)
func CGImageSourceUpdateDataProvider(isrc unsafe.Pointer, provider CGDataProviderRef, final unsafe.Pointer) {
	_CGImageSourceUpdateDataProvider(isrc, provider, final)
	}




