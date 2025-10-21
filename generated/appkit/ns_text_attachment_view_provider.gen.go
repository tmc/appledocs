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
	TextAttachmentViewProviderClass     _TextAttachmentViewProviderClass
	TextAttachmentViewProviderClassOnce sync.Once
)

func getTextAttachmentViewProviderClass() _TextAttachmentViewProviderClass {
	TextAttachmentViewProviderClassOnce.Do(func() {
		TextAttachmentViewProviderClass = _TextAttachmentViewProviderClass{objc.GetClass("NSTextAttachmentViewProvider")}
	})
	return TextAttachmentViewProviderClass
}

type _TextAttachmentViewProviderClass struct {
	class objc.Class
}

// An interface definition for the [TextAttachmentViewProvider] class.
type ITextAttachmentViewProvider interface {
	objectivec.IObject
}

// A container object that associates a text attachment at a particular document location with a view object.
//
// Use when you need to represent document locations in terms of an  or an  or you want to support view-based text attachments. The view provider controls the view placement and layout without requiring view classes to be aware of the text attachment coordination using a in macOS 12 or iOS 15 and later.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The location that indicates the start of the text attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/location
func (t_ TextAttachmentViewProvider) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
// The location that indicates the start of the text attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/location
func (t_ TextAttachmentViewProvider) SetLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocation:"), value)
}

// The text attachment for this view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/textattachment
func (t_ TextAttachmentViewProvider) TextAttachment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textAttachment"))
	return rv
}


// SetTextAttachment sets the value of the textAttachment property.
// The text attachment for this view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/textattachment
func (t_ TextAttachmentViewProvider) SetTextAttachment(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextAttachment:"), value)
}

// The text layout manager for this view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/textlayoutmanager
func (t_ TextAttachmentViewProvider) TextLayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}


// SetTextLayoutManager sets the value of the textLayoutManager property.
// The text layout manager for this view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/textlayoutmanager
func (t_ TextAttachmentViewProvider) SetTextLayoutManager(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLayoutManager:"), value)
}

// A Boolean value that determines the text attachment’s bounds policy.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/trackstextattachmentviewbounds
func (t_ TextAttachmentViewProvider) TracksTextAttachmentViewBounds() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("tracksTextAttachmentViewBounds"))
	return rv
}


// SetTracksTextAttachmentViewBounds sets the value of the tracksTextAttachmentViewBounds property.
// A Boolean value that determines the text attachment’s bounds policy.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/trackstextattachmentviewbounds
func (t_ TextAttachmentViewProvider) SetTracksTextAttachmentViewBounds(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTracksTextAttachmentViewBounds:"), value)
}

// The text attachment’s view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/view
func (t_ TextAttachmentViewProvider) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The text attachment’s view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachmentviewprovider/view
func (t_ TextAttachmentViewProvider) SetView(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}



