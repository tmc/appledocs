// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLImageConstraint */


/* debug [class_header]: Header for MLImageConstraint */
// The class instance for the [ImageConstraint] class.
var (
	ImageConstraintClass     _ImageConstraintClass
	ImageConstraintClassOnce sync.Once
)

func getImageConstraintClass() _ImageConstraintClass {
	ImageConstraintClassOnce.Do(func() {
		ImageConstraintClass = _ImageConstraintClass{objc.GetClass("MLImageConstraint")}
	})
	return ImageConstraintClass
}

type _ImageConstraintClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageConstraint */
// An interface definition for the [ImageConstraint] class.
type IImageConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ImageConstraint */
	// properties:
	PixelFormatType() uint32 /* not a class type */
	PixelsHigh() int
	PixelsWide() int
	SizeConstraint() IMLImageSizeConstraint
	ImageConstraint() IMLImageConstraint
	SetImageConstraint(value IMLImageConstraint)
	Type() FeatureType
	SetType(value FeatureType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageConstraint */
// Alloc allocates a new instance without initialization.
func (ic _ImageConstraintClass) Alloc() ImageConstraint {
	rv := objc.Send[ImageConstraint](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageConstraintClass) New() ImageConstraint {
	rv := objc.Send[ImageConstraint](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageConstraint) Init() ImageConstraint {
	rv := objc.Send[ImageConstraint](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageConstraint) Autorelease() ImageConstraint {
	rv := objc.Send[ImageConstraint](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageConstraint creates a new ImageConstraint instance.
func NewImageConstraint() ImageConstraint {
	return getImageConstraintClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageConstraint */
// The width, height, and pixel format constraints of an image feature.
//
// In CoreML, an is a collection of pixels represented by (Swift) or (Objective-C). An is a model input or output that accepts or produces, respectively, an image bundled in an . defines the image feature’s limitations for the images within an . If a model has an image feature for an input or output, the model author uses an by creating an . The feature description for an image input or output has: Its property set to Its property set to an instance configured to the image feature’s size and format Image features that support additional image sizes provide a range of sizes, or a list of discrete sizes, in their image constraint’s property.


// The width, height, and pixel format constraints of an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageConstraint
type ImageConstraint struct {
	objectivec.Object
}

// ImageConstraintFrom constructs a [ImageConstraint] from an unsafe.Pointer.
//
// The width, height, and pixel format constraints of an image feature.
func ImageConstraintFrom(ptr unsafe.Pointer) ImageConstraint {
	return ImageConstraint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageConstraint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageConstraint */

// The model’s pixel format for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageConstraint/pixelFormatType
func (i_ ImageConstraint) PixelFormatType() uint32 /* not a class type */ {
	rv := objc.Send[uint32](i_.ID, objc.Sel("pixelFormatType"))
	return rv
}/* debug [instance_properties/getter]: pixelFormatType */


// The model’s default height for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageConstraint/pixelsHigh
func (i_ ImageConstraint) PixelsHigh() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsHigh"))
	return rv
}/* debug [instance_properties/getter]: pixelsHigh */


// The model’s default width for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageConstraint/pixelsWide
func (i_ ImageConstraint) PixelsWide() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsWide"))
	return rv
}/* debug [instance_properties/getter]: pixelsWide */


// Additional sizes this image feature supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageConstraint/sizeConstraint
func (i_ ImageConstraint) SizeConstraint() IMLImageSizeConstraint {
	rv := objc.Send[ImageSizeConstraint](i_.ID, objc.Sel("sizeConstraint"))
	return rv
}/* debug [instance_properties/getter]: sizeConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (i_ ImageConstraint) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[ImageConstraint](i_.ID, objc.Sel("imageConstraint"))
	return rv
}/* debug [instance_properties/getter]: imageConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (i_ ImageConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageConstraint:"), value)
}/* debug [instance_properties/setter]: imageConstraint */


// The type of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/type
func (i_ ImageConstraint) Type() FeatureType {
	rv := objc.Send[FeatureType](i_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The type of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/type
func (i_ ImageConstraint) SetType(value FeatureType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLImageConstraint */



