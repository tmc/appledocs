// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureSlider */


/* debug [class_header]: Header for AVCaptureSlider */
// The class instance for the [CaptureSlider] class.
var (
	CaptureSliderClass     _CaptureSliderClass
	CaptureSliderClassOnce sync.Once
)

func getCaptureSliderClass() _CaptureSliderClass {
	CaptureSliderClassOnce.Do(func() {
		CaptureSliderClass = _CaptureSliderClass{objc.GetClass("AVCaptureSlider")}
	})
	return CaptureSliderClass
}

type _CaptureSliderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSlider */
// An interface definition for the [CaptureSlider] class.
type ICaptureSlider interface {
	ICaptureControl
	
/* debug [class_interface_properties]: Properties for CaptureSlider */
	// properties:
	AccessibilityIdentifier() objc.IObject /* cross-framework: NSString */
	SetAccessibilityIdentifier(value objc.IObject /* cross-framework: NSString */)
	LocalizedTitle() objc.IObject /* cross-framework: NSString */
	LocalizedValueFormat() objc.IObject /* cross-framework: NSString */
	SetLocalizedValueFormat(value objc.IObject /* cross-framework: NSString */)
	ProminentValues() []foundation.Number
	SetProminentValues(value []foundation.Number)
	SymbolName() objc.IObject /* cross-framework: NSString */
	Value() float32
	SetValue(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSlider */
	// methods:
	SetActionQueueAction(actionQueue objectivec.IObject, action unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSlider */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSliderClass) Alloc() CaptureSlider {
	rv := objc.Send[CaptureSlider](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSliderClass) New() CaptureSlider {
	rv := objc.Send[CaptureSlider](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSlider) Init() CaptureSlider {
	rv := objc.Send[CaptureSlider](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSlider) Autorelease() CaptureSlider {
	rv := objc.Send[CaptureSlider](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSlider creates a new CaptureSlider instance.
func NewCaptureSlider() CaptureSlider {
	return getCaptureSliderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSlider */
// A slider control that selects a value from a bounded range.
//
// Sliders are appropriate for controls that provide a single float value.


// A slider control that selects a value from a bounded range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider
type CaptureSlider struct {
	CaptureControl
}

// CaptureSliderFrom constructs a [CaptureSlider] from an unsafe.Pointer.
//
// A slider control that selects a value from a bounded range.
func CaptureSliderFrom(ptr unsafe.Pointer) CaptureSlider {
	return CaptureSlider{
		CaptureControl: CaptureControlFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSlider */

// Creates a continuous slider control that selects a value from a bounded range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:minValue:maxValue:
func NewCaptureSliderWithLocalizedTitleSymbolNameMinValueMaxValue(localizedTitle objc.IObject /* cross-framework: NSString */, symbolName objc.IObject /* cross-framework: NSString */, minValue float32, maxValue float32) CaptureSlider {
	instance := getCaptureSliderClass().Alloc()
	rv := objc.Send[CaptureSlider](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:minValue:maxValue:"), localizedTitle, symbolName, minValue, maxValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureSliderWithLocalizedTitleSymbolNameMinValueMaxValue */


// Creates a discrete slider control that selects a stepped value from a bounded range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:minValue:maxValue:step:
func NewCaptureSliderWithLocalizedTitleSymbolNameMinValueMaxValueStep(localizedTitle objc.IObject /* cross-framework: NSString */, symbolName objc.IObject /* cross-framework: NSString */, minValue float32, maxValue float32, step float32) CaptureSlider {
	instance := getCaptureSliderClass().Alloc()
	rv := objc.Send[CaptureSlider](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:minValue:maxValue:step:"), localizedTitle, symbolName, minValue, maxValue, step)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureSliderWithLocalizedTitleSymbolNameMinValueMaxValueStep */


// Creates a discrete slider control that selects a value from a list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:values:
func NewCaptureSliderWithLocalizedTitleSymbolNameValues(localizedTitle objc.IObject /* cross-framework: NSString */, symbolName objc.IObject /* cross-framework: NSString */, values []foundation.Number) CaptureSlider {
	instance := getCaptureSliderClass().Alloc()
	rv := objc.Send[CaptureSlider](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:values:"), localizedTitle, symbolName, values)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureSliderWithLocalizedTitleSymbolNameValues */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSlider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSlider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSlider */

// Sets the action to perform on the specified dispatch queue when the slider’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/setActionQueue:action:
func (c_ CaptureSlider) SetActionQueueAction(actionQueue objectivec.IObject, action unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActionQueue:action:"), actionQueue, action)
}/* debug [instance_methods/method]: SetActionQueueAction */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSlider */

// A string identifier for the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/accessibilityIdentifier
func (c_ CaptureSlider) AccessibilityIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("accessibilityIdentifier"))
	return rv
}/* debug [instance_properties/getter]: accessibilityIdentifier */


// A string identifier for the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/accessibilityIdentifier
func (c_ CaptureSlider) SetAccessibilityIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccessibilityIdentifier:"), value)
}/* debug [instance_properties/setter]: accessibilityIdentifier */


// A localized title that describes the control’s action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/localizedTitle
func (c_ CaptureSlider) LocalizedTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedTitle"))
	return rv
}/* debug [instance_properties/getter]: localizedTitle */


// A localized string that defines the presentation of the slider’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/localizedValueFormat
func (c_ CaptureSlider) LocalizedValueFormat() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedValueFormat"))
	return rv
}/* debug [instance_properties/getter]: localizedValueFormat */


// A localized string that defines the presentation of the slider’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/localizedValueFormat
func (c_ CaptureSlider) SetLocalizedValueFormat(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedValueFormat:"), value)
}/* debug [instance_properties/setter]: localizedValueFormat */


// Values in this array may receive unique visual representations or behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/prominentValues-7usgc
func (c_ CaptureSlider) ProminentValues() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("prominentValues"))
	return rv
}/* debug [instance_properties/getter]: prominentValues */


// Values in this array may receive unique visual representations or behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/prominentValues-7usgc
func (c_ CaptureSlider) SetProminentValues(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setProminentValues:"), nsArray)
}/* debug [instance_properties/setter]: prominentValues */


// The name of the SF Symbol that represents this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/symbolName
func (c_ CaptureSlider) SymbolName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("symbolName"))
	return rv
}/* debug [instance_properties/getter]: symbolName */


// The current value of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/value
func (c_ CaptureSlider) Value() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The current value of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/value
func (c_ CaptureSlider) SetValue(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSlider */


