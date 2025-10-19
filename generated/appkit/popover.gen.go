// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Popover] class.
var (
	popoverClass     _PopoverClass
	popoverClassOnce sync.Once
)

func getPopoverClass() _PopoverClass {
	popoverClassOnce.Do(func() {
		popoverClass = _PopoverClass{objc.GetClass("NSPopover")}
	})
	return popoverClass
}

type _PopoverClass struct {
	class objc.Class
}

// An interface definition for the [Popover] class.
type IPopover interface {
	IResponder
}

// A means to display additional content related to existing content on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover

type Popover struct {
	Responder
}

// PopoverFrom constructs a [Popover] from an unsafe.Pointer.
//
// A means to display additional content related to existing content on the screen.
func PopoverFrom(ptr unsafe.Pointer) Popover {
	return Popover{
		Responder: ResponderFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PopoverClass) Alloc() Popover {
	rv := objc.Send[Popover](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PopoverClass) New() Popover {
	rv := objc.Send[Popover](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Popover) Init() Popover {
	rv := objc.Send[Popover](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Popover) Autorelease() Popover {
	rv := objc.Send[Popover](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPopover creates a new Popover instance.
func NewPopover() Popover {
	return getPopoverClass().New()
}




