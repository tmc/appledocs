// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ICScannerFeatureTemplate */


/* debug [class_header]: Header for ICScannerFeatureTemplate */
// The class instance for the [ICScannerFeatureTemplate] class.
var (
	ICScannerFeatureTemplateClass     _ICScannerFeatureTemplateClass
	ICScannerFeatureTemplateClassOnce sync.Once
)

func getICScannerFeatureTemplateClass() _ICScannerFeatureTemplateClass {
	ICScannerFeatureTemplateClassOnce.Do(func() {
		ICScannerFeatureTemplateClass = _ICScannerFeatureTemplateClass{objc.GetClass("ICScannerFeatureTemplate")}
	})
	return ICScannerFeatureTemplateClass
}

type _ICScannerFeatureTemplateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFeatureTemplate */
// An interface definition for the [ICScannerFeatureTemplate] class.
type IICScannerFeatureTemplate interface {
	IICScannerFeature
	
/* debug [class_interface_properties]: Properties for ICScannerFeatureTemplate */
	// properties:
	Targets() foundation.MutableArray
	SetTargets(value foundation.MutableArray)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFeatureTemplate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFeatureTemplate */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFeatureTemplateClass) Alloc() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFeatureTemplateClass) New() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFeatureTemplate) Init() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFeatureTemplate) Autorelease() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureTemplate creates a new ICScannerFeatureTemplate instance.
func NewICScannerFeatureTemplate() ICScannerFeatureTemplate {
	return getICScannerFeatureTemplateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFeatureTemplate */
// A group of one or more rectangular scan areas that can be used with a scanner functional unit.


// A group of one or more rectangular scan areas that can be used with a scanner functional unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureTemplate
type ICScannerFeatureTemplate struct {
	ICScannerFeature
}

// ICScannerFeatureTemplateFrom constructs a [ICScannerFeatureTemplate] from an unsafe.Pointer.
//
// A group of one or more rectangular scan areas that can be used with a scanner functional unit.
func ICScannerFeatureTemplateFrom(ptr unsafe.Pointer) ICScannerFeatureTemplate {
	return ICScannerFeatureTemplate{
		ICScannerFeature: ICScannerFeatureFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFeatureTemplate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFeatureTemplate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFeatureTemplate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFeatureTemplate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFeatureTemplate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturetemplate/1508048-targets
func (i_ ICScannerFeatureTemplate) Targets() foundation.MutableArray {
	rv := objc.Send[foundation.MutableArray](i_.ID, objc.Sel("targets"))
	return rv
}/* debug [instance_properties/getter]: targets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturetemplate/1508048-targets
func (i_ ICScannerFeatureTemplate) SetTargets(value foundation.MutableArray) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTargets:"), value)
}/* debug [instance_properties/setter]: targets */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFeatureTemplate */



