// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

/* debug [functions.gen.go]: Generating 58 functions for ImageIO */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ImageIO Functions (58 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CGAnimateImageAtURLWithBlock func(URLRef, DictionaryRef, ImageSourceAnimationBlock) unsafe.Pointer
	_CGAnimateImageDataWithBlock func(DataRef, DictionaryRef, ImageSourceAnimationBlock) unsafe.Pointer
	_CGImageDestinationAddAuxiliaryDataInfo func(ImageDestinationRef, StringRef, DictionaryRef)
	_CGImageDestinationAddImage func(ImageDestinationRef, ImageRef, DictionaryRef)
	_CGImageDestinationAddImageAndMetadata func(ImageDestinationRef, ImageRef, ImageMetadataRef, DictionaryRef)
	_CGImageDestinationAddImageFromSource func(ImageDestinationRef, ImageSourceRef, uintptr, DictionaryRef)
	_CGImageDestinationCopyImageSource func(ImageDestinationRef, ImageSourceRef, DictionaryRef, unsafe.Pointer) bool
	_CGImageDestinationCopyTypeIdentifiers func() ArrayRef
	_CGImageDestinationCreateWithData func(MutableDataRef, StringRef, uintptr, DictionaryRef) ImageDestinationRef
	_CGImageDestinationCreateWithDataConsumer func(DataConsumerRef, StringRef, uintptr, DictionaryRef) ImageDestinationRef
	_CGImageDestinationCreateWithURL func(URLRef, StringRef, uintptr, DictionaryRef) ImageDestinationRef
	_CGImageDestinationFinalize func(ImageDestinationRef) bool
	_CGImageDestinationGetTypeID func() TypeID
	_CGImageDestinationSetProperties func(ImageDestinationRef, DictionaryRef)
	_CGImageMetadataCopyStringValueWithPath func(ImageMetadataRef, ImageMetadataTagRef, StringRef) StringRef
	_CGImageMetadataCopyTagMatchingImageProperty func(ImageMetadataRef, StringRef, StringRef) ImageMetadataTagRef
	_CGImageMetadataCopyTags func(ImageMetadataRef) ArrayRef
	_CGImageMetadataCopyTagWithPath func(ImageMetadataRef, ImageMetadataTagRef, StringRef) ImageMetadataTagRef
	_CGImageMetadataCreateFromXMPData func(DataRef) ImageMetadataRef
	_CGImageMetadataCreateMutable func() MutableImageMetadataRef
	_CGImageMetadataCreateMutableCopy func(ImageMetadataRef) MutableImageMetadataRef
	_CGImageMetadataCreateXMPData func(ImageMetadataRef, DictionaryRef) DataRef
	_CGImageMetadataEnumerateTagsUsingBlock func(ImageMetadataRef, StringRef, DictionaryRef, ImageMetadataTagBlock)
	_CGImageMetadataGetTypeID func() TypeID
	_CGImageMetadataRegisterNamespaceForPrefix func(MutableImageMetadataRef, StringRef, StringRef, unsafe.Pointer) bool
	_CGImageMetadataRemoveTagWithPath func(MutableImageMetadataRef, ImageMetadataTagRef, StringRef) bool
	_CGImageMetadataSetTagWithPath func(MutableImageMetadataRef, ImageMetadataTagRef, StringRef, ImageMetadataTagRef) bool
	_CGImageMetadataSetValueMatchingImageProperty func(MutableImageMetadataRef, StringRef, StringRef, TypeRef) bool
	_CGImageMetadataSetValueWithPath func(MutableImageMetadataRef, ImageMetadataTagRef, StringRef, TypeRef) bool
	_CGImageMetadataTagCopyName func(ImageMetadataTagRef) StringRef
	_CGImageMetadataTagCopyNamespace func(ImageMetadataTagRef) StringRef
	_CGImageMetadataTagCopyPrefix func(ImageMetadataTagRef) StringRef
	_CGImageMetadataTagCopyQualifiers func(ImageMetadataTagRef) ArrayRef
	_CGImageMetadataTagCopyValue func(ImageMetadataTagRef) TypeRef
	_CGImageMetadataTagCreate func(StringRef, StringRef, StringRef, ImageMetadataType, TypeRef) ImageMetadataTagRef
	_CGImageMetadataTagGetType func(ImageMetadataTagRef) ImageMetadataType
	_CGImageMetadataTagGetTypeID func() TypeID
	_CGImageSourceCopyAuxiliaryDataInfoAtIndex func(ImageSourceRef, uintptr, StringRef) DictionaryRef
	_CGImageSourceCopyMetadataAtIndex func(ImageSourceRef, uintptr, DictionaryRef) ImageMetadataRef
	_CGImageSourceCopyProperties func(ImageSourceRef, DictionaryRef) DictionaryRef
	_CGImageSourceCopyPropertiesAtIndex func(ImageSourceRef, uintptr, DictionaryRef) DictionaryRef
	_CGImageSourceCopyTypeIdentifiers func() ArrayRef
	_CGImageSourceCreateImageAtIndex func(ImageSourceRef, uintptr, DictionaryRef) ImageRef
	_CGImageSourceCreateIncremental func(DictionaryRef) ImageSourceRef
	_CGImageSourceCreateThumbnailAtIndex func(ImageSourceRef, uintptr, DictionaryRef) ImageRef
	_CGImageSourceCreateWithData func(DataRef, DictionaryRef) ImageSourceRef
	_CGImageSourceCreateWithDataProvider func(DataProviderRef, DictionaryRef) ImageSourceRef
	_CGImageSourceCreateWithURL func(URLRef, DictionaryRef) ImageSourceRef
	_CGImageSourceGetCount func(ImageSourceRef) uintptr
	_CGImageSourceGetPrimaryImageIndex func(ImageSourceRef) uintptr
	_CGImageSourceGetStatus func(ImageSourceRef) ImageSourceStatus
	_CGImageSourceGetStatusAtIndex func(ImageSourceRef, uintptr) ImageSourceStatus
	_CGImageSourceGetType func(ImageSourceRef) StringRef
	_CGImageSourceGetTypeID func() TypeID
	_CGImageSourceRemoveCacheAtIndex func(ImageSourceRef, uintptr)
	_CGImageSourceSetAllowableTypes func(ArrayRef) unsafe.Pointer
	_CGImageSourceUpdateData func(ImageSourceRef, DataRef, bool)
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
	tryRegister(&_CGImageMetadataCopyTags, lib, "CGImageMetadataCopyTags")
	tryRegister(&_CGImageMetadataCopyTagWithPath, lib, "CGImageMetadataCopyTagWithPath")
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



// Animate the sequence of images in the Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file at the specified URL.
//
// Added in macOS 10.15.
// Animate the sequence of images in the Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGAnimateImageAtURLWithBlock(_:_:_:)
func CGAnimateImageAtURLWithBlock(url URLRef, options DictionaryRef, block ImageSourceAnimationBlock) unsafe.Pointer {
	return _CGAnimateImageAtURLWithBlock(url, options, block)
}/* debug [functions.gen.go/function]: CGAnimateImageAtURLWithBlock */

// Animate the sequence of images using data from a Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file file.
//
// Added in macOS 10.15.
// Animate the sequence of images using data from a Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGAnimateImageDataWithBlock(_:_:_:)
func CGAnimateImageDataWithBlock(data DataRef, options DictionaryRef, block ImageSourceAnimationBlock) unsafe.Pointer {
	return _CGAnimateImageDataWithBlock(data, options, block)
}/* debug [functions.gen.go/function]: CGAnimateImageDataWithBlock */

// Sets the auxiliary data, such as mattes and depth information, that accompany the image.
//
// Added in macOS 10.13.
// Sets the auxiliary data, such as mattes and depth information, that accompany the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)
func CGImageDestinationAddAuxiliaryDataInfo(idst ImageDestinationRef, auxiliaryImageDataType StringRef, auxiliaryDataInfoDictionary DictionaryRef) {
	_CGImageDestinationAddAuxiliaryDataInfo(idst, auxiliaryImageDataType, auxiliaryDataInfoDictionary)
}/* debug [functions.gen.go/function]: CGImageDestinationAddAuxiliaryDataInfo */

// Adds an image to an image destination.
//
// Added in macOS 10.4.
// Adds an image to an image destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImage(_:_:_:)
func CGImageDestinationAddImage(idst ImageDestinationRef, image ImageRef, properties DictionaryRef) {
	_CGImageDestinationAddImage(idst, image, properties)
}/* debug [functions.gen.go/function]: CGImageDestinationAddImage */

// CGImageDestinationAddImageAndMetadata is a ImageIO function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImageAndMetadata(_:_:_:_:)
func CGImageDestinationAddImageAndMetadata(idst ImageDestinationRef, image ImageRef, metadata ImageMetadataRef, options DictionaryRef) {
	_CGImageDestinationAddImageAndMetadata(idst, image, metadata, options)
}/* debug [functions.gen.go/function]: CGImageDestinationAddImageAndMetadata */

// Adds an image from an image source to an image destination.
//
// Added in macOS 10.4.
// Adds an image from an image source to an image destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImageFromSource(_:_:_:_:)
func CGImageDestinationAddImageFromSource(idst ImageDestinationRef, isrc ImageSourceRef, index uintptr, properties DictionaryRef) {
	_CGImageDestinationAddImageFromSource(idst, isrc, index, properties)
}/* debug [functions.gen.go/function]: CGImageDestinationAddImageFromSource */

// CGImageDestinationCopyImageSource is a ImageIO function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyImageSource(_:_:_:_:)
func CGImageDestinationCopyImageSource(idst ImageDestinationRef, isrc ImageSourceRef, options DictionaryRef, err unsafe.Pointer) bool {
	return _CGImageDestinationCopyImageSource(idst, isrc, options, err)
}/* debug [functions.gen.go/function]: CGImageDestinationCopyImageSource */

// Returns an array of the uniform type identifiers that are supported for image destinations.
//
// Added in macOS 10.4.
// Returns an array of the uniform type identifiers that are supported for image destinations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyTypeIdentifiers()
func CGImageDestinationCopyTypeIdentifiers() ArrayRef {
	return _CGImageDestinationCopyTypeIdentifiers()
}/* debug [functions.gen.go/function]: CGImageDestinationCopyTypeIdentifiers */

// Creates an image destination that writes to a Core Foundation mutable data object.
//
// Added in macOS 10.4.
// Creates an image destination that writes to a Core Foundation mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithData(_:_:_:_:)
func CGImageDestinationCreateWithData(data MutableDataRef, type_ StringRef, count uintptr, options DictionaryRef) ImageDestinationRef {
	return _CGImageDestinationCreateWithData(data, type_, count, options)
}/* debug [functions.gen.go/function]: CGImageDestinationCreateWithData */

// Creates an image destination that writes to the specified data consumer.
//
// Added in macOS 10.4.
// Creates an image destination that writes to the specified data consumer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithDataConsumer(_:_:_:_:)
func CGImageDestinationCreateWithDataConsumer(consumer DataConsumerRef, type_ StringRef, count uintptr, options DictionaryRef) ImageDestinationRef {
	return _CGImageDestinationCreateWithDataConsumer(consumer, type_, count, options)
}/* debug [functions.gen.go/function]: CGImageDestinationCreateWithDataConsumer */

// Creates an image destination that writes image data to the specified URL.
//
// Added in macOS 10.4.
// Creates an image destination that writes image data to the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithURL(_:_:_:_:)
func CGImageDestinationCreateWithURL(url URLRef, type_ StringRef, count uintptr, options DictionaryRef) ImageDestinationRef {
	return _CGImageDestinationCreateWithURL(url, type_, count, options)
}/* debug [functions.gen.go/function]: CGImageDestinationCreateWithURL */

// Writes image data and properties to the data, URL, or data consumer associated with the image destination.
//
// Added in macOS 10.4.
// Writes image data and properties to the data, URL, or data consumer associated with the image destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationFinalize(_:)
func CGImageDestinationFinalize(idst ImageDestinationRef) bool {
	return _CGImageDestinationFinalize(idst)
}/* debug [functions.gen.go/function]: CGImageDestinationFinalize */

// Returns the unique type identifier of an image destination opaque type.
//
// Added in macOS 10.4.
// Returns the unique type identifier of an image destination opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationGetTypeID()
func CGImageDestinationGetTypeID() TypeID {
	return _CGImageDestinationGetTypeID()
}/* debug [functions.gen.go/function]: CGImageDestinationGetTypeID */

// Applies one or more properties to all images in an image destination.
//
// Added in macOS 10.4.
// Applies one or more properties to all images in an image destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationSetProperties(_:_:)
func CGImageDestinationSetProperties(idst ImageDestinationRef, properties DictionaryRef) {
	_CGImageDestinationSetProperties(idst, properties)
}/* debug [functions.gen.go/function]: CGImageDestinationSetProperties */

// Searches the metadata for the specified tag, and returns its string value if it exists.
//
// Added in macOS 10.8.
// Searches the metadata for the specified tag, and returns its string value if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyStringValueWithPath(_:_:_:)
func CGImageMetadataCopyStringValueWithPath(metadata ImageMetadataRef, parent ImageMetadataTagRef, path StringRef) StringRef {
	return _CGImageMetadataCopyStringValueWithPath(metadata, parent, path)
}/* debug [functions.gen.go/function]: CGImageMetadataCopyStringValueWithPath */

// Searches for the specified image property and, if found, returns the corresponding tag object.
//
// Added in macOS 10.8.
// Searches for the specified image property and, if found, returns the corresponding tag object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTagMatchingImageProperty(_:_:_:)
func CGImageMetadataCopyTagMatchingImageProperty(metadata ImageMetadataRef, dictionaryName StringRef, propertyName StringRef) ImageMetadataTagRef {
	return _CGImageMetadataCopyTagMatchingImageProperty(metadata, dictionaryName, propertyName)
}/* debug [functions.gen.go/function]: CGImageMetadataCopyTagMatchingImageProperty */

// Returns an array of root-level metadata tags from the specified metadata object.
//
// Added in macOS 10.8.
// Returns an array of root-level metadata tags from the specified metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTags(_:)
func CGImageMetadataCopyTags(metadata ImageMetadataRef) ArrayRef {
	return _CGImageMetadataCopyTags(metadata)
}/* debug [functions.gen.go/function]: CGImageMetadataCopyTags */

// Searches for a specific metadata tag within a metadata collection.
//
// Added in macOS 10.8.
// Searches for a specific metadata tag within a metadata collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTagWithPath(_:_:_:)
func CGImageMetadataCopyTagWithPath(metadata ImageMetadataRef, parent ImageMetadataTagRef, path StringRef) ImageMetadataTagRef {
	return _CGImageMetadataCopyTagWithPath(metadata, parent, path)
}/* debug [functions.gen.go/function]: CGImageMetadataCopyTagWithPath */

// Creates a collection of metadata tags from the specified XMP data.
//
// Added in macOS 10.8.
// Creates a collection of metadata tags from the specified XMP data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateFromXMPData(_:)
func CGImageMetadataCreateFromXMPData(data DataRef) ImageMetadataRef {
	return _CGImageMetadataCreateFromXMPData(data)
}/* debug [functions.gen.go/function]: CGImageMetadataCreateFromXMPData */

// Creates an empty, mutable image metdata opaque type.
//
// Added in macOS 10.8.
// Creates an empty, mutable image metdata opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateMutable()
func CGImageMetadataCreateMutable() MutableImageMetadataRef {
	return _CGImageMetadataCreateMutable()
}/* debug [functions.gen.go/function]: CGImageMetadataCreateMutable */

// Creates a deep, mutable copy of the specified metadata information.
//
// Added in macOS 10.8.
// Creates a deep, mutable copy of the specified metadata information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateMutableCopy(_:)
func CGImageMetadataCreateMutableCopy(metadata ImageMetadataRef) MutableImageMetadataRef {
	return _CGImageMetadataCreateMutableCopy(metadata)
}/* debug [functions.gen.go/function]: CGImageMetadataCreateMutableCopy */

// Returns a data object that contains the metadata object’s contents serialized into the XMP format.
//
// Added in macOS 10.8.
// Returns a data object that contains the metadata object’s contents serialized into the XMP format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateXMPData(_:_:)
func CGImageMetadataCreateXMPData(metadata ImageMetadataRef, options DictionaryRef) DataRef {
	return _CGImageMetadataCreateXMPData(metadata, options)
}/* debug [functions.gen.go/function]: CGImageMetadataCreateXMPData */

// Enumerates the tags of a metadata object and executes the specified block on each tag.
//
// Added in macOS 10.8.
// Enumerates the tags of a metadata object and executes the specified block on each tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataEnumerateTagsUsingBlock(_:_:_:_:)
func CGImageMetadataEnumerateTagsUsingBlock(metadata ImageMetadataRef, rootPath StringRef, options DictionaryRef, block ImageMetadataTagBlock) {
	_CGImageMetadataEnumerateTagsUsingBlock(metadata, rootPath, options, block)
}/* debug [functions.gen.go/function]: CGImageMetadataEnumerateTagsUsingBlock */

// Returns the type identifier for metadata objects.
//
// Added in macOS 10.8.
// Returns the type identifier for metadata objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataGetTypeID()
func CGImageMetadataGetTypeID() TypeID {
	return _CGImageMetadataGetTypeID()
}/* debug [functions.gen.go/function]: CGImageMetadataGetTypeID */

// Registers the specified namespace and prefix with the metadata object.
//
// Added in macOS 10.8.
// Registers the specified namespace and prefix with the metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataRegisterNamespaceForPrefix(_:_:_:_:)
func CGImageMetadataRegisterNamespaceForPrefix(metadata MutableImageMetadataRef, xmlns StringRef, prefix StringRef, err unsafe.Pointer) bool {
	return _CGImageMetadataRegisterNamespaceForPrefix(metadata, xmlns, prefix, err)
}/* debug [functions.gen.go/function]: CGImageMetadataRegisterNamespaceForPrefix */

// Removes the tag at the specified path from the metadata object.
//
// Added in macOS 10.8.
// Removes the tag at the specified path from the metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataRemoveTagWithPath(_:_:_:)
func CGImageMetadataRemoveTagWithPath(metadata MutableImageMetadataRef, parent ImageMetadataTagRef, path StringRef) bool {
	return _CGImageMetadataRemoveTagWithPath(metadata, parent, path)
}/* debug [functions.gen.go/function]: CGImageMetadataRemoveTagWithPath */

// Sets the tag at the specified path in the metadata object.
//
// Added in macOS 10.8.
// Sets the tag at the specified path in the metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetTagWithPath(_:_:_:_:)
func CGImageMetadataSetTagWithPath(metadata MutableImageMetadataRef, parent ImageMetadataTagRef, path StringRef, tag ImageMetadataTagRef) bool {
	return _CGImageMetadataSetTagWithPath(metadata, parent, path, tag)
}/* debug [functions.gen.go/function]: CGImageMetadataSetTagWithPath */

// Updates the value of the metadata tag assigned to the specified image property.
//
// Added in macOS 10.8.
// Updates the value of the metadata tag assigned to the specified image property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetValueMatchingImageProperty(_:_:_:_:)
func CGImageMetadataSetValueMatchingImageProperty(metadata MutableImageMetadataRef, dictionaryName StringRef, propertyName StringRef, value TypeRef) bool {
	return _CGImageMetadataSetValueMatchingImageProperty(metadata, dictionaryName, propertyName, value)
}/* debug [functions.gen.go/function]: CGImageMetadataSetValueMatchingImageProperty */

// Update the value of an existing metadata tag, or create a new tag using the specified information.
//
// Added in macOS 10.8.
// Update the value of an existing metadata tag, or create a new tag using the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetValueWithPath(_:_:_:_:)
func CGImageMetadataSetValueWithPath(metadata MutableImageMetadataRef, parent ImageMetadataTagRef, path StringRef, value TypeRef) bool {
	return _CGImageMetadataSetValueWithPath(metadata, parent, path, value)
}/* debug [functions.gen.go/function]: CGImageMetadataSetValueWithPath */

// Returns an immutable copy of the tag’s name.
//
// Added in macOS 10.8.
// Returns an immutable copy of the tag’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyName(_:)
func CGImageMetadataTagCopyName(tag ImageMetadataTagRef) StringRef {
	return _CGImageMetadataTagCopyName(tag)
}/* debug [functions.gen.go/function]: CGImageMetadataTagCopyName */

// Returns an immutable copy of the tag’s XMP namespace.
//
// Added in macOS 10.8.
// Returns an immutable copy of the tag’s XMP namespace.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyNamespace(_:)
func CGImageMetadataTagCopyNamespace(tag ImageMetadataTagRef) StringRef {
	return _CGImageMetadataTagCopyNamespace(tag)
}/* debug [functions.gen.go/function]: CGImageMetadataTagCopyNamespace */

// Returns an immutable copy of the tag’s prefix.
//
// Added in macOS 10.8.
// Returns an immutable copy of the tag’s prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyPrefix(_:)
func CGImageMetadataTagCopyPrefix(tag ImageMetadataTagRef) StringRef {
	return _CGImageMetadataTagCopyPrefix(tag)
}/* debug [functions.gen.go/function]: CGImageMetadataTagCopyPrefix */

// Returns a shallow copy of the metadata tags that act as qualifiers for the current tag.
//
// Added in macOS 10.8.
// Returns a shallow copy of the metadata tags that act as qualifiers for the current tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyQualifiers(_:)
func CGImageMetadataTagCopyQualifiers(tag ImageMetadataTagRef) ArrayRef {
	return _CGImageMetadataTagCopyQualifiers(tag)
}/* debug [functions.gen.go/function]: CGImageMetadataTagCopyQualifiers */

// Returns a shallow copy of the tag’s value, which is suitable only for reading.
//
// Added in macOS 10.8.
// Returns a shallow copy of the tag’s value, which is suitable only for reading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyValue(_:)
func CGImageMetadataTagCopyValue(tag ImageMetadataTagRef) TypeRef {
	return _CGImageMetadataTagCopyValue(tag)
}/* debug [functions.gen.go/function]: CGImageMetadataTagCopyValue */

// Creates a new image metadata tag, and fills it with the specified information.
//
// Added in macOS 10.8.
// Creates a new image metadata tag, and fills it with the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCreate(_:_:_:_:_:)
func CGImageMetadataTagCreate(xmlns StringRef, prefix StringRef, name StringRef, type_ ImageMetadataType, value TypeRef) ImageMetadataTagRef {
	return _CGImageMetadataTagCreate(xmlns, prefix, name, type_, value)
}/* debug [functions.gen.go/function]: CGImageMetadataTagCreate */

// Returns the type of the metadata tag’s value.
//
// Added in macOS 10.8.
// Returns the type of the metadata tag’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagGetType(_:)
func CGImageMetadataTagGetType(tag ImageMetadataTagRef) ImageMetadataType {
	return _CGImageMetadataTagGetType(tag)
}/* debug [functions.gen.go/function]: CGImageMetadataTagGetType */

// Returns the type identifier for the image metadata tag opaque type
//
// Added in macOS 10.8.
// Returns the type identifier for the image metadata tag opaque type
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagGetTypeID()
func CGImageMetadataTagGetTypeID() TypeID {
	return _CGImageMetadataTagGetTypeID()
}/* debug [functions.gen.go/function]: CGImageMetadataTagGetTypeID */

// Returns auxiliary data, such as mattes and depth information, that accompany the image.
//
// Added in macOS 10.13.
// Returns auxiliary data, such as mattes and depth information, that accompany the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
func CGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc ImageSourceRef, index uintptr, auxiliaryImageDataType StringRef) DictionaryRef {
	return _CGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc, index, auxiliaryImageDataType)
}/* debug [functions.gen.go/function]: CGImageSourceCopyAuxiliaryDataInfoAtIndex */

// CGImageSourceCopyMetadataAtIndex is a ImageIO function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyMetadataAtIndex(_:_:_:)
func CGImageSourceCopyMetadataAtIndex(isrc ImageSourceRef, index uintptr, options DictionaryRef) ImageMetadataRef {
	return _CGImageSourceCopyMetadataAtIndex(isrc, index, options)
}/* debug [functions.gen.go/function]: CGImageSourceCopyMetadataAtIndex */

// Returns the properties of the image source.
//
// Added in macOS 10.4.
// Returns the properties of the image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyProperties(_:_:)
func CGImageSourceCopyProperties(isrc ImageSourceRef, options DictionaryRef) DictionaryRef {
	return _CGImageSourceCopyProperties(isrc, options)
}/* debug [functions.gen.go/function]: CGImageSourceCopyProperties */

// Returns the properties of the image at a specified location in an image source.
//
// Added in macOS 10.4.
// Returns the properties of the image at a specified location in an image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyPropertiesAtIndex(_:_:_:)
func CGImageSourceCopyPropertiesAtIndex(isrc ImageSourceRef, index uintptr, options DictionaryRef) DictionaryRef {
	return _CGImageSourceCopyPropertiesAtIndex(isrc, index, options)
}/* debug [functions.gen.go/function]: CGImageSourceCopyPropertiesAtIndex */

// Returns an array of uniform type identifiers that are supported for image sources.
//
// Added in macOS 10.4.
// Returns an array of uniform type identifiers that are supported for image sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyTypeIdentifiers()
func CGImageSourceCopyTypeIdentifiers() ArrayRef {
	return _CGImageSourceCopyTypeIdentifiers()
}/* debug [functions.gen.go/function]: CGImageSourceCopyTypeIdentifiers */

// Creates an image object from the data at the specified index in an image source.
//
// Added in macOS 10.4.
// Creates an image object from the data at the specified index in an image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateImageAtIndex(_:_:_:)
func CGImageSourceCreateImageAtIndex(isrc ImageSourceRef, index uintptr, options DictionaryRef) ImageRef {
	return _CGImageSourceCreateImageAtIndex(isrc, index, options)
}/* debug [functions.gen.go/function]: CGImageSourceCreateImageAtIndex */

// Creates an empty image source that you can use to accumulate incremental image data.
//
// Added in macOS 10.4.
// Creates an empty image source that you can use to accumulate incremental image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateIncremental(_:)
func CGImageSourceCreateIncremental(options DictionaryRef) ImageSourceRef {
	return _CGImageSourceCreateIncremental(options)
}/* debug [functions.gen.go/function]: CGImageSourceCreateIncremental */

// Creates a thumbnail version of the image at the specified index in an image source.
//
// Added in macOS 10.4.
// Creates a thumbnail version of the image at the specified index in an image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateThumbnailAtIndex(_:_:_:)
func CGImageSourceCreateThumbnailAtIndex(isrc ImageSourceRef, index uintptr, options DictionaryRef) ImageRef {
	return _CGImageSourceCreateThumbnailAtIndex(isrc, index, options)
}/* debug [functions.gen.go/function]: CGImageSourceCreateThumbnailAtIndex */

// Creates an image source that reads from a Core Foundation data object.
//
// Added in macOS 10.4.
// Creates an image source that reads from a Core Foundation data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithData(_:_:)
func CGImageSourceCreateWithData(data DataRef, options DictionaryRef) ImageSourceRef {
	return _CGImageSourceCreateWithData(data, options)
}/* debug [functions.gen.go/function]: CGImageSourceCreateWithData */

// Creates an image source that reads data from the specified data provider.
//
// Added in macOS 10.4.
// Creates an image source that reads data from the specified data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithDataProvider(_:_:)
func CGImageSourceCreateWithDataProvider(provider DataProviderRef, options DictionaryRef) ImageSourceRef {
	return _CGImageSourceCreateWithDataProvider(provider, options)
}/* debug [functions.gen.go/function]: CGImageSourceCreateWithDataProvider */

// Creates an image source that reads from a location specified by a URL.
//
// Added in macOS 10.4.
// Creates an image source that reads from a location specified by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithURL(_:_:)
func CGImageSourceCreateWithURL(url URLRef, options DictionaryRef) ImageSourceRef {
	return _CGImageSourceCreateWithURL(url, options)
}/* debug [functions.gen.go/function]: CGImageSourceCreateWithURL */

// Returns the number of images (not including thumbnails) in the image source.
//
// Added in macOS 10.4.
// Returns the number of images (not including thumbnails) in the image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetCount(_:)
func CGImageSourceGetCount(isrc ImageSourceRef) uintptr {
	return _CGImageSourceGetCount(isrc)
}/* debug [functions.gen.go/function]: CGImageSourceGetCount */

// Returns the index of the primary image for an High Efficiency Image File Format (HEIF) image.
//
// Added in macOS 10.14.
// Returns the index of the primary image for an High Efficiency Image File Format (HEIF) image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetPrimaryImageIndex(_:)
func CGImageSourceGetPrimaryImageIndex(isrc ImageSourceRef) uintptr {
	return _CGImageSourceGetPrimaryImageIndex(isrc)
}/* debug [functions.gen.go/function]: CGImageSourceGetPrimaryImageIndex */

// Return the status of an image source.
//
// Added in macOS 10.4.
// Return the status of an image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetStatus(_:)
func CGImageSourceGetStatus(isrc ImageSourceRef) ImageSourceStatus {
	return _CGImageSourceGetStatus(isrc)
}/* debug [functions.gen.go/function]: CGImageSourceGetStatus */

// Returns the current status of an image at the specified location in the image source.
//
// Added in macOS 10.4.
// Returns the current status of an image at the specified location in the image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetStatusAtIndex(_:_:)
func CGImageSourceGetStatusAtIndex(isrc ImageSourceRef, index uintptr) ImageSourceStatus {
	return _CGImageSourceGetStatusAtIndex(isrc, index)
}/* debug [functions.gen.go/function]: CGImageSourceGetStatusAtIndex */

// Returns the uniform type identifier of the source container.
//
// Added in macOS 10.4.
// Returns the uniform type identifier of the source container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetType(_:)
func CGImageSourceGetType(isrc ImageSourceRef) StringRef {
	return _CGImageSourceGetType(isrc)
}/* debug [functions.gen.go/function]: CGImageSourceGetType */

// Returns the unique type identifier of an image source opaque type.
//
// Added in macOS 10.4.
// Returns the unique type identifier of an image source opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetTypeID()
func CGImageSourceGetTypeID() TypeID {
	return _CGImageSourceGetTypeID()
}/* debug [functions.gen.go/function]: CGImageSourceGetTypeID */

// CGImageSourceRemoveCacheAtIndex is a ImageIO function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceRemoveCacheAtIndex(_:_:)
func CGImageSourceRemoveCacheAtIndex(isrc ImageSourceRef, index uintptr) {
	_CGImageSourceRemoveCacheAtIndex(isrc, index)
}/* debug [functions.gen.go/function]: CGImageSourceRemoveCacheAtIndex */

// CGImageSourceSetAllowableTypes is a ImageIO function.
//
// Added in macOS 14.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceSetAllowableTypes(_:)
func CGImageSourceSetAllowableTypes(allowableTypes ArrayRef) unsafe.Pointer {
	return _CGImageSourceSetAllowableTypes(allowableTypes)
}/* debug [functions.gen.go/function]: CGImageSourceSetAllowableTypes */

// Updates the data in an incremental image source.
//
// Added in macOS 10.4.
// Updates the data in an incremental image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceUpdateData(_:_:_:)
func CGImageSourceUpdateData(isrc ImageSourceRef, data DataRef, final bool) {
	_CGImageSourceUpdateData(isrc, data, final)
}/* debug [functions.gen.go/function]: CGImageSourceUpdateData */

// Updates an incremental image source with a new data provider.
//
// Added in macOS 10.4.
// Updates an incremental image source with a new data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceUpdateDataProvider(_:_:_:)
func CGImageSourceUpdateDataProvider(isrc ImageSourceRef, provider DataProviderRef, final bool) {
	_CGImageSourceUpdateDataProvider(isrc, provider, final)
}/* debug [functions.gen.go/function]: CGImageSourceUpdateDataProvider */




