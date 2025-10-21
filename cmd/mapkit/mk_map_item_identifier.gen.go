// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapItemIdentifier] class.
var (
	MKMapItemIdentifierClass     _MKMapItemIdentifierClass
	MKMapItemIdentifierClassOnce sync.Once
)

func getMKMapItemIdentifierClass() _MKMapItemIdentifierClass {
	MKMapItemIdentifierClassOnce.Do(func() {
		MKMapItemIdentifierClass = _MKMapItemIdentifierClass{objc.GetClass("MKMapItemIdentifier")}
	})
	return MKMapItemIdentifierClass
}

type _MKMapItemIdentifierClass struct {
	class objc.Class
}

// An interface definition for the [MKMapItemIdentifier] class.
type IMKMapItemIdentifier interface {
	objectivec.IObject
}

// A unique identifier for a place.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/Identifier-swift.class
type MKMapItemIdentifier struct {
	objectivec.Object
}

// MKMapItemIdentifierFrom constructs a [MKMapItemIdentifier] from an unsafe.Pointer.
//
// A unique identifier for a place.
func MKMapItemIdentifierFrom(ptr unsafe.Pointer) MKMapItemIdentifier {
	return MKMapItemIdentifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapItemIdentifierClass) Alloc() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapItemIdentifierClass) New() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemIdentifier) Init() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemIdentifier) Autorelease() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemIdentifier creates a new MKMapItemIdentifier instance.
func NewMKMapItemIdentifier() MKMapItemIdentifier {
	return getMKMapItemIdentifierClass().New()
}




