// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKPointAnnotation] class.
var (
	MKPointAnnotationClass     _MKPointAnnotationClass
	MKPointAnnotationClassOnce sync.Once
)

func getMKPointAnnotationClass() _MKPointAnnotationClass {
	MKPointAnnotationClassOnce.Do(func() {
		MKPointAnnotationClass = _MKPointAnnotationClass{objc.GetClass("MKPointAnnotation")}
	})
	return MKPointAnnotationClass
}

type _MKPointAnnotationClass struct {
	class objc.Class
}

// An interface definition for the [MKPointAnnotation] class.
type IMKPointAnnotation interface {
	IMKShape
}

// A string-based piece of location-specific data that you apply to a specific point on a map.
//
// You use this class, rather than define a custom annotation object, in situations where all you want to do is display a title string at the specified point on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPointAnnotation
type MKPointAnnotation struct {
	MKShape
}

// MKPointAnnotationFrom constructs a [MKPointAnnotation] from an unsafe.Pointer.
//
// A string-based piece of location-specific data that you apply to a specific point on a map.
func MKPointAnnotationFrom(ptr unsafe.Pointer) MKPointAnnotation {
	return MKPointAnnotation{
		MKShape: MKShapeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKPointAnnotationClass) Alloc() MKPointAnnotation {
	rv := objc.Send[MKPointAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKPointAnnotationClass) New() MKPointAnnotation {
	rv := objc.Send[MKPointAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPointAnnotation) Init() MKPointAnnotation {
	rv := objc.Send[MKPointAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPointAnnotation) Autorelease() MKPointAnnotation {
	rv := objc.Send[MKPointAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPointAnnotation creates a new MKPointAnnotation instance.
func NewMKPointAnnotation() MKPointAnnotation {
	return getMKPointAnnotationClass().New()
}




