// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ThumbnailGenerator] class.
var (
	ThumbnailGeneratorClass     _ThumbnailGeneratorClass
	ThumbnailGeneratorClassOnce sync.Once
)

func getThumbnailGeneratorClass() _ThumbnailGeneratorClass {
	ThumbnailGeneratorClassOnce.Do(func() {
		ThumbnailGeneratorClass = _ThumbnailGeneratorClass{objc.GetClass("QLThumbnailGenerator")}
	})
	return ThumbnailGeneratorClass
}

type _ThumbnailGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [ThumbnailGenerator] class.
type IThumbnailGenerator interface {
	objectivec.IObject
	CancelRequest(request unsafe.Pointer)
	GenerateBestRepresentationForRequestCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer)
	GenerateRepresentationsForRequestUpdateHandler(request unsafe.Pointer, updateHandler unsafe.Pointer)
	SaveBestRepresentationForRequestToFileAtURLAsContentTypeCompletionHandler(request unsafe.Pointer, fileURL unsafe.Pointer, contentType unsafe.Pointer, completionHandler unsafe.Pointer)
	SaveBestRepresentationForRequestToFileAtURLWithContentTypeCompletionHandler(request unsafe.Pointer, fileURL unsafe.Pointer, contentType string, completionHandler unsafe.Pointer)
}

// An object that generates thumbnail images based on provided requirements.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator
type ThumbnailGenerator struct {
	objectivec.Object
}

// ThumbnailGeneratorFrom constructs a [ThumbnailGenerator] from an unsafe.Pointer.
//
// An object that generates thumbnail images based on provided requirements.
func ThumbnailGeneratorFrom(ptr unsafe.Pointer) ThumbnailGenerator {
	return ThumbnailGenerator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ThumbnailGeneratorClass) Alloc() ThumbnailGenerator {
	rv := objc.Send[ThumbnailGenerator](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ThumbnailGeneratorClass) New() ThumbnailGenerator {
	rv := objc.Send[ThumbnailGenerator](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ThumbnailGenerator) Init() ThumbnailGenerator {
	rv := objc.Send[ThumbnailGenerator](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ThumbnailGenerator) Autorelease() ThumbnailGenerator {
	rv := objc.Send[ThumbnailGenerator](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewThumbnailGenerator creates a new ThumbnailGenerator instance.
func NewThumbnailGenerator() ThumbnailGenerator {
	return getThumbnailGeneratorClass().New()
}


// The singleton thumbnail generator instance.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/shared
func (tc _ThumbnailGeneratorClass) SharedGenerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("sharedGenerator"))
	return rv
}
// Cancels the generation of a thumbnail for a given request.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/cancel(_:)
func (t_ ThumbnailGenerator) CancelRequest(request unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancelRequest:"), request)
}

// Generates the best possible thumbnail representation for a file and calls a handler upon completion.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/generateBestRepresentation(for:completion:)
func (t_ ThumbnailGenerator) GenerateBestRepresentationForRequestCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("generateBestRepresentationForRequest:completionHandler:"), request, completionHandler)
}

// Generates various thumbnail representations for a file and calls the update handler for each thumbnail representation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/generateRepresentations(for:update:)
func (t_ ThumbnailGenerator) GenerateRepresentationsForRequestUpdateHandler(request unsafe.Pointer, updateHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("generateRepresentationsForRequest:updateHandler:"), request, updateHandler)
}

// Saves a thumbnail for the request on disk at fileURL. The file saved at fileURL has to be deleted when it is not used anymore. This is primarily intended for file provider extensions which need to upload thumbnails and have a small memory limit.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/saveBestRepresentation(for:to:as:completion:)
func (t_ ThumbnailGenerator) SaveBestRepresentationForRequestToFileAtURLAsContentTypeCompletionHandler(request unsafe.Pointer, fileURL unsafe.Pointer, contentType unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("saveBestRepresentationForRequest:toFileAtURL:asContentType:completionHandler:"), request, fileURL, contentType, completionHandler)
}

// Saves the best representation of thumbnail for a specific request to the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/saveBestRepresentation(for:to:contentType:completion:)
func (t_ ThumbnailGenerator) SaveBestRepresentationForRequestToFileAtURLWithContentTypeCompletionHandler(request unsafe.Pointer, fileURL unsafe.Pointer, contentType string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("saveBestRepresentationForRequest:toFileAtURL:withContentType:completionHandler:"), request, fileURL, objc.String(contentType), completionHandler)
}

// The singleton thumbnail generator instance.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/shared
func (t_ ThumbnailGenerator) SharedGenerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("sharedGenerator"))
	return rv
}



