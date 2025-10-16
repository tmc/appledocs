
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [defaultMenu] class.
var defaultMenuClass _defaultMenuClass

func init() {
	defaultMenuClass = _defaultMenuClass{objc.GetClass("defaultMenu")}
}

type _defaultMenuClass struct {
	objc.Class
}

// An interface definition for the [defaultMenu] class.
type IdefaultMenu interface {
	ID() objc.ID
}

type defaultMenu struct {
	id objc.ID
}

func defaultMenuFrom(ptr unsafe.Pointer) defaultMenu {
	return defaultMenu{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ defaultMenu) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _defaultMenuClass) Alloc() defaultMenu {
	rv := objc.Send[defaultMenu](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _defaultMenuClass) New() defaultMenu {
	rv := objc.Send[defaultMenu](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdefaultMenu creates and returns a new initialized instance.
func NewdefaultMenu() defaultMenu {
	return defaultMenuClass.New()
}

// Init initializes the instance.
func (d_ defaultMenu) Init() defaultMenu {
	rv := objc.Send[defaultMenu](d_.ID(), selInit)
	return rv
}
