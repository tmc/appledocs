// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAsynchronousCIImageFilteringRequest */


/* debug [class_header]: Header for AVAsynchronousCIImageFilteringRequest */
// The class instance for the [AsynchronousCIImageFilteringRequest] class.
var (
	AsynchronousCIImageFilteringRequestClass     _AsynchronousCIImageFilteringRequestClass
	AsynchronousCIImageFilteringRequestClassOnce sync.Once
)

func getAsynchronousCIImageFilteringRequestClass() _AsynchronousCIImageFilteringRequestClass {
	AsynchronousCIImageFilteringRequestClassOnce.Do(func() {
		AsynchronousCIImageFilteringRequestClass = _AsynchronousCIImageFilteringRequestClass{objc.GetClass("AVAsynchronousCIImageFilteringRequest")}
	})
	return AsynchronousCIImageFilteringRequestClass
}

type _AsynchronousCIImageFilteringRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AsynchronousCIImageFilteringRequest */
// An interface definition for the [AsynchronousCIImageFilteringRequest] class.
type IAsynchronousCIImageFilteringRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AsynchronousCIImageFilteringRequest */
	// properties:
	CompositionTime() objc.IObject /* cross-framework: Time */
	RenderSize() corefoundation.CGSize
	SourceImage() appkit.Image
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AsynchronousCIImageFilteringRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AsynchronousCIImageFilteringRequest */
// Alloc allocates a new instance without initialization.
func (ac _AsynchronousCIImageFilteringRequestClass) Alloc() AsynchronousCIImageFilteringRequest {
	rv := objc.Send[AsynchronousCIImageFilteringRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AsynchronousCIImageFilteringRequestClass) New() AsynchronousCIImageFilteringRequest {
	rv := objc.Send[AsynchronousCIImageFilteringRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AsynchronousCIImageFilteringRequest) Init() AsynchronousCIImageFilteringRequest {
	rv := objc.Send[AsynchronousCIImageFilteringRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AsynchronousCIImageFilteringRequest) Autorelease() AsynchronousCIImageFilteringRequest {
	rv := objc.Send[AsynchronousCIImageFilteringRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAsynchronousCIImageFilteringRequest creates a new AsynchronousCIImageFilteringRequest instance.
func NewAsynchronousCIImageFilteringRequest() AsynchronousCIImageFilteringRequest {
	return getAsynchronousCIImageFilteringRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AsynchronousCIImageFilteringRequest */
// An object that supports using Core Image filters to process an individual video frame in a video composition.
//
// You use this class when creating a composition for Core Image filtering with the method. In that method call, you provide a block to be called by AVFoundation as it processes each frame of video, and the block’s sole parameter is a object. Use that object both to the video frame image to be filtered and allows you to return a filtered image to AVFoundation for display or export. The code listing below shows an example of applying a filter to an asset.


// An object that supports using Core Image filters to process an individual video frame in a video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest
type AsynchronousCIImageFilteringRequest struct {
	objectivec.Object
}

// AsynchronousCIImageFilteringRequestFrom constructs a [AsynchronousCIImageFilteringRequest] from an unsafe.Pointer.
//
// An object that supports using Core Image filters to process an individual video frame in a video composition.
func AsynchronousCIImageFilteringRequestFrom(ptr unsafe.Pointer) AsynchronousCIImageFilteringRequest {
	return AsynchronousCIImageFilteringRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AsynchronousCIImageFilteringRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AsynchronousCIImageFilteringRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AsynchronousCIImageFilteringRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AsynchronousCIImageFilteringRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AsynchronousCIImageFilteringRequest */

// The time in the video composition corresponding to the frame being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest/compositionTime
func (a_ AsynchronousCIImageFilteringRequest) CompositionTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("compositionTime"))
	return rv
}/* debug [instance_properties/getter]: compositionTime */


// The width and height, in pixels, of the frame being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest/renderSize
func (a_ AsynchronousCIImageFilteringRequest) RenderSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("renderSize"))
	return rv
}/* debug [instance_properties/getter]: renderSize */


// The current video frame image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest/sourceImage
func (a_ AsynchronousCIImageFilteringRequest) SourceImage() appkit.Image {
	rv := objc.Send[appkit.Image](a_.ID, objc.Sel("sourceImage"))
	return rv
}/* debug [instance_properties/getter]: sourceImage */


// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/videocomposition
func (a_ AsynchronousCIImageFilteringRequest) VideoComposition() IAVVideoComposition {
	rv := objc.Send[VideoComposition](a_.ID, objc.Sel("videoComposition"))
	return rv
}/* debug [instance_properties/getter]: videoComposition */


// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/videocomposition
func (a_ AsynchronousCIImageFilteringRequest) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoComposition:"), value)
}/* debug [instance_properties/setter]: videoComposition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAsynchronousCIImageFilteringRequest */



