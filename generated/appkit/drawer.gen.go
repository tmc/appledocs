// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Drawer] class.
var drawerClass = _DrawerClass{objc.GetClass("NSDrawer")}

type _DrawerClass struct {
	class objc.Class
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



