// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKMultiPoint] class.
var (
	MKMultiPointClass     _MKMultiPointClass
	MKMultiPointClassOnce sync.Once
)

func getMKMultiPointClass() _MKMultiPointClass {
	MKMultiPointClassOnce.Do(func() {
		MKMultiPointClass = _MKMultiPointClass{objc.GetClass("MKMultiPoint")}
	})
	return MKMultiPointClass
}

type _MKMultiPointClass struct {
	class objc.Class
}

// An interface definition for the [MKMultiPoint] class.
type IMKMultiPoint interface {
	IMKShape
}

// An abstract class that defines the common behavior that open and closed polygon overlays share.
//
// Don’t create instances of this class directly. Instead, create instances of the or classes. However, you can use the methods and property of this class to access information about the specific points associated with the line or polygon.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPoint
type MKMultiPoint struct {
	MKShape
}

// MKMultiPointFrom constructs a [MKMultiPoint] from an unsafe.Pointer.
//
// An abstract class that defines the common behavior that open and closed polygon overlays share.
func MKMultiPointFrom(ptr unsafe.Pointer) MKMultiPoint {
	return MKMultiPoint{
		MKShape: MKShapeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMultiPointClass) Alloc() MKMultiPoint {
	rv := objc.Send[MKMultiPoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMultiPointClass) New() MKMultiPoint {
	rv := objc.Send[MKMultiPoint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMultiPoint) Init() MKMultiPoint {
	rv := objc.Send[MKMultiPoint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMultiPoint) Autorelease() MKMultiPoint {
	rv := objc.Send[MKMultiPoint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMultiPoint creates a new MKMultiPoint instance.
func NewMKMultiPoint() MKMultiPoint {
	return getMKMultiPointClass().New()
}




