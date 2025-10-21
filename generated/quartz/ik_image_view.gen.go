// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [IKImageView] class.
var (
	IKImageViewClass     _IKImageViewClass
	IKImageViewClassOnce sync.Once
)

func getIKImageViewClass() _IKImageViewClass {
	IKImageViewClassOnce.Do(func() {
		IKImageViewClass = _IKImageViewClass{objc.GetClass("IKImageView")}
	})
	return IKImageViewClass
}

type _IKImageViewClass struct {
	class objc.Class
}

// An interface definition for the [IKImageView] class.
type IIKImageView interface {
	appkit.IView
	ConvertImageRectToViewRect(imageRect foundation.Rect) foundation.Rect
}

// A view that allows displaying and minor editing of an image.
//
// The class provides an efficient way to display images in a view while at the same time supporting a number of image editing operations such as rotating, zooming, and cropping. If possible, image rendering uses hardware acceleration to achieve optimal performance. The class is implemented as a subclass of . Similar to , the class is used to display a single image. You can provide an images for the view in any of these formats: File reference ( , , or a path) Data ( or ) Image ( or ) Providing a file reference is the preferred way to set the the image for a view because in addition to the actual image data, also handles the image metadata embedded in the file. The image view automatically fetches the metadata from a file reference, whereas for the other sources (except for a source), it cannot. For images set from other sources, you need to set the metadata separately. supports multi-frame images (TIFF, GIF, and so forth) and animated images.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView
type IKImageView struct {
	appkit.View
}

// IKImageViewFrom constructs a [IKImageView] from an unsafe.Pointer.
//
// A view that allows displaying and minor editing of an image.
func IKImageViewFrom(ptr unsafe.Pointer) IKImageView {
	return IKImageView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKImageViewClass) Alloc() IKImageView {
	rv := objc.Send[IKImageView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKImageViewClass) New() IKImageView {
	rv := objc.Send[IKImageView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKImageView) Init() IKImageView {
	rv := objc.Send[IKImageView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKImageView) Autorelease() IKImageView {
	rv := objc.Send[IKImageView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKImageView creates a new IKImageView instance.
func NewIKImageView() IKImageView {
	return getIKImageViewClass().New()
}


// Converts an image rectangle to an image view rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/convertImageRect(toViewRect:)
func (i_ IKImageView) ConvertImageRectToViewRect(imageRect foundation.Rect) foundation.Rect {
	rv := objc.Send[foundation.Rect](i_.ID, objc.Sel("convertImageRectToViewRect:"), imageRect)
	return rv
}



