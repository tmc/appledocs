// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ICScannerFeature */


/* debug [class_header]: Header for ICScannerFeature */
// The class instance for the [ICScannerFeature] class.
var (
	ICScannerFeatureClass     _ICScannerFeatureClass
	ICScannerFeatureClassOnce sync.Once
)

func getICScannerFeatureClass() _ICScannerFeatureClass {
	ICScannerFeatureClassOnce.Do(func() {
		ICScannerFeatureClass = _ICScannerFeatureClass{objc.GetClass("ICScannerFeature")}
	})
	return ICScannerFeatureClass
}

type _ICScannerFeatureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFeature */
// An interface definition for the [ICScannerFeature] class.
type IICScannerFeature interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ICScannerFeature */
	// properties:
	Tooltip() unsafe.Pointer
	SetTooltip(value unsafe.Pointer)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	InternalName() unsafe.Pointer
	SetInternalName(value unsafe.Pointer)
	HumanReadableName() unsafe.Pointer
	SetHumanReadableName(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFeature */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFeature */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFeatureClass) Alloc() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFeatureClass) New() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFeature) Init() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFeature) Autorelease() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeature creates a new ICScannerFeature instance.
func NewICScannerFeature() ICScannerFeature {
	return getICScannerFeatureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFeature */
// An abstract class that describes a scanner feature.
//
// The ImageCaptureCore framework defines three concrete subclasses of scanner features: , , and . Scanner functional units may have one or more instances of these classes to allow users to choose scanner-specific settings or operations before performing a scan.


// An abstract class that describes a scanner feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeature
type ICScannerFeature struct {
	objectivec.Object
}

// ICScannerFeatureFrom constructs a [ICScannerFeature] from an unsafe.Pointer.
//
// An abstract class that describes a scanner feature.
func ICScannerFeatureFrom(ptr unsafe.Pointer) ICScannerFeature {
	return ICScannerFeature{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFeature *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFeature */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFeature */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFeature */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFeature */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeature/1507600-tooltip
func (i_ ICScannerFeature) Tooltip() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("tooltip"))
	return rv
}/* debug [instance_properties/getter]: tooltip */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeature/1507600-tooltip
func (i_ ICScannerFeature) SetTooltip(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTooltip:"), value)
}/* debug [instance_properties/setter]: tooltip */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeature/1507807-type
func (i_ ICScannerFeature) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeature/1507807-type
func (i_ ICScannerFeature) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeature/1508021-internalname
func (i_ ICScannerFeature) InternalName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("internalName"))
	return rv
}/* debug [instance_properties/getter]: internalName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeature/1508021-internalname
func (i_ ICScannerFeature) SetInternalName(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInternalName:"), value)
}/* debug [instance_properties/setter]: internalName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeature/1508051-humanreadablename
func (i_ ICScannerFeature) HumanReadableName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("humanReadableName"))
	return rv
}/* debug [instance_properties/getter]: humanReadableName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeature/1508051-humanreadablename
func (i_ ICScannerFeature) SetHumanReadableName(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHumanReadableName:"), value)
}/* debug [instance_properties/setter]: humanReadableName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFeature */



