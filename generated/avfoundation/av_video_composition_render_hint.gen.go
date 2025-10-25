// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVVideoCompositionRenderHint */


/* debug [class_header]: Header for AVVideoCompositionRenderHint */
// The class instance for the [VideoCompositionRenderHint] class.
var (
	VideoCompositionRenderHintClass     _VideoCompositionRenderHintClass
	VideoCompositionRenderHintClassOnce sync.Once
)

func getVideoCompositionRenderHintClass() _VideoCompositionRenderHintClass {
	VideoCompositionRenderHintClassOnce.Do(func() {
		VideoCompositionRenderHintClass = _VideoCompositionRenderHintClass{objc.GetClass("AVVideoCompositionRenderHint")}
	})
	return VideoCompositionRenderHintClass
}

type _VideoCompositionRenderHintClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VideoCompositionRenderHint */
// An interface definition for the [VideoCompositionRenderHint] class.
type IVideoCompositionRenderHint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VideoCompositionRenderHint */
	// properties:
	EndCompositionTime() objc.IObject /* cross-framework: Time */
	StartCompositionTime() objc.IObject /* cross-framework: Time */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VideoCompositionRenderHint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VideoCompositionRenderHint */
// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionRenderHintClass) Alloc() VideoCompositionRenderHint {
	rv := objc.Send[VideoCompositionRenderHint](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoCompositionRenderHintClass) New() VideoCompositionRenderHint {
	rv := objc.Send[VideoCompositionRenderHint](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoCompositionRenderHint) Init() VideoCompositionRenderHint {
	rv := objc.Send[VideoCompositionRenderHint](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoCompositionRenderHint) Autorelease() VideoCompositionRenderHint {
	rv := objc.Send[VideoCompositionRenderHint](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoCompositionRenderHint creates a new VideoCompositionRenderHint instance.
func NewVideoCompositionRenderHint() VideoCompositionRenderHint {
	return getVideoCompositionRenderHintClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VideoCompositionRenderHint */
// Information about upcoming composition requests, such as composition start time and end time.


// Information about upcoming composition requests, such as composition start time and end time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint
type VideoCompositionRenderHint struct {
	objectivec.Object
}

// VideoCompositionRenderHintFrom constructs a [VideoCompositionRenderHint] from an unsafe.Pointer.
//
// Information about upcoming composition requests, such as composition start time and end time.
func VideoCompositionRenderHintFrom(ptr unsafe.Pointer) VideoCompositionRenderHint {
	return VideoCompositionRenderHint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VideoCompositionRenderHint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VideoCompositionRenderHint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VideoCompositionRenderHint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VideoCompositionRenderHint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VideoCompositionRenderHint */

// The end time of the upcoming composition requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint/endCompositionTime
func (v_ VideoCompositionRenderHint) EndCompositionTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](v_.ID, objc.Sel("endCompositionTime"))
	return rv
}/* debug [instance_properties/getter]: endCompositionTime */


// The start time of the upcoming composition requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint/startCompositionTime
func (v_ VideoCompositionRenderHint) StartCompositionTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](v_.ID, objc.Sel("startCompositionTime"))
	return rv
}/* debug [instance_properties/getter]: startCompositionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVVideoCompositionRenderHint */



