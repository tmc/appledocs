// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXChartDescriptor] class.
var (
	AXChartDescriptorClass     _AXChartDescriptorClass
	AXChartDescriptorClassOnce sync.Once
)

func getAXChartDescriptorClass() _AXChartDescriptorClass {
	AXChartDescriptorClassOnce.Do(func() {
		AXChartDescriptorClass = _AXChartDescriptorClass{objc.GetClass("AXChartDescriptor")}
	})
	return AXChartDescriptorClass
}

type _AXChartDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AXChartDescriptor] class.
type IAXChartDescriptor interface {
	objectivec.IObject
	AdditionalAxes() []objc.ID
	SetAdditionalAxes(value []objc.ID)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	ContentDirection() AXChartDescriptorContentDirection
	SetContentDirection(value IAXChartDescriptorContentDirection)
	ContentFrame() coregraphics.CGRect
	SetContentFrame(value coregraphics.CGRect)
	Series() []AXDataSeriesDescriptor
	SetSeries(value []AXDataSeriesDescriptor)
	Summary() string
	SetSummary(value string)
	Title() string
	SetTitle(value string)
	XAxis() objc.ID
	SetXAxis(value objc.ID)
	YAxis() AXNumericDataAxisDescriptor
	SetYAxis(value IAXNumericDataAxisDescriptor)
}

// An object that contains all the semantic information about an accessible chart.


// An object that contains all the semantic information about an accessible chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor
type AXChartDescriptor struct {
	objectivec.Object
}

// AXChartDescriptorFrom constructs a [AXChartDescriptor] from an unsafe.Pointer.
//
// An object that contains all the semantic information about an accessible chart.
func AXChartDescriptorFrom(ptr unsafe.Pointer) AXChartDescriptor {
	return AXChartDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXChartDescriptorClass) Alloc() AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXChartDescriptorClass) New() AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXChartDescriptor) Init() AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXChartDescriptor) Autorelease() AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXChartDescriptor creates a new AXChartDescriptor instance.
func NewAXChartDescriptor() AXChartDescriptor {
	return getAXChartDescriptorClass().New()
}



// Creates a chart descriptor with the specified attributed title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(attributedTitle foundation.IAttributedString, summary string, xAxis objectivec.IObject, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objc.ID, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), attributedTitle, objc.String(summary), xAxis, yAxis, additionalAxes, series)
	rv.Autorelease()
	return rv
}


// Creates a chart descriptor with the specified attributed title, summary, x-axis descriptor, y-axis descriptor, and array of data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorSeries(attributedTitle foundation.IAttributedString, summary string, xAxis objectivec.IObject, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), attributedTitle, objc.String(summary), xAxis, yAxis, series)
	rv.Autorelease()
	return rv
}


// Creates a chart descriptor with the specified title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(title string, summary string, xAxis objectivec.IObject, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objc.ID, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), objc.String(title), objc.String(summary), xAxis, yAxis, additionalAxes, series)
	rv.Autorelease()
	return rv
}


// Creates a chart descriptor with the specified title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorSeries(title string, summary string, xAxis objectivec.IObject, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), objc.String(title), objc.String(summary), xAxis, yAxis, series)
	rv.Autorelease()
	return rv
}



// The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/additionalAxes-9ldc0
func (a_ AXChartDescriptor) AdditionalAxes() []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("additionalAxes"))
	return rv
}


// The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/additionalAxes-9ldc0
func (a_ AXChartDescriptor) SetAdditionalAxes(value []objc.ID) {
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
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdditionalAxes:"), nsArray)
}


// An attributed version of the chart title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/attributedTitle
func (a_ AXChartDescriptor) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedTitle"))
	return rv
}


// An attributed version of the chart title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/attributedTitle
func (a_ AXChartDescriptor) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedTitle:"), value)
}


// The direction of the content in the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentDirection-swift.property
func (a_ AXChartDescriptor) ContentDirection() AXChartDescriptorContentDirection {
	rv := objc.Send[AXChartDescriptorContentDirection](a_.ID, objc.Sel("contentDirection"))
	return rv
}


// The direction of the content in the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentDirection-swift.property
func (a_ AXChartDescriptor) SetContentDirection(value IAXChartDescriptorContentDirection) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentDirection:"), value)
}


// The bounds of the view, in screen coordinates, for visually rendering data values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentFrame
func (a_ AXChartDescriptor) ContentFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](a_.ID, objc.Sel("contentFrame"))
	return rv
}


// The bounds of the view, in screen coordinates, for visually rendering data values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentFrame
func (a_ AXChartDescriptor) SetContentFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentFrame:"), value)
}


// The descriptors for each data series in the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/series
func (a_ AXChartDescriptor) Series() []AXDataSeriesDescriptor {
	rv := objc.Send[[]AXDataSeriesDescriptor](a_.ID, objc.Sel("series"))
	return rv
}


// The descriptors for each data series in the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/series
func (a_ AXChartDescriptor) SetSeries(value []AXDataSeriesDescriptor) {
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
	objc.Send[objc.ID](a_.ID, objc.Sel("setSeries:"), nsArray)
}


// A description of the key takeaways or features of the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/summary
func (a_ AXChartDescriptor) Summary() string {
	rv := objc.Send[string](a_.ID, objc.Sel("summary"))
	return rv
}


// A description of the key takeaways or features of the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/summary
func (a_ AXChartDescriptor) SetSummary(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSummary:"), objc.String(value))
}


// The title of the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/title
func (a_ AXChartDescriptor) Title() string {
	rv := objc.Send[string](a_.ID, objc.Sel("title"))
	return rv
}


// The title of the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/title
func (a_ AXChartDescriptor) SetTitle(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The axis descriptor for the chart’s x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/xAxis-6dnxd
func (a_ AXChartDescriptor) XAxis() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("xAxis"))
	return rv
}


// The axis descriptor for the chart’s x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/xAxis-6dnxd
func (a_ AXChartDescriptor) SetXAxis(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setXAxis:"), value)
}


// The axis descriptor for the chart’s y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/yAxis
func (a_ AXChartDescriptor) YAxis() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](a_.ID, objc.Sel("yAxis"))
	return rv
}


// The axis descriptor for the chart’s y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/yAxis
func (a_ AXChartDescriptor) SetYAxis(value IAXNumericDataAxisDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setYAxis:"), value)
}


