// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKAnnotationView] class.
var (
	MKAnnotationViewClass     _MKAnnotationViewClass
	MKAnnotationViewClassOnce sync.Once
)

func getMKAnnotationViewClass() _MKAnnotationViewClass {
	MKAnnotationViewClassOnce.Do(func() {
		MKAnnotationViewClass = _MKAnnotationViewClass{objc.GetClass("MKAnnotationView")}
	})
	return MKAnnotationViewClass
}

type _MKAnnotationViewClass struct {
	class objc.Class
}

// An interface definition for the [MKAnnotationView] class.
type IMKAnnotationView interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MapKit classes.


// A parent class referenced by other MapKit classes. [Full Topic]
type MKAnnotationView struct {
	objectivec.Object
}

// MKAnnotationViewFrom constructs a [MKAnnotationView] from an unsafe.Pointer.
//
// A parent class referenced by other MapKit classes.
func MKAnnotationViewFrom(ptr unsafe.Pointer) MKAnnotationView {
	return MKAnnotationView{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKAnnotationViewClass) Alloc() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKAnnotationViewClass) New() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKAnnotationView) Init() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKAnnotationView) Autorelease() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKAnnotationView creates a new MKAnnotationView instance.
func NewMKAnnotationView() MKAnnotationView {
	return getMKAnnotationViewClass().New()
}




