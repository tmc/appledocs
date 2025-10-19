// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Drawer] class.
var (
	drawerClass     _DrawerClass
	drawerClassOnce sync.Once
)

func getDrawerClass() _DrawerClass {
	drawerClassOnce.Do(func() {
		drawerClass = _DrawerClass{objc.GetClass("NSDrawer")}
	})
	return drawerClass
}

type _DrawerClass struct {
	class objc.Class
}

// An interface definition for the [Drawer] class.
type IDrawer interface {
	IResponder
}

// A user interface element that contains and displays text, scroll, and browser views, in addition to other view subclasses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer

type Drawer struct {
	Responder
}

// DrawerFrom constructs a [Drawer] from an unsafe.Pointer.
//
// A user interface element that contains and displays text, scroll, and browser views, in addition to other view subclasses.
func DrawerFrom(ptr unsafe.Pointer) Drawer {
	return Drawer{
		Responder: ResponderFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (dc _DrawerClass) Alloc() Drawer {
	rv := objc.Send[Drawer](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DrawerClass) New() Drawer {
	rv := objc.Send[Drawer](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Drawer) Init() Drawer {
	rv := objc.Send[Drawer](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Drawer) Autorelease() Drawer {
	rv := objc.Send[Drawer](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDrawer creates a new Drawer instance.
func NewDrawer() Drawer {
	return getDrawerClass().New()
}




