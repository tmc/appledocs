// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNBoundsPrediction] class.
var (
	CNBoundsPredictionClass     _CNBoundsPredictionClass
	CNBoundsPredictionClassOnce sync.Once
)

func getCNBoundsPredictionClass() _CNBoundsPredictionClass {
	CNBoundsPredictionClassOnce.Do(func() {
		CNBoundsPredictionClass = _CNBoundsPredictionClass{objc.GetClass("CNBoundsPrediction")}
	})
	return CNBoundsPredictionClass
}

type _CNBoundsPredictionClass struct {
	class objc.Class
}

// An interface definition for the [CNBoundsPrediction] class.
type ICNBoundsPrediction interface {
	objectivec.IObject
	Confidence() float32
	SetConfidence(value float32)
	NormalizedBounds() coregraphics.CGRect
	SetNormalizedBounds(value coregraphics.CGRect)
}

// An object representing the bounds of the predicted subject.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class
type CNBoundsPrediction struct {
	objectivec.Object
}

// CNBoundsPredictionFrom constructs a [CNBoundsPrediction] from an unsafe.Pointer.
//
// An object representing the bounds of the predicted subject.
func CNBoundsPredictionFrom(ptr unsafe.Pointer) CNBoundsPrediction {
	return CNBoundsPrediction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNBoundsPredictionClass) Alloc() CNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNBoundsPredictionClass) New() CNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNBoundsPrediction) Init() CNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNBoundsPrediction) Autorelease() CNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNBoundsPrediction creates a new CNBoundsPrediction instance.
func NewCNBoundsPrediction() CNBoundsPrediction {
	return getCNBoundsPredictionClass().New()
}


// A number between 0.0 and 1.0 representing the probability that a defined object is within the bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class/confidence
func (c_ CNBoundsPrediction) Confidence() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("confidence"))
	return rv
}


// SetConfidence sets the value of the confidence property.
// A number between 0.0 and 1.0 representing the probability that a defined object is within the bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class/confidence
func (c_ CNBoundsPrediction) SetConfidence(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfidence:"), value)
}

// The bounds of the detected object in normalized coordinates where (0.0, 0.0) is the upper-left corner, and (1.0, 1.0) is the lower-right.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class/normalizedBounds
func (c_ CNBoundsPrediction) NormalizedBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("normalizedBounds"))
	return rv
}


// SetNormalizedBounds sets the value of the normalizedBounds property.
// The bounds of the detected object in normalized coordinates where (0.0, 0.0) is the upper-left corner, and (1.0, 1.0) is the lower-right.

//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class/normalizedBounds
func (c_ CNBoundsPrediction) SetNormalizedBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNormalizedBounds:"), value)
}



