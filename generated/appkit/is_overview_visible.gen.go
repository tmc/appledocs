
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isOverviewVisible] class.
var isOverviewVisibleClass _isOverviewVisibleClass

func init() {
	isOverviewVisibleClass = _isOverviewVisibleClass{objc.GetClass("isOverviewVisible")}
}

type _isOverviewVisibleClass struct {
	objc.Class
}

// An interface definition for the [isOverviewVisible] class.
type IisOverviewVisible interface {
	ID() objc.ID
}

type isOverviewVisible struct {
	id objc.ID
}

func isOverviewVisibleFrom(ptr unsafe.Pointer) isOverviewVisible {
	return isOverviewVisible{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isOverviewVisible) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isOverviewVisibleClass) Alloc() isOverviewVisible {
	rv := objc.Send[isOverviewVisible](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isOverviewVisibleClass) New() isOverviewVisible {
	rv := objc.Send[isOverviewVisible](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisOverviewVisible creates and returns a new initialized instance.
func NewisOverviewVisible() isOverviewVisible {
	return isOverviewVisibleClass.New()
}

// Init initializes the instance.
func (i_ isOverviewVisible) Init() isOverviewVisible {
	rv := objc.Send[isOverviewVisible](i_.ID(), selInit)
	return rv
}
