// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Pasteboard] class.
var (
	pasteboardClass     _PasteboardClass
	pasteboardClassOnce sync.Once
)

func getPasteboardClass() _PasteboardClass {
	pasteboardClassOnce.Do(func() {
		pasteboardClass = _PasteboardClass{objc.GetClass("NSPasteboard")}
	})
	return pasteboardClass
}

type _PasteboardClass struct {
	class objc.Class
}

// An interface definition for the [Pasteboard] class.
type IPasteboard interface {
	objectivec.IObject
	SetStringForType(string string, dataType unsafe.Pointer) bool
}

// An object that transfers data to and from the pasteboard server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard

type Pasteboard struct {
	objectivec.Object
}

// PasteboardFrom constructs a [Pasteboard] from an unsafe.Pointer.
//
// An object that transfers data to and from the pasteboard server.
func PasteboardFrom(ptr unsafe.Pointer) Pasteboard {
	return Pasteboard{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _PasteboardClass) Alloc() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PasteboardClass) New() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Pasteboard) Init() Pasteboard {
	rv := objc.Send[Pasteboard](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Pasteboard) Autorelease() Pasteboard {
	rv := objc.Send[Pasteboard](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasteboard creates a new Pasteboard instance.
func NewPasteboard() Pasteboard {
	return getPasteboardClass().New()
}


// Sets the given string as the representation for the specified type for the first item on the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setString(_:forType:)
func (p_ Pasteboard) SetStringForType(string string, dataType unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setString:forType:"), objc.String(string), dataType)
	return rv
}


