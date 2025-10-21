// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKClusterAnnotation] class.
var (
	MKClusterAnnotationClass     _MKClusterAnnotationClass
	MKClusterAnnotationClassOnce sync.Once
)

func getMKClusterAnnotationClass() _MKClusterAnnotationClass {
	MKClusterAnnotationClassOnce.Do(func() {
		MKClusterAnnotationClass = _MKClusterAnnotationClass{objc.GetClass("MKClusterAnnotation")}
	})
	return MKClusterAnnotationClass
}

type _MKClusterAnnotationClass struct {
	class objc.Class
}

// An interface definition for the [MKClusterAnnotation] class.
type IMKClusterAnnotation interface {
	objectivec.IObject
}

// An annotation that groups two or more distinct annotations into a single entity.
//
// A cluster annotation object stands in for the group of annotations. Cluster views promote legibility of the underlying annotations by displaying a single annotation that takes it’s title from one annotation and includes a subtitle that indicates how many additional annotations belong to the group. MapKit automatically creates cluster annotations when two or more annotation views group too closely together on the map surface. To customize the cluster annotations that display on your map, implement the method in your map’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKClusterAnnotation
type MKClusterAnnotation struct {
	objectivec.Object
}

// MKClusterAnnotationFrom constructs a [MKClusterAnnotation] from an unsafe.Pointer.
//
// An annotation that groups two or more distinct annotations into a single entity.
func MKClusterAnnotationFrom(ptr unsafe.Pointer) MKClusterAnnotation {
	return MKClusterAnnotation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKClusterAnnotationClass) Alloc() MKClusterAnnotation {
	rv := objc.Send[MKClusterAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKClusterAnnotationClass) New() MKClusterAnnotation {
	rv := objc.Send[MKClusterAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKClusterAnnotation) Init() MKClusterAnnotation {
	rv := objc.Send[MKClusterAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKClusterAnnotation) Autorelease() MKClusterAnnotation {
	rv := objc.Send[MKClusterAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKClusterAnnotation creates a new MKClusterAnnotation instance.
func NewMKClusterAnnotation() MKClusterAnnotation {
	return getMKClusterAnnotationClass().New()
}




