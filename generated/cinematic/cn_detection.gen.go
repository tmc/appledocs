// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNDetection */


/* debug [class_header]: Header for CNDetection */
// The class instance for the [CNDetection] class.
var (
	CNDetectionClass     _CNDetectionClass
	CNDetectionClassOnce sync.Once
)

func getCNDetectionClass() _CNDetectionClass {
	CNDetectionClassOnce.Do(func() {
		CNDetectionClass = _CNDetectionClass{objc.GetClass("CNDetection")}
	})
	return CNDetectionClass
}

type _CNDetectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNDetection */
// An interface definition for the [CNDetection] class.
type ICNDetection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNDetection */
	// properties:
	DetectionGroupID() CNDetectionGroupID /* typedef */
	DetectionID() CNDetectionID /* typedef */
	DetectionType() CNDetectionType
	FocusDisparity() float32
	NormalizedRect() corefoundation.CGRect
	Time() objc.IObject /* cross-framework: Time */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNDetection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNDetection */
// Alloc allocates a new instance without initialization.
func (cc _CNDetectionClass) Alloc() CNDetection {
	rv := objc.Send[CNDetection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNDetectionClass) New() CNDetection {
	rv := objc.Send[CNDetection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNDetection) Init() CNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNDetection) Autorelease() CNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNDetection creates a new CNDetection instance.
func NewCNDetection() CNDetection {
	return getCNDetectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNDetection */
// A structure that represents a detected subject, face, torso or pet at a particular time.
//
// Specifies the type, distance bounds, and time of the detection. Detections obtained from the Cinematic script include a unique number that can tracks the detection over time. Some types of detections also include a unique group number that associates related detections (for example, the face and torso of the same person).


// A structure that represents a detected subject, face, torso or pet at a particular time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class
type CNDetection struct {
	objectivec.Object
}

// CNDetectionFrom constructs a [CNDetection] from an unsafe.Pointer.
//
// A structure that represents a detected subject, face, torso or pet at a particular time.
func CNDetectionFrom(ptr unsafe.Pointer) CNDetection {
	return CNDetection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNDetection */

// Creates a Cinematic detection of a subject.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/initWithTime:detectionType:normalizedRect:focusDisparity:
func NewCNDetectionWithTimeDetectionTypeNormalizedRectFocusDisparity(time objc.IObject /* cross-framework: Time */, detectionType CNDetectionType, normalizedRect corefoundation.CGRect, focusDisparity float32) CNDetection {
	instance := getCNDetectionClass().Alloc()
	rv := objc.Send[CNDetection](instance.ID, objc.Sel("initWithTime:detectionType:normalizedRect:focusDisparity:"), time, detectionType, normalizedRect, focusDisparity)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNDetectionWithTimeDetectionTypeNormalizedRectFocusDisparity */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNDetection */

// A localized accessibility label converting a specific detection type into a broad category such as a person, pet, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/accessibilityLabelForDetectionType:
func (cc _CNDetectionClass) AccessibilityLabelForDetectionType(detectionType CNDetectionType) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("accessibilityLabelForDetectionType:"), detectionType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AccessibilityLabelForDetectionType) */


// Determines the disparity to use to focus on the object in the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/disparityInNormalizedRect:sourceDisparity:detectionType:priorDisparity:
func (cc _CNDetectionClass) DisparityInNormalizedRectSourceDisparityDetectionTypePriorDisparity(normalizedRect corefoundation.CGRect, sourceDisparity PixelBufferRef /* not a class type */, detectionType CNDetectionType, priorDisparity float32) float32 {
	rv := objc.Send[float32](objc.ID(cc.class), objc.Sel("disparityInNormalizedRect:sourceDisparity:detectionType:priorDisparity:"), normalizedRect, sourceDisparity, detectionType, priorDisparity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DisparityInNormalizedRectSourceDisparityDetectionTypePriorDisparity) */


// Determines whether a given detection group ID is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/isValidDetectionGroupID:
func (cc _CNDetectionClass) IsValidDetectionGroupID(detectionGroupID CNDetectionGroupID /* typedef */) bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("isValidDetectionGroupID:"), detectionGroupID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsValidDetectionGroupID) */


// Determines whether a given detection ID is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/isValidDetectionID:
func (cc _CNDetectionClass) IsValidDetectionID(detectionID CNDetectionID /* typedef */) bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("isValidDetectionID:"), detectionID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsValidDetectionID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNDetection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNDetection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNDetection */

// A unique number representing the detection to focus on if this is a group decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/detectionGroupID
func (c_ CNDetection) DetectionGroupID() CNDetectionGroupID /* typedef */ {
	rv := objc.Send[int64](c_.ID, objc.Sel("detectionGroupID"))
	return rv
}/* debug [instance_properties/getter]: detectionGroupID */


// An unique identifier assigned by the Cinematic script to all detections of the same subject and detection type across time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/detectionID
func (c_ CNDetection) DetectionID() CNDetectionID /* typedef */ {
	rv := objc.Send[int64](c_.ID, objc.Sel("detectionID"))
	return rv
}/* debug [instance_properties/getter]: detectionID */


// The type of object detected, such as the face, torso, cat, dog, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/detectionType
func (c_ CNDetection) DetectionType() CNDetectionType {
	rv := objc.Send[CNDetectionType](c_.ID, objc.Sel("detectionType"))
	return rv
}/* debug [instance_properties/getter]: detectionType */


// The disparity to use in order to focus on the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/focusDisparity
func (c_ CNDetection) FocusDisparity() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("focusDisparity"))
	return rv
}/* debug [instance_properties/getter]: focusDisparity */


// The rectangle within the image where the object occurs, normalized such that (0.0, 0.0) is the top-left and (1.0, 1.0) is the bottom-right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/normalizedRect
func (c_ CNDetection) NormalizedRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("normalizedRect"))
	return rv
}/* debug [instance_properties/getter]: normalizedRect */


// The first presentation time which the subject should be in focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/time
func (c_ CNDetection) Time() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("time"))
	return rv
}/* debug [instance_properties/getter]: time */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNDetection */


