// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXDataSeriesDescriptor */


/* debug [class_header]: Header for AXDataSeriesDescriptor */
// The class instance for the [AXDataSeriesDescriptor] class.
var (
	AXDataSeriesDescriptorClass     _AXDataSeriesDescriptorClass
	AXDataSeriesDescriptorClassOnce sync.Once
)

func getAXDataSeriesDescriptorClass() _AXDataSeriesDescriptorClass {
	AXDataSeriesDescriptorClassOnce.Do(func() {
		AXDataSeriesDescriptorClass = _AXDataSeriesDescriptorClass{objc.GetClass("AXDataSeriesDescriptor")}
	})
	return AXDataSeriesDescriptorClass
}

type _AXDataSeriesDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXDataSeriesDescriptor */
// An interface definition for the [AXDataSeriesDescriptor] class.
type IAXDataSeriesDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXDataSeriesDescriptor */
	// properties:
	AttributedName() foundation.AttributedString
	SetAttributedName(value foundation.AttributedString)
	DataPoints() []AXDataPoint
	SetDataPoints(value []AXDataPoint)
	IsContinuous() bool
	SetIsContinuous(value bool)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXDataSeriesDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXDataSeriesDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AXDataSeriesDescriptorClass) Alloc() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXDataSeriesDescriptorClass) New() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXDataSeriesDescriptor) Init() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXDataSeriesDescriptor) Autorelease() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXDataSeriesDescriptor creates a new AXDataSeriesDescriptor instance.
func NewAXDataSeriesDescriptor() AXDataSeriesDescriptor {
	return getAXDataSeriesDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXDataSeriesDescriptor */
// An object that represents a series of data points.


// An object that represents a series of data points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor
type AXDataSeriesDescriptor struct {
	objectivec.Object
}

// AXDataSeriesDescriptorFrom constructs a [AXDataSeriesDescriptor] from an unsafe.Pointer.
//
// An object that represents a series of data points.
func AXDataSeriesDescriptorFrom(ptr unsafe.Pointer) AXDataSeriesDescriptor {
	return AXDataSeriesDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXDataSeriesDescriptor */

// Creates a data series with the specified attributed name, a Boolean value that indicates whether the series is continuous, and data points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/init(attributedName:isContinuous:dataPoints:)
func NewAXDataSeriesDescriptorWithAttributedNameIsContinuousDataPoints(attributedName foundation.AttributedString, isContinuous bool, dataPoints []AXDataPoint) AXDataSeriesDescriptor {
	instance := getAXDataSeriesDescriptorClass().Alloc()
	rv := objc.Send[AXDataSeriesDescriptor](instance.ID, objc.Sel("initWithAttributedName:isContinuous:dataPoints:"), attributedName, isContinuous, dataPoints)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXDataSeriesDescriptorWithAttributedNameIsContinuousDataPoints */


// Creates a data series with the specified name, a Boolean value that indicates whether the series is continuous, and data points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/init(name:isContinuous:dataPoints:)
func NewAXDataSeriesDescriptorWithNameIsContinuousDataPoints(name objc.IObject /* cross-framework: NSString */, isContinuous bool, dataPoints []AXDataPoint) AXDataSeriesDescriptor {
	instance := getAXDataSeriesDescriptorClass().Alloc()
	rv := objc.Send[AXDataSeriesDescriptor](instance.ID, objc.Sel("initWithName:isContinuous:dataPoints:"), name, isContinuous, dataPoints)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXDataSeriesDescriptorWithNameIsContinuousDataPoints */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXDataSeriesDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXDataSeriesDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXDataSeriesDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXDataSeriesDescriptor */

// An attributed version of the data series name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/attributedName
func (a_ AXDataSeriesDescriptor) AttributedName() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedName"))
	return rv
}/* debug [instance_properties/getter]: attributedName */


// An attributed version of the data series name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/attributedName
func (a_ AXDataSeriesDescriptor) SetAttributedName(value foundation.AttributedString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedName:"), value)
}/* debug [instance_properties/setter]: attributedName */


// The data points that the series contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/dataPoints
func (a_ AXDataSeriesDescriptor) DataPoints() []AXDataPoint {
	rv := objc.Send[[]AXDataPoint](a_.ID, objc.Sel("dataPoints"))
	return rv
}/* debug [instance_properties/getter]: dataPoints */


// The data points that the series contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/dataPoints
func (a_ AXDataSeriesDescriptor) SetDataPoints(value []AXDataPoint) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataPoints:"), nsArray)
}/* debug [instance_properties/setter]: dataPoints */


// A Boolean value that determines whether the data series is continuous.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/isContinuous
func (a_ AXDataSeriesDescriptor) IsContinuous() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isContinuous"))
	return rv
}/* debug [instance_properties/getter]: isContinuous */


// A Boolean value that determines whether the data series is continuous.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/isContinuous
func (a_ AXDataSeriesDescriptor) SetIsContinuous(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsContinuous:"), value)
}/* debug [instance_properties/setter]: isContinuous */


// The name of the data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/name
func (a_ AXDataSeriesDescriptor) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the data series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/name
func (a_ AXDataSeriesDescriptor) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXDataSeriesDescriptor */


