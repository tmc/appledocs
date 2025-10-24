// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ICScannerFeatureBoolean */


/* debug [class_header]: Header for ICScannerFeatureBoolean */
// The class instance for the [ICScannerFeatureBoolean] class.
var (
	ICScannerFeatureBooleanClass     _ICScannerFeatureBooleanClass
	ICScannerFeatureBooleanClassOnce sync.Once
)

func getICScannerFeatureBooleanClass() _ICScannerFeatureBooleanClass {
	ICScannerFeatureBooleanClassOnce.Do(func() {
		ICScannerFeatureBooleanClass = _ICScannerFeatureBooleanClass{objc.GetClass("ICScannerFeatureBoolean")}
	})
	return ICScannerFeatureBooleanClass
}

type _ICScannerFeatureBooleanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFeatureBoolean */
// An interface definition for the [ICScannerFeatureBoolean] class.
type IICScannerFeatureBoolean interface {
	IICScannerFeature
	
/* debug [class_interface_properties]: Properties for ICScannerFeatureBoolean */
	// properties:
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFeatureBoolean */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFeatureBoolean */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFeatureBooleanClass) Alloc() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFeatureBooleanClass) New() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFeatureBoolean) Init() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFeatureBoolean) Autorelease() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureBoolean creates a new ICScannerFeatureBoolean instance.
func NewICScannerFeatureBoolean() ICScannerFeatureBoolean {
	return getICScannerFeatureBooleanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFeatureBoolean */
// A feature with a value of or .


// A feature with a value of or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureBoolean
type ICScannerFeatureBoolean struct {
	ICScannerFeature
}

// ICScannerFeatureBooleanFrom constructs a [ICScannerFeatureBoolean] from an unsafe.Pointer.
//
// A feature with a value of or .
func ICScannerFeatureBooleanFrom(ptr unsafe.Pointer) ICScannerFeatureBoolean {
	return ICScannerFeatureBoolean{
		ICScannerFeature: ICScannerFeatureFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFeatureBoolean *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFeatureBoolean */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFeatureBoolean */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFeatureBoolean */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFeatureBoolean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureboolean/1507884-value
func (i_ ICScannerFeatureBoolean) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeatureboolean/1507884-value
func (i_ ICScannerFeatureBoolean) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFeatureBoolean */



