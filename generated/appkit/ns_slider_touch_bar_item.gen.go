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
func (s_ SliderTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](s_.ID, objc.Sel("customizationLabel"))
	return rv
}


// SetCustomizationLabel sets the value of the customizationLabel property.
// The user-visible string identifying this item during bar customization.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/customizationLabel
func (s_ SliderTouchBarItem) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
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

// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/action
func (s_ SliderTouchBarItem) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/action
func (s_ SliderTouchBarItem) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}

// The double value of the slider.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/doublevalue
func (s_ SliderTouchBarItem) DoubleValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("doubleValue"))
	return rv
}


// SetDoubleValue sets the value of the doubleValue property.
// The double value of the slider.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/doublevalue
func (s_ SliderTouchBarItem) SetDoubleValue(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDoubleValue:"), value)
}

// The text displayed alongside the slider.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/label
func (s_ SliderTouchBarItem) Label() string {
	rv := objc.Send[string](s_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// The text displayed alongside the slider.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/label
func (s_ SliderTouchBarItem) SetLabel(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// The maximum width of the slider’s track.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/maximumsliderwidth
func (s_ SliderTouchBarItem) MaximumSliderWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumSliderWidth"))
	return rv
}


// SetMaximumSliderWidth sets the value of the maximumSliderWidth property.
// The maximum width of the slider’s track.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/maximumsliderwidth
func (s_ SliderTouchBarItem) SetMaximumSliderWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumSliderWidth:"), value)
}

// The accessory that appears at the end of the slider with the maximum value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/maximumvalueaccessory
func (s_ SliderTouchBarItem) MaximumValueAccessory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("maximumValueAccessory"))
	return rv
}


// SetMaximumValueAccessory sets the value of the maximumValueAccessory property.
// The accessory that appears at the end of the slider with the maximum value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/maximumvalueaccessory
func (s_ SliderTouchBarItem) SetMaximumValueAccessory(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumValueAccessory:"), value)
}

// The minimum width of the slider’s track.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/minimumsliderwidth
func (s_ SliderTouchBarItem) MinimumSliderWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumSliderWidth"))
	return rv
}


// SetMinimumSliderWidth sets the value of the minimumSliderWidth property.
// The minimum width of the slider’s track.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/minimumsliderwidth
func (s_ SliderTouchBarItem) SetMinimumSliderWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumSliderWidth:"), value)
}

// The slider displayed by the bar item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/slider
func (s_ SliderTouchBarItem) Slider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("slider"))
	return rv
}


// SetSlider sets the value of the slider property.
// The slider displayed by the bar item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/slider
func (s_ SliderTouchBarItem) SetSlider(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSlider:"), value)
}

// An object that is notified when a user interacts with the slider or either of the accessories.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/target
func (s_ SliderTouchBarItem) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// An object that is notified when a user interacts with the slider or either of the accessories.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/target
func (s_ SliderTouchBarItem) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTarget:"), value)
}

// The width of the value accessories that appear at either end of the slider.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/valueaccessorywidth
func (s_ SliderTouchBarItem) ValueAccessoryWidth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("valueAccessoryWidth"))
	return rv
}


// SetValueAccessoryWidth sets the value of the valueAccessoryWidth property.
// The width of the value accessories that appear at either end of the slider.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/valueaccessorywidth
func (s_ SliderTouchBarItem) SetValueAccessoryWidth(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValueAccessoryWidth:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/view
func (s_ SliderTouchBarItem) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidertouchbaritem/view
func (s_ SliderTouchBarItem) SetView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setView:"), value)
}



