// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLImageSizeConstraint */


/* debug [class_header]: Header for MLImageSizeConstraint */
// The class instance for the [ImageSizeConstraint] class.
var (
	ImageSizeConstraintClass     _ImageSizeConstraintClass
	ImageSizeConstraintClassOnce sync.Once
)

func getImageSizeConstraintClass() _ImageSizeConstraintClass {
	ImageSizeConstraintClassOnce.Do(func() {
		ImageSizeConstraintClass = _ImageSizeConstraintClass{objc.GetClass("MLImageSizeConstraint")}
	})
	return ImageSizeConstraintClass
}

type _ImageSizeConstraintClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageSizeConstraint */
// An interface definition for the [ImageSizeConstraint] class.
type IImageSizeConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ImageSizeConstraint */
	// properties:
	EnumeratedImageSizes() []ImageSize
	PixelsHighRange() corefoundation.Range
	PixelsWideRange() corefoundation.Range
	Type() ImageSizeConstraintType
	PixelsHigh() int
	SetPixelsHigh(value int)
	PixelsWide() int
	SetPixelsWide(value int)
	SizeConstraint() IMLImageSizeConstraint
	SetSizeConstraint(value IMLImageSizeConstraint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageSizeConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageSizeConstraint */
// Alloc allocates a new instance without initialization.
func (ic _ImageSizeConstraintClass) Alloc() ImageSizeConstraint {
	rv := objc.Send[ImageSizeConstraint](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageSizeConstraintClass) New() ImageSizeConstraint {
	rv := objc.Send[ImageSizeConstraint](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageSizeConstraint) Init() ImageSizeConstraint {
	rv := objc.Send[ImageSizeConstraint](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageSizeConstraint) Autorelease() ImageSizeConstraint {
	rv := objc.Send[ImageSizeConstraint](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageSizeConstraint creates a new ImageSizeConstraint instance.
func NewImageSizeConstraint() ImageSizeConstraint {
	return getImageSizeConstraintClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageSizeConstraint */
// A list or range of sizes that augment an image constraint’s default size.
//
// You use an to express what image sizes of an image feature a model will accept as input or produce as output. Use to determine which properties describe what image sizes the model’s image feature expects as input or produces as output. If is: , the image feature accepts any image that has a width in and a height in . , the image feature accepts any image size listed in . , the instance is not configured and should be ignored. Instead, use the image feature’s default image size constraint, defined by and .


// A list or range of sizes that augment an image constraint’s default size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint
type ImageSizeConstraint struct {
	objectivec.Object
}

// ImageSizeConstraintFrom constructs a [ImageSizeConstraint] from an unsafe.Pointer.
//
// A list or range of sizes that augment an image constraint’s default size.
func ImageSizeConstraintFrom(ptr unsafe.Pointer) ImageSizeConstraint {
	return ImageSizeConstraint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageSizeConstraint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageSizeConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageSizeConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageSizeConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageSizeConstraint */

// An array of image sizes a model’s image feature accepts as input or produces as output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/enumeratedImageSizes
func (i_ ImageSizeConstraint) EnumeratedImageSizes() []ImageSize {
	rv := objc.Send[[]ImageSize](i_.ID, objc.Sel("enumeratedImageSizes"))
	return rv
}/* debug [instance_properties/getter]: enumeratedImageSizes */


// The range of heights a model’s image feature accepts as input or produces as output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/pixelsHighRange
func (i_ ImageSizeConstraint) PixelsHighRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](i_.ID, objc.Sel("pixelsHighRange"))
	return rv
}/* debug [instance_properties/getter]: pixelsHighRange */


// The range of widths a model’s image feature accepts as input or produces as output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/pixelsWideRange
func (i_ ImageSizeConstraint) PixelsWideRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](i_.ID, objc.Sel("pixelsWideRange"))
	return rv
}/* debug [instance_properties/getter]: pixelsWideRange */


// Indicator of which properties to inspect for this image size constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/type
func (i_ ImageSizeConstraint) Type() ImageSizeConstraintType {
	rv := objc.Send[ImageSizeConstraintType](i_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The model’s default height for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelshigh
func (i_ ImageSizeConstraint) PixelsHigh() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsHigh"))
	return rv
}/* debug [instance_properties/getter]: pixelsHigh */


// The model’s default height for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelshigh
func (i_ ImageSizeConstraint) SetPixelsHigh(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsHigh:"), value)
}/* debug [instance_properties/setter]: pixelsHigh */


// The model’s default width for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelswide
func (i_ ImageSizeConstraint) PixelsWide() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsWide"))
	return rv
}/* debug [instance_properties/getter]: pixelsWide */


// The model’s default width for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelswide
func (i_ ImageSizeConstraint) SetPixelsWide(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsWide:"), value)
}/* debug [instance_properties/setter]: pixelsWide */


// Additional sizes this image feature supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/sizeconstraint
func (i_ ImageSizeConstraint) SizeConstraint() IMLImageSizeConstraint {
	rv := objc.Send[ImageSizeConstraint](i_.ID, objc.Sel("sizeConstraint"))
	return rv
}/* debug [instance_properties/getter]: sizeConstraint */


// Additional sizes this image feature supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/sizeconstraint
func (i_ ImageSizeConstraint) SetSizeConstraint(value IMLImageSizeConstraint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSizeConstraint:"), value)
}/* debug [instance_properties/setter]: sizeConstraint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLImageSizeConstraint */



