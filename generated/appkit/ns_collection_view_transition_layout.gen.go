// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionViewTransitionLayout */


/* debug [class_header]: Header for NSCollectionViewTransitionLayout */
// The class instance for the [CollectionViewTransitionLayout] class.
var (
	CollectionViewTransitionLayoutClass     _CollectionViewTransitionLayoutClass
	CollectionViewTransitionLayoutClassOnce sync.Once
)

func getCollectionViewTransitionLayoutClass() _CollectionViewTransitionLayoutClass {
	CollectionViewTransitionLayoutClassOnce.Do(func() {
		CollectionViewTransitionLayoutClass = _CollectionViewTransitionLayoutClass{objc.GetClass("NSCollectionViewTransitionLayout")}
	})
	return CollectionViewTransitionLayoutClass
}

type _CollectionViewTransitionLayoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewTransitionLayout */
// An interface definition for the [CollectionViewTransitionLayout] class.
type ICollectionViewTransitionLayout interface {
	ICollectionViewLayout
	
/* debug [class_interface_properties]: Properties for CollectionViewTransitionLayout */
	// properties:
	CurrentLayout() ICollectionViewLayout
	NextLayout() ICollectionViewLayout
	TransitionProgress() float64
	SetTransitionProgress(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewTransitionLayout */
	// methods:
	UpdateValueForAnimatedKey(value float64, key CollectionViewTransitionLayoutAnimatedKey /* typedef */)
	ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey /* typedef */) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewTransitionLayout */
// Alloc allocates a new instance without initialization.
func (cc _CollectionViewTransitionLayoutClass) Alloc() CollectionViewTransitionLayout {
	rv := objc.Send[CollectionViewTransitionLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewTransitionLayoutClass) New() CollectionViewTransitionLayout {
	rv := objc.Send[CollectionViewTransitionLayout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewTransitionLayout) Init() CollectionViewTransitionLayout {
	rv := objc.Send[CollectionViewTransitionLayout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewTransitionLayout) Autorelease() CollectionViewTransitionLayout {
	rv := objc.Send[CollectionViewTransitionLayout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewTransitionLayout creates a new CollectionViewTransitionLayout instance.
func NewCollectionViewTransitionLayout() CollectionViewTransitionLayout {
	return getCollectionViewTransitionLayoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewTransitionLayout */
// An object that implements custom behaviors when changing from one layout to another in a collection view.
//
// Transition layout objects are commonly used to implement interactive transitions between layouts, where the transition itself is driven by a gesture recognizer.


// An object that implements custom behaviors when changing from one layout to another in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout
type CollectionViewTransitionLayout struct {
	CollectionViewLayout
}

// CollectionViewTransitionLayoutFrom constructs a [CollectionViewTransitionLayout] from an unsafe.Pointer.
//
// An object that implements custom behaviors when changing from one layout to another in a collection view.
func CollectionViewTransitionLayoutFrom(ptr unsafe.Pointer) CollectionViewTransitionLayout {
	return CollectionViewTransitionLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewTransitionLayout */

// Initializes and returns the transition layout object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/init(currentLayout:nextLayout:)
func NewCollectionViewTransitionLayoutWithCurrentLayoutNextLayout(currentLayout ICollectionViewLayout, newLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	instance := getCollectionViewTransitionLayoutClass().Alloc()
	rv := objc.Send[CollectionViewTransitionLayout](instance.ID, objc.Sel("initWithCurrentLayout:nextLayout:"), currentLayout, newLayout)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionViewTransitionLayoutWithCurrentLayoutNextLayout */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewTransitionLayout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewTransitionLayout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewTransitionLayout */

// Sets the value of a key whose value you use during the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/updateValue(_:forAnimatedKey:)
func (c_ CollectionViewTransitionLayout) UpdateValueForAnimatedKey(value float64, key CollectionViewTransitionLayoutAnimatedKey /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateValue:forAnimatedKey:"), value, key)
}/* debug [instance_methods/method]: UpdateValueForAnimatedKey */


// Returns the most recently set value for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/value(forAnimatedKey:)
func (c_ CollectionViewTransitionLayout) ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey /* typedef */) float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("valueForAnimatedKey:"), key)
	return rv
}/* debug [instance_methods/method]: ValueForAnimatedKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewTransitionLayout */

// The collection view’s current layout object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/currentLayout
func (c_ CollectionViewTransitionLayout) CurrentLayout() ICollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("currentLayout"))
	return rv
}/* debug [instance_properties/getter]: currentLayout */


// The collection view’s new layout object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/nextLayout
func (c_ CollectionViewTransitionLayout) NextLayout() ICollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("nextLayout"))
	return rv
}/* debug [instance_properties/getter]: nextLayout */


// The completion percentage of the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/transitionProgress
func (c_ CollectionViewTransitionLayout) TransitionProgress() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("transitionProgress"))
	return rv
}/* debug [instance_properties/getter]: transitionProgress */


// The completion percentage of the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/transitionProgress
func (c_ CollectionViewTransitionLayout) SetTransitionProgress(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTransitionProgress:"), value)
}/* debug [instance_properties/setter]: transitionProgress */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewTransitionLayout */


