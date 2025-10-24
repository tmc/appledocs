// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNStatefulRequest */


/* debug [class_header]: Header for VNStatefulRequest */
// The class instance for the [StatefulRequest] class.
var (
	StatefulRequestClass     _StatefulRequestClass
	StatefulRequestClassOnce sync.Once
)

func getStatefulRequestClass() _StatefulRequestClass {
	StatefulRequestClassOnce.Do(func() {
		StatefulRequestClass = _StatefulRequestClass{objc.GetClass("VNStatefulRequest")}
	})
	return StatefulRequestClass
}

type _StatefulRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StatefulRequest */
// An interface definition for the [StatefulRequest] class.
type IStatefulRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for StatefulRequest */
	// properties:
	FrameAnalysisSpacing() objc.IObject /* cross-framework: Time */
	MinimumLatencyFrameCount() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StatefulRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StatefulRequest */
// Alloc allocates a new instance without initialization.
func (sc _StatefulRequestClass) Alloc() StatefulRequest {
	rv := objc.Send[StatefulRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StatefulRequestClass) New() StatefulRequest {
	rv := objc.Send[StatefulRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StatefulRequest) Init() StatefulRequest {
	rv := objc.Send[StatefulRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StatefulRequest) Autorelease() StatefulRequest {
	rv := objc.Send[StatefulRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStatefulRequest creates a new StatefulRequest instance.
func NewStatefulRequest() StatefulRequest {
	return getStatefulRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StatefulRequest */
// An abstract request type that builds evidence of a condition over time.


// An abstract request type that builds evidence of a condition over time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNStatefulRequest
type StatefulRequest struct {
	ImageBasedRequest
}

// StatefulRequestFrom constructs a [StatefulRequest] from an unsafe.Pointer.
//
// An abstract request type that builds evidence of a condition over time.
func StatefulRequestFrom(ptr unsafe.Pointer) StatefulRequest {
	return StatefulRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StatefulRequest */

// Initializes a video-based request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNStatefulRequest/init(frameAnalysisSpacing:completionHandler:)
func NewStatefulRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing objc.IObject /* cross-framework: Time */, completionHandler RequestCompletionHandler /* not a class type */) StatefulRequest {
	instance := getStatefulRequestClass().Alloc()
	rv := objc.Send[StatefulRequest](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStatefulRequestWithFrameAnalysisSpacingCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StatefulRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StatefulRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StatefulRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StatefulRequest */

// A time value that indicates the interval between analysis operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNStatefulRequest/frameAnalysisSpacing
func (s_ StatefulRequest) FrameAnalysisSpacing() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](s_.ID, objc.Sel("frameAnalysisSpacing"))
	return rv
}/* debug [instance_properties/getter]: frameAnalysisSpacing */


// The minimum number of frames a request processes before reporting an observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNStatefulRequest/minimumLatencyFrameCount
func (s_ StatefulRequest) MinimumLatencyFrameCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("minimumLatencyFrameCount"))
	return rv
}/* debug [instance_properties/getter]: minimumLatencyFrameCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNStatefulRequest */


