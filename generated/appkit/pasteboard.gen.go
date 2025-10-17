
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Pasteboard] class.
var PasteboardClass _PasteboardClass

func init() {
	PasteboardClass = _PasteboardClass{objc.GetClass("NSPasteboard")}
}

type _PasteboardClass struct {
	objc.Class
}

// An interface definition for the [Pasteboard] class.
type IPasteboard interface {
	ID() objc.ID
	SetStringForType(string string, dataType unsafe.Pointer) bool
}

type Pasteboard struct {
	id objc.ID
}

func PasteboardFrom(ptr unsafe.Pointer) Pasteboard {
	return Pasteboard{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ Pasteboard) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PasteboardClass) Alloc() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PasteboardClass) New() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPasteboard creates and returns a new initialized instance.
func NewPasteboard() Pasteboard {
	return PasteboardClass.New()
}

// Init initializes the instance.
func (p_ Pasteboard) Init() Pasteboard {
	rv := objc.Send[Pasteboard](p_.ID(), selInit)
	return rv
}
// Sets the given string as the representation for the specified type for the first item on the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPasteboard/setString(_:forType:)
func (p_ Pasteboard) SetStringForType(string string, dataType unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID(), objc.RegisterName("setString:forType:"), string, dataType)
	return rv
}
// The current pasteboard access behavior. The user can customize this behavior per-app in System Settings for any app that has triggered a pasteboard access alert in the past. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPasteboard/accessBehavior-86972
func (p_ Pasteboard) AccessBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("accessBehavior"))
	return rv
}
