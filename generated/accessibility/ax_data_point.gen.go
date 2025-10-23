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
	// properties:
	AttributedLabel() objc.IObject /* cross-framework: AttributedString */
	SetAttributedLabel(value objc.IObject /* cross-framework: AttributedString */)
	Label() string /* primitive/slice/pointer. */
	SetLabel(value string /* primitive/slice/pointer. */)
	XValue() objc.IObject /* cross-framework: AXDataPointValue */
	SetXValue(value objc.IObject /* cross-framework: AXDataPointValue */)
	YValue() objc.IObject /* cross-framework: AXDataPointValue */
	SetYValue(value objc.IObject /* cross-framework: AXDataPointValue */)
	// methods:
}

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



// An attributed version of the label for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/attributedlabel
func (a_ AXDataPoint) AttributedLabel() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("attributedLabel"))
	return rv
}


// An attributed version of the label for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/attributedlabel
func (a_ AXDataPoint) SetAttributedLabel(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedLabel:"), value)
}


// The label for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/label
func (a_ AXDataPoint) Label() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("label"))
	return rv
}


// The label for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/label
func (a_ AXDataPoint) SetLabel(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), objc.String(value))
}


// The value of the x-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/xvalue
func (a_ AXDataPoint) XValue() objc.IObject /* cross-framework: AXDataPointValue */ {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("xValue"))
	return rv
}


// The value of the x-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/xvalue
func (a_ AXDataPoint) SetXValue(value objc.IObject /* cross-framework: AXDataPointValue */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setXValue:"), value)
}


// The value of the y-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/yvalue
func (a_ AXDataPoint) YValue() objc.IObject /* cross-framework: AXDataPointValue */ {
	rv := objc.Send[AXDataPointValue](a_.ID, objc.Sel("yValue"))
	return rv
}


// The value of the y-axis for the data point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axdatapoint/yvalue
func (a_ AXDataPoint) SetYValue(value objc.IObject /* cross-framework: AXDataPointValue */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setYValue:"), value)
}



