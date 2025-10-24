// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScrollView] class.
var (
	ScrollViewClass     _ScrollViewClass
	ScrollViewClassOnce sync.Once
)

func getScrollViewClass() _ScrollViewClass {
	ScrollViewClassOnce.Do(func() {
		ScrollViewClass = _ScrollViewClass{objc.GetClass("UIScrollView")}
	})
	return ScrollViewClass
}

type _ScrollViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrollView] class.
type IScrollView interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other PencilKit classes.


// A parent class referenced by other PencilKit classes. [Full Topic]
type ScrollView struct {
	objectivec.Object
}

// ScrollViewFrom constructs a [ScrollView] from an unsafe.Pointer.
//
// A parent class referenced by other PencilKit classes.
func ScrollViewFrom(ptr unsafe.Pointer) ScrollView {
	return ScrollView{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrollViewClass) Alloc() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrollViewClass) New() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrollView) Init() ScrollView {
	rv := objc.Send[ScrollView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrollView) Autorelease() ScrollView {
	rv := objc.Send[ScrollView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrollView creates a new ScrollView instance.
func NewScrollView() ScrollView {
	return getScrollViewClass().New()
}




