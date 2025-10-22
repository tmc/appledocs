// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AXDataPointValue] class.
type IAXDataPointValue interface {
	objectivec.IObject
	XValue() AXDataPointValue
	SetXValue(value IAXDataPointValue)
	YValue() AXDataPointValue
	SetYValue(value IAXDataPointValue)
}

// A single data value.
//
// An can be either numeric or categorical. Data points in a numeric axis use the property, and data points in a categorical axis use the property.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AXDataPointValueClass) Alloc() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The value of the x-axis for the data point.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/xvalue
func (a_ AXDataPointValue) XValue() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("xValue"))
	return rv
}


// SetXValue sets the value of the xValue property.
// The value of the x-axis for the data point.

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/xvalue
func (a_ AXDataPointValue) SetXValue(value IAXDataPointValue) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setXValue:"), value)
}

// The value of the y-axis for the data point.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/yvalue
func (a_ AXDataPointValue) YValue() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("yValue"))
	return rv
}


// SetYValue sets the value of the yValue property.
// The value of the y-axis for the data point.

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/yvalue
func (a_ AXDataPointValue) SetYValue(value IAXDataPointValue) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setYValue:"), value)
}



