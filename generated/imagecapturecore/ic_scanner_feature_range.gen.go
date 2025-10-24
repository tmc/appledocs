// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ICScannerFeatureRange */


/* debug [class_header]: Header for ICScannerFeatureRange */
// The class instance for the [ICScannerFeatureRange] class.
var (
	ICScannerFeatureRangeClass     _ICScannerFeatureRangeClass
	ICScannerFeatureRangeClassOnce sync.Once
)

func getICScannerFeatureRangeClass() _ICScannerFeatureRangeClass {
	ICScannerFeatureRangeClassOnce.Do(func() {
		ICScannerFeatureRangeClass = _ICScannerFeatureRangeClass{objc.GetClass("ICScannerFeatureRange")}
	})
	return ICScannerFeatureRangeClass
}

type _ICScannerFeatureRangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFeatureRange */
// An interface definition for the [ICScannerFeatureRange] class.
type IICScannerFeatureRange interface {
	IICScannerFeature
	
/* debug [class_interface_properties]: Properties for ICScannerFeatureRange */
	// properties:
	StepSize() Float get /* not a class type */
	SetStepSize(value Float get /* not a class type */)
	MinValue() Float get /* not a class type */
	SetMinValue(value Float get /* not a class type */)
	CurrentValue() Float get set /* not a class type */
	SetCurrentValue(value Float get set /* not a class type */)
	MaxValue() Float get /* not a class type */
	SetMaxValue(value Float get /* not a class type */)
	DefaultValue() Float get /* not a class type */
	SetDefaultValue(value Float get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFeatureRange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFeatureRange */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFeatureRangeClass) Alloc() ICScannerFeatureRange {
	rv := objc.Send[ICScannerFeatureRange](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFeatureRangeClass) New() ICScannerFeatureRange {
	rv := objc.Send[ICScannerFeatureRange](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFeatureRange) Init() ICScannerFeatureRange {
	rv := objc.Send[ICScannerFeatureRange](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFeatureRange) Autorelease() ICScannerFeatureRange {
	rv := objc.Send[ICScannerFeatureRange](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureRange creates a new ICScannerFeatureRange instance.
func NewICScannerFeatureRange() ICScannerFeatureRange {
	return getICScannerFeatureRangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFeatureRange */
// A feature with a value that lies within a range.


// A feature with a value that lies within a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureRange
type ICScannerFeatureRange struct {
	ICScannerFeature
}

// ICScannerFeatureRangeFrom constructs a [ICScannerFeatureRange] from an unsafe.Pointer.
//
// A feature with a value that lies within a range.
func ICScannerFeatureRangeFrom(ptr unsafe.Pointer) ICScannerFeatureRange {
	return ICScannerFeatureRange{
		ICScannerFeature: ICScannerFeatureFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFeatureRange *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFeatureRange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFeatureRange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFeatureRange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFeatureRange */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1507697-stepsize
func (i_ ICScannerFeatureRange) StepSize() Float get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("stepSize"))
	return rv
}/* debug [instance_properties/getter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1507697-stepsize
func (i_ ICScannerFeatureRange) SetStepSize(value Float get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStepSize:"), value)
}/* debug [instance_properties/setter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1507774-minvalue
func (i_ ICScannerFeatureRange) MinValue() Float get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1507774-minvalue
func (i_ ICScannerFeatureRange) SetMinValue(value Float get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1507778-currentvalue
func (i_ ICScannerFeatureRange) CurrentValue() Float get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("currentValue"))
	return rv
}/* debug [instance_properties/getter]: currentValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1507778-currentvalue
func (i_ ICScannerFeatureRange) SetCurrentValue(value Float get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrentValue:"), value)
}/* debug [instance_properties/setter]: currentValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1507839-maxvalue
func (i_ ICScannerFeatureRange) MaxValue() Float get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1507839-maxvalue
func (i_ ICScannerFeatureRange) SetMaxValue(value Float get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1508018-defaultvalue
func (i_ ICScannerFeatureRange) DefaultValue() Float get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("defaultValue"))
	return rv
}/* debug [instance_properties/getter]: defaultValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturerange/1508018-defaultvalue
func (i_ ICScannerFeatureRange) SetDefaultValue(value Float get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDefaultValue:"), value)
}/* debug [instance_properties/setter]: defaultValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFeatureRange */



