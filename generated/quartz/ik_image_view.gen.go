// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coreimage"
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
	ConvertImageRectToViewRect(imageRect foundation.IRect) foundation.Rect
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
func (i_ IKImageView) ConvertImageRectToViewRect(imageRect foundation.IRect) foundation.Rect {
	rv := objc.Send[foundation.Rect](i_.ID, objc.Sel("convertImageRectToViewRect:"), imageRect)
	return rv
}

// Specifies the automatic-hiding scroll bar state for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/autohidesscrollers
func (i_ IKImageView) AutohidesScrollers() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("autohidesScrollers"))
	return rv
}


// SetAutohidesScrollers sets the value of the autohidesScrollers property.
// Specifies the automatic-hiding scroll bar state for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/autohidesscrollers
func (i_ IKImageView) SetAutohidesScrollers(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutohidesScrollers:"), value)
}

// Specifies the automatic resizing state for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/autoresizes
func (i_ IKImageView) Autoresizes() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("autoresizes"))
	return rv
}


// SetAutoresizes sets the value of the autoresizes property.
// Specifies the automatic resizing state for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/autoresizes
func (i_ IKImageView) SetAutoresizes(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutoresizes:"), value)
}

// Specifies the background color for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/backgroundcolor
func (i_ IKImageView) BackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// Specifies the background color for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/backgroundcolor
func (i_ IKImageView) SetBackgroundColor(value appkit.IColor) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBackgroundColor:"), value)
}

// Specifies the current tool mode for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/currenttoolmode
func (i_ IKImageView) CurrentToolMode() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("currentToolMode"))
	return rv
}


// SetCurrentToolMode sets the value of the currentToolMode property.
// Specifies the current tool mode for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/currenttoolmode
func (i_ IKImageView) SetCurrentToolMode(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrentToolMode:"), value)
}

// Specifies the delegate object of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/delegate
func (i_ IKImageView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// Specifies the delegate object of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/delegate
func (i_ IKImageView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}

// Specifies the image-opening state of the editing pane in the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/doubleclickopensimageeditpanel
func (i_ IKImageView) DoubleClickOpensImageEditPanel() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("doubleClickOpensImageEditPanel"))
	return rv
}


// SetDoubleClickOpensImageEditPanel sets the value of the doubleClickOpensImageEditPanel property.
// Specifies the image-opening state of the editing pane in the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/doubleclickopensimageeditpanel
func (i_ IKImageView) SetDoubleClickOpensImageEditPanel(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDoubleClickOpensImageEditPanel:"), value)
}

// Specifies the editable state for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/editable
func (i_ IKImageView) Editable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("editable"))
	return rv
}


// SetEditable sets the value of the editable property.
// Specifies the editable state for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/editable
func (i_ IKImageView) SetEditable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEditable:"), value)
}

// Specifies the horizontal scroll bar state for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/hashorizontalscroller
func (i_ IKImageView) HasHorizontalScroller() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}


// SetHasHorizontalScroller sets the value of the hasHorizontalScroller property.
// Specifies the horizontal scroll bar state for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/hashorizontalscroller
func (i_ IKImageView) SetHasHorizontalScroller(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasHorizontalScroller:"), value)
}

// Specifies the vertical scroll bar state for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/hasverticalscroller
func (i_ IKImageView) HasVerticalScroller() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}


// SetHasVerticalScroller sets the value of the hasVerticalScroller property.
// Specifies the vertical scroll bar state for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/hasverticalscroller
func (i_ IKImageView) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasVerticalScroller:"), value)
}

// Specifies a Core Image filter for image correction.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/imagecorrection
func (i_ IKImageView) ImageCorrection() coreimage.Filter {
	rv := objc.Send[coreimage.Filter](i_.ID, objc.Sel("imageCorrection"))
	return rv
}


// SetImageCorrection sets the value of the imageCorrection property.
// Specifies a Core Image filter for image correction.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/imagecorrection
func (i_ IKImageView) SetImageCorrection(value coreimage.IFilter) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageCorrection:"), value)
}

// Specifies the rotation angle for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/rotationangle
func (i_ IKImageView) RotationAngle() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("rotationAngle"))
	return rv
}


// SetRotationAngle sets the value of the rotationAngle property.
// Specifies the rotation angle for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/rotationangle
func (i_ IKImageView) SetRotationAngle(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRotationAngle:"), value)
}

// Specifies the drag-and-drop support state for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/supportsdraganddrop
func (i_ IKImageView) SupportsDragAndDrop() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("supportsDragAndDrop"))
	return rv
}


// SetSupportsDragAndDrop sets the value of the supportsDragAndDrop property.
// Specifies the drag-and-drop support state for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/supportsdraganddrop
func (i_ IKImageView) SetSupportsDragAndDrop(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportsDragAndDrop:"), value)
}

// Specifies the zoom factor for the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/zoomfactor
func (i_ IKImageView) ZoomFactor() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("zoomFactor"))
	return rv
}


// SetZoomFactor sets the value of the zoomFactor property.
// Specifies the zoom factor for the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/zoomfactor
func (i_ IKImageView) SetZoomFactor(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZoomFactor:"), value)
}



