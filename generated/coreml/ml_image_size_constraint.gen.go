// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ImageSizeConstraint] class.
type IImageSizeConstraint interface {
	objectivec.IObject
	EnumeratedImageSizes() []ImageSize
	PixelsHighRange() foundation.Range
	PixelsWideRange() foundation.Range
	Type() ImageSizeConstraintType
	PixelsHigh() int
	SetPixelsHigh(value int)
	PixelsWide() int
	SetPixelsWide(value int)
	SizeConstraint() MLImageSizeConstraint
	SetSizeConstraint(value IMLImageSizeConstraint)
}

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

// Alloc allocates a new instance without initialization.
func (ic _ImageSizeConstraintClass) Alloc() ImageSizeConstraint {
	rv := objc.Send[ImageSizeConstraint](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An array of image sizes a model’s image feature accepts as input or produces as output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/enumeratedImageSizes

func (i_ ImageSizeConstraint) EnumeratedImageSizes() []ImageSize {
	rv := objc.Send[[]ImageSize](i_.ID, objc.Sel("enumeratedImageSizes"))
	return rv
}


// The range of heights a model’s image feature accepts as input or produces as output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/pixelsHighRange

func (i_ ImageSizeConstraint) PixelsHighRange() foundation.Range {
	rv := objc.Send[foundation.Range](i_.ID, objc.Sel("pixelsHighRange"))
	return rv
}


// The range of widths a model’s image feature accepts as input or produces as output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/pixelsWideRange

func (i_ ImageSizeConstraint) PixelsWideRange() foundation.Range {
	rv := objc.Send[foundation.Range](i_.ID, objc.Sel("pixelsWideRange"))
	return rv
}


// Indicator of which properties to inspect for this image size constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/type

func (i_ ImageSizeConstraint) Type() ImageSizeConstraintType {
	rv := objc.Send[ImageSizeConstraintType](i_.ID, objc.Sel("type"))
	return rv
}


// The model’s default height for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelshigh

func (i_ ImageSizeConstraint) PixelsHigh() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsHigh"))
	return rv
}


// The model’s default height for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelshigh

func (i_ ImageSizeConstraint) SetPixelsHigh(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsHigh:"), value)
}


// The model’s default width for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelswide

func (i_ ImageSizeConstraint) PixelsWide() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsWide"))
	return rv
}


// The model’s default width for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelswide

func (i_ ImageSizeConstraint) SetPixelsWide(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPixelsWide:"), value)
}


// Additional sizes this image feature supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/sizeconstraint

func (i_ ImageSizeConstraint) SizeConstraint() MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](i_.ID, objc.Sel("sizeConstraint"))
	return rv
}


// Additional sizes this image feature supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimageconstraint/sizeconstraint

func (i_ ImageSizeConstraint) SetSizeConstraint(value IMLImageSizeConstraint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSizeConstraint:"), value)
}



