// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SliderTouchBarItem] class.
var (
	SliderTouchBarItemClass     _SliderTouchBarItemClass
	SliderTouchBarItemClassOnce sync.Once
)

func getSliderTouchBarItemClass() _SliderTouchBarItemClass {
	SliderTouchBarItemClassOnce.Do(func() {
		SliderTouchBarItemClass = _SliderTouchBarItemClass{objc.GetClass("NSSliderTouchBarItem")}
	})
	return SliderTouchBarItemClass
}

type _SliderTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [SliderTouchBarItem] class.
type ISliderTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a slider control for choosing a value in a range.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem
type SliderTouchBarItem struct {
	TouchBarItem
}

// SliderTouchBarItemFrom constructs a [SliderTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a slider control for choosing a value in a range.
func SliderTouchBarItemFrom(ptr unsafe.Pointer) SliderTouchBarItem {
	return SliderTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SliderTouchBarItemClass) Alloc() SliderTouchBarItem {
	rv := objc.Send[SliderTouchBarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SliderTouchBarItemClass) New() SliderTouchBarItem {
	rv := objc.Send[SliderTouchBarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SliderTouchBarItem) Init() SliderTouchBarItem {
	rv := objc.Send[SliderTouchBarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SliderTouchBarItem) Autorelease() SliderTouchBarItem {
	rv := objc.Send[SliderTouchBarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSliderTouchBarItem creates a new SliderTouchBarItem instance.
func NewSliderTouchBarItem() SliderTouchBarItem {
	return getSliderTouchBarItemClass().New()
}

// The user-visible string identifying this item during bar customization.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/customizationLabel
func (s_ SliderTouchBarItem) CustomizationLabel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("customizationLabel"))
	return rv
}

// SetCustomizationLabel sets the value of the customizationLabel property.
// The user-visible string identifying this item during bar customization.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/customizationLabel
func (s_ SliderTouchBarItem) SetCustomizationLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizationLabel:"), value)
}

// The accessory that appears at the end of the slider with the minimum value.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumValueAccessory
func (s_ SliderTouchBarItem) MinimumValueAccessory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("minimumValueAccessory"))
	return rv
}

// SetMinimumValueAccessory sets the value of the minimumValueAccessory property.
// The accessory that appears at the end of the slider with the minimum value.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumValueAccessory
func (s_ SliderTouchBarItem) SetMinimumValueAccessory(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumValueAccessory:"), value)
}
