// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CollectionViewTransitionLayout] class.
type ICollectionViewTransitionLayout interface {
	ICollectionViewLayout
	// properties:
	CurrentLayout() ICollectionViewLayout
	NextLayout() ICollectionViewLayout
	TransitionProgress() float64
	SetTransitionProgress(value float64)
	// methods:
	UpdateValueForAnimatedKey(value float64, key objc.IObject /* cross-framework: CollectionViewTransitionLayoutAnimatedKey */)
	ValueForAnimatedKey(key objc.IObject /* cross-framework: CollectionViewTransitionLayoutAnimatedKey */) float64
}

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

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewTransitionLayoutClass) Alloc() CollectionViewTransitionLayout {
	rv := objc.Send[CollectionViewTransitionLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes and returns the transition layout object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/init(currentLayout:nextLayout:)
func NewCollectionViewTransitionLayoutWithCurrentLayoutNextLayout(currentLayout ICollectionViewLayout, newLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	instance := getCollectionViewTransitionLayoutClass().Alloc()
	rv := objc.Send[CollectionViewTransitionLayout](instance.ID, objc.Sel("initWithCurrentLayout:nextLayout:"), currentLayout, newLayout)
	rv.Autorelease()
	return rv
}



// Sets the value of a key whose value you use during the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/updateValue(_:forAnimatedKey:)
func (c_ CollectionViewTransitionLayout) UpdateValueForAnimatedKey(value float64, key objc.IObject /* cross-framework: CollectionViewTransitionLayoutAnimatedKey */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateValue:forAnimatedKey:"), value, key)
}


// Returns the most recently set value for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/value(forAnimatedKey:)
func (c_ CollectionViewTransitionLayout) ValueForAnimatedKey(key objc.IObject /* cross-framework: CollectionViewTransitionLayoutAnimatedKey */) float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("valueForAnimatedKey:"), key)
	return rv
}


// The collection view’s current layout object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/currentLayout
func (c_ CollectionViewTransitionLayout) CurrentLayout() ICollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("currentLayout"))
	return rv
}


// The collection view’s new layout object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/nextLayout
func (c_ CollectionViewTransitionLayout) NextLayout() ICollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("nextLayout"))
	return rv
}


// The completion percentage of the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/transitionProgress
func (c_ CollectionViewTransitionLayout) TransitionProgress() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("transitionProgress"))
	return rv
}


// The completion percentage of the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/transitionProgress
func (c_ CollectionViewTransitionLayout) SetTransitionProgress(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTransitionProgress:"), value)
}


