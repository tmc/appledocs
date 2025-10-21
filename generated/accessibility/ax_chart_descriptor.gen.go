// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// An object that contains all the semantic information about an accessible chart.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(attributedTitle unsafe.Pointer, summary string, xAxis objc.ID, yAxis unsafe.Pointer, additionalAxes unsafe.Pointer, series unsafe.Pointer) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), attributedTitle, objc.String(summary), xAxis, yAxis, additionalAxes, series)
	rv.Autorelease()
	return rv
}

// Creates a chart descriptor with the specified attributed title, summary, x-axis descriptor, y-axis descriptor, and array of data series.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorSeries(attributedTitle unsafe.Pointer, summary string, xAxis objc.ID, yAxis unsafe.Pointer, series unsafe.Pointer) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), attributedTitle, objc.String(summary), xAxis, yAxis, series)
	rv.Autorelease()
	return rv
}

// Creates a chart descriptor with the specified title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(title string, summary string, xAxis objc.ID, yAxis unsafe.Pointer, additionalAxes unsafe.Pointer, series unsafe.Pointer) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), objc.String(title), objc.String(summary), xAxis, yAxis, additionalAxes, series)
	rv.Autorelease()
	return rv
}

// Creates a chart descriptor with the specified title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorSeries(title string, summary string, xAxis objc.ID, yAxis unsafe.Pointer, series unsafe.Pointer) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), objc.String(title), objc.String(summary), xAxis, yAxis, series)
	rv.Autorelease()
	return rv
}


// The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/additionalAxes-9ldc0
func (a_ AXChartDescriptor) AdditionalAxes() []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("additionalAxes"))
	return rv
}


// SetAdditionalAxes sets the value of the additionalAxes property.
// The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.

//
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
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/attributedTitle
func (a_ AXChartDescriptor) AttributedTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributedTitle"))
	return rv
}


// SetAttributedTitle sets the value of the attributedTitle property.
// An attributed version of the chart title.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/attributedTitle
func (a_ AXChartDescriptor) SetAttributedTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedTitle:"), value)
}
// The direction of the content in the chart.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentDirection-swift.property
func (a_ AXChartDescriptor) ContentDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("contentDirection"))
	return rv
}


// SetContentDirection sets the value of the contentDirection property.
// The direction of the content in the chart.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentDirection-swift.property
func (a_ AXChartDescriptor) SetContentDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentDirection:"), value)
}
// The bounds of the view, in screen coordinates, for visually rendering data values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentFrame
func (a_ AXChartDescriptor) ContentFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](a_.ID, objc.Sel("contentFrame"))
	return rv
}


// SetContentFrame sets the value of the contentFrame property.
// The bounds of the view, in screen coordinates, for visually rendering data values.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentFrame
func (a_ AXChartDescriptor) SetContentFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentFrame:"), value)
}
// The descriptors for each data series in the chart.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/series
func (a_ AXChartDescriptor) Series() []AXDataSeriesDescriptor {
	rv := objc.Send[[]AXDataSeriesDescriptor](a_.ID, objc.Sel("series"))
	return rv
}


// SetSeries sets the value of the series property.
// The descriptors for each data series in the chart.

//
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
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/summary
func (a_ AXChartDescriptor) Summary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("summary"))
	return rv
}


// SetSummary sets the value of the summary property.
// A description of the key takeaways or features of the chart.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/summary
func (a_ AXChartDescriptor) SetSummary(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSummary:"), value)
}
// The title of the chart.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/title
func (a_ AXChartDescriptor) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the chart.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/title
func (a_ AXChartDescriptor) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTitle:"), value)
}
// The axis descriptor for the chart’s x-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/xAxis-6dnxd
func (a_ AXChartDescriptor) XAxis() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("xAxis"))
	return rv
}


// SetXAxis sets the value of the xAxis property.
// The axis descriptor for the chart’s x-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/xAxis-6dnxd
func (a_ AXChartDescriptor) SetXAxis(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setXAxis:"), value)
}
// The axis descriptor for the chart’s y-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/yAxis
func (a_ AXChartDescriptor) YAxis() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("yAxis"))
	return rv
}


// SetYAxis sets the value of the yAxis property.
// The axis descriptor for the chart’s y-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/yAxis
func (a_ AXChartDescriptor) SetYAxis(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setYAxis:"), value)
}

