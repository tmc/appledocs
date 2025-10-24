// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMediaQuerySection */


/* debug [class_header]: Header for MPMediaQuerySection */
// The class instance for the [MediaQuerySection] class.
var (
	MediaQuerySectionClass     _MediaQuerySectionClass
	MediaQuerySectionClassOnce sync.Once
)

func getMediaQuerySectionClass() _MediaQuerySectionClass {
	MediaQuerySectionClassOnce.Do(func() {
		MediaQuerySectionClass = _MediaQuerySectionClass{objc.GetClass("MPMediaQuerySection")}
	})
	return MediaQuerySectionClass
}

type _MediaQuerySectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaQuerySection */
// An interface definition for the [MediaQuerySection] class.
type IMediaQuerySection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaQuerySection */
	// properties:
	CollectionSections() IMPMediaQuerySection
	SetCollectionSections(value IMPMediaQuerySection)
	ItemSections() IMPMediaQuerySection
	SetItemSections(value IMPMediaQuerySection)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaQuerySection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaQuerySection */
// Alloc allocates a new instance without initialization.
func (mc _MediaQuerySectionClass) Alloc() MediaQuerySection {
	rv := objc.Send[MediaQuerySection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaQuerySectionClass) New() MediaQuerySection {
	rv := objc.Send[MediaQuerySection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaQuerySection) Init() MediaQuerySection {
	rv := objc.Send[MediaQuerySection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaQuerySection) Autorelease() MediaQuerySection {
	rv := objc.Send[MediaQuerySection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaQuerySection creates a new MediaQuerySection instance.
func NewMediaQuerySection() MediaQuerySection {
	return getMediaQuerySectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaQuerySection */
// A range of media items or media item collections from within a media query.
//
// You can use sections when displaying a query’s items or collections in your app’s user interface. You obtain an array of media query sections by using the or properties of a media query (an instance of the class). The property values of a media query section are read-only.


// A range of media items or media item collections from within a media query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuerySection
type MediaQuerySection struct {
	objectivec.Object
}

// MediaQuerySectionFrom constructs a [MediaQuerySection] from an unsafe.Pointer.
//
// A range of media items or media item collections from within a media query.
func MediaQuerySectionFrom(ptr unsafe.Pointer) MediaQuerySection {
	return MediaQuerySection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaQuerySection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaQuerySection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaQuerySection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaQuerySection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaQuerySection */

// An array representing the section grouping of the query’s specified media item collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaquery/collectionsections
func (m_ MediaQuerySection) CollectionSections() IMPMediaQuerySection {
	rv := objc.Send[MediaQuerySection](m_.ID, objc.Sel("collectionSections"))
	return rv
}/* debug [instance_properties/getter]: collectionSections */


// An array representing the section grouping of the query’s specified media item collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaquery/collectionsections
func (m_ MediaQuerySection) SetCollectionSections(value IMPMediaQuerySection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCollectionSections:"), value)
}/* debug [instance_properties/setter]: collectionSections */


// An array representing the section grouping of the query’s specified media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaquery/itemsections
func (m_ MediaQuerySection) ItemSections() IMPMediaQuerySection {
	rv := objc.Send[MediaQuerySection](m_.ID, objc.Sel("itemSections"))
	return rv
}/* debug [instance_properties/getter]: itemSections */


// An array representing the section grouping of the query’s specified media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaquery/itemsections
func (m_ MediaQuerySection) SetItemSections(value IMPMediaQuerySection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setItemSections:"), value)
}/* debug [instance_properties/setter]: itemSections */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaQuerySection */


