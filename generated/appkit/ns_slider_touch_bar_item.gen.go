// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSSliderTouchBarItem */


/* debug [class_header]: Header for NSSliderTouchBarItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SliderTouchBarItem */
// An interface definition for the [SliderTouchBarItem] class.
type ISliderTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for SliderTouchBarItem */
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	DoubleValue() float64
	SetDoubleValue(value float64)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
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
	ValueAccessoryWidth() SliderAccessoryWidth /* typedef */
	SetValueAccessoryWidth(value SliderAccessoryWidth /* typedef */)
	View() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SliderTouchBarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SliderTouchBarItem */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SliderTouchBarItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SliderTouchBarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SliderTouchBarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SliderTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SliderTouchBarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SliderTouchBarItem */

// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/action
func (s_ SliderTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/action
func (s_ SliderTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/customizationLabel
func (s_ SliderTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/customizationLabel
func (s_ SliderTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// The double value of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/doubleValue
func (s_ SliderTouchBarItem) DoubleValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// The double value of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/doubleValue
func (s_ SliderTouchBarItem) SetDoubleValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDoubleValue:"), value)
}/* debug [instance_properties/setter]: doubleValue */


// The text displayed alongside the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/label
func (s_ SliderTouchBarItem) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// The text displayed alongside the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/label
func (s_ SliderTouchBarItem) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The maximum width of the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumSliderWidth
func (s_ SliderTouchBarItem) MaximumSliderWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumSliderWidth"))
	return rv
}/* debug [instance_properties/getter]: maximumSliderWidth */


// The maximum width of the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumSliderWidth
func (s_ SliderTouchBarItem) SetMaximumSliderWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumSliderWidth:"), value)
}/* debug [instance_properties/setter]: maximumSliderWidth */


// The accessory that appears at the end of the slider with the maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumValueAccessory
func (s_ SliderTouchBarItem) MaximumValueAccessory() ISliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID, objc.Sel("maximumValueAccessory"))
	return rv
}/* debug [instance_properties/getter]: maximumValueAccessory */


// The accessory that appears at the end of the slider with the maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumValueAccessory
func (s_ SliderTouchBarItem) SetMaximumValueAccessory(value ISliderAccessory) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumValueAccessory:"), value)
}/* debug [instance_properties/setter]: maximumValueAccessory */


// The minimum width of the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumSliderWidth
func (s_ SliderTouchBarItem) MinimumSliderWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumSliderWidth"))
	return rv
}/* debug [instance_properties/getter]: minimumSliderWidth */


// The minimum width of the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumSliderWidth
func (s_ SliderTouchBarItem) SetMinimumSliderWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumSliderWidth:"), value)
}/* debug [instance_properties/setter]: minimumSliderWidth */


// The accessory that appears at the end of the slider with the minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumValueAccessory
func (s_ SliderTouchBarItem) MinimumValueAccessory() ISliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID, objc.Sel("minimumValueAccessory"))
	return rv
}/* debug [instance_properties/getter]: minimumValueAccessory */


// The accessory that appears at the end of the slider with the minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumValueAccessory
func (s_ SliderTouchBarItem) SetMinimumValueAccessory(value ISliderAccessory) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumValueAccessory:"), value)
}/* debug [instance_properties/setter]: minimumValueAccessory */


// The slider displayed by the bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/slider
func (s_ SliderTouchBarItem) Slider() ISlider {
	rv := objc.Send[Slider](s_.ID, objc.Sel("slider"))
	return rv
}/* debug [instance_properties/getter]: slider */


// The slider displayed by the bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/slider
func (s_ SliderTouchBarItem) SetSlider(value ISlider) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSlider:"), value)
}/* debug [instance_properties/setter]: slider */


// An object that is notified when a user interacts with the slider or either of the accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/target
func (s_ SliderTouchBarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// An object that is notified when a user interacts with the slider or either of the accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/target
func (s_ SliderTouchBarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// The width of the value accessories that appear at either end of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/valueAccessoryWidth
func (s_ SliderTouchBarItem) ValueAccessoryWidth() SliderAccessoryWidth /* typedef */ {
	rv := objc.Send[CGFloat](s_.ID, objc.Sel("valueAccessoryWidth"))
	return rv
}/* debug [instance_properties/getter]: valueAccessoryWidth */


// The width of the value accessories that appear at either end of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/valueAccessoryWidth
func (s_ SliderTouchBarItem) SetValueAccessoryWidth(value SliderAccessoryWidth /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValueAccessoryWidth:"), value)
}/* debug [instance_properties/setter]: valueAccessoryWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/view
func (s_ SliderTouchBarItem) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSliderTouchBarItem */



