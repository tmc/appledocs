// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKClusterAnnotation */


/* debug [class_header]: Header for MKClusterAnnotation */
// The class instance for the [MKClusterAnnotation] class.
var (
	MKClusterAnnotationClass     _MKClusterAnnotationClass
	MKClusterAnnotationClassOnce sync.Once
)

func getMKClusterAnnotationClass() _MKClusterAnnotationClass {
	MKClusterAnnotationClassOnce.Do(func() {
		MKClusterAnnotationClass = _MKClusterAnnotationClass{objc.GetClass("MKClusterAnnotation")}
	})
	return MKClusterAnnotationClass
}

type _MKClusterAnnotationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKClusterAnnotation */
// An interface definition for the [MKClusterAnnotation] class.
type IMKClusterAnnotation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKClusterAnnotation */
	// properties:
	MemberAnnotations() []objc.ID
	Subtitle() objc.IObject /* cross-framework: NSString */
	SetSubtitle(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKClusterAnnotation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKClusterAnnotation */
// Alloc allocates a new instance without initialization.
func (mc _MKClusterAnnotationClass) Alloc() MKClusterAnnotation {
	rv := objc.Send[MKClusterAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKClusterAnnotationClass) New() MKClusterAnnotation {
	rv := objc.Send[MKClusterAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKClusterAnnotation) Init() MKClusterAnnotation {
	rv := objc.Send[MKClusterAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKClusterAnnotation) Autorelease() MKClusterAnnotation {
	rv := objc.Send[MKClusterAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKClusterAnnotation creates a new MKClusterAnnotation instance.
func NewMKClusterAnnotation() MKClusterAnnotation {
	return getMKClusterAnnotationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKClusterAnnotation */
// An annotation that groups two or more distinct annotations into a single entity.
//
// A cluster annotation object stands in for the group of annotations. Cluster views promote legibility of the underlying annotations by displaying a single annotation that takes it’s title from one annotation and includes a subtitle that indicates how many additional annotations belong to the group. MapKit automatically creates cluster annotations when two or more annotation views group too closely together on the map surface. To customize the cluster annotations that display on your map, implement the method in your map’s delegate.


// An annotation that groups two or more distinct annotations into a single entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKClusterAnnotation
type MKClusterAnnotation struct {
	objectivec.Object
}

// MKClusterAnnotationFrom constructs a [MKClusterAnnotation] from an unsafe.Pointer.
//
// An annotation that groups two or more distinct annotations into a single entity.
func MKClusterAnnotationFrom(ptr unsafe.Pointer) MKClusterAnnotation {
	return MKClusterAnnotation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKClusterAnnotation */

// Creates a cluster annotation with the specified individual annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKClusterAnnotation/init(memberAnnotations:)
func NewMKClusterAnnotationWithMemberAnnotations(memberAnnotations []objc.ID) MKClusterAnnotation {
	instance := getMKClusterAnnotationClass().Alloc()
	rv := objc.Send[MKClusterAnnotation](instance.ID, objc.Sel("initWithMemberAnnotations:"), memberAnnotations)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKClusterAnnotationWithMemberAnnotations */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKClusterAnnotation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKClusterAnnotation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKClusterAnnotation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKClusterAnnotation */

// The annotations that the cluster contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKClusterAnnotation/memberAnnotations
func (m_ MKClusterAnnotation) MemberAnnotations() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("memberAnnotations"))
	return rv
}/* debug [instance_properties/getter]: memberAnnotations */


// The subtitle string to display for the group of annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKClusterAnnotation/subtitle
func (m_ MKClusterAnnotation) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// The subtitle string to display for the group of annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKClusterAnnotation/subtitle
func (m_ MKClusterAnnotation) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), value)
}/* debug [instance_properties/setter]: subtitle */


// The title string to display for the group of annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKClusterAnnotation/title
func (m_ MKClusterAnnotation) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title string to display for the group of annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKClusterAnnotation/title
func (m_ MKClusterAnnotation) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKClusterAnnotation */


