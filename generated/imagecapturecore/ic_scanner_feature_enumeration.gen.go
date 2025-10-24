// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ICScannerFeatureEnumeration */


/* debug [class_header]: Header for ICScannerFeatureEnumeration */
// The class instance for the [ICScannerFeatureEnumeration] class.
var (
	ICScannerFeatureEnumerationClass     _ICScannerFeatureEnumerationClass
	ICScannerFeatureEnumerationClassOnce sync.Once
)

func getICScannerFeatureEnumerationClass() _ICScannerFeatureEnumerationClass {
	ICScannerFeatureEnumerationClassOnce.Do(func() {
		ICScannerFeatureEnumerationClass = _ICScannerFeatureEnumerationClass{objc.GetClass("ICScannerFeatureEnumeration")}
	})
	return ICScannerFeatureEnumerationClass
}

type _ICScannerFeatureEnumerationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFeatureEnumeration */
// An interface definition for the [ICScannerFeatureEnumeration] class.
type IICScannerFeatureEnumeration interface {
	IICScannerFeature
	
/* debug [class_interface_properties]: Properties for ICScannerFeatureEnumeration */
	// properties:
	MenuItemLabels() unsafe.Pointer
	SetMenuItemLabels(value unsafe.Pointer)
	DefaultValue() unsafe.Pointer
	SetDefaultValue(value unsafe.Pointer)
	MenuItemLabelsTooltips() unsafe.Pointer
	SetMenuItemLabelsTooltips(value unsafe.Pointer)
	Values() objc.IObject /* cross-framework: NSNumber */
	SetValues(value objc.IObject /* cross-framework: NSNumber */)
	CurrentValue() objc.ID
	SetCurrentValue(value objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFeatureEnumeration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFeatureEnumeration */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFeatureEnumerationClass) Alloc() ICScannerFeatureEnumeration {
	rv := objc.Send[ICScannerFeatureEnumeration](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFeatureEnumerationClass) New() ICScannerFeatureEnumeration {
	rv := objc.Send[ICScannerFeatureEnumeration](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFeatureEnumeration) Init() ICScannerFeatureEnumeration {
	rv := objc.Send[ICScannerFeatureEnumeration](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFeatureEnumeration) Autorelease() ICScannerFeatureEnumeration {
	rv := objc.Send[ICScannerFeatureEnumeration](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureEnumeration creates a new ICScannerFeatureEnumeration instance.
func NewICScannerFeatureEnumeration() ICScannerFeatureEnumeration {
	return getICScannerFeatureEnumerationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFeatureEnumeration */
// A feature that can have one of several discrete values, strings or numbers.


// A feature that can have one of several discrete values, strings or numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration
type ICScannerFeatureEnumeration struct {
	ICScannerFeature
}

// ICScannerFeatureEnumerationFrom constructs a [ICScannerFeatureEnumeration] from an unsafe.Pointer.
//
// A feature that can have one of several discrete values, strings or numbers.
func ICScannerFeatureEnumerationFrom(ptr unsafe.Pointer) ICScannerFeatureEnumeration {
	return ICScannerFeatureEnumeration{
		ICScannerFeature: ICScannerFeatureFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFeatureEnumeration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFeatureEnumeration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFeatureEnumeration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFeatureEnumeration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFeatureEnumeration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureenumeration/1507712-menuitemlabels
func (i_ ICScannerFeatureEnumeration) MenuItemLabels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("menuItemLabels"))
	return rv
}/* debug [instance_properties/getter]: menuItemLabels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureenumeration/1507712-menuitemlabels
func (i_ ICScannerFeatureEnumeration) SetMenuItemLabels(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMenuItemLabels:"), value)
}/* debug [instance_properties/setter]: menuItemLabels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureenumeration/1507821-defaultvalue
func (i_ ICScannerFeatureEnumeration) DefaultValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("defaultValue"))
	return rv
}/* debug [instance_properties/getter]: defaultValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureenumeration/1507821-defaultvalue
func (i_ ICScannerFeatureEnumeration) SetDefaultValue(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDefaultValue:"), value)
}/* debug [instance_properties/setter]: defaultValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureenumeration/1507843-menuitemlabelstooltips
func (i_ ICScannerFeatureEnumeration) MenuItemLabelsTooltips() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("menuItemLabelsTooltips"))
	return rv
}/* debug [instance_properties/getter]: menuItemLabelsTooltips */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureenumeration/1507843-menuitemlabelstooltips
func (i_ ICScannerFeatureEnumeration) SetMenuItemLabelsTooltips(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMenuItemLabelsTooltips:"), value)
}/* debug [instance_properties/setter]: menuItemLabelsTooltips */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureenumeration/1508144-values
func (i_ ICScannerFeatureEnumeration) Values() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](i_.ID, objc.Sel("values"))
	return rv
}/* debug [instance_properties/getter]: values */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureenumeration/1508144-values
func (i_ ICScannerFeatureEnumeration) SetValues(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setValues:"), value)
}/* debug [instance_properties/setter]: values */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration/currentValue
func (i_ ICScannerFeatureEnumeration) CurrentValue() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("currentValue"))
	return rv
}/* debug [instance_properties/getter]: currentValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration/currentValue
func (i_ ICScannerFeatureEnumeration) SetCurrentValue(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrentValue:"), value)
}/* debug [instance_properties/setter]: currentValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFeatureEnumeration */



