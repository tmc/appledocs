// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SliderTouchBarItem] class.
var (
	sliderTouchBarItemClass     _SliderTouchBarItemClass
	sliderTouchBarItemClassOnce sync.Once
)

func getSliderTouchBarItemClass() _SliderTouchBarItemClass {
	sliderTouchBarItemClassOnce.Do(func() {
		sliderTouchBarItemClass = _SliderTouchBarItemClass{objc.GetClass("NSSliderTouchBarItem")}
	})
	return sliderTouchBarItemClass
}

type _SliderTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [SliderTouchBarItem] class.
type ISliderTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a slider control for choosing a value in a range. [Full Topic]
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




