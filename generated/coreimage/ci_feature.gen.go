// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIFeature */


/* debug [class_header]: Header for CIFeature */
// The class instance for the [Feature] class.
var (
	FeatureClass     _FeatureClass
	FeatureClassOnce sync.Once
)

func getFeatureClass() _FeatureClass {
	FeatureClassOnce.Do(func() {
		FeatureClass = _FeatureClass{objc.GetClass("CIFeature")}
	})
	return FeatureClass
}

type _FeatureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Feature */
// An interface definition for the [Feature] class.
type IFeature interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Feature */
	// properties:
	Bounds() corefoundation.CGRect
	Type() objc.IObject /* cross-framework: NSString */
	CIFeatureTypeFace() objc.IObject /* cross-framework: NSString */
	CIFeatureTypeQRCode() objc.IObject /* cross-framework: NSString */
	CIFeatureTypeRectangle() objc.IObject /* cross-framework: NSString */
	CIFeatureTypeText() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Feature */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Feature */
// Alloc allocates a new instance without initialization.
func (fc _FeatureClass) Alloc() Feature {
	rv := objc.Send[Feature](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FeatureClass) New() Feature {
	rv := objc.Send[Feature](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Feature) Init() Feature {
	rv := objc.Send[Feature](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Feature) Autorelease() Feature {
	rv := objc.Send[Feature](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeature creates a new Feature instance.
func NewFeature() Feature {
	return getFeatureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Feature */
// The abstract superclass for objects representing notable features detected in an image.
//
// A object represents a portion of an image that a detector believes matches its criteria. Subclasses of CIFeature holds additional information specific to the detector that discovered the feature.


// The abstract superclass for objects representing notable features detected in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFeature
type Feature struct {
	objectivec.Object
}

// FeatureFrom constructs a [Feature] from an unsafe.Pointer.
//
// The abstract superclass for objects representing notable features detected in an image.
func FeatureFrom(ptr unsafe.Pointer) Feature {
	return Feature{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Feature *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Feature */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Feature */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Feature */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Feature */

// The rectangle that holds discovered feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFeature/bounds
func (f_ Feature) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The type of feature that was discovered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFeature/type
func (f_ Feature) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A Core Image feature type for person’s face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifeaturetypeface
func (f_ Feature) CIFeatureTypeFace() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("CIFeatureTypeFace"))
	return rv
}/* debug [instance_properties/getter]: CIFeatureTypeFace */


// A Core Image feature type for QR code object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifeaturetypeqrcode
func (f_ Feature) CIFeatureTypeQRCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("CIFeatureTypeQRCode"))
	return rv
}/* debug [instance_properties/getter]: CIFeatureTypeQRCode */


// A Core Image feature type for rectangular object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifeaturetyperectangle
func (f_ Feature) CIFeatureTypeRectangle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("CIFeatureTypeRectangle"))
	return rv
}/* debug [instance_properties/getter]: CIFeatureTypeRectangle */


// A Core Image feature type for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifeaturetypetext
func (f_ Feature) CIFeatureTypeText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("CIFeatureTypeText"))
	return rv
}/* debug [instance_properties/getter]: CIFeatureTypeText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIFeature */



