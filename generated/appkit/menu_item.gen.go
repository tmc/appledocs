// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MenuItem] class.
var menuItemClass = _MenuItemClass{objc.GetClass("NSMenuItem")}

type _MenuItemClass struct {
	class objc.Class
}

// A command item in an app menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem

type MenuItem struct {
	objectivec.Object
}

// MenuItemFrom constructs a [MenuItem] from an unsafe.Pointer.
//
// A command item in an app menu.
func MenuItemFrom(ptr unsafe.Pointer) MenuItem {
	return MenuItem{objectivec.Object{objc.ID(ptr)}}
}



