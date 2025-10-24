// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class IKImageView */


/* debug [class_header]: Header for IKImageView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKImageView */
// An interface definition for the [IKImageView] class.
type IIKImageView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for IKImageView */
	// properties:
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)
	Autoresizes() bool
	SetAutoresizes(value bool)
	BackgroundColor() appkit.Color
	SetBackgroundColor(value appkit.Color)
	CurrentToolMode() objc.IObject /* cross-framework: NSString */
	SetCurrentToolMode(value objc.IObject /* cross-framework: NSString */)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DoubleClickOpensImageEditPanel() bool
	SetDoubleClickOpensImageEditPanel(value bool)
	Editable() bool
	SetEditable(value bool)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	ImageCorrection() coreimage.Filter
	SetImageCorrection(value coreimage.Filter)
	RotationAngle() float64
	SetRotationAngle(value float64)
	SupportsDragAndDrop() bool
	SetSupportsDragAndDrop(value bool)
	ZoomFactor() float64
	SetZoomFactor(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKImageView */
	// methods:
	ConvertImagePointToViewPoint(imagePoint vision.Point) vision.Point
	ConvertImageRectToViewRect(imageRect Rect /* not a class type */) Rect /* not a class type */
	ConvertViewPointToImagePoint(viewPoint vision.Point) vision.Point
	ConvertViewRectToImageRect(viewRect Rect /* not a class type */) Rect /* not a class type */
	Crop(sender objc.IObject)
	FlipImageHorizontal(sender objc.IObject)
	FlipImageVertical(sender objc.IObject)
	Image() ImageRef /* not a class type */
	ImageProperties() foundation.Dictionary
	ImageSize() Size /* not a class type */
	OverlayForType(layerType objc.IObject /* cross-framework: NSString */) avfoundation.Layer
	RotateImageLeft(sender objc.IObject)
	RotateImageRight(sender objc.IObject)
	ScrollToRect(rect Rect /* not a class type */)
	ScrollToPoint(point vision.Point)
	SetImageImageProperties(image ImageRef /* not a class type */, metaData objc.IObject /* cross-framework: NSDictionary */)
	SetImageWithURL(url objc.IObject /* cross-framework: NSURL */)
	SetImageZoomFactorCenterPoint(zoomFactor float64, centerPoint vision.Point)
	SetOverlayForType(layer avfoundation.Layer, layerType objc.IObject /* cross-framework: NSString */)
	SetRotationAngleCenterPoint(rotationAngle float64, centerPoint vision.Point)
	ZoomImageToRect(rect Rect /* not a class type */)
	ZoomImageToActualSize(sender objc.IObject)
	ZoomImageToFit(sender objc.IObject)
	ZoomIn(sender objc.IObject)
	ZoomOut(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKImageView */
// Alloc allocates a new instance without initialization.
func (ic _IKImageViewClass) Alloc() IKImageView {
	rv := objc.Send[IKImageView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKImageView */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKImageView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKImageView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKImageView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKImageView */

// Converts an image coordinate to an image view coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/convertImagePoint(toViewPoint:)
func (i_ IKImageView) ConvertImagePointToViewPoint(imagePoint vision.Point) vision.Point {
	rv := objc.Send[vision.Point](i_.ID, objc.Sel("convertImagePointToViewPoint:"), imagePoint)
	return rv
}/* debug [instance_methods/method]: ConvertImagePointToViewPoint */


// Converts an image rectangle to an image view rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/convertImageRect(toViewRect:)
func (i_ IKImageView) ConvertImageRectToViewRect(imageRect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("convertImageRectToViewRect:"), imageRect)
	return rv
}/* debug [instance_methods/method]: ConvertImageRectToViewRect */


// Converts an image view coordinate to an image coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/convertPoint(toImagePoint:)
func (i_ IKImageView) ConvertViewPointToImagePoint(viewPoint vision.Point) vision.Point {
	rv := objc.Send[vision.Point](i_.ID, objc.Sel("convertViewPointToImagePoint:"), viewPoint)
	return rv
}/* debug [instance_methods/method]: ConvertViewPointToImagePoint */


// Converts an image view rectangle to an image rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/convertRect(toImageRect:)
func (i_ IKImageView) ConvertViewRectToImageRect(viewRect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("convertViewRectToImageRect:"), viewRect)
	return rv
}/* debug [instance_methods/method]: ConvertViewRectToImageRect */


// Crops the image using the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/crop(_:)
func (i_ IKImageView) Crop(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("crop:"), sender)
}/* debug [instance_methods/method]: Crop */


// Flips an image along the horizontal axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/flipImageHorizontal(_:)
func (i_ IKImageView) FlipImageHorizontal(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("flipImageHorizontal:"), sender)
}/* debug [instance_methods/method]: FlipImageHorizontal */


// Flips an image along the vertical axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/flipImageVertical(_:)
func (i_ IKImageView) FlipImageVertical(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("flipImageVertical:"), sender)
}/* debug [instance_methods/method]: FlipImageVertical */


// Returns the image associated with the view, after any image corrections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/image()
func (i_ IKImageView) Image() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](i_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_methods/method]: Image */


// Returns the metadata for the image in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/imageProperties()
func (i_ IKImageView) ImageProperties() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](i_.ID, objc.Sel("imageProperties"))
	return rv
}/* debug [instance_methods/method]: ImageProperties */


// Returns the size of the image in the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/imageSize()
func (i_ IKImageView) ImageSize() Size /* not a class type */ {
	rv := objc.Send[Size](i_.ID, objc.Sel("imageSize"))
	return rv
}/* debug [instance_methods/method]: ImageSize */


// Returns the Core Animation layer associated with a layer type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/overlay(forType:)
func (i_ IKImageView) OverlayForType(layerType objc.IObject /* cross-framework: NSString */) avfoundation.Layer {
	rv := objc.Send[avfoundation.Layer](i_.ID, objc.Sel("overlayForType:"), layerType)
	return rv
}/* debug [instance_methods/method]: OverlayForType */


// Rotates the image left (counter-clockwise).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/rotateImageLeft(_:)
func (i_ IKImageView) RotateImageLeft(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("rotateImageLeft:"), sender)
}/* debug [instance_methods/method]: RotateImageLeft */


// Rotates the image right (clockwise).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/rotateImageRight(_:)
func (i_ IKImageView) RotateImageRight(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("rotateImageRight:"), sender)
}/* debug [instance_methods/method]: RotateImageRight */


// Scrolls the view so that it includes the provided rectangular area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/scroll(to:)-535q6
func (i_ IKImageView) ScrollToRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("scrollToRect:"), rect)
}/* debug [instance_methods/method]: ScrollToRect */


// Scrolls the view to the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/scroll(to:)-myqk
func (i_ IKImageView) ScrollToPoint(point vision.Point) {
	objc.Send[objc.ID](i_.ID, objc.Sel("scrollToPoint:"), point)
}/* debug [instance_methods/method]: ScrollToPoint */


// Sets the image to display in an image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/setImage(_:imageProperties:)
func (i_ IKImageView) SetImageImageProperties(image ImageRef /* not a class type */, metaData objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImage:imageProperties:"), image, metaData)
}/* debug [instance_methods/method]: SetImageImageProperties */


// Initializes an image view with the image specified by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/setImageWith(_:)
func (i_ IKImageView) SetImageWithURL(url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageWithURL:"), url)
}/* debug [instance_methods/method]: SetImageWithURL */


// Sets the zoom factor at the provided origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/setImageZoomFactor(_:center:)
func (i_ IKImageView) SetImageZoomFactorCenterPoint(zoomFactor float64, centerPoint vision.Point) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageZoomFactor:centerPoint:"), zoomFactor, centerPoint)
}/* debug [instance_methods/method]: SetImageZoomFactorCenterPoint */


// Sets an overlay type for a Core Animation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/setOverlay(_:forType:)
func (i_ IKImageView) SetOverlayForType(layer avfoundation.Layer, layerType objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOverlay:forType:"), layer, layerType)
}/* debug [instance_methods/method]: SetOverlayForType */


// Sets the rotation angle at the provided origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/setRotationAngle(_:center:)
func (i_ IKImageView) SetRotationAngleCenterPoint(rotationAngle float64, centerPoint vision.Point) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRotationAngle:centerPoint:"), rotationAngle, centerPoint)
}/* debug [instance_methods/method]: SetRotationAngleCenterPoint */


// Zooms the image so that it fits in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/zoomImage(to:)
func (i_ IKImageView) ZoomImageToRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("zoomImageToRect:"), rect)
}/* debug [instance_methods/method]: ZoomImageToRect */


// Zooms the image so that it is displayed using its true size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/zoomImageToActualSize(_:)
func (i_ IKImageView) ZoomImageToActualSize(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("zoomImageToActualSize:"), sender)
}/* debug [instance_methods/method]: ZoomImageToActualSize */


// Zooms the image so that it fits in the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/zoomImageToFit(_:)
func (i_ IKImageView) ZoomImageToFit(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("zoomImageToFit:"), sender)
}/* debug [instance_methods/method]: ZoomImageToFit */


// Zooms the image in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/zoomIn(_:)
func (i_ IKImageView) ZoomIn(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("zoomIn:"), sender)
}/* debug [instance_methods/method]: ZoomIn */


// Zooms the image out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/zoomOut(_:)
func (i_ IKImageView) ZoomOut(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("zoomOut:"), sender)
}/* debug [instance_methods/method]: ZoomOut */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKImageView */

// Specifies the automatic-hiding scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/autohidesScrollers
func (i_ IKImageView) AutohidesScrollers() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("autohidesScrollers"))
	return rv
}/* debug [instance_properties/getter]: autohidesScrollers */


// Specifies the automatic-hiding scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/autohidesScrollers
func (i_ IKImageView) SetAutohidesScrollers(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutohidesScrollers:"), value)
}/* debug [instance_properties/setter]: autohidesScrollers */


// Specifies the automatic resizing state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/autoresizes
func (i_ IKImageView) Autoresizes() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("autoresizes"))
	return rv
}/* debug [instance_properties/getter]: autoresizes */


// Specifies the automatic resizing state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/autoresizes
func (i_ IKImageView) SetAutoresizes(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutoresizes:"), value)
}/* debug [instance_properties/setter]: autoresizes */


// Specifies the background color for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/backgroundColor
func (i_ IKImageView) BackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// Specifies the background color for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/backgroundColor
func (i_ IKImageView) SetBackgroundColor(value appkit.Color) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// Specifies the current tool mode for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/currentToolMode
func (i_ IKImageView) CurrentToolMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("currentToolMode"))
	return rv
}/* debug [instance_properties/getter]: currentToolMode */


// Specifies the current tool mode for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/currentToolMode
func (i_ IKImageView) SetCurrentToolMode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrentToolMode:"), value)
}/* debug [instance_properties/setter]: currentToolMode */


// Specifies the delegate object of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/delegate
func (i_ IKImageView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Specifies the delegate object of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/delegate
func (i_ IKImageView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Specifies the image-opening state of the editing pane in the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/doubleClickOpensImageEditPanel
func (i_ IKImageView) DoubleClickOpensImageEditPanel() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("doubleClickOpensImageEditPanel"))
	return rv
}/* debug [instance_properties/getter]: doubleClickOpensImageEditPanel */


// Specifies the image-opening state of the editing pane in the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/doubleClickOpensImageEditPanel
func (i_ IKImageView) SetDoubleClickOpensImageEditPanel(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDoubleClickOpensImageEditPanel:"), value)
}/* debug [instance_properties/setter]: doubleClickOpensImageEditPanel */


// Specifies the editable state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/editable
func (i_ IKImageView) Editable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("editable"))
	return rv
}/* debug [instance_properties/getter]: editable */


// Specifies the editable state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/editable
func (i_ IKImageView) SetEditable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEditable:"), value)
}/* debug [instance_properties/setter]: editable */


// Specifies the horizontal scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/hasHorizontalScroller
func (i_ IKImageView) HasHorizontalScroller() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}/* debug [instance_properties/getter]: hasHorizontalScroller */


// Specifies the horizontal scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/hasHorizontalScroller
func (i_ IKImageView) SetHasHorizontalScroller(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasHorizontalScroller:"), value)
}/* debug [instance_properties/setter]: hasHorizontalScroller */


// Specifies the vertical scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/hasVerticalScroller
func (i_ IKImageView) HasVerticalScroller() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}/* debug [instance_properties/getter]: hasVerticalScroller */


// Specifies the vertical scroll bar state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/hasVerticalScroller
func (i_ IKImageView) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasVerticalScroller:"), value)
}/* debug [instance_properties/setter]: hasVerticalScroller */


// Specifies a Core Image filter for image correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/imageCorrection
func (i_ IKImageView) ImageCorrection() coreimage.Filter {
	rv := objc.Send[coreimage.Filter](i_.ID, objc.Sel("imageCorrection"))
	return rv
}/* debug [instance_properties/getter]: imageCorrection */


// Specifies a Core Image filter for image correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/imageCorrection
func (i_ IKImageView) SetImageCorrection(value coreimage.Filter) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageCorrection:"), value)
}/* debug [instance_properties/setter]: imageCorrection */


// Specifies the rotation angle for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/rotationAngle
func (i_ IKImageView) RotationAngle() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("rotationAngle"))
	return rv
}/* debug [instance_properties/getter]: rotationAngle */


// Specifies the rotation angle for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/rotationAngle
func (i_ IKImageView) SetRotationAngle(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRotationAngle:"), value)
}/* debug [instance_properties/setter]: rotationAngle */


// Specifies the drag-and-drop support state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/supportsDragAndDrop
func (i_ IKImageView) SupportsDragAndDrop() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("supportsDragAndDrop"))
	return rv
}/* debug [instance_properties/getter]: supportsDragAndDrop */


// Specifies the drag-and-drop support state for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/supportsDragAndDrop
func (i_ IKImageView) SetSupportsDragAndDrop(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportsDragAndDrop:"), value)
}/* debug [instance_properties/setter]: supportsDragAndDrop */


// Specifies the zoom factor for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/zoomFactor
func (i_ IKImageView) ZoomFactor() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("zoomFactor"))
	return rv
}/* debug [instance_properties/getter]: zoomFactor */


// Specifies the zoom factor for the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageView/zoomFactor
func (i_ IKImageView) SetZoomFactor(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZoomFactor:"), value)
}/* debug [instance_properties/setter]: zoomFactor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKImageView */



