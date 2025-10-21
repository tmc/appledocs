// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AXDataPoint] class.
type IAXDataPoint interface {
	objectivec.IObject
}

// An object that represents a single data point in a chart.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AXDataPointClass) Alloc() AXDataPoint {
	rv := objc.Send[AXDataPoint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a data point with the specified x-value, y-value, additional values, and label.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:additionalValues:label:
func NewAXDataPointWithXYAdditionalValuesLabel(xValue unsafe.Pointer, yValue unsafe.Pointer, additionalValues unsafe.Pointer, label string) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[AXDataPoint](instance.ID, objc.Sel("initWithX:y:additionalValues:label:"), xValue, yValue, additionalValues, objc.String(label))
	rv.Autorelease()
	return rv
}

// Creates a data point with the specified x- and y-values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:
func NewAXDataPointWithXY(xValue unsafe.Pointer, yValue unsafe.Pointer) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[AXDataPoint](instance.ID, objc.Sel("initWithX:y:"), xValue, yValue)
	rv.Autorelease()
	return rv
}

// Creates a data point with the specified x-value, y-value, and additional values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:additionalValues:
func NewAXDataPointWithXYAdditionalValues(xValue unsafe.Pointer, yValue unsafe.Pointer, additionalValues unsafe.Pointer) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[AXDataPoint](instance.ID, objc.Sel("initWithX:y:additionalValues:"), xValue, yValue, additionalValues)
	rv.Autorelease()
	return rv
}


// An array of values for additional axes for the data point.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/additionalValues
func (a_ AXDataPoint) AdditionalValues() []AXDataPointValue {
	rv := objc.Send[[]AXDataPointValue](a_.ID, objc.Sel("additionalValues"))
	return rv
}


// SetAdditionalValues sets the value of the additionalValues property.
// An array of values for additional axes for the data point.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/additionalValues
func (a_ AXDataPoint) SetAdditionalValues(value []AXDataPointValue) {
	// Convert Go slice to NSArray
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
}
// An attributed version of the label for the data point.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/attributedLabel
func (a_ AXDataPoint) AttributedLabel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributedLabel"))
	return rv
}


// SetAttributedLabel sets the value of the attributedLabel property.
// An attributed version of the label for the data point.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/attributedLabel
func (a_ AXDataPoint) SetAttributedLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedLabel:"), value)
}
// The label for the data point.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/label
func (a_ AXDataPoint) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// The label for the data point.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/label
func (a_ AXDataPoint) SetLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), value)
}
// The value of the x-axis for the data point.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/xValue
func (a_ AXDataPoint) XValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("xValue"))
	return rv
}


// SetXValue sets the value of the xValue property.
// The value of the x-axis for the data point.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/xValue
func (a_ AXDataPoint) SetXValue(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setXValue:"), value)
}
// The value of the y-axis for the data point.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/yValue
func (a_ AXDataPoint) YValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("yValue"))
	return rv
}


// SetYValue sets the value of the yValue property.
// The value of the y-axis for the data point.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataPoint/yValue
func (a_ AXDataPoint) SetYValue(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setYValue:"), value)
}

