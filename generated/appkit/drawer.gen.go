
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Drawer] class.
var DrawerClass _DrawerClass

func init() {
	DrawerClass = _DrawerClass{objc.GetClass("NSDrawer")}
}

type _DrawerClass struct {
	objc.Class
}

// An interface definition for the [Drawer] class.
type IDrawer interface {
	ID() objc.ID
}

type Drawer struct {
	id objc.ID
}

func DrawerFrom(ptr unsafe.Pointer) Drawer {
	return Drawer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ Drawer) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DrawerClass) Alloc() Drawer {
	rv := objc.Send[Drawer](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DrawerClass) New() Drawer {
	rv := objc.Send[Drawer](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDrawer creates and returns a new initialized instance.
func NewDrawer() Drawer {
	return DrawerClass.New()
}

// Init initializes the instance.
func (d_ Drawer) Init() Drawer {
	rv := objc.Send[Drawer](d_.ID(), selInit)
	return rv
}
