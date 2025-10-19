// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SliderAccessory] class.
var sliderAccessoryClass = _SliderAccessoryClass{objc.GetClass("NSSliderAccessory")}

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

// New creates and returns a new instance with a +1 retain count.
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
	return sliderAccessoryClass.New()
}




