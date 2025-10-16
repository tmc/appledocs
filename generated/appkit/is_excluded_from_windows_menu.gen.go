
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isExcludedFromWindowsMenu] class.
var isExcludedFromWindowsMenuClass _isExcludedFromWindowsMenuClass

func init() {
	isExcludedFromWindowsMenuClass = _isExcludedFromWindowsMenuClass{objc.GetClass("isExcludedFromWindowsMenu")}
}

type _isExcludedFromWindowsMenuClass struct {
	objc.Class
}

// An interface definition for the [isExcludedFromWindowsMenu] class.
type IisExcludedFromWindowsMenu interface {
	ID() objc.ID
}

type isExcludedFromWindowsMenu struct {
	id objc.ID
}

func isExcludedFromWindowsMenuFrom(ptr unsafe.Pointer) isExcludedFromWindowsMenu {
	return isExcludedFromWindowsMenu{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isExcludedFromWindowsMenu) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isExcludedFromWindowsMenuClass) Alloc() isExcludedFromWindowsMenu {
	rv := objc.Send[isExcludedFromWindowsMenu](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isExcludedFromWindowsMenuClass) New() isExcludedFromWindowsMenu {
	rv := objc.Send[isExcludedFromWindowsMenu](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisExcludedFromWindowsMenu creates and returns a new initialized instance.
func NewisExcludedFromWindowsMenu() isExcludedFromWindowsMenu {
	return isExcludedFromWindowsMenuClass.New()
}

// Init initializes the instance.
func (i_ isExcludedFromWindowsMenu) Init() isExcludedFromWindowsMenu {
	rv := objc.Send[isExcludedFromWindowsMenu](i_.ID(), selInit)
	return rv
}
