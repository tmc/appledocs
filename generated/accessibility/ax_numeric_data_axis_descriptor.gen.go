// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AXNumericDataAxisDescriptor] class.
type IAXNumericDataAxisDescriptor interface {
	objectivec.IObject
}

// An object that represents an axis of numerical data.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AXNumericDataAxisDescriptorClass) Alloc() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a numeric data axis with the specified attributed title, lower bound value, upper bound value, gridline positions, and value description provider block.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/initWithAttributedTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:
func NewAXNumericDataAxisDescriptorWithAttributedTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider(attributedTitle foundation.IAttributedString, lowerbound unsafe.Pointer, upperBound unsafe.Pointer, gridlinePositions []foundation.INumber, valueDescriptionProvider unsafe.Pointer) AXNumericDataAxisDescriptor {
	instance := getAXNumericDataAxisDescriptorClass().Alloc()
	rv := objc.Send[AXNumericDataAxisDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:"), attributedTitle, lowerbound, upperBound, gridlinePositions, valueDescriptionProvider)
	rv.Autorelease()
	return rv
}



// Creates a numeric data axis with the specified title, lower bound value, upper bound value, gridline positions, and value description provider block.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/initWithTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:
func NewAXNumericDataAxisDescriptorWithTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider(title appkit.string, lowerbound unsafe.Pointer, upperBound unsafe.Pointer, gridlinePositions []foundation.INumber, valueDescriptionProvider unsafe.Pointer) AXNumericDataAxisDescriptor {
	instance := getAXNumericDataAxisDescriptorClass().Alloc()
	rv := objc.Send[AXNumericDataAxisDescriptor](instance.ID, objc.Sel("initWithTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:"), title, lowerbound, upperBound, gridlinePositions, valueDescriptionProvider)
	rv.Autorelease()
	return rv
}


// The positions of the gridlines along the axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/gridlinePositions-9z10e
func (a_ AXNumericDataAxisDescriptor) GridlinePositions() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("gridlinePositions"))
	return rv
}


// SetGridlinePositions sets the value of the gridlinePositions property.
// The positions of the gridlines along the axis.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/gridlinePositions-9z10e
func (a_ AXNumericDataAxisDescriptor) SetGridlinePositions(value []foundation.INumber) {
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
	objc.Send[objc.ID](a_.ID, objc.Sel("setGridlinePositions:"), nsArray)
}

// The minimum displayable value for the axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/lowerBound
func (a_ AXNumericDataAxisDescriptor) LowerBound() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("lowerBound"))
	return rv
}


// SetLowerBound sets the value of the lowerBound property.
// The minimum displayable value for the axis.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/lowerBound
func (a_ AXNumericDataAxisDescriptor) SetLowerBound(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLowerBound:"), value)
}

// The scale for the axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/scaleType-swift.property
func (a_ AXNumericDataAxisDescriptor) ScaleType() AXNumericDataAxisDescriptorScale {
	rv := objc.Send[AXNumericDataAxisDescriptorScale](a_.ID, objc.Sel("scaleType"))
	return rv
}


// SetScaleType sets the value of the scaleType property.
// The scale for the axis.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/scaleType-swift.property
func (a_ AXNumericDataAxisDescriptor) SetScaleType(value IAXNumericDataAxisDescriptorScale) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setScaleType:"), value)
}

// The maximum displayable value for the axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/upperBound
func (a_ AXNumericDataAxisDescriptor) UpperBound() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("upperBound"))
	return rv
}


// SetUpperBound sets the value of the upperBound property.
// The maximum displayable value for the axis.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/upperBound
func (a_ AXNumericDataAxisDescriptor) SetUpperBound(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUpperBound:"), value)
}

// A description to speak for a particular data value on the axis.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/valueDescriptionProvider
func (a_ AXNumericDataAxisDescriptor) ValueDescriptionProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("valueDescriptionProvider"))
	return rv
}


// SetValueDescriptionProvider sets the value of the valueDescriptionProvider property.
// A description to speak for a particular data value on the axis.

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/valueDescriptionProvider
func (a_ AXNumericDataAxisDescriptor) SetValueDescriptionProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValueDescriptionProvider:"), value)
}

// A range that defines the minimum and maximum displayable values for the axis.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axnumericdataaxisdescriptor/range
func (a_ AXNumericDataAxisDescriptor) Range() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("range"))
	return rv
}


// SetRange sets the value of the range property.
// A range that defines the minimum and maximum displayable values for the axis.

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axnumericdataaxisdescriptor/range
func (a_ AXNumericDataAxisDescriptor) SetRange(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRange:"), value)
}


