// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutSection */


/* debug [class_header]: Header for NSCollectionLayoutSection */
// The class instance for the [CollectionLayoutSection] class.
var (
	CollectionLayoutSectionClass     _CollectionLayoutSectionClass
	CollectionLayoutSectionClassOnce sync.Once
)

func getCollectionLayoutSectionClass() _CollectionLayoutSectionClass {
	CollectionLayoutSectionClassOnce.Do(func() {
		CollectionLayoutSectionClass = _CollectionLayoutSectionClass{objc.GetClass("NSCollectionLayoutSection")}
	})
	return CollectionLayoutSectionClass
}

type _CollectionLayoutSectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutSection */
// An interface definition for the [CollectionLayoutSection] class.
type ICollectionLayoutSection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionLayoutSection */
	// properties:
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem)
	ContentInsets() objc.IObject /* cross-framework: DirectionalEdgeInsets */
	SetContentInsets(value objc.IObject /* cross-framework: DirectionalEdgeInsets */)
	DecorationItems() []CollectionLayoutDecorationItem
	SetDecorationItems(value []CollectionLayoutDecorationItem)
	InterGroupSpacing() float64
	SetInterGroupSpacing(value float64)
	OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior
	SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior)
	SupplementariesFollowContentInsets() bool
	SetSupplementariesFollowContentInsets(value bool)
	VisibleItemsInvalidationHandler() CollectionLayoutSectionVisibleItemsInvalidationHandler /* not a class type */
	SetVisibleItemsInvalidationHandler(value CollectionLayoutSectionVisibleItemsInvalidationHandler /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutSection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutSection */
// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutSectionClass) Alloc() CollectionLayoutSection {
	rv := objc.Send[CollectionLayoutSection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionLayoutSectionClass) New() CollectionLayoutSection {
	rv := objc.Send[CollectionLayoutSection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutSection) Init() CollectionLayoutSection {
	rv := objc.Send[CollectionLayoutSection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutSection) Autorelease() CollectionLayoutSection {
	rv := objc.Send[CollectionLayoutSection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutSection creates a new CollectionLayoutSection instance.
func NewCollectionLayoutSection() CollectionLayoutSection {
	return getCollectionLayoutSectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutSection */
// A container that combines a set of groups into distinct visual groupings.
//
// A collection view layout has one or more sections. Sections provide a way to separate the layout into distinct pieces. Each section can have the same layout or a different layout than the other sections in the collection view. A section’s layout is determined by the properties of the group ( ) that’s used to create the section. In the Photos app, each section in the Years page uses the same layout. In the App Store, the Apps page displays several sections with different content arrangements. Each section can have its own background, header, and footer to distinguish it from other sections.


// A container that combines a set of groups into distinct visual groupings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection
type CollectionLayoutSection struct {
	objectivec.Object
}

// CollectionLayoutSectionFrom constructs a [CollectionLayoutSection] from an unsafe.Pointer.
//
// A container that combines a set of groups into distinct visual groupings.
func CollectionLayoutSectionFrom(ptr unsafe.Pointer) CollectionLayoutSection {
	return CollectionLayoutSection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutSection */

// Creates a section containing the specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/init(group:)
func NewCollectionLayoutSectionWithGroup(group ICollectionLayoutGroup) CollectionLayoutSection {
	rv := objc.Send[CollectionLayoutSection](objc.ID(getCollectionLayoutSectionClass().class), objc.Sel("sectionWithGroup:"), group)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutSectionWithGroup */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutSection */

// Creates a section containing the specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/init(group:)
func (cc _CollectionLayoutSectionClass) SectionWithGroup(group ICollectionLayoutGroup) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("sectionWithGroup:"), group)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SectionWithGroup) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutSection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutSection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutSection */

// An array of the supplementary items that are associated with the boundary edges of the section, such as headers and footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/boundarySupplementaryItems
func (c_ CollectionLayoutSection) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[[]CollectionLayoutBoundarySupplementaryItem](c_.ID, objc.Sel("boundarySupplementaryItems"))
	return rv
}/* debug [instance_properties/getter]: boundarySupplementaryItems */


// An array of the supplementary items that are associated with the boundary edges of the section, such as headers and footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/boundarySupplementaryItems
func (c_ CollectionLayoutSection) SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setBoundarySupplementaryItems:"), nsArray)
}/* debug [instance_properties/setter]: boundarySupplementaryItems */


// The amount of space between the content of the section and its boundaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/contentInsets
func (c_ CollectionLayoutSection) ContentInsets() objc.IObject /* cross-framework: DirectionalEdgeInsets */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("contentInsets"))
	return rv
}/* debug [instance_properties/getter]: contentInsets */


// The amount of space between the content of the section and its boundaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/contentInsets
func (c_ CollectionLayoutSection) SetContentInsets(value objc.IObject /* cross-framework: DirectionalEdgeInsets */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentInsets:"), value)
}/* debug [instance_properties/setter]: contentInsets */


// An array of the decoration items that are anchored to the section, such as background decoration views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/decorationItems
func (c_ CollectionLayoutSection) DecorationItems() []CollectionLayoutDecorationItem {
	rv := objc.Send[[]CollectionLayoutDecorationItem](c_.ID, objc.Sel("decorationItems"))
	return rv
}/* debug [instance_properties/getter]: decorationItems */


// An array of the decoration items that are anchored to the section, such as background decoration views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/decorationItems
func (c_ CollectionLayoutSection) SetDecorationItems(value []CollectionLayoutDecorationItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setDecorationItems:"), nsArray)
}/* debug [instance_properties/setter]: decorationItems */


// The amount of space between the groups in the section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/interGroupSpacing
func (c_ CollectionLayoutSection) InterGroupSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("interGroupSpacing"))
	return rv
}/* debug [instance_properties/getter]: interGroupSpacing */


// The amount of space between the groups in the section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/interGroupSpacing
func (c_ CollectionLayoutSection) SetInterGroupSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInterGroupSpacing:"), value)
}/* debug [instance_properties/setter]: interGroupSpacing */


// The section’s scrolling behavior in relation to the main layout axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/orthogonalScrollingBehavior
func (c_ CollectionLayoutSection) OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior {
	rv := objc.Send[CollectionLayoutSectionOrthogonalScrollingBehavior](c_.ID, objc.Sel("orthogonalScrollingBehavior"))
	return rv
}/* debug [instance_properties/getter]: orthogonalScrollingBehavior */


// The section’s scrolling behavior in relation to the main layout axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/orthogonalScrollingBehavior
func (c_ CollectionLayoutSection) SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrthogonalScrollingBehavior:"), value)
}/* debug [instance_properties/setter]: orthogonalScrollingBehavior */


// A Boolean value that indicates whether the section’s supplementary items follow the specified content insets for the section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/supplementariesFollowContentInsets
func (c_ CollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supplementariesFollowContentInsets"))
	return rv
}/* debug [instance_properties/getter]: supplementariesFollowContentInsets */


// A Boolean value that indicates whether the section’s supplementary items follow the specified content insets for the section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/supplementariesFollowContentInsets
func (c_ CollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupplementariesFollowContentInsets:"), value)
}/* debug [instance_properties/setter]: supplementariesFollowContentInsets */


// A closure called before each layout cycle to allow modification of the items in the section immediately before they’re displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/visibleItemsInvalidationHandler
func (c_ CollectionLayoutSection) VisibleItemsInvalidationHandler() CollectionLayoutSectionVisibleItemsInvalidationHandler /* not a class type */ {
	rv := objc.Send[CollectionLayoutSectionVisibleItemsInvalidationHandler](c_.ID, objc.Sel("visibleItemsInvalidationHandler"))
	return rv
}/* debug [instance_properties/getter]: visibleItemsInvalidationHandler */


// A closure called before each layout cycle to allow modification of the items in the section immediately before they’re displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/visibleItemsInvalidationHandler
func (c_ CollectionLayoutSection) SetVisibleItemsInvalidationHandler(value CollectionLayoutSectionVisibleItemsInvalidationHandler /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVisibleItemsInvalidationHandler:"), value)
}/* debug [instance_properties/setter]: visibleItemsInvalidationHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutSection */


