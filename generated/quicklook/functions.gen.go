// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

/* debug [functions.gen.go]: Generating 40 functions for QuickLook */
import (
	"unsafe"

	"github.com/ebitengine/purego"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
)


// QuickLook Functions (40 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_QLPreviewRequestCopyContentUTI func(PreviewRequestRef) StringRef
	_QLPreviewRequestCopyOptions func(PreviewRequestRef) DictionaryRef
	_QLPreviewRequestCopyURL func(PreviewRequestRef) URLRef
	_QLPreviewRequestCreateContext func(PreviewRequestRef, corefoundation.CGSize, unsafe.Pointer, DictionaryRef) ContextRef
	_QLPreviewRequestCreatePDFContext func(PreviewRequestRef, unsafe.Pointer, DictionaryRef, DictionaryRef) ContextRef
	_QLPreviewRequestFlushContext func(PreviewRequestRef, ContextRef)
	_QLPreviewRequestGetDocumentObject func(PreviewRequestRef) unsafe.Pointer
	_QLPreviewRequestGetGeneratorBundle func(PreviewRequestRef) BundleRef
	_QLPreviewRequestGetTypeID func() TypeID
	_QLPreviewRequestIsCancelled func(PreviewRequestRef) unsafe.Pointer
	_QLPreviewRequestSetDataRepresentation func(PreviewRequestRef, DataRef, StringRef, DictionaryRef)
	_QLPreviewRequestSetDocumentObject func(PreviewRequestRef, unsafe.Pointer, unsafe.Pointer)
	_QLPreviewRequestSetURLRepresentation func(PreviewRequestRef, URLRef, StringRef, DictionaryRef)
	_QLThumbnailCancel func(ThumbnailRef)
	_QLThumbnailCopyDocumentURL func(ThumbnailRef) URLRef
	_QLThumbnailCopyImage func(ThumbnailRef) ImageRef
	_QLThumbnailCopyOptions func(ThumbnailRef) DictionaryRef
	_QLThumbnailCreate func(AllocatorRef, URLRef, corefoundation.CGSize, DictionaryRef) ThumbnailRef
	_QLThumbnailDispatchAsync func(ThumbnailRef, unsafe.Pointer, unsafe.Pointer)
	_QLThumbnailGetContentRect func(ThumbnailRef) corefoundation.CGRect
	_QLThumbnailGetMaximumSize func(ThumbnailRef) corefoundation.CGSize
	_QLThumbnailGetTypeID func() TypeID
	_QLThumbnailImageCreate func(AllocatorRef, URLRef, corefoundation.CGSize, DictionaryRef) ImageRef
	_QLThumbnailIsCancelled func(ThumbnailRef) unsafe.Pointer
	_QLThumbnailRequestCopyContentUTI func(ThumbnailRequestRef) StringRef
	_QLThumbnailRequestCopyOptions func(ThumbnailRequestRef) DictionaryRef
	_QLThumbnailRequestCopyURL func(ThumbnailRequestRef) URLRef
	_QLThumbnailRequestCreateContext func(ThumbnailRequestRef, corefoundation.CGSize, unsafe.Pointer, DictionaryRef) ContextRef
	_QLThumbnailRequestFlushContext func(ThumbnailRequestRef, ContextRef)
	_QLThumbnailRequestGetDocumentObject func(ThumbnailRequestRef) unsafe.Pointer
	_QLThumbnailRequestGetGeneratorBundle func(ThumbnailRequestRef) BundleRef
	_QLThumbnailRequestGetMaximumSize func(ThumbnailRequestRef) corefoundation.CGSize
	_QLThumbnailRequestGetTypeID func() TypeID
	_QLThumbnailRequestIsCancelled func(ThumbnailRequestRef) unsafe.Pointer
	_QLThumbnailRequestSetDocumentObject func(ThumbnailRequestRef, unsafe.Pointer, unsafe.Pointer)
	_QLThumbnailRequestSetImage func(ThumbnailRequestRef, ImageRef, DictionaryRef)
	_QLThumbnailRequestSetImageAtURL func(ThumbnailRequestRef, URLRef, DictionaryRef)
	_QLThumbnailRequestSetImageWithData func(ThumbnailRequestRef, DataRef, DictionaryRef)
	_QLThumbnailRequestSetThumbnailWithDataRepresentation func(ThumbnailRequestRef, DataRef, StringRef, DictionaryRef, DictionaryRef)
	_QLThumbnailRequestSetThumbnailWithURLRepresentation func(ThumbnailRequestRef, URLRef, StringRef, DictionaryRef, DictionaryRef)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_QLPreviewRequestCopyContentUTI, lib, "QLPreviewRequestCopyContentUTI")
	tryRegister(&_QLPreviewRequestCopyOptions, lib, "QLPreviewRequestCopyOptions")
	tryRegister(&_QLPreviewRequestCopyURL, lib, "QLPreviewRequestCopyURL")
	tryRegister(&_QLPreviewRequestCreateContext, lib, "QLPreviewRequestCreateContext")
	tryRegister(&_QLPreviewRequestCreatePDFContext, lib, "QLPreviewRequestCreatePDFContext")
	tryRegister(&_QLPreviewRequestFlushContext, lib, "QLPreviewRequestFlushContext")
	tryRegister(&_QLPreviewRequestGetDocumentObject, lib, "QLPreviewRequestGetDocumentObject")
	tryRegister(&_QLPreviewRequestGetGeneratorBundle, lib, "QLPreviewRequestGetGeneratorBundle")
	tryRegister(&_QLPreviewRequestGetTypeID, lib, "QLPreviewRequestGetTypeID")
	tryRegister(&_QLPreviewRequestIsCancelled, lib, "QLPreviewRequestIsCancelled")
	tryRegister(&_QLPreviewRequestSetDataRepresentation, lib, "QLPreviewRequestSetDataRepresentation")
	tryRegister(&_QLPreviewRequestSetDocumentObject, lib, "QLPreviewRequestSetDocumentObject")
	tryRegister(&_QLPreviewRequestSetURLRepresentation, lib, "QLPreviewRequestSetURLRepresentation")
	tryRegister(&_QLThumbnailCancel, lib, "QLThumbnailCancel")
	tryRegister(&_QLThumbnailCopyDocumentURL, lib, "QLThumbnailCopyDocumentURL")
	tryRegister(&_QLThumbnailCopyImage, lib, "QLThumbnailCopyImage")
	tryRegister(&_QLThumbnailCopyOptions, lib, "QLThumbnailCopyOptions")
	tryRegister(&_QLThumbnailCreate, lib, "QLThumbnailCreate")
	tryRegister(&_QLThumbnailDispatchAsync, lib, "QLThumbnailDispatchAsync")
	tryRegister(&_QLThumbnailGetContentRect, lib, "QLThumbnailGetContentRect")
	tryRegister(&_QLThumbnailGetMaximumSize, lib, "QLThumbnailGetMaximumSize")
	tryRegister(&_QLThumbnailGetTypeID, lib, "QLThumbnailGetTypeID")
	tryRegister(&_QLThumbnailImageCreate, lib, "QLThumbnailImageCreate")
	tryRegister(&_QLThumbnailIsCancelled, lib, "QLThumbnailIsCancelled")
	tryRegister(&_QLThumbnailRequestCopyContentUTI, lib, "QLThumbnailRequestCopyContentUTI")
	tryRegister(&_QLThumbnailRequestCopyOptions, lib, "QLThumbnailRequestCopyOptions")
	tryRegister(&_QLThumbnailRequestCopyURL, lib, "QLThumbnailRequestCopyURL")
	tryRegister(&_QLThumbnailRequestCreateContext, lib, "QLThumbnailRequestCreateContext")
	tryRegister(&_QLThumbnailRequestFlushContext, lib, "QLThumbnailRequestFlushContext")
	tryRegister(&_QLThumbnailRequestGetDocumentObject, lib, "QLThumbnailRequestGetDocumentObject")
	tryRegister(&_QLThumbnailRequestGetGeneratorBundle, lib, "QLThumbnailRequestGetGeneratorBundle")
	tryRegister(&_QLThumbnailRequestGetMaximumSize, lib, "QLThumbnailRequestGetMaximumSize")
	tryRegister(&_QLThumbnailRequestGetTypeID, lib, "QLThumbnailRequestGetTypeID")
	tryRegister(&_QLThumbnailRequestIsCancelled, lib, "QLThumbnailRequestIsCancelled")
	tryRegister(&_QLThumbnailRequestSetDocumentObject, lib, "QLThumbnailRequestSetDocumentObject")
	tryRegister(&_QLThumbnailRequestSetImage, lib, "QLThumbnailRequestSetImage")
	tryRegister(&_QLThumbnailRequestSetImageAtURL, lib, "QLThumbnailRequestSetImageAtURL")
	tryRegister(&_QLThumbnailRequestSetImageWithData, lib, "QLThumbnailRequestSetImageWithData")
	tryRegister(&_QLThumbnailRequestSetThumbnailWithDataRepresentation, lib, "QLThumbnailRequestSetThumbnailWithDataRepresentation")
	tryRegister(&_QLThumbnailRequestSetThumbnailWithURLRepresentation, lib, "QLThumbnailRequestSetThumbnailWithURLRepresentation")
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



// Returns the UTI for the preview request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns the UTI for the preview request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestCopyContentUTI(_:)
func QLPreviewRequestCopyContentUTI(preview PreviewRequestRef) StringRef {
	return _QLPreviewRequestCopyContentUTI(preview)
}/* debug [functions.gen.go/function]: QLPreviewRequestCopyContentUTI */

// Returns the options specified for the preview request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns the options specified for the preview request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestCopyOptions(_:)
func QLPreviewRequestCopyOptions(preview PreviewRequestRef) DictionaryRef {
	return _QLPreviewRequestCopyOptions(preview)
}/* debug [functions.gen.go/function]: QLPreviewRequestCopyOptions */

// Returns the URL of the document for which a preview is requested.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns the URL of the document for which a preview is requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestCopyURL(_:)
func QLPreviewRequestCopyURL(preview PreviewRequestRef) URLRef {
	return _QLPreviewRequestCopyURL(preview)
}/* debug [functions.gen.go/function]: QLPreviewRequestCopyURL */

// Creates a graphics context to draw the preview in.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Creates a graphics context to draw the preview in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestCreateContext(_:_:_:_:)
func QLPreviewRequestCreateContext(preview PreviewRequestRef, size corefoundation.CGSize, isBitmap unsafe.Pointer, properties DictionaryRef) ContextRef {
	return _QLPreviewRequestCreateContext(preview, size, isBitmap, properties)
}/* debug [functions.gen.go/function]: QLPreviewRequestCreateContext */

// Creates a PDF context suitable to draw a multi-page preview.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Creates a PDF context suitable to draw a multi-page preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestCreatePDFContext(_:_:_:_:)
func QLPreviewRequestCreatePDFContext(preview PreviewRequestRef, mediaBox unsafe.Pointer, auxiliaryInfo DictionaryRef, properties DictionaryRef) ContextRef {
	return _QLPreviewRequestCreatePDFContext(preview, mediaBox, auxiliaryInfo, properties)
}/* debug [functions.gen.go/function]: QLPreviewRequestCreatePDFContext */

// Flush the context and sets the preview response.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Flush the context and sets the preview response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestFlushContext(_:_:)
func QLPreviewRequestFlushContext(preview PreviewRequestRef, context ContextRef) {
	_QLPreviewRequestFlushContext(preview, context)
}/* debug [functions.gen.go/function]: QLPreviewRequestFlushContext */

// Returns the object that’s stored as part of a preview request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.6.
// Returns the object that’s stored as part of a preview request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestGetDocumentObject(_:)
func QLPreviewRequestGetDocumentObject(preview PreviewRequestRef) unsafe.Pointer {
	return _QLPreviewRequestGetDocumentObject(preview)
}/* debug [functions.gen.go/function]: QLPreviewRequestGetDocumentObject */

// Get the bundle of the generator receiving the preview request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Get the bundle of the generator receiving the preview request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestGetGeneratorBundle(_:)
func QLPreviewRequestGetGeneratorBundle(preview PreviewRequestRef) BundleRef {
	return _QLPreviewRequestGetGeneratorBundle(preview)
}/* debug [functions.gen.go/function]: QLPreviewRequestGetGeneratorBundle */

// Gets the type identifier for the opaque type.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Gets the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestGetTypeID()
func QLPreviewRequestGetTypeID() TypeID {
	return _QLPreviewRequestGetTypeID()
}/* debug [functions.gen.go/function]: QLPreviewRequestGetTypeID */

// Returns whether the preview request has been cancelled by the client.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns whether the preview request has been cancelled by the client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestIsCancelled(_:)
func QLPreviewRequestIsCancelled(preview PreviewRequestRef) unsafe.Pointer {
	return _QLPreviewRequestIsCancelled(preview)
}/* debug [functions.gen.go/function]: QLPreviewRequestIsCancelled */

// Sets the preview request to data saved within the document or to dynamically generated data.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Sets the preview request to data saved within the document or to dynamically generated data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestSetDataRepresentation(_:_:_:_:)
func QLPreviewRequestSetDataRepresentation(preview PreviewRequestRef, data DataRef, contentTypeUTI StringRef, properties DictionaryRef) {
	_QLPreviewRequestSetDataRepresentation(preview, data, contentTypeUTI, properties)
}/* debug [functions.gen.go/function]: QLPreviewRequestSetDataRepresentation */

// Stores an object as part of a preview request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.6.
// Stores an object as part of a preview request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestSetDocumentObject(_:_:_:)
func QLPreviewRequestSetDocumentObject(preview PreviewRequestRef, object unsafe.Pointer, callbacks unsafe.Pointer) {
	_QLPreviewRequestSetDocumentObject(preview, object, callbacks)
}/* debug [functions.gen.go/function]: QLPreviewRequestSetDocumentObject */

// Sets the contents of the file at the given URL as the response to the preview request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Sets the contents of the file at the given URL as the response to the preview request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewRequestSetURLRepresentation(_:_:_:_:)
func QLPreviewRequestSetURLRepresentation(preview PreviewRequestRef, url URLRef, contentTypeUTI StringRef, properties DictionaryRef) {
	_QLPreviewRequestSetURLRepresentation(preview, url, contentTypeUTI, properties)
}/* debug [functions.gen.go/function]: QLPreviewRequestSetURLRepresentation */

// Cancels the computation of the thumbnail.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Cancels the computation of the thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailCancel(_:)
func QLThumbnailCancel(thumbnail ThumbnailRef) {
	_QLThumbnailCancel(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailCancel */

// Returns the URL of the document that you’re requesting a thumbnail for.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Returns the URL of the document that you’re requesting a thumbnail for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailCopyDocumentURL(_:)
func QLThumbnailCopyDocumentURL(thumbnail ThumbnailRef) URLRef {
	return _QLThumbnailCopyDocumentURL(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailCopyDocumentURL */

// Returns a thumbnail image.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Returns a thumbnail image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailCopyImage(_:)
func QLThumbnailCopyImage(thumbnail ThumbnailRef) ImageRef {
	return _QLThumbnailCopyImage(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailCopyImage */

// Returns the options for the requested thumbnail.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Returns the options for the requested thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailCopyOptions(_:)
func QLThumbnailCopyOptions(thumbnail ThumbnailRef) DictionaryRef {
	return _QLThumbnailCopyOptions(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailCopyOptions */

// Returns a thumbnail that’s generated in the background.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Returns a thumbnail that’s generated in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailCreate(_:_:_:_:)
func QLThumbnailCreate(allocator AllocatorRef, url URLRef, maxThumbnailSize corefoundation.CGSize, options DictionaryRef) ThumbnailRef {
	return _QLThumbnailCreate(allocator, url, maxThumbnailSize, options)
}/* debug [functions.gen.go/function]: QLThumbnailCreate */

// Creates a thumbnail in the background on the provided background queue.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Creates a thumbnail in the background on the provided background queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailDispatchAsync(_:_:_:)
func QLThumbnailDispatchAsync(thumbnail ThumbnailRef, queue unsafe.Pointer, completion unsafe.Pointer) {
	_QLThumbnailDispatchAsync(thumbnail, queue, completion)
}/* debug [functions.gen.go/function]: QLThumbnailDispatchAsync */

// Returns the rectangle of the provided thumbnail image that represents the content of the document.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Returns the rectangle of the provided thumbnail image that represents the content of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailGetContentRect(_:)
func QLThumbnailGetContentRect(thumbnail ThumbnailRef) corefoundation.CGRect {
	return _QLThumbnailGetContentRect(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailGetContentRect */

// Returns the maximum allowed size for the provided thumbnail image.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Returns the maximum allowed size for the provided thumbnail image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailGetMaximumSize(_:)
func QLThumbnailGetMaximumSize(thumbnail ThumbnailRef) corefoundation.CGSize {
	return _QLThumbnailGetMaximumSize(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailGetMaximumSize */

// Returns the type identifier for the thumbnail’s opaque type.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Returns the type identifier for the thumbnail’s opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailGetTypeID()
func QLThumbnailGetTypeID() TypeID {
	return _QLThumbnailGetTypeID()
}/* debug [functions.gen.go/function]: QLThumbnailGetTypeID */

// Creates a thumbnail image for the specified file.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Creates a thumbnail image for the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailImageCreate(_:_:_:_:)
func QLThumbnailImageCreate(allocator AllocatorRef, url URLRef, maxThumbnailSize corefoundation.CGSize, options DictionaryRef) ImageRef {
	return _QLThumbnailImageCreate(allocator, url, maxThumbnailSize, options)
}/* debug [functions.gen.go/function]: QLThumbnailImageCreate */

// Returns whether the creation of the thumbnail was canceled.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.6.
// Returns whether the creation of the thumbnail was canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailIsCancelled(_:)
func QLThumbnailIsCancelled(thumbnail ThumbnailRef) unsafe.Pointer {
	return _QLThumbnailIsCancelled(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailIsCancelled */

// Returns the UTI for the thumbnail request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns the UTI for the thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestCopyContentUTI(_:)
func QLThumbnailRequestCopyContentUTI(thumbnail ThumbnailRequestRef) StringRef {
	return _QLThumbnailRequestCopyContentUTI(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailRequestCopyContentUTI */

// Returns the options specified for the thumbnail request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns the options specified for the thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestCopyOptions(_:)
func QLThumbnailRequestCopyOptions(thumbnail ThumbnailRequestRef) DictionaryRef {
	return _QLThumbnailRequestCopyOptions(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailRequestCopyOptions */

// Returns the URL of the document for which the thumbnail request is requested.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns the URL of the document for which the thumbnail request is requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestCopyURL(_:)
func QLThumbnailRequestCopyURL(thumbnail ThumbnailRequestRef) URLRef {
	return _QLThumbnailRequestCopyURL(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailRequestCopyURL */

// Creates a graphics context to draw the thumbnail in.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Creates a graphics context to draw the thumbnail in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestCreateContext(_:_:_:_:)
func QLThumbnailRequestCreateContext(thumbnail ThumbnailRequestRef, size corefoundation.CGSize, isBitmap unsafe.Pointer, properties DictionaryRef) ContextRef {
	return _QLThumbnailRequestCreateContext(thumbnail, size, isBitmap, properties)
}/* debug [functions.gen.go/function]: QLThumbnailRequestCreateContext */

// Flush the graphics context and sets the thumbnail response.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Flush the graphics context and sets the thumbnail response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestFlushContext(_:_:)
func QLThumbnailRequestFlushContext(thumbnail ThumbnailRequestRef, context ContextRef) {
	_QLThumbnailRequestFlushContext(thumbnail, context)
}/* debug [functions.gen.go/function]: QLThumbnailRequestFlushContext */

// Returns the object that’s stored as part of a thumbnail request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.6.
// Returns the object that’s stored as part of a thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestGetDocumentObject(_:)
func QLThumbnailRequestGetDocumentObject(thumbnail ThumbnailRequestRef) unsafe.Pointer {
	return _QLThumbnailRequestGetDocumentObject(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailRequestGetDocumentObject */

// Get the bundle of the generator receiving the thumbnail request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Get the bundle of the generator receiving the thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestGetGeneratorBundle(_:)
func QLThumbnailRequestGetGeneratorBundle(thumbnail ThumbnailRequestRef) BundleRef {
	return _QLThumbnailRequestGetGeneratorBundle(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailRequestGetGeneratorBundle */

// Returns the maximum size (in points) specified for the thumbnail image.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns the maximum size (in points) specified for the thumbnail image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestGetMaximumSize(_:)
func QLThumbnailRequestGetMaximumSize(thumbnail ThumbnailRequestRef) corefoundation.CGSize {
	return _QLThumbnailRequestGetMaximumSize(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailRequestGetMaximumSize */

// Gets the type identifier for the opaque type.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Gets the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestGetTypeID()
func QLThumbnailRequestGetTypeID() TypeID {
	return _QLThumbnailRequestGetTypeID()
}/* debug [functions.gen.go/function]: QLThumbnailRequestGetTypeID */

// Returns whether the thumbnail request has been cancelled by the client.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Returns whether the thumbnail request has been cancelled by the client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestIsCancelled(_:)
func QLThumbnailRequestIsCancelled(thumbnail ThumbnailRequestRef) unsafe.Pointer {
	return _QLThumbnailRequestIsCancelled(thumbnail)
}/* debug [functions.gen.go/function]: QLThumbnailRequestIsCancelled */

// Stores an object as part of a thumbnail request.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.6.
// Stores an object as part of a thumbnail request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestSetDocumentObject(_:_:_:)
func QLThumbnailRequestSetDocumentObject(thumbnail ThumbnailRequestRef, object unsafe.Pointer, callbacks unsafe.Pointer) {
	_QLThumbnailRequestSetDocumentObject(thumbnail, object, callbacks)
}/* debug [functions.gen.go/function]: QLThumbnailRequestSetDocumentObject */

// Sets the thumbnail request to a specified image.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Sets the thumbnail request to a specified image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestSetImage(_:_:_:)
func QLThumbnailRequestSetImage(thumbnail ThumbnailRequestRef, image ImageRef, properties DictionaryRef) {
	_QLThumbnailRequestSetImage(thumbnail, image, properties)
}/* debug [functions.gen.go/function]: QLThumbnailRequestSetImage */

// Sets the thumbnail request to contain the image at a given URL.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Sets the thumbnail request to contain the image at a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestSetImageAtURL(_:_:_:)
func QLThumbnailRequestSetImageAtURL(thumbnail ThumbnailRequestRef, url URLRef, properties DictionaryRef) {
	_QLThumbnailRequestSetImageAtURL(thumbnail, url, properties)
}/* debug [functions.gen.go/function]: QLThumbnailRequestSetImageAtURL */

// Sets the response to the thumbnail request to image data saved within the document.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Sets the response to the thumbnail request to image data saved within the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestSetImageWithData(_:_:_:)
func QLThumbnailRequestSetImageWithData(thumbnail ThumbnailRequestRef, data DataRef, properties DictionaryRef) {
	_QLThumbnailRequestSetImageWithData(thumbnail, data, properties)
}/* debug [functions.gen.go/function]: QLThumbnailRequestSetImageWithData */

// Sets the default image representation for an item with the provided data and specified file type.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.6.
// Sets the default image representation for an item with the provided data and specified file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestSetThumbnailWithDataRepresentation(_:_:_:_:_:)
func QLThumbnailRequestSetThumbnailWithDataRepresentation(thumbnail ThumbnailRequestRef, data DataRef, contentTypeUTI StringRef, previewProperties DictionaryRef, properties DictionaryRef) {
	_QLThumbnailRequestSetThumbnailWithDataRepresentation(thumbnail, data, contentTypeUTI, previewProperties, properties)
}/* debug [functions.gen.go/function]: QLThumbnailRequestSetThumbnailWithDataRepresentation */

// Sets the default image representation for an item of a given type at the specified URL.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.6.
// Sets the default image representation for an item of a given type at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLThumbnailRequestSetThumbnailWithURLRepresentation(_:_:_:_:_:)
func QLThumbnailRequestSetThumbnailWithURLRepresentation(thumbnail ThumbnailRequestRef, url URLRef, contentTypeUTI StringRef, previewProperties DictionaryRef, properties DictionaryRef) {
	_QLThumbnailRequestSetThumbnailWithURLRepresentation(thumbnail, url, contentTypeUTI, previewProperties, properties)
}/* debug [functions.gen.go/function]: QLThumbnailRequestSetThumbnailWithURLRepresentation */




