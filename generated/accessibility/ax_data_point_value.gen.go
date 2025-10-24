// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXDataPointValue */


/* debug [class_header]: Header for AXDataPointValue */
// The class instance for the [AXDataPointValue] class.
var (
	AXDataPointValueClass     _AXDataPointValueClass
	AXDataPointValueClassOnce sync.Once
)

func getAXDataPointValueClass() _AXDataPointValueClass {
	AXDataPointValueClassOnce.Do(func() {
		AXDataPointValueClass = _AXDataPointValueClass{objc.GetClass("AXDataPointValue")}
	})
	return AXDataPointValueClass
}

type _AXDataPointValueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXDataPointValue */
// An interface definition for the [AXDataPointValue] class.
type IAXDataPointValue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXDataPointValue */
	// properties:
	Category() objc.IObject /* cross-framework: NSString */
	SetCategory(value objc.IObject /* cross-framework: NSString */)
	Number() float64
	SetNumber(value float64)
	XValue() IAXDataPointValue
	SetXValue(value IAXDataPointValue)
	YValue() IAXDataPointValue
	SetYValue(value IAXDataPointValue)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXDataPointValue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXDataPointValue */
// Alloc allocates a new instance without initialization.
func (ac _AXDataPointValueClass) Alloc() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXDataPointValueClass) New() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXDataPointValue) Init() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXDataPointValue) Autorelease() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXDataPointValue creates a new AXDataPointValue instance.
func NewAXDataPointValue() AXDataPointValue {
	return getAXDataPointValueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXDataPointValue */
// A single data value.
//
// An can be either numeric or categorical. Data points in a numeric axis use the property, and data points in a categorical axis use the property.


// A single data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPointValue
type AXDataPointValue struct {
	objectivec.Object
}

// AXDataPointValueFrom constructs a [AXDataPointValue] from an unsafe.Pointer.
//
// A single data value.
func AXDataPointValueFrom(ptr unsafe.Pointer) AXDataPointValue {
	return AXDataPointValue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXDataPointValue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXDataPointValue */

// Creates a categorical data value with the specified category string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/valueWithCategory:
func (ac _AXDataPointValueClass) ValueWithCategory(category objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("valueWithCategory:"), category)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithCategory) */


// Creates a numeric data value with the specified number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/valueWithNumber:
func (ac _AXDataPointValueClass) ValueWithNumber(number float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("valueWithNumber:"), number)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNumber) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXDataPointValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXDataPointValue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXDataPointValue */

// A string that represents the categorical data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/category
func (a_ AXDataPointValue) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// A string that represents the categorical data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/category
func (a_ AXDataPointValue) SetCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCategory:"), value)
}/* debug [instance_properties/setter]: category */


// A number that represents the numeric data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/number
func (a_ AXDataPointValue) Number() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("number"))
	return rv
}/* debug [instance_properties/getter]: number */


// A number that represents the numeric data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/number
func (a_ AXDataPointValue) SetNumber(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumber:"), value)
}/* debug [instance_properties/setter]: number */


// The value of the x-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/xvalue
func (a_ AXDataPointValue) XValue() IAXDataPointValue {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("xValue"))
	return rv
}/* debug [instance_properties/getter]: xValue */


// The value of the x-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/xvalue
func (a_ AXDataPointValue) SetXValue(value IAXDataPointValue) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setXValue:"), value)
}/* debug [instance_properties/setter]: xValue */


// The value of the y-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/yvalue
func (a_ AXDataPointValue) YValue() IAXDataPointValue {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("yValue"))
	return rv
}/* debug [instance_properties/getter]: yValue */


// The value of the y-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/yvalue
func (a_ AXDataPointValue) SetYValue(value IAXDataPointValue) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setYValue:"), value)
}/* debug [instance_properties/setter]: yValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXDataPointValue */



