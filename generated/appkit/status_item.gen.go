
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StatusItem] class.
var StatusItemClass _StatusItemClass

func init() {
	StatusItemClass = _StatusItemClass{objc.GetClass("NSStatusItem")}
}

type _StatusItemClass struct {
	objc.Class
}

// An interface definition for the [StatusItem] class.
type IStatusItem interface {
	ID() objc.ID
}

type StatusItem struct {
	id objc.ID
}

func StatusItemFrom(ptr unsafe.Pointer) StatusItem {
	return StatusItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ StatusItem) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StatusItemClass) Alloc() StatusItem {
	rv := objc.Send[StatusItem](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StatusItemClass) New() StatusItem {
	rv := objc.Send[StatusItem](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStatusItem creates and returns a new initialized instance.
func NewStatusItem() StatusItem {
	return StatusItemClass.New()
}

// Init initializes the instance.
func (s_ StatusItem) Init() StatusItem {
	rv := objc.Send[StatusItem](s_.ID(), selInit)
	return rv
}
