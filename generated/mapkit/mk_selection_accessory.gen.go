// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKSelectionAccessory] class.
var (
	MKSelectionAccessoryClass     _MKSelectionAccessoryClass
	MKSelectionAccessoryClassOnce sync.Once
)

func getMKSelectionAccessoryClass() _MKSelectionAccessoryClass {
	MKSelectionAccessoryClassOnce.Do(func() {
		MKSelectionAccessoryClass = _MKSelectionAccessoryClass{objc.GetClass("MKSelectionAccessory")}
	})
	return MKSelectionAccessoryClass
}

type _MKSelectionAccessoryClass struct {
	class objc.Class
}

// An interface definition for the [MKSelectionAccessory] class.
type IMKSelectionAccessory interface {
	objectivec.IObject
}

// The type of accessory to display for a selected annotation.
//
// Implement in your map view delegate to specify a selection accessory for annotation content.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory
type MKSelectionAccessory struct {
	objectivec.Object
}

// MKSelectionAccessoryFrom constructs a [MKSelectionAccessory] from an unsafe.Pointer.
//
// The type of accessory to display for a selected annotation.
func MKSelectionAccessoryFrom(ptr unsafe.Pointer) MKSelectionAccessory {
	return MKSelectionAccessory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKSelectionAccessoryClass) Alloc() MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKSelectionAccessoryClass) New() MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKSelectionAccessory) Init() MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKSelectionAccessory) Autorelease() MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKSelectionAccessory creates a new MKSelectionAccessory instance.
func NewMKSelectionAccessory() MKSelectionAccessory {
	return getMKSelectionAccessoryClass().New()
}




