// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVVideoCompositionRenderContext */


/* debug [class_header]: Header for AVVideoCompositionRenderContext */
// The class instance for the [VideoCompositionRenderContext] class.
var (
	VideoCompositionRenderContextClass     _VideoCompositionRenderContextClass
	VideoCompositionRenderContextClassOnce sync.Once
)

func getVideoCompositionRenderContextClass() _VideoCompositionRenderContextClass {
	VideoCompositionRenderContextClassOnce.Do(func() {
		VideoCompositionRenderContextClass = _VideoCompositionRenderContextClass{objc.GetClass("AVVideoCompositionRenderContext")}
	})
	return VideoCompositionRenderContextClass
}

type _VideoCompositionRenderContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VideoCompositionRenderContext */
// An interface definition for the [VideoCompositionRenderContext] class.
type IVideoCompositionRenderContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VideoCompositionRenderContext */
	// properties:
	EdgeWidths() objc.IObject /* cross-framework: AVEdgeWidths */
	HighQualityRendering() bool
	PixelAspectRatio() objc.IObject /* cross-framework: AVPixelAspectRatio */
	RenderScale() float32
	RenderTransform() corefoundation.CGAffineTransform
	Size() corefoundation.CGSize
	VideoComposition() IAVVideoComposition
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VideoCompositionRenderContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VideoCompositionRenderContext */
// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionRenderContextClass) Alloc() VideoCompositionRenderContext {
	rv := objc.Send[VideoCompositionRenderContext](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoCompositionRenderContextClass) New() VideoCompositionRenderContext {
	rv := objc.Send[VideoCompositionRenderContext](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoCompositionRenderContext) Init() VideoCompositionRenderContext {
	rv := objc.Send[VideoCompositionRenderContext](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoCompositionRenderContext) Autorelease() VideoCompositionRenderContext {
	rv := objc.Send[VideoCompositionRenderContext](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoCompositionRenderContext creates a new VideoCompositionRenderContext instance.
func NewVideoCompositionRenderContext() VideoCompositionRenderContext {
	return getVideoCompositionRenderContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VideoCompositionRenderContext */
// An object that defines the context in which custom compositors render pixel buffers.
//
// A render context provides size and scaling information and offers a service for efficiently providing pixel buffers from a managed pool of buffers.


// An object that defines the context in which custom compositors render pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext
type VideoCompositionRenderContext struct {
	objectivec.Object
}

// VideoCompositionRenderContextFrom constructs a [VideoCompositionRenderContext] from an unsafe.Pointer.
//
// An object that defines the context in which custom compositors render pixel buffers.
func VideoCompositionRenderContextFrom(ptr unsafe.Pointer) VideoCompositionRenderContext {
	return VideoCompositionRenderContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VideoCompositionRenderContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VideoCompositionRenderContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VideoCompositionRenderContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VideoCompositionRenderContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VideoCompositionRenderContext */

// The width of the edge processing region on the left, top, right, and bottom edges, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/edgeWidths
func (v_ VideoCompositionRenderContext) EdgeWidths() objc.IObject /* cross-framework: AVEdgeWidths */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("edgeWidths"))
	return rv
}/* debug [instance_properties/getter]: edgeWidths */


// The rendering quality to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/highQualityRendering
func (v_ VideoCompositionRenderContext) HighQualityRendering() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("highQualityRendering"))
	return rv
}/* debug [instance_properties/getter]: highQualityRendering */


// The pixel aspect ratio for rendered frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/pixelAspectRatio
func (v_ VideoCompositionRenderContext) PixelAspectRatio() objc.IObject /* cross-framework: AVPixelAspectRatio */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("pixelAspectRatio"))
	return rv
}/* debug [instance_properties/getter]: pixelAspectRatio */


// A scaling ratio that is applied when rendering frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/renderScale
func (v_ VideoCompositionRenderContext) RenderScale() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("renderScale"))
	return rv
}/* debug [instance_properties/getter]: renderScale */


// A transform to apply to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/renderTransform
func (v_ VideoCompositionRenderContext) RenderTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](v_.ID, objc.Sel("renderTransform"))
	return rv
}/* debug [instance_properties/getter]: renderTransform */


// The width and height for the rendering frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/size
func (v_ VideoCompositionRenderContext) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The video composition being rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/videoComposition
func (v_ VideoCompositionRenderContext) VideoComposition() IAVVideoComposition {
	rv := objc.Send[VideoComposition](v_.ID, objc.Sel("videoComposition"))
	return rv
}/* debug [instance_properties/getter]: videoComposition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVVideoCompositionRenderContext */



