// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNScriptFrame */


/* debug [class_header]: Header for CNScriptFrame */
// The class instance for the [CNScriptFrame] class.
var (
	CNScriptFrameClass     _CNScriptFrameClass
	CNScriptFrameClassOnce sync.Once
)

func getCNScriptFrameClass() _CNScriptFrameClass {
	CNScriptFrameClassOnce.Do(func() {
		CNScriptFrameClass = _CNScriptFrameClass{objc.GetClass("CNScriptFrame")}
	})
	return CNScriptFrameClass
}

type _CNScriptFrameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNScriptFrame */
// An interface definition for the [CNScriptFrame] class.
type ICNScriptFrame interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNScriptFrame */
	// properties:
	AllDetections() []CNDetection
	FocusDetection() ICNDetection
	FocusDisparity() float32
	Time() objc.IObject /* cross-framework: Time */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNScriptFrame */
	// methods:
	BestDetectionForGroupID(detectionGroupID CNDetectionGroupID /* typedef */) ICNDetection
	DetectionForID(detectionID CNDetectionID /* typedef */) ICNDetection
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNScriptFrame */
// Alloc allocates a new instance without initialization.
func (cc _CNScriptFrameClass) Alloc() CNScriptFrame {
	rv := objc.Send[CNScriptFrame](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNScriptFrameClass) New() CNScriptFrame {
	rv := objc.Send[CNScriptFrame](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNScriptFrame) Init() CNScriptFrame {
	rv := objc.Send[CNScriptFrame](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNScriptFrame) Autorelease() CNScriptFrame {
	rv := objc.Send[CNScriptFrame](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNScriptFrame creates a new CNScriptFrame instance.
func NewCNScriptFrame() CNScriptFrame {
	return getCNScriptFrameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNScriptFrame */
// An object that represents what to focus on, and where to focus, in a given movie frame.


// An object that represents what to focus on, and where to focus, in a given movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame
type CNScriptFrame struct {
	objectivec.Object
}

// CNScriptFrameFrom constructs a [CNScriptFrame] from an unsafe.Pointer.
//
// An object that represents what to focus on, and where to focus, in a given movie frame.
func CNScriptFrameFrom(ptr unsafe.Pointer) CNScriptFrame {
	return CNScriptFrame{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNScriptFrame *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNScriptFrame */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNScriptFrame */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNScriptFrame */

// The best detection to focus on in a frame among those within the given detection group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/bestDetectionForGroupID:
func (c_ CNScriptFrame) BestDetectionForGroupID(detectionGroupID CNDetectionGroupID /* typedef */) ICNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("bestDetectionForGroupID:"), detectionGroupID)
	return rv
}/* debug [instance_methods/method]: BestDetectionForGroupID */


// The detection in the frame with the given detection ID, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/detectionForID:
func (c_ CNScriptFrame) DetectionForID(detectionID CNDetectionID /* typedef */) ICNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("detectionForID:"), detectionID)
	return rv
}/* debug [instance_methods/method]: DetectionForID */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNScriptFrame */

// All detections for the Cinematic movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/allDetections
func (c_ CNScriptFrame) AllDetections() []CNDetection {
	rv := objc.Send[[]CNDetection](c_.ID, objc.Sel("allDetections"))
	return rv
}/* debug [instance_properties/getter]: allDetections */


// What to focus on in a given frame of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/focusDetection
func (c_ CNScriptFrame) FocusDetection() ICNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("focusDetection"))
	return rv
}/* debug [instance_properties/getter]: focusDetection */


// Where to focus in a given frame of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/focusDisparity
func (c_ CNScriptFrame) FocusDisparity() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("focusDisparity"))
	return rv
}/* debug [instance_properties/getter]: focusDisparity */


// The time of the focus transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/time
func (c_ CNScriptFrame) Time() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("time"))
	return rv
}/* debug [instance_properties/getter]: time */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNScriptFrame */






