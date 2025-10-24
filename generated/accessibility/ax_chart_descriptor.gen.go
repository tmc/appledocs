// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXChartDescriptor */


/* debug [class_header]: Header for AXChartDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXChartDescriptor */
// An interface definition for the [AXChartDescriptor] class.
type IAXChartDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXChartDescriptor */
	// properties:
	AdditionalAxes() []objc.ID
	SetAdditionalAxes(value []objc.ID)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	ContentDirection() AXChartDescriptorContentDirection
	SetContentDirection(value AXChartDescriptorContentDirection)
	ContentFrame() corefoundation.CGRect
	SetContentFrame(value corefoundation.CGRect)
	Series() []AXDataSeriesDescriptor
	SetSeries(value []AXDataSeriesDescriptor)
	Summary() objc.IObject /* cross-framework: NSString */
	SetSummary(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	XAxis() unsafe.Pointer
	SetXAxis(value unsafe.Pointer)
	YAxis() IAXNumericDataAxisDescriptor
	SetYAxis(value IAXNumericDataAxisDescriptor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXChartDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXChartDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AXChartDescriptorClass) Alloc() AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXChartDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXChartDescriptor */

// Creates a chart descriptor with the specified attributed title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(attributedTitle foundation.AttributedString, summary objc.IObject /* cross-framework: NSString */, xAxis unsafe.Pointer, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objc.ID, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), attributedTitle, summary, xAxis, yAxis, additionalAxes, series)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries */


// Creates a chart descriptor with the specified attributed title, summary, x-axis descriptor, y-axis descriptor, and array of data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorSeries(attributedTitle foundation.AttributedString, summary objc.IObject /* cross-framework: NSString */, xAxis unsafe.Pointer, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), attributedTitle, summary, xAxis, yAxis, series)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorSeries */


// Creates a chart descriptor with the specified title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(title objc.IObject /* cross-framework: NSString */, summary objc.IObject /* cross-framework: NSString */, xAxis unsafe.Pointer, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objc.ID, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), title, summary, xAxis, yAxis, additionalAxes, series)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries */


// Creates a chart descriptor with the specified title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorSeries(title objc.IObject /* cross-framework: NSString */, summary objc.IObject /* cross-framework: NSString */, xAxis unsafe.Pointer, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[AXChartDescriptor](instance.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), title, summary, xAxis, yAxis, series)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorSeries */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXChartDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXChartDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXChartDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXChartDescriptor */

// The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/additionalAxes-9ldc0
func (a_ AXChartDescriptor) AdditionalAxes() []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("additionalAxes"))
	return rv
}/* debug [instance_properties/getter]: additionalAxes */


// The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/additionalAxes-9ldc0
func (a_ AXChartDescriptor) SetAdditionalAxes(value []objc.ID) {
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
}/* debug [instance_properties/setter]: additionalAxes */


// An attributed version of the chart title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/attributedTitle
func (a_ AXChartDescriptor) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedTitle"))
	return rv
}/* debug [instance_properties/getter]: attributedTitle */


// An attributed version of the chart title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/attributedTitle
func (a_ AXChartDescriptor) SetAttributedTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedTitle:"), value)
}/* debug [instance_properties/setter]: attributedTitle */


// The direction of the content in the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentDirection-swift.property
func (a_ AXChartDescriptor) ContentDirection() AXChartDescriptorContentDirection {
	rv := objc.Send[AXChartDescriptorContentDirection](a_.ID, objc.Sel("contentDirection"))
	return rv
}/* debug [instance_properties/getter]: contentDirection */


// The direction of the content in the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentDirection-swift.property
func (a_ AXChartDescriptor) SetContentDirection(value AXChartDescriptorContentDirection) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentDirection:"), value)
}/* debug [instance_properties/setter]: contentDirection */


// The bounds of the view, in screen coordinates, for visually rendering data values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentFrame
func (a_ AXChartDescriptor) ContentFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a_.ID, objc.Sel("contentFrame"))
	return rv
}/* debug [instance_properties/getter]: contentFrame */


// The bounds of the view, in screen coordinates, for visually rendering data values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentFrame
func (a_ AXChartDescriptor) SetContentFrame(value corefoundation.CGRect) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentFrame:"), value)
}/* debug [instance_properties/setter]: contentFrame */


// The descriptors for each data series in the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/series
func (a_ AXChartDescriptor) Series() []AXDataSeriesDescriptor {
	rv := objc.Send[[]AXDataSeriesDescriptor](a_.ID, objc.Sel("series"))
	return rv
}/* debug [instance_properties/getter]: series */


// The descriptors for each data series in the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/series
func (a_ AXChartDescriptor) SetSeries(value []AXDataSeriesDescriptor) {
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
}/* debug [instance_properties/setter]: series */


// A description of the key takeaways or features of the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/summary
func (a_ AXChartDescriptor) Summary() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("summary"))
	return rv
}/* debug [instance_properties/getter]: summary */


// A description of the key takeaways or features of the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/summary
func (a_ AXChartDescriptor) SetSummary(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSummary:"), value)
}/* debug [instance_properties/setter]: summary */


// The title of the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/title
func (a_ AXChartDescriptor) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title of the chart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/title
func (a_ AXChartDescriptor) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The axis descriptor for the chart’s x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/xAxis-6dnxd
func (a_ AXChartDescriptor) XAxis() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("xAxis"))
	return rv
}/* debug [instance_properties/getter]: xAxis */


// The axis descriptor for the chart’s x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/xAxis-6dnxd
func (a_ AXChartDescriptor) SetXAxis(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setXAxis:"), value)
}/* debug [instance_properties/setter]: xAxis */


// The axis descriptor for the chart’s y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/yAxis
func (a_ AXChartDescriptor) YAxis() IAXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](a_.ID, objc.Sel("yAxis"))
	return rv
}/* debug [instance_properties/getter]: yAxis */


// The axis descriptor for the chart’s y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/yAxis
func (a_ AXChartDescriptor) SetYAxis(value IAXNumericDataAxisDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setYAxis:"), value)
}/* debug [instance_properties/setter]: yAxis */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXChartDescriptor */


