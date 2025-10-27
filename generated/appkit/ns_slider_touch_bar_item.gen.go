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
	

	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	CustomizationLabel() foundation.foundation.INSString
	SetCustomizationLabel(value foundation.foundation.INSString)
	DoubleValue() float64
	SetDoubleValue(value float64)
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	MaximumSliderWidth() float64
	SetMaximumSliderWidth(value float64)
	MaximumValueAccessory() ISliderAccessory
	SetMaximumValueAccessory(value ISliderAccessory)
	MinimumSliderWidth() float64
	SetMinimumSliderWidth(value float64)
	MinimumValueAccessory() ISliderAccessory
	SetMinimumValueAccessory(value ISliderAccessory)
	Slider() ISlider
	SetSlider(value ISlider)
	Target() objc.ID
	SetTarget(value objc.ID)
	ValueAccessoryWidth() SliderAccessoryWidth
	SetValueAccessoryWidth(value SliderAccessoryWidth)
	View() unsafe.Pointer


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SliderTouchBarItemClass) Alloc() SliderTouchBarItem {
	rv := objc.Send[SliderTouchBarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A bar item that provides a slider control for choosing a value in a range.


// A bar item that provides a slider control for choosing a value in a range.
//
// [Full Topic]
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

























// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/action
func (s_ SliderTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("action"))
	return rv
}


// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/action
func (s_ SliderTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/customizationLabel
func (s_ SliderTouchBarItem) CustomizationLabel() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/customizationLabel
func (s_ SliderTouchBarItem) SetCustomizationLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// The double value of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/doubleValue
func (s_ SliderTouchBarItem) DoubleValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("doubleValue"))
	return rv
}


// The double value of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/doubleValue
func (s_ SliderTouchBarItem) SetDoubleValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDoubleValue:"), value)
}


// The text displayed alongside the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/label
func (s_ SliderTouchBarItem) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("label"))
	return rv
}


// The text displayed alongside the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/label
func (s_ SliderTouchBarItem) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), value)
}


// The maximum width of the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumSliderWidth
func (s_ SliderTouchBarItem) MaximumSliderWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumSliderWidth"))
	return rv
}


// The maximum width of the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumSliderWidth
func (s_ SliderTouchBarItem) SetMaximumSliderWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumSliderWidth:"), value)
}


// The accessory that appears at the end of the slider with the maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumValueAccessory
func (s_ SliderTouchBarItem) MaximumValueAccessory() ISliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID, objc.Sel("maximumValueAccessory"))
	return rv
}


// The accessory that appears at the end of the slider with the maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumValueAccessory
func (s_ SliderTouchBarItem) SetMaximumValueAccessory(value ISliderAccessory) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumValueAccessory:"), value)
}


// The minimum width of the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumSliderWidth
func (s_ SliderTouchBarItem) MinimumSliderWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumSliderWidth"))
	return rv
}


// The minimum width of the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumSliderWidth
func (s_ SliderTouchBarItem) SetMinimumSliderWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumSliderWidth:"), value)
}


// The accessory that appears at the end of the slider with the minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumValueAccessory
func (s_ SliderTouchBarItem) MinimumValueAccessory() ISliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID, objc.Sel("minimumValueAccessory"))
	return rv
}


// The accessory that appears at the end of the slider with the minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumValueAccessory
func (s_ SliderTouchBarItem) SetMinimumValueAccessory(value ISliderAccessory) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumValueAccessory:"), value)
}


// The slider displayed by the bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/slider
func (s_ SliderTouchBarItem) Slider() ISlider {
	rv := objc.Send[Slider](s_.ID, objc.Sel("slider"))
	return rv
}


// The slider displayed by the bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/slider
func (s_ SliderTouchBarItem) SetSlider(value ISlider) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSlider:"), value)
}


// An object that is notified when a user interacts with the slider or either of the accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/target
func (s_ SliderTouchBarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("target"))
	return rv
}


// An object that is notified when a user interacts with the slider or either of the accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/target
func (s_ SliderTouchBarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTarget:"), value)
}


// The width of the value accessories that appear at either end of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/valueAccessoryWidth
func (s_ SliderTouchBarItem) ValueAccessoryWidth() SliderAccessoryWidth {
	rv := objc.Send[SliderAccessoryWidth](s_.ID, objc.Sel("valueAccessoryWidth"))
	return rv
}


// The width of the value accessories that appear at either end of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/valueAccessoryWidth
func (s_ SliderTouchBarItem) SetValueAccessoryWidth(value SliderAccessoryWidth) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValueAccessoryWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/view
func (s_ SliderTouchBarItem) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("view"))
	return rv
}








