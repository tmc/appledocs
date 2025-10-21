// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
}

// The visual representation of one of your annotation objects.
//
// are loosely coupled to a corresponding , which is an object that conforms to the protocol. When an annotation’s coordinate point is in the map’s visible region, the map view asks its delegate to provide a corresponding annotation view. MapKit may recycle annotation views and put them into a reuse queue that the map view maintains. The most efficient way to provide the content for an annotation view is to set its property. The annotation view sizes itself automatically to the image you specify and draws that image for its contents. Because it’s a view, you can also override the method and draw your view’s content manually. If you choose to override directly and you don’t specify a custom image in the property, the annotation view sets the width and height of the annotation view’s frame to by default. Before the framework can draw your custom content, you need to set the width and height to nonzero values by modifying the view’s property. In general, if your content consists entirely of static images, it’s more efficient to set the property and change it as necessary than to draw the images yourself. Annotation views anchor to the map at the point that their associated annotation object specifies. Although they scroll with the map contents, annotation views reside in a separate display layer and don’t scale when the size of the visible map region changes. Additionally, annotation views support the concept of a , which determines whether the map displays the annotation view as unselected, selected, or selected and displaying a standard callout view. The user toggles between the selection states through interactions with the annotation view. In the unselected state, the map displays the annotation view, but doesn’t highlight it. In the selected state, the framework highlights the annotation, but doesn’t display the callout. Finally, the map view can display the annotation with both a highlight and a callout. The callout view displays additional information, such as a title string and controls for viewing more information. The annotation object provides the title information, but your annotation view is responsible for providing any custom controls. For more information, see the section below.
//
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


// An identifier that determines whether the annotation view participates in clustering.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/clusteringIdentifier
func (m_ MKAnnotationView) ClusteringIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("clusteringIdentifier"))
	return rv
}


// SetClusteringIdentifier sets the value of the clusteringIdentifier property.
// An identifier that determines whether the annotation view participates in clustering.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/clusteringIdentifier
func (m_ MKAnnotationView) SetClusteringIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClusteringIdentifier:"), objc.String(value))
}



