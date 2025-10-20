// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextSelectionNavigation] class.
var (
	TextSelectionNavigationClass     _TextSelectionNavigationClass
	TextSelectionNavigationClassOnce sync.Once
)

func getTextSelectionNavigationClass() _TextSelectionNavigationClass {
	TextSelectionNavigationClassOnce.Do(func() {
		TextSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}
	})
	return TextSelectionNavigationClass
}

type _TextSelectionNavigationClass struct {
	class objc.Class
}

// An interface definition for the [TextSelectionNavigation] class.
type ITextSelectionNavigation interface {
	objectivec.IObject
}

// An interface you use to expose methods for obtaining results from actions performed on text selections.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation
type TextSelectionNavigation struct {
	objectivec.Object
}

// TextSelectionNavigationFrom constructs a [TextSelectionNavigation] from an unsafe.Pointer.
//
// An interface you use to expose methods for obtaining results from actions performed on text selections.
func TextSelectionNavigationFrom(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextSelectionNavigationClass) Alloc() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextSelectionNavigationClass) New() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextSelectionNavigation) Init() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextSelectionNavigation) Autorelease() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextSelectionNavigation creates a new TextSelectionNavigation instance.
func NewTextSelectionNavigation() TextSelectionNavigation {
	return getTextSelectionNavigationClass().New()
}
