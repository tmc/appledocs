// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CollectionViewCompositionalLayout] class.
var (
	CollectionViewCompositionalLayoutClass     _CollectionViewCompositionalLayoutClass
	CollectionViewCompositionalLayoutClassOnce sync.Once
)

func getCollectionViewCompositionalLayoutClass() _CollectionViewCompositionalLayoutClass {
	CollectionViewCompositionalLayoutClassOnce.Do(func() {
		CollectionViewCompositionalLayoutClass = _CollectionViewCompositionalLayoutClass{objc.GetClass("NSCollectionViewCompositionalLayout")}
	})
	return CollectionViewCompositionalLayoutClass
}

type _CollectionViewCompositionalLayoutClass struct {
	class objc.Class
}





// An interface definition for the [CollectionViewCompositionalLayout] class.
type ICollectionViewCompositionalLayout interface {
	ICollectionViewLayout
	

	// properties:
	Configuration() ICollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionViewCompositionalLayoutClass) Alloc() CollectionViewCompositionalLayout {
	rv := objc.Send[CollectionViewCompositionalLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewCompositionalLayoutClass) New() CollectionViewCompositionalLayout {
	rv := objc.Send[CollectionViewCompositionalLayout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewCompositionalLayout) Init() CollectionViewCompositionalLayout {
	rv := objc.Send[CollectionViewCompositionalLayout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewCompositionalLayout) Autorelease() CollectionViewCompositionalLayout {
	rv := objc.Send[CollectionViewCompositionalLayout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewCompositionalLayout creates a new CollectionViewCompositionalLayout instance.
func NewCollectionViewCompositionalLayout() CollectionViewCompositionalLayout {
	return getCollectionViewCompositionalLayoutClass().New()
}





// A layout object that lets you combine items in highly adaptive and flexible visual arrangements.
//
// A compositional layout is a type of collection view layout. It’s designed to be composable, flexible, and fast, letting you build any kind of visual arrangement for your content by combining—or compositing—each smaller component into a full layout. A compositional layout is composed of one or more sections that break up the layout into distinct visual groupings. Each section is composed of groups of individual items, the smallest unit of data you want to present. A group might lay out its items in a horizontal row, a vertical column, or a custom arrangement. You combine the components by building up from items into a group, from groups into a section, and finally into a full layout, like in this example of a basic list layout:


// A layout object that lets you combine items in highly adaptive and flexible visual arrangements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout
type CollectionViewCompositionalLayout struct {
	CollectionViewLayout
}

// CollectionViewCompositionalLayoutFrom constructs a [CollectionViewCompositionalLayout] from an unsafe.Pointer.
//
// A layout object that lets you combine items in highly adaptive and flexible visual arrangements.
func CollectionViewCompositionalLayoutFrom(ptr unsafe.Pointer) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}






// Creates a compositional layout object with a single section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(section:)
func NewCollectionViewCompositionalLayoutWithSection(section ICollectionLayoutSection) CollectionViewCompositionalLayout {
	instance := getCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[CollectionViewCompositionalLayout](instance.ID, objc.Sel("initWithSection:"), section)
	rv.Autorelease()
	return rv
}


// Creates a compositional layout object with a single section and an additional configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(section:configuration:)
func NewCollectionViewCompositionalLayoutWithSectionConfiguration(section ICollectionLayoutSection, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	instance := getCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[CollectionViewCompositionalLayout](instance.ID, objc.Sel("initWithSection:configuration:"), section, configuration)
	rv.Autorelease()
	return rv
}


// Creates a compositional layout object with a section provider to supply the layout’s sections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(sectionProvider:)
func NewCollectionViewCompositionalLayoutWithSectionProvider(sectionProvider CollectionViewCompositionalLayoutSectionProvider /* not a class type */) CollectionViewCompositionalLayout {
	instance := getCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[CollectionViewCompositionalLayout](instance.ID, objc.Sel("initWithSectionProvider:"), sectionProvider)
	rv.Autorelease()
	return rv
}


// Creates a compositional layout object with a section provider and an additional configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(sectionProvider:configuration:)
func NewCollectionViewCompositionalLayoutWithSectionProviderConfiguration(sectionProvider CollectionViewCompositionalLayoutSectionProvider /* not a class type */, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	instance := getCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[CollectionViewCompositionalLayout](instance.ID, objc.Sel("initWithSectionProvider:configuration:"), sectionProvider, configuration)
	rv.Autorelease()
	return rv
}






















// The layout’s configuration, such as its scroll direction and section spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/configuration
func (c_ CollectionViewCompositionalLayout) Configuration() ICollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[CollectionViewCompositionalLayoutConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}


// The layout’s configuration, such as its scroll direction and section spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/configuration
func (c_ CollectionViewCompositionalLayout) SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}







