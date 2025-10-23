// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [MKAnnotationView] class.
var (
	MKAnnotationViewClass     _MKAnnotationViewClass
	MKAnnotationViewClassOnce sync.Once
)

func getMKAnnotationViewClass() _MKAnnotationViewClass {
	MKAnnotationViewClassOnce.Do(func() {
		MKAnnotationViewClass = _MKAnnotationViewClass{objc.GetClass("MKAnnotationView")}
	})
	return MKAnnotationViewClass
}

type _MKAnnotationViewClass struct {
	class objc.Class
}

// An interface definition for the [MKAnnotationView] class.
type IMKAnnotationView interface {
	appkit.IView
	ClusterAnnotationView() MKAnnotationView
	ClusteringIdentifier() string
	SetClusteringIdentifier(value string)
	AccessoryOffset() coregraphics.CGPoint
	SetAccessoryOffset(value coregraphics.CGPoint)
	Annotation() unsafe.Pointer
	SetAnnotation(value unsafe.Pointer)
	CalloutOffset() coregraphics.CGPoint
	SetCalloutOffset(value coregraphics.CGPoint)
	CanShowCallout() bool
	SetCanShowCallout(value bool)
	CenterOffset() coregraphics.CGPoint
	SetCenterOffset(value coregraphics.CGPoint)
	Cluster() MKAnnotationView
	SetCluster(value IMKAnnotationView)
	CollisionMode() unsafe.Pointer
	SetCollisionMode(value unsafe.Pointer)
	DetailCalloutAccessoryView() appkit.View
	SetDetailCalloutAccessoryView(value appkit.IView)
	DisplayPriority() unsafe.Pointer
	SetDisplayPriority(value unsafe.Pointer)
	DragState() unsafe.Pointer
	SetDragState(value unsafe.Pointer)
	Image() appkit.Image
	SetImage(value appkit.IImage)
	IsDraggable() bool
	SetIsDraggable(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	IsSelected() bool
	SetIsSelected(value bool)
	LeftCalloutAccessoryView() appkit.View
	SetLeftCalloutAccessoryView(value appkit.IView)
	LeftCalloutOffset() coregraphics.CGPoint
	SetLeftCalloutOffset(value coregraphics.CGPoint)
	ReuseIdentifier() string
	SetReuseIdentifier(value string)
	RightCalloutAccessoryView() appkit.View
	SetRightCalloutAccessoryView(value appkit.IView)
	RightCalloutOffset() coregraphics.CGPoint
	SetRightCalloutOffset(value coregraphics.CGPoint)
	SelectedZPriority() unsafe.Pointer
	SetSelectedZPriority(value unsafe.Pointer)
	ZPriority() unsafe.Pointer
	SetZPriority(value unsafe.Pointer)
}

// The visual representation of one of your annotation objects.
//
// are loosely coupled to a corresponding , which is an object that conforms to the protocol. When an annotation’s coordinate point is in the map’s visible region, the map view asks its delegate to provide a corresponding annotation view. MapKit may recycle annotation views and put them into a reuse queue that the map view maintains. The most efficient way to provide the content for an annotation view is to set its property. The annotation view sizes itself automatically to the image you specify and draws that image for its contents. Because it’s a view, you can also override the method and draw your view’s content manually. If you choose to override directly and you don’t specify a custom image in the property, the annotation view sets the width and height of the annotation view’s frame to by default. Before the framework can draw your custom content, you need to set the width and height to nonzero values by modifying the view’s property. In general, if your content consists entirely of static images, it’s more efficient to set the property and change it as necessary than to draw the images yourself. Annotation views anchor to the map at the point that their associated annotation object specifies. Although they scroll with the map contents, annotation views reside in a separate display layer and don’t scale when the size of the visible map region changes. Additionally, annotation views support the concept of a , which determines whether the map displays the annotation view as unselected, selected, or selected and displaying a standard callout view. The user toggles between the selection states through interactions with the annotation view. In the unselected state, the map displays the annotation view, but doesn’t highlight it. In the selected state, the framework highlights the annotation, but doesn’t display the callout. Finally, the map view can display the annotation with both a highlight and a callout. The callout view displays additional information, such as a title string and controls for viewing more information. The annotation object provides the title information, but your annotation view is responsible for providing any custom controls. For more information, see the section below.


// The visual representation of one of your annotation objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView
type MKAnnotationView struct {
	appkit.View
}

// MKAnnotationViewFrom constructs a [MKAnnotationView] from an unsafe.Pointer.
//
// The visual representation of one of your annotation objects.
func MKAnnotationViewFrom(ptr unsafe.Pointer) MKAnnotationView {
	return MKAnnotationView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKAnnotationViewClass) Alloc() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKAnnotationViewClass) New() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKAnnotationView) Init() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKAnnotationView) Autorelease() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKAnnotationView creates a new MKAnnotationView instance.
func NewMKAnnotationView() MKAnnotationView {
	return getMKAnnotationViewClass().New()
}



// The clustering annotation view that replaces the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/cluster
func (m_ MKAnnotationView) ClusterAnnotationView() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](m_.ID, objc.Sel("clusterAnnotationView"))
	return rv
}


// An identifier that determines whether the annotation view participates in clustering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/clusteringIdentifier
func (m_ MKAnnotationView) ClusteringIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("clusteringIdentifier"))
	return rv
}


// An identifier that determines whether the annotation view participates in clustering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/clusteringIdentifier
func (m_ MKAnnotationView) SetClusteringIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClusteringIdentifier:"), objc.String(value))
}


// An offset that changes the accessory’s default anchor point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/accessoryoffset
func (m_ MKAnnotationView) AccessoryOffset() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("accessoryOffset"))
	return rv
}


// An offset that changes the accessory’s default anchor point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/accessoryoffset
func (m_ MKAnnotationView) SetAccessoryOffset(value coregraphics.CGPoint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccessoryOffset:"), value)
}


// The annotation object associated with the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/annotation
func (m_ MKAnnotationView) Annotation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("annotation"))
	return rv
}


// The annotation object associated with the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/annotation
func (m_ MKAnnotationView) SetAnnotation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnnotation:"), value)
}


// The offset (in points) at which to place the callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/calloutoffset
func (m_ MKAnnotationView) CalloutOffset() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("calloutOffset"))
	return rv
}


// The offset (in points) at which to place the callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/calloutoffset
func (m_ MKAnnotationView) SetCalloutOffset(value coregraphics.CGPoint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCalloutOffset:"), value)
}


// A Boolean value that indicates whether the annotation view is able to display extra information in a callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/canshowcallout
func (m_ MKAnnotationView) CanShowCallout() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canShowCallout"))
	return rv
}


// A Boolean value that indicates whether the annotation view is able to display extra information in a callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/canshowcallout
func (m_ MKAnnotationView) SetCanShowCallout(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCanShowCallout:"), value)
}


// The offset (in points) at which to display the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/centeroffset
func (m_ MKAnnotationView) CenterOffset() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("centerOffset"))
	return rv
}


// The offset (in points) at which to display the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/centeroffset
func (m_ MKAnnotationView) SetCenterOffset(value coregraphics.CGPoint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterOffset:"), value)
}


// The clustering annotation view that replaces the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/cluster
func (m_ MKAnnotationView) Cluster() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](m_.ID, objc.Sel("cluster"))
	return rv
}


// The clustering annotation view that replaces the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/cluster
func (m_ MKAnnotationView) SetCluster(value IMKAnnotationView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}


// The collision mode to use when interpreting the collision frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/collisionmode-swift.property
func (m_ MKAnnotationView) CollisionMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("collisionMode"))
	return rv
}


// The collision mode to use when interpreting the collision frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/collisionmode-swift.property
func (m_ MKAnnotationView) SetCollisionMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCollisionMode:"), value)
}


// The detail accessory view to use in the standard callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/detailcalloutaccessoryview
func (m_ MKAnnotationView) DetailCalloutAccessoryView() appkit.View {
	rv := objc.Send[appkit.View](m_.ID, objc.Sel("detailCalloutAccessoryView"))
	return rv
}


// The detail accessory view to use in the standard callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/detailcalloutaccessoryview
func (m_ MKAnnotationView) SetDetailCalloutAccessoryView(value appkit.IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDetailCalloutAccessoryView:"), value)
}


// The display priority of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/displaypriority
func (m_ MKAnnotationView) DisplayPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("displayPriority"))
	return rv
}


// The display priority of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/displaypriority
func (m_ MKAnnotationView) SetDisplayPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplayPriority:"), value)
}


// The drag state of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/dragstate-swift.property
func (m_ MKAnnotationView) DragState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dragState"))
	return rv
}


// The drag state of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/dragstate-swift.property
func (m_ MKAnnotationView) SetDragState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDragState:"), value)
}


// The image the annotation view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/image
func (m_ MKAnnotationView) Image() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("image"))
	return rv
}


// The image the annotation view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/image
func (m_ MKAnnotationView) SetImage(value appkit.IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImage:"), value)
}


// A Boolean value that indicates whether the annotation view is draggable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/isdraggable
func (m_ MKAnnotationView) IsDraggable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isDraggable"))
	return rv
}


// A Boolean value that indicates whether the annotation view is draggable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/isdraggable
func (m_ MKAnnotationView) SetIsDraggable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDraggable:"), value)
}


// A Boolean value that indicates whether the annotation is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/isenabled
func (m_ MKAnnotationView) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the annotation is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/isenabled
func (m_ MKAnnotationView) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the map view highlights the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/ishighlighted
func (m_ MKAnnotationView) IsHighlighted() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isHighlighted"))
	return rv
}


// A Boolean value that indicates whether the map view highlights the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/ishighlighted
func (m_ MKAnnotationView) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsHighlighted:"), value)
}


// A Boolean value that indicates whether the annotation view is in a selected state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/isselected
func (m_ MKAnnotationView) IsSelected() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSelected"))
	return rv
}


// A Boolean value that indicates whether the annotation view is in a selected state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/isselected
func (m_ MKAnnotationView) SetIsSelected(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSelected:"), value)
}


// The view to display on the left side of the standard callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/leftcalloutaccessoryview
func (m_ MKAnnotationView) LeftCalloutAccessoryView() appkit.View {
	rv := objc.Send[appkit.View](m_.ID, objc.Sel("leftCalloutAccessoryView"))
	return rv
}


// The view to display on the left side of the standard callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/leftcalloutaccessoryview
func (m_ MKAnnotationView) SetLeftCalloutAccessoryView(value appkit.IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLeftCalloutAccessoryView:"), value)
}


// The offset in points from the middle-left of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/leftcalloutoffset
func (m_ MKAnnotationView) LeftCalloutOffset() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("leftCalloutOffset"))
	return rv
}


// The offset in points from the middle-left of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/leftcalloutoffset
func (m_ MKAnnotationView) SetLeftCalloutOffset(value coregraphics.CGPoint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLeftCalloutOffset:"), value)
}


// The string that identifies that the annotation view is reusable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/reuseidentifier
func (m_ MKAnnotationView) ReuseIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("reuseIdentifier"))
	return rv
}


// The string that identifies that the annotation view is reusable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/reuseidentifier
func (m_ MKAnnotationView) SetReuseIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReuseIdentifier:"), objc.String(value))
}


// The view to display on the right side of the standard callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/rightcalloutaccessoryview
func (m_ MKAnnotationView) RightCalloutAccessoryView() appkit.View {
	rv := objc.Send[appkit.View](m_.ID, objc.Sel("rightCalloutAccessoryView"))
	return rv
}


// The view to display on the right side of the standard callout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/rightcalloutaccessoryview
func (m_ MKAnnotationView) SetRightCalloutAccessoryView(value appkit.IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRightCalloutAccessoryView:"), value)
}


// The offset in points from the middle-right of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/rightcalloutoffset
func (m_ MKAnnotationView) RightCalloutOffset() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("rightCalloutOffset"))
	return rv
}


// The offset in points from the middle-right of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/rightcalloutoffset
func (m_ MKAnnotationView) SetRightCalloutOffset(value coregraphics.CGPoint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRightCalloutOffset:"), value)
}


// The relative importance of the annotation view when in a selected state with respect to its ordering along the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/selectedzpriority
func (m_ MKAnnotationView) SelectedZPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("selectedZPriority"))
	return rv
}


// The relative importance of the annotation view when in a selected state with respect to its ordering along the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/selectedzpriority
func (m_ MKAnnotationView) SetSelectedZPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedZPriority:"), value)
}


// The relative importance of the annotation view when in an unselected state with respect to its ordering along the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/zpriority
func (m_ MKAnnotationView) ZPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("zPriority"))
	return rv
}


// The relative importance of the annotation view when in an unselected state with respect to its ordering along the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/zpriority
func (m_ MKAnnotationView) SetZPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setZPriority:"), value)
}



