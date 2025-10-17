
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BackgroundExtensionView] class.
var BackgroundExtensionViewClass _BackgroundExtensionViewClass

func init() {
	BackgroundExtensionViewClass = _BackgroundExtensionViewClass{objc.GetClass("NSBackgroundExtensionView")}
}

type _BackgroundExtensionViewClass struct {
	objc.Class
}

// An interface definition for the [BackgroundExtensionView] class.
type IBackgroundExtensionView interface {
	ID() objc.ID
}

type BackgroundExtensionView struct {
	id objc.ID
}

func BackgroundExtensionViewFrom(ptr unsafe.Pointer) BackgroundExtensionView {
	return BackgroundExtensionView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ BackgroundExtensionView) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _BackgroundExtensionViewClass) Alloc() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _BackgroundExtensionViewClass) New() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewBackgroundExtensionView creates and returns a new initialized instance.
func NewBackgroundExtensionView() BackgroundExtensionView {
	return BackgroundExtensionViewClass.New()
}

// Init initializes the instance.
func (b_ BackgroundExtensionView) Init() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](b_.ID(), selInit)
	return rv
}
// Controls the automatic safe area placement of the   within the   container. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBackgroundExtensionView/automaticallyPlacesContentView
func (b_ BackgroundExtensionView) AutomaticallyPlacesContentView() bool {
	rv := objc.Send[bool](b_.ID(), objc.RegisterName("automaticallyPlacesContentView"))
	return rv
}
// SetAutomaticallyPlacesContentView sets the value of the automaticallyPlacesContentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBackgroundExtensionView/automaticallyPlacesContentView
func (b_ BackgroundExtensionView) SetAutomaticallyPlacesContentView(value bool) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setAutomaticallyPlacesContentView:"), value)
}
// The content view to extend to fill the  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBackgroundExtensionView/contentView
func (b_ BackgroundExtensionView) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("contentView"))
	return rv
}
// SetContentView sets the value of the contentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBackgroundExtensionView/contentView
func (b_ BackgroundExtensionView) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setContentView:"), value)
}
