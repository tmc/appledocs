// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionViewCompositionalLayoutConfiguration */


/* debug [class_header]: Header for NSCollectionViewCompositionalLayoutConfiguration */
// The class instance for the [CollectionViewCompositionalLayoutConfiguration] class.
var (
	CollectionViewCompositionalLayoutConfigurationClass     _CollectionViewCompositionalLayoutConfigurationClass
	CollectionViewCompositionalLayoutConfigurationClassOnce sync.Once
)

func getCollectionViewCompositionalLayoutConfigurationClass() _CollectionViewCompositionalLayoutConfigurationClass {
	CollectionViewCompositionalLayoutConfigurationClassOnce.Do(func() {
		CollectionViewCompositionalLayoutConfigurationClass = _CollectionViewCompositionalLayoutConfigurationClass{objc.GetClass("NSCollectionViewCompositionalLayoutConfiguration")}
	})
	return CollectionViewCompositionalLayoutConfigurationClass
}

type _CollectionViewCompositionalLayoutConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewCompositionalLayoutConfiguration */
// An interface definition for the [CollectionViewCompositionalLayoutConfiguration] class.
type ICollectionViewCompositionalLayoutConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionViewCompositionalLayoutConfiguration */
	// properties:
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem)
	InterSectionSpacing() float64
	SetInterSectionSpacing(value float64)
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	Configuration() ICollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewCompositionalLayoutConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewCompositionalLayoutConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _CollectionViewCompositionalLayoutConfigurationClass) Alloc() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[CollectionViewCompositionalLayoutConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewCompositionalLayoutConfigurationClass) New() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[CollectionViewCompositionalLayoutConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewCompositionalLayoutConfiguration) Init() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[CollectionViewCompositionalLayoutConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewCompositionalLayoutConfiguration) Autorelease() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[CollectionViewCompositionalLayoutConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewCompositionalLayoutConfiguration creates a new CollectionViewCompositionalLayoutConfiguration instance.
func NewCollectionViewCompositionalLayoutConfiguration() CollectionViewCompositionalLayoutConfiguration {
	return getCollectionViewCompositionalLayoutConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewCompositionalLayoutConfiguration */
// An object that defines scroll direction, section spacing, and headers or footers for the layout.
//
// You use a layout configuration to modify a collection view layout’s default scroll direction, add extra spacing between each section of the layout, and add headers or footers to the entire layout. You can pass in this configuration when creating an , or you can set the property on an existing layout. If you modify the configuration on an existing layout, the system invalidates the layout so that it will be updated with the new configuration.


// An object that defines scroll direction, section spacing, and headers or footers for the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration
type CollectionViewCompositionalLayoutConfiguration struct {
	objectivec.Object
}

// CollectionViewCompositionalLayoutConfigurationFrom constructs a [CollectionViewCompositionalLayoutConfiguration] from an unsafe.Pointer.
//
// An object that defines scroll direction, section spacing, and headers or footers for the layout.
func CollectionViewCompositionalLayoutConfigurationFrom(ptr unsafe.Pointer) CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewCompositionalLayoutConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewCompositionalLayoutConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewCompositionalLayoutConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewCompositionalLayoutConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewCompositionalLayoutConfiguration */

// An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/boundarySupplementaryItems
func (c_ CollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[[]CollectionLayoutBoundarySupplementaryItem](c_.ID, objc.Sel("boundarySupplementaryItems"))
	return rv
}/* debug [instance_properties/getter]: boundarySupplementaryItems */


// An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/boundarySupplementaryItems
func (c_ CollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem) {
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


// The amount of space between the sections in the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/interSectionSpacing
func (c_ CollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("interSectionSpacing"))
	return rv
}/* debug [instance_properties/getter]: interSectionSpacing */


// The amount of space between the sections in the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/interSectionSpacing
func (c_ CollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInterSectionSpacing:"), value)
}/* debug [instance_properties/setter]: interSectionSpacing */


// The axis that the content in the collection view layout scrolls along.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/scrollDirection
func (c_ CollectionViewCompositionalLayoutConfiguration) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.Send[CollectionViewScrollDirection](c_.ID, objc.Sel("scrollDirection"))
	return rv
}/* debug [instance_properties/getter]: scrollDirection */


// The axis that the content in the collection view layout scrolls along.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/scrollDirection
func (c_ CollectionViewCompositionalLayoutConfiguration) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScrollDirection:"), value)
}/* debug [instance_properties/setter]: scrollDirection */


// The layout’s configuration, such as its scroll direction and section spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayout/configuration
func (c_ CollectionViewCompositionalLayoutConfiguration) Configuration() ICollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[CollectionViewCompositionalLayoutConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The layout’s configuration, such as its scroll direction and section spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayout/configuration
func (c_ CollectionViewCompositionalLayoutConfiguration) SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewCompositionalLayoutConfiguration */



