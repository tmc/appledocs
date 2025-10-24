// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXNumericDataAxisDescriptor */


/* debug [class_header]: Header for AXNumericDataAxisDescriptor */
// The class instance for the [AXNumericDataAxisDescriptor] class.
var (
	AXNumericDataAxisDescriptorClass     _AXNumericDataAxisDescriptorClass
	AXNumericDataAxisDescriptorClassOnce sync.Once
)

func getAXNumericDataAxisDescriptorClass() _AXNumericDataAxisDescriptorClass {
	AXNumericDataAxisDescriptorClassOnce.Do(func() {
		AXNumericDataAxisDescriptorClass = _AXNumericDataAxisDescriptorClass{objc.GetClass("AXNumericDataAxisDescriptor")}
	})
	return AXNumericDataAxisDescriptorClass
}

type _AXNumericDataAxisDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXNumericDataAxisDescriptor */
// An interface definition for the [AXNumericDataAxisDescriptor] class.
type IAXNumericDataAxisDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXNumericDataAxisDescriptor */
	// properties:
	GridlinePositions() []foundation.Number
	SetGridlinePositions(value []foundation.Number)
	LowerBound() float64
	SetLowerBound(value float64)
	ScaleType() AXNumericDataAxisDescriptorScale
	SetScaleType(value AXNumericDataAxisDescriptorScale)
	UpperBound() float64
	SetUpperBound(value float64)
	ValueDescriptionProvider() unsafe.Pointer
	SetValueDescriptionProvider(value unsafe.Pointer)
	Range() float64
	SetRange(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXNumericDataAxisDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXNumericDataAxisDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AXNumericDataAxisDescriptorClass) Alloc() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXNumericDataAxisDescriptorClass) New() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXNumericDataAxisDescriptor) Init() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXNumericDataAxisDescriptor) Autorelease() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXNumericDataAxisDescriptor creates a new AXNumericDataAxisDescriptor instance.
func NewAXNumericDataAxisDescriptor() AXNumericDataAxisDescriptor {
	return getAXNumericDataAxisDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXNumericDataAxisDescriptor */
// An object that represents an axis of numerical data.


// An object that represents an axis of numerical data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor
type AXNumericDataAxisDescriptor struct {
	objectivec.Object
}

// AXNumericDataAxisDescriptorFrom constructs a [AXNumericDataAxisDescriptor] from an unsafe.Pointer.
//
// An object that represents an axis of numerical data.
func AXNumericDataAxisDescriptorFrom(ptr unsafe.Pointer) AXNumericDataAxisDescriptor {
	return AXNumericDataAxisDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXNumericDataAxisDescriptor */

// Creates a numeric data axis with the specified attributed title, lower bound value, upper bound value, gridline positions, and value description provider block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/initWithAttributedTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:
func NewAXNumericDataAxisDescriptorWithAttributedTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider(attributedTitle foundation.AttributedString, lowerbound float64, upperBound float64, gridlinePositions []foundation.Number, valueDescriptionProvider unsafe.Pointer) AXNumericDataAxisDescriptor {
	instance := getAXNumericDataAxisDescriptorClass().Alloc()
	rv := objc.Send[AXNumericDataAxisDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:"), attributedTitle, lowerbound, upperBound, gridlinePositions, valueDescriptionProvider)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXNumericDataAxisDescriptorWithAttributedTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider */


// Creates a numeric data axis with the specified title, lower bound value, upper bound value, gridline positions, and value description provider block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/initWithTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:
func NewAXNumericDataAxisDescriptorWithTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider(title objc.IObject /* cross-framework: NSString */, lowerbound float64, upperBound float64, gridlinePositions []foundation.Number, valueDescriptionProvider unsafe.Pointer) AXNumericDataAxisDescriptor {
	instance := getAXNumericDataAxisDescriptorClass().Alloc()
	rv := objc.Send[AXNumericDataAxisDescriptor](instance.ID, objc.Sel("initWithTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:"), title, lowerbound, upperBound, gridlinePositions, valueDescriptionProvider)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXNumericDataAxisDescriptorWithTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXNumericDataAxisDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXNumericDataAxisDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXNumericDataAxisDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXNumericDataAxisDescriptor */

// The positions of the gridlines along the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/gridlinePositions-9z10e
func (a_ AXNumericDataAxisDescriptor) GridlinePositions() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("gridlinePositions"))
	return rv
}/* debug [instance_properties/getter]: gridlinePositions */


// The positions of the gridlines along the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/gridlinePositions-9z10e
func (a_ AXNumericDataAxisDescriptor) SetGridlinePositions(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setGridlinePositions:"), nsArray)
}/* debug [instance_properties/setter]: gridlinePositions */


// The minimum displayable value for the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/lowerBound
func (a_ AXNumericDataAxisDescriptor) LowerBound() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("lowerBound"))
	return rv
}/* debug [instance_properties/getter]: lowerBound */


// The minimum displayable value for the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/lowerBound
func (a_ AXNumericDataAxisDescriptor) SetLowerBound(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLowerBound:"), value)
}/* debug [instance_properties/setter]: lowerBound */


// The scale for the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/scaleType-swift.property
func (a_ AXNumericDataAxisDescriptor) ScaleType() AXNumericDataAxisDescriptorScale {
	rv := objc.Send[AXNumericDataAxisDescriptorScale](a_.ID, objc.Sel("scaleType"))
	return rv
}/* debug [instance_properties/getter]: scaleType */


// The scale for the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/scaleType-swift.property
func (a_ AXNumericDataAxisDescriptor) SetScaleType(value AXNumericDataAxisDescriptorScale) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setScaleType:"), value)
}/* debug [instance_properties/setter]: scaleType */


// The maximum displayable value for the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/upperBound
func (a_ AXNumericDataAxisDescriptor) UpperBound() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("upperBound"))
	return rv
}/* debug [instance_properties/getter]: upperBound */


// The maximum displayable value for the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/upperBound
func (a_ AXNumericDataAxisDescriptor) SetUpperBound(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUpperBound:"), value)
}/* debug [instance_properties/setter]: upperBound */


// A description to speak for a particular data value on the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/valueDescriptionProvider
func (a_ AXNumericDataAxisDescriptor) ValueDescriptionProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("valueDescriptionProvider"))
	return rv
}/* debug [instance_properties/getter]: valueDescriptionProvider */


// A description to speak for a particular data value on the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/valueDescriptionProvider
func (a_ AXNumericDataAxisDescriptor) SetValueDescriptionProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValueDescriptionProvider:"), value)
}/* debug [instance_properties/setter]: valueDescriptionProvider */


// A range that defines the minimum and maximum displayable values for the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axnumericdataaxisdescriptor/range
func (a_ AXNumericDataAxisDescriptor) Range() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("range"))
	return rv
}/* debug [instance_properties/getter]: range */


// A range that defines the minimum and maximum displayable values for the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axnumericdataaxisdescriptor/range
func (a_ AXNumericDataAxisDescriptor) SetRange(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRange:"), value)
}/* debug [instance_properties/setter]: range */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXNumericDataAxisDescriptor */


