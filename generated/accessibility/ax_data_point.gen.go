// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXDataPoint */


/* debug [class_header]: Header for AXDataPoint */
// The class instance for the [AXDataPoint] class.
var (
	AXDataPointClass     _AXDataPointClass
	AXDataPointClassOnce sync.Once
)

func getAXDataPointClass() _AXDataPointClass {
	AXDataPointClassOnce.Do(func() {
		AXDataPointClass = _AXDataPointClass{objc.GetClass("AXDataPoint")}
	})
	return AXDataPointClass
}

type _AXDataPointClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXDataPoint */
// An interface definition for the [AXDataPoint] class.
type IAXDataPoint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXDataPoint */
	// properties:
	AdditionalValues() []AXDataPointValue
	SetAdditionalValues(value []AXDataPointValue)
	AttributedLabel() foundation.AttributedString
	SetAttributedLabel(value foundation.AttributedString)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	XValue() IAXDataPointValue
	SetXValue(value IAXDataPointValue)
	YValue() IAXDataPointValue
	SetYValue(value IAXDataPointValue)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXDataPoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXDataPoint */
// Alloc allocates a new instance without initialization.
func (ac _AXDataPointClass) Alloc() AXDataPoint {
	rv := objc.Send[AXDataPoint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXDataPointClass) New() AXDataPoint {
	rv := objc.Send[AXDataPoint](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXDataPoint) Init() AXDataPoint {
	rv := objc.Send[AXDataPoint](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXDataPoint) Autorelease() AXDataPoint {
	rv := objc.Send[AXDataPoint](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXDataPoint creates a new AXDataPoint instance.
func NewAXDataPoint() AXDataPoint {
	return getAXDataPointClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXDataPoint */
// An object that represents a single data point in a chart.


// An object that represents a single data point in a chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint
type AXDataPoint struct {
	objectivec.Object
}

// AXDataPointFrom constructs a [AXDataPoint] from an unsafe.Pointer.
//
// An object that represents a single data point in a chart.
func AXDataPointFrom(ptr unsafe.Pointer) AXDataPoint {
	return AXDataPoint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXDataPoint */

// Creates a data point with the specified x- and y-values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:
func NewAXDataPointWithXY(xValue IAXDataPointValue, yValue IAXDataPointValue) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[AXDataPoint](instance.ID, objc.Sel("initWithX:y:"), xValue, yValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXDataPointWithXY */


// Creates a data point with the specified x-value, y-value, and additional values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:additionalValues:
func NewAXDataPointWithXYAdditionalValues(xValue IAXDataPointValue, yValue IAXDataPointValue, additionalValues []AXDataPointValue) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[AXDataPoint](instance.ID, objc.Sel("initWithX:y:additionalValues:"), xValue, yValue, additionalValues)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXDataPointWithXYAdditionalValues */


// Creates a data point with the specified x-value, y-value, additional values, and label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:additionalValues:label:
func NewAXDataPointWithXYAdditionalValuesLabel(xValue IAXDataPointValue, yValue IAXDataPointValue, additionalValues []AXDataPointValue, label objc.IObject /* cross-framework: NSString */) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[AXDataPoint](instance.ID, objc.Sel("initWithX:y:additionalValues:label:"), xValue, yValue, additionalValues, label)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXDataPointWithXYAdditionalValuesLabel */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXDataPoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXDataPoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXDataPoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXDataPoint */

// An array of values for additional axes for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/additionalValues
func (a_ AXDataPoint) AdditionalValues() []AXDataPointValue {
	rv := objc.Send[[]AXDataPointValue](a_.ID, objc.Sel("additionalValues"))
	return rv
}/* debug [instance_properties/getter]: additionalValues */


// An array of values for additional axes for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/additionalValues
func (a_ AXDataPoint) SetAdditionalValues(value []AXDataPointValue) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdditionalValues:"), nsArray)
}/* debug [instance_properties/setter]: additionalValues */


// An attributed version of the label for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/attributedLabel
func (a_ AXDataPoint) AttributedLabel() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedLabel"))
	return rv
}/* debug [instance_properties/getter]: attributedLabel */


// An attributed version of the label for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/attributedLabel
func (a_ AXDataPoint) SetAttributedLabel(value foundation.AttributedString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedLabel:"), value)
}/* debug [instance_properties/setter]: attributedLabel */


// The label for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/label
func (a_ AXDataPoint) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// The label for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/label
func (a_ AXDataPoint) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The value of the x-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/xValue
func (a_ AXDataPoint) XValue() IAXDataPointValue {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("xValue"))
	return rv
}/* debug [instance_properties/getter]: xValue */


// The value of the x-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/xValue
func (a_ AXDataPoint) SetXValue(value IAXDataPointValue) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setXValue:"), value)
}/* debug [instance_properties/setter]: xValue */


// The value of the y-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/yValue
func (a_ AXDataPoint) YValue() IAXDataPointValue {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("yValue"))
	return rv
}/* debug [instance_properties/getter]: yValue */


// The value of the y-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/yValue
func (a_ AXDataPoint) SetYValue(value IAXDataPointValue) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setYValue:"), value)
}/* debug [instance_properties/setter]: yValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXDataPoint */


