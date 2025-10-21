// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SliderAccessory] class.
var (
	SliderAccessoryClass     _SliderAccessoryClass
	SliderAccessoryClassOnce sync.Once
)

func getSliderAccessoryClass() _SliderAccessoryClass {
	SliderAccessoryClassOnce.Do(func() {
		SliderAccessoryClass = _SliderAccessoryClass{objc.GetClass("NSSliderAccessory")}
	})
	return SliderAccessoryClass
}

type _SliderAccessoryClass struct {
	class objc.Class
}

// An interface definition for the [SliderAccessory] class.
type ISliderAccessory interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory
type SliderAccessory struct {
	objectivec.Object
}

// SliderAccessoryFrom constructs a [SliderAccessory] from an unsafe.Pointer.
func SliderAccessoryFrom(ptr unsafe.Pointer) SliderAccessory {
	return SliderAccessory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SliderAccessoryClass) Alloc() SliderAccessory {
	rv := objc.Send[SliderAccessory](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SliderAccessoryClass) New() SliderAccessory {
	rv := objc.Send[SliderAccessory](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SliderAccessory) Init() SliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SliderAccessory) Autorelease() SliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSliderAccessory creates a new SliderAccessory instance.
func NewSliderAccessory() SliderAccessory {
	return getSliderAccessoryClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessory/behavior
func (s_ SliderAccessory) Behavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("behavior"))
	return rv
}


// SetBehavior sets the value of the behavior property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessory/behavior
func (s_ SliderAccessory) SetBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBehavior:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessory/isenabled
func (s_ SliderAccessory) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslideraccessory/isenabled
func (s_ SliderAccessory) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}



