
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isContextHelpModeActive] class.
var isContextHelpModeActiveClass _isContextHelpModeActiveClass

func init() {
	isContextHelpModeActiveClass = _isContextHelpModeActiveClass{objc.GetClass("isContextHelpModeActive")}
}

type _isContextHelpModeActiveClass struct {
	objc.Class
}

// An interface definition for the [isContextHelpModeActive] class.
type IisContextHelpModeActive interface {
	ID() objc.ID
}

type isContextHelpModeActive struct {
	id objc.ID
}

func isContextHelpModeActiveFrom(ptr unsafe.Pointer) isContextHelpModeActive {
	return isContextHelpModeActive{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isContextHelpModeActive) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isContextHelpModeActiveClass) Alloc() isContextHelpModeActive {
	rv := objc.Send[isContextHelpModeActive](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isContextHelpModeActiveClass) New() isContextHelpModeActive {
	rv := objc.Send[isContextHelpModeActive](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisContextHelpModeActive creates and returns a new initialized instance.
func NewisContextHelpModeActive() isContextHelpModeActive {
	return isContextHelpModeActiveClass.New()
}

// Init initializes the instance.
func (i_ isContextHelpModeActive) Init() isContextHelpModeActive {
	rv := objc.Send[isContextHelpModeActive](i_.ID(), selInit)
	return rv
}
