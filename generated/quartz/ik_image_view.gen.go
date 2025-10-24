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
	// properties:
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)
	Autoresizes() bool
	SetAutoresizes(value bool)
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	CurrentToolMode() objc.IObject /* cross-framework: NSString */
	SetCurrentToolMode(value objc.IObject /* cross-framework: NSString */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DoubleClickOpensImageEditPanel() bool
	SetDoubleClickOpensImageEditPanel(value bool)
	Editable() bool
	SetEditable(value bool)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	ImageCorrection() objc.IObject /* cross-framework: Filter */
	SetImageCorrection(value objc.IObject /* cross-framework: Filter */)
	RotationAngle() float64
	SetRotationAngle(value float64)
	SupportsDragAndDrop() bool
	SetSupportsDragAndDrop(value bool)
	ZoomFactor() float64
	SetZoomFactor(value float64)
	// methods:
}

// A view that allows displaying and minor editing of an image.
//
// The class provides an efficient way to display images in a view while at the same time supporting a number of image editing operations such as rotating, zooming, and cropping. If possible, image rendering uses hardware acceleration to achieve optimal performance. The class is implemented as a subclass of . Similar to , the class is used to display a single image. You can provide an images for the view in any of these formats: File reference ( , , or a path) Data ( or ) Image ( or ) Providing a file reference is the preferred way to set the the image for a view because in addition to the actual image data, also handles the image metadata embedded in the file. The image view automatically fetches the metadata from a file reference, whereas for the other sources (except for a source), it cannot. For images set from other sources, you need to set the metadata separately. supports multi-frame images (TIFF, GIF, and so forth) and animated images.


// A view that allows displaying and minor editing of an image.
//
// [Full Topic]
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



// Specifies the automatic-hiding scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/autohidesscrollers
func (i_ IKImageView) AutohidesScrollers() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("autohidesScrollers"))
	return rv
}


// Specifies the automatic-hiding scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/autohidesscrollers
func (i_ IKImageView) SetAutohidesScrollers(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutohidesScrollers:"), value)
}


// Specifies the automatic resizing state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/autoresizes
func (i_ IKImageView) Autoresizes() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("autoresizes"))
	return rv
}


// Specifies the automatic resizing state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/autoresizes
func (i_ IKImageView) SetAutoresizes(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutoresizes:"), value)
}


// Specifies the background color for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/backgroundcolor
func (i_ IKImageView) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("backgroundColor"))
	return rv
}


// Specifies the background color for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/backgroundcolor
func (i_ IKImageView) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBackgroundColor:"), value)
}


// Specifies the current tool mode for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/currenttoolmode
func (i_ IKImageView) CurrentToolMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("currentToolMode"))
	return rv
}


// Specifies the current tool mode for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/currenttoolmode
func (i_ IKImageView) SetCurrentToolMode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrentToolMode:"), value)
}


// Specifies the delegate object of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/delegate
func (i_ IKImageView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// Specifies the delegate object of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/delegate
func (i_ IKImageView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}


// Specifies the image-opening state of the editing pane in the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/doubleclickopensimageeditpanel
func (i_ IKImageView) DoubleClickOpensImageEditPanel() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("doubleClickOpensImageEditPanel"))
	return rv
}


// Specifies the image-opening state of the editing pane in the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/doubleclickopensimageeditpanel
func (i_ IKImageView) SetDoubleClickOpensImageEditPanel(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDoubleClickOpensImageEditPanel:"), value)
}


// Specifies the editable state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/editable
func (i_ IKImageView) Editable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("editable"))
	return rv
}


// Specifies the editable state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/editable
func (i_ IKImageView) SetEditable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEditable:"), value)
}


// Specifies the horizontal scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/hashorizontalscroller
func (i_ IKImageView) HasHorizontalScroller() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}


// Specifies the horizontal scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/hashorizontalscroller
func (i_ IKImageView) SetHasHorizontalScroller(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasHorizontalScroller:"), value)
}


// Specifies the vertical scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/hasverticalscroller
func (i_ IKImageView) HasVerticalScroller() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}


// Specifies the vertical scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/hasverticalscroller
func (i_ IKImageView) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasVerticalScroller:"), value)
}


// Specifies a Core Image filter for image correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/imagecorrection
func (i_ IKImageView) ImageCorrection() objc.IObject /* cross-framework: Filter */ {
	rv := objc.Send[coreimage.Filter](i_.ID, objc.Sel("imageCorrection"))
	return rv
}


// Specifies a Core Image filter for image correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/imagecorrection
func (i_ IKImageView) SetImageCorrection(value objc.IObject /* cross-framework: Filter */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageCorrection:"), value)
}


// Specifies the rotation angle for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/rotationangle
func (i_ IKImageView) RotationAngle() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("rotationAngle"))
	return rv
}


// Specifies the rotation angle for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/rotationangle
func (i_ IKImageView) SetRotationAngle(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRotationAngle:"), value)
}


// Specifies the drag-and-drop support state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/supportsdraganddrop
func (i_ IKImageView) SupportsDragAndDrop() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("supportsDragAndDrop"))
	return rv
}


// Specifies the drag-and-drop support state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/supportsdraganddrop
func (i_ IKImageView) SetSupportsDragAndDrop(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportsDragAndDrop:"), value)
}


// Specifies the zoom factor for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/zoomfactor
func (i_ IKImageView) ZoomFactor() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("zoomFactor"))
	return rv
}


// Specifies the zoom factor for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/zoomfactor
func (i_ IKImageView) SetZoomFactor(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZoomFactor:"), value)
}



