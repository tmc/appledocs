
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasTitleBar] class.
var hasTitleBarClass _hasTitleBarClass

func init() {
	hasTitleBarClass = _hasTitleBarClass{objc.GetClass("hasTitleBar")}
}

type _hasTitleBarClass struct {
	objc.Class
}

// An interface definition for the [hasTitleBar] class.
type IhasTitleBar interface {
	ID() objc.ID
}

type hasTitleBar struct {
	id objc.ID
}

func hasTitleBarFrom(ptr unsafe.Pointer) hasTitleBar {
	return hasTitleBar{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasTitleBar) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasTitleBarClass) Alloc() hasTitleBar {
	rv := objc.Send[hasTitleBar](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasTitleBarClass) New() hasTitleBar {
	rv := objc.Send[hasTitleBar](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasTitleBar creates and returns a new initialized instance.
func NewhasTitleBar() hasTitleBar {
	return hasTitleBarClass.New()
}

// Init initializes the instance.
func (h_ hasTitleBar) Init() hasTitleBar {
	rv := objc.Send[hasTitleBar](h_.ID(), selInit)
	return rv
}
