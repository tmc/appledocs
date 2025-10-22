// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RecognizedPoint] class.
var (
	RecognizedPointClass     _RecognizedPointClass
	RecognizedPointClassOnce sync.Once
)

func getRecognizedPointClass() _RecognizedPointClass {
	RecognizedPointClassOnce.Do(func() {
		RecognizedPointClass = _RecognizedPointClass{objc.GetClass("VNRecognizedPoint")}
	})
	return RecognizedPointClass
}

type _RecognizedPointClass struct {
	class objc.Class
}

// An interface definition for the [RecognizedPoint] class.
type IRecognizedPoint interface {
	IDetectedPoint
	Identifier() RecognizedPointKey
}

// An object that represents a normalized point in an image, along with an identifier label and a confidence value.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoint
type RecognizedPoint struct {
	DetectedPoint
}

// RecognizedPointFrom constructs a [RecognizedPoint] from an unsafe.Pointer.
//
// An object that represents a normalized point in an image, along with an identifier label and a confidence value.
func RecognizedPointFrom(ptr unsafe.Pointer) RecognizedPoint {
	return RecognizedPoint{
		DetectedPoint: DetectedPointFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecognizedPointClass) Alloc() RecognizedPoint {
	rv := objc.Send[RecognizedPoint](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecognizedPointClass) New() RecognizedPoint {
	rv := objc.Send[RecognizedPoint](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedPoint) Init() RecognizedPoint {
	rv := objc.Send[RecognizedPoint](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedPoint) Autorelease() RecognizedPoint {
	rv := objc.Send[RecognizedPoint](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedPoint creates a new RecognizedPoint instance.
func NewRecognizedPoint() RecognizedPoint {
	return getRecognizedPointClass().New()
}


// The point’s identifier label.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoint/identifier
func (r_ RecognizedPoint) Identifier() RecognizedPointKey {
	rv := objc.Send[RecognizedPointKey](r_.ID, objc.Sel("identifier"))
	return rv
}



