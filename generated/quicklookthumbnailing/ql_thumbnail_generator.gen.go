// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class QLThumbnailGenerator */


/* debug [class_header]: Header for QLThumbnailGenerator */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ThumbnailGenerator */
// An interface definition for the [ThumbnailGenerator] class.
type IThumbnailGenerator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ThumbnailGenerator */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ThumbnailGenerator */
	// methods:
	CancelRequest(request IQLThumbnailGenerationRequest)
	GenerateBestRepresentationForRequestCompletionHandler(request IQLThumbnailGenerationRequest, completionHandler unsafe.Pointer)
	GenerateRepresentationsForRequestUpdateHandler(request IQLThumbnailGenerationRequest, updateHandler unsafe.Pointer)
	SaveBestRepresentationForRequestToFileAtURLAsContentTypeCompletionHandler(request IQLThumbnailGenerationRequest, fileURL objc.IObject /* cross-framework: NSURL */, contentType uniformtypeidentifiers.UTType, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ThumbnailGenerator */
// Alloc allocates a new instance without initialization.
func (tc _ThumbnailGeneratorClass) Alloc() ThumbnailGenerator {
	rv := objc.Send[ThumbnailGenerator](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ThumbnailGenerator */
// An object that generates thumbnail images based on provided requirements.


// An object that generates thumbnail images based on provided requirements.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ThumbnailGenerator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ThumbnailGenerator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ThumbnailGenerator */

// The singleton thumbnail generator instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/shared
func (tc _ThumbnailGeneratorClass) SharedGenerator() ThumbnailGenerator {
	rv := objc.Send[ThumbnailGenerator](objc.ID(tc.class), objc.Sel("sharedGenerator"))
	return rv
}/* debug [class_properties_class/property]: sharedGenerator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ThumbnailGenerator */

// Cancels the generation of a thumbnail for a given request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/cancel(_:)
func (t_ ThumbnailGenerator) CancelRequest(request IQLThumbnailGenerationRequest) {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancelRequest:"), request)
}/* debug [instance_methods/method]: CancelRequest */


// Generates the best possible thumbnail representation for a file and calls a handler upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/generateBestRepresentation(for:completion:)
func (t_ ThumbnailGenerator) GenerateBestRepresentationForRequestCompletionHandler(request IQLThumbnailGenerationRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("generateBestRepresentationForRequest:completionHandler:"), request, completionHandler)
}/* debug [instance_methods/method]: GenerateBestRepresentationForRequestCompletionHandler */


// Generates various thumbnail representations for a file and calls the update handler for each thumbnail representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/generateRepresentations(for:update:)
func (t_ ThumbnailGenerator) GenerateRepresentationsForRequestUpdateHandler(request IQLThumbnailGenerationRequest, updateHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("generateRepresentationsForRequest:updateHandler:"), request, updateHandler)
}/* debug [instance_methods/method]: GenerateRepresentationsForRequestUpdateHandler */


// Saves a thumbnail for the request on disk at fileURL. The file saved at fileURL has to be deleted when it is not used anymore. This is primarily intended for file provider extensions which need to upload thumbnails and have a small memory limit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/saveBestRepresentation(for:to:as:completion:)
func (t_ ThumbnailGenerator) SaveBestRepresentationForRequestToFileAtURLAsContentTypeCompletionHandler(request IQLThumbnailGenerationRequest, fileURL objc.IObject /* cross-framework: NSURL */, contentType uniformtypeidentifiers.UTType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("saveBestRepresentationForRequest:toFileAtURL:asContentType:completionHandler:"), request, fileURL, contentType, completionHandler)
}/* debug [instance_methods/method]: SaveBestRepresentationForRequestToFileAtURLAsContentTypeCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ThumbnailGenerator */

// The singleton thumbnail generator instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/shared
func (t_ ThumbnailGenerator) SharedGenerator() IQLThumbnailGenerator {
	rv := objc.Send[ThumbnailGenerator](t_.ID, objc.Sel("sharedGenerator"))
	return rv
}/* debug [instance_properties/getter]: sharedGenerator */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLThumbnailGenerator */



