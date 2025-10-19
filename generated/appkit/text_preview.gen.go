// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextPreview] class.
var (
	textPreviewClass     _TextPreviewClass
	textPreviewClassOnce sync.Once
)

func getTextPreviewClass() _TextPreviewClass {
	textPreviewClassOnce.Do(func() {
		textPreviewClass = _TextPreviewClass{objc.GetClass("NSTextPreview")}
	})
	return textPreviewClass
}

type _TextPreviewClass struct {
	class objc.Class
}

// An interface definition for the [TextPreview] class.
type ITextPreview interface {
	objectivec.IObject
}

// A snapshot of the text in your view, which the system uses to create user-visible effects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextPreview
type TextPreview struct {
	objectivec.Object
}

// TextPreviewFrom constructs a [TextPreview] from an unsafe.Pointer.
//
// A snapshot of the text in your view, which the system uses to create user-visible effects.
func TextPreviewFrom(ptr unsafe.Pointer) TextPreview {
	return TextPreview{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextPreviewClass) Alloc() TextPreview {
	rv := objc.Send[TextPreview](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextPreviewClass) New() TextPreview {
	rv := objc.Send[TextPreview](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextPreview) Init() TextPreview {
	rv := objc.Send[TextPreview](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextPreview) Autorelease() TextPreview {
	rv := objc.Send[TextPreview](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextPreview creates a new TextPreview instance.
func NewTextPreview() TextPreview {
	return getTextPreviewClass().New()
}




