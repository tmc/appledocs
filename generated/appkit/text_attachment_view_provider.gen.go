// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextAttachmentViewProvider] class.
var (
	textAttachmentViewProviderClass     _TextAttachmentViewProviderClass
	textAttachmentViewProviderClassOnce sync.Once
)

func getTextAttachmentViewProviderClass() _TextAttachmentViewProviderClass {
	textAttachmentViewProviderClassOnce.Do(func() {
		textAttachmentViewProviderClass = _TextAttachmentViewProviderClass{objc.GetClass("NSTextAttachmentViewProvider")}
	})
	return textAttachmentViewProviderClass
}

type _TextAttachmentViewProviderClass struct {
	class objc.Class
}

// An interface definition for the [TextAttachmentViewProvider] class.
type ITextAttachmentViewProvider interface {
	objectivec.IObject
}

// A container object that associates a text attachment at a particular document location with a view object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider

type TextAttachmentViewProvider struct {
	objectivec.Object
}

// TextAttachmentViewProviderFrom constructs a [TextAttachmentViewProvider] from an unsafe.Pointer.
//
// A container object that associates a text attachment at a particular document location with a view object.
func TextAttachmentViewProviderFrom(ptr unsafe.Pointer) TextAttachmentViewProvider {
	return TextAttachmentViewProvider{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextAttachmentViewProviderClass) Alloc() TextAttachmentViewProvider {
	rv := objc.Send[TextAttachmentViewProvider](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextAttachmentViewProviderClass) New() TextAttachmentViewProvider {
	rv := objc.Send[TextAttachmentViewProvider](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextAttachmentViewProvider) Init() TextAttachmentViewProvider {
	rv := objc.Send[TextAttachmentViewProvider](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextAttachmentViewProvider) Autorelease() TextAttachmentViewProvider {
	rv := objc.Send[TextAttachmentViewProvider](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextAttachmentViewProvider creates a new TextAttachmentViewProvider instance.
func NewTextAttachmentViewProvider() TextAttachmentViewProvider {
	return getTextAttachmentViewProviderClass().New()
}




