// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapItemAnnotation] class.
var (
	MKMapItemAnnotationClass     _MKMapItemAnnotationClass
	MKMapItemAnnotationClassOnce sync.Once
)

func getMKMapItemAnnotationClass() _MKMapItemAnnotationClass {
	MKMapItemAnnotationClassOnce.Do(func() {
		MKMapItemAnnotationClass = _MKMapItemAnnotationClass{objc.GetClass("MKMapItemAnnotation")}
	})
	return MKMapItemAnnotationClass
}

type _MKMapItemAnnotationClass struct {
	class objc.Class
}

// An interface definition for the [MKMapItemAnnotation] class.
type IMKMapItemAnnotation interface {
	objectivec.IObject
}

// An annotation that represents a map item
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemAnnotation
type MKMapItemAnnotation struct {
	objectivec.Object
}

// MKMapItemAnnotationFrom constructs a [MKMapItemAnnotation] from an unsafe.Pointer.
//
// An annotation that represents a map item
func MKMapItemAnnotationFrom(ptr unsafe.Pointer) MKMapItemAnnotation {
	return MKMapItemAnnotation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapItemAnnotationClass) Alloc() MKMapItemAnnotation {
	rv := objc.Send[MKMapItemAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapItemAnnotationClass) New() MKMapItemAnnotation {
	rv := objc.Send[MKMapItemAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemAnnotation) Init() MKMapItemAnnotation {
	rv := objc.Send[MKMapItemAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemAnnotation) Autorelease() MKMapItemAnnotation {
	rv := objc.Send[MKMapItemAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemAnnotation creates a new MKMapItemAnnotation instance.
func NewMKMapItemAnnotation() MKMapItemAnnotation {
	return getMKMapItemAnnotationClass().New()
}




// Creates a map item annotation
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemAnnotation/init(mapItem:)
func NewMKMapItemAnnotationWithMapItem(mapItem unsafe.Pointer) MKMapItemAnnotation {
	instance := getMKMapItemAnnotationClass().Alloc()
	rv := objc.Send[MKMapItemAnnotation](instance.ID, objc.Sel("initWithMapItem:"), mapItem)
	rv.Autorelease()
	return rv
}


// The map item represented by this annotation
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemAnnotation/mapItem
func (m_ MKMapItemAnnotation) MapItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapItem"))
	return rv
}


