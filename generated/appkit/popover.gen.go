
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Popover] class.
var PopoverClass _PopoverClass

func init() {
	PopoverClass = _PopoverClass{objc.GetClass("NSPopover")}
}

type _PopoverClass struct {
	objc.Class
}

// An interface definition for the [Popover] class.
type IPopover interface {
	ID() objc.ID
}

type Popover struct {
	id objc.ID
}

func PopoverFrom(ptr unsafe.Pointer) Popover {
	return Popover{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ Popover) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PopoverClass) Alloc() Popover {
	rv := objc.Send[Popover](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PopoverClass) New() Popover {
	rv := objc.Send[Popover](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPopover creates and returns a new initialized instance.
func NewPopover() Popover {
	return PopoverClass.New()
}

// Init initializes the instance.
func (p_ Popover) Init() Popover {
	rv := objc.Send[Popover](p_.ID(), selInit)
	return rv
}
