// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TapDescription] class.
var (
	tapDescriptionClass     _TapDescriptionClass
	tapDescriptionClassOnce sync.Once
)

func getTapDescriptionClass() _TapDescriptionClass {
	tapDescriptionClassOnce.Do(func() {
		tapDescriptionClass = _TapDescriptionClass{objc.GetClass("CATapDescription")}
	})
	return tapDescriptionClass
}

type _TapDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [TapDescription] class.
type ITapDescription interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription
type TapDescription struct {
	objectivec.Object
}

// TapDescriptionFrom constructs a [TapDescription] from an unsafe.Pointer.
func TapDescriptionFrom(ptr unsafe.Pointer) TapDescription {
	return TapDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TapDescriptionClass) Alloc() TapDescription {
	rv := objc.Send[TapDescription](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TapDescriptionClass) New() TapDescription {
	rv := objc.Send[TapDescription](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TapDescription) Init() TapDescription {
	rv := objc.Send[TapDescription](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TapDescription) Autorelease() TapDescription {
	rv := objc.Send[TapDescription](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTapDescription creates a new TapDescription instance.
func NewTapDescription() TapDescription {
	return getTapDescriptionClass().New()
}



