// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CNDetection] class.
type ICNDetection interface {
	objectivec.IObject
}

// A structure that represents a detected subject, face, torso or pet at a particular time.
//
// Specifies the type, distance bounds, and time of the detection. Detections obtained from the Cinematic script include a unique number that can tracks the detection over time. Some types of detections also include a unique group number that associates related detections (for example, the face and torso of the same person).
//
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

// Alloc allocates a new instance without initialization.
func (cc _CNDetectionClass) Alloc() CNDetection {
	rv := objc.Send[CNDetection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Determines the disparity to use to focus on the object in the rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/disparityInNormalizedRect:sourceDisparity:detectionType:priorDisparity:
func (cc _CNDetectionClass) DisparityInNormalizedRectSourceDisparityDetectionTypePriorDisparity(normalizedRect coregraphics.CGRect, sourceDisparity unsafe.Pointer, detectionType unsafe.Pointer, priorDisparity unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("disparityInNormalizedRect:sourceDisparity:detectionType:priorDisparity:"), normalizedRect, sourceDisparity, detectionType, priorDisparity)
	return rv
}

// Determines whether a given detection ID is valid.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/isValidDetectionID:
func (cc _CNDetectionClass) IsValidDetectionID(detectionID unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("isValidDetectionID:"), detectionID)
	return rv
}

// The disparity to use in order to focus on the object.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/focusDisparity
func (c_ CNDetection) FocusDisparity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("focusDisparity"))
	return rv
}

// The rectangle within the image where the object occurs, normalized such that (0.0, 0.0) is the top-left and (1.0, 1.0) is the bottom-right.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetection-c.class/normalizedRect
func (c_ CNDetection) NormalizedRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("normalizedRect"))
	return rv
}



