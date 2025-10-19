// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WritingToolsCoordinatorContext] class.
var (
	writingToolsCoordinatorContextClass     _WritingToolsCoordinatorContextClass
	writingToolsCoordinatorContextClassOnce sync.Once
)

func getWritingToolsCoordinatorContextClass() _WritingToolsCoordinatorContextClass {
	writingToolsCoordinatorContextClassOnce.Do(func() {
		writingToolsCoordinatorContextClass = _WritingToolsCoordinatorContextClass{objc.GetClass("NSWritingToolsCoordinatorContext")}
	})
	return writingToolsCoordinatorContextClass
}

type _WritingToolsCoordinatorContextClass struct {
	class objc.Class
}

// An interface definition for the [WritingToolsCoordinatorContext] class.
type IWritingToolsCoordinatorContext interface {
	objectivec.IObject
}

// A data object that you use to share your custom view’s text with Writing Tools. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context

type WritingToolsCoordinatorContext struct {
	objectivec.Object
}

// WritingToolsCoordinatorContextFrom constructs a [WritingToolsCoordinatorContext] from an unsafe.Pointer.
//
// A data object that you use to share your custom view’s text with Writing Tools.
func WritingToolsCoordinatorContextFrom(ptr unsafe.Pointer) WritingToolsCoordinatorContext {
	return WritingToolsCoordinatorContext{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorContextClass) Alloc() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WritingToolsCoordinatorContextClass) New() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WritingToolsCoordinatorContext) Init() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WritingToolsCoordinatorContext) Autorelease() WritingToolsCoordinatorContext {
	rv := objc.Send[WritingToolsCoordinatorContext](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWritingToolsCoordinatorContext creates a new WritingToolsCoordinatorContext instance.
func NewWritingToolsCoordinatorContext() WritingToolsCoordinatorContext {
	return getWritingToolsCoordinatorContextClass().New()
}




