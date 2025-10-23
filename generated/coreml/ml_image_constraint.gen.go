// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ImageConstraint] class.
type IImageConstraint interface {
	objectivec.IObject
	ImageConstraint() IMLImageConstraint
	SetImageConstraint(value IMLImageConstraint)
	Type() MLFeatureType
	SetType(value MLFeatureType)
	PixelFormatType() unsafe.Pointer
	SetPixelFormatType(value unsafe.Pointer)
	PixelsHigh() int
	SetPixelsHigh(value int)
	PixelsWide() int
	SetPixelsWide(value int)
	SizeConstraint() ImageSizeConstraint
	SetSizeConstraint(value ImageSizeConstraint)
}

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

// Alloc allocates a new instance without initialization.
func (ic _ImageConstraintClass) Alloc() ImageConstraint {
	rv := objc.Send[ImageConstraint](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (i_ ImageConstraint) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[ImageConstraint](i_.ID, objc.Sel("imageConstraint"))
	return rv
}


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (i_ ImageConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageConstraint:"), value)
}


// The type of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/type
func (i_ ImageConstraint) Type() MLFeatureType {
	rv := objc.Send[MLFeatureType](i_.ID, objc.Sel("type"))
	return rv
}


// The type of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/type
func (i_ ImageConstraint) SetType(value MLFeatureType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setType:"), value)
}


// The model’s pixel format for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelformattype
func (i_ ImageConstraint) PixelFormatType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("pixelFormatType"))
	return rv
}


// The model’s pixel format for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelformattype
func (i_ ImageConstraint) SetPixelFormatType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelFormatType:"), value)
}


// The model’s default height for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelshigh
func (i_ ImageConstraint) PixelsHigh() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsHigh"))
	return rv
}


// The model’s default height for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelshigh
func (i_ ImageConstraint) SetPixelsHigh(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsHigh:"), value)
}


// The model’s default width for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelswide
func (i_ ImageConstraint) PixelsWide() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsWide"))
	return rv
}


// The model’s default width for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelswide
func (i_ ImageConstraint) SetPixelsWide(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsWide:"), value)
}


// Additional sizes this image feature supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/sizeconstraint
func (i_ ImageConstraint) SizeConstraint() ImageSizeConstraint {
	rv := objc.Send[ImageSizeConstraint](i_.ID, objc.Sel("sizeConstraint"))
	return rv
}


// Additional sizes this image feature supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/sizeconstraint
func (i_ ImageConstraint) SetSizeConstraint(value ImageSizeConstraint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSizeConstraint:"), value)
}



