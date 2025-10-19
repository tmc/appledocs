// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLayoutFragment] class.
var (
	textLayoutFragmentClass     _TextLayoutFragmentClass
	textLayoutFragmentClassOnce sync.Once
)

func getTextLayoutFragmentClass() _TextLayoutFragmentClass {
	textLayoutFragmentClassOnce.Do(func() {
		textLayoutFragmentClass = _TextLayoutFragmentClass{objc.GetClass("NSTextLayoutFragment")}
	})
	return textLayoutFragmentClass
}

type _TextLayoutFragmentClass struct {
	class objc.Class
}

// An interface definition for the [TextLayoutFragment] class.
type ITextLayoutFragment interface {
	objectivec.IObject
}

// A class that represents the layout fragment typically corresponding to a rendering surface, such as a layer or view subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment
type TextLayoutFragment struct {
	objectivec.Object
}

// TextLayoutFragmentFrom constructs a [TextLayoutFragment] from an unsafe.Pointer.
//
// A class that represents the layout fragment typically corresponding to a rendering surface, such as a layer or view subclass.
func TextLayoutFragmentFrom(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextLayoutFragmentClass) Alloc() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextLayoutFragmentClass) New() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLayoutFragment) Init() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLayoutFragment) Autorelease() TextLayoutFragment {
	rv := objc.Send[TextLayoutFragment](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLayoutFragment creates a new TextLayoutFragment instance.
func NewTextLayoutFragment() TextLayoutFragment {
	return getTextLayoutFragmentClass().New()
}




