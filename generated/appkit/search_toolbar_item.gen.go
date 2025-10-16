
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SearchToolbarItem] class.
var SearchToolbarItemClass _SearchToolbarItemClass

func init() {
	SearchToolbarItemClass = _SearchToolbarItemClass{objc.GetClass("NSSearchToolbarItem")}
}

type _SearchToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [SearchToolbarItem] class.
type ISearchToolbarItem interface {
	ID() objc.ID
}

type SearchToolbarItem struct {
	id objc.ID
}

func SearchToolbarItemFrom(ptr unsafe.Pointer) SearchToolbarItem {
	return SearchToolbarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SearchToolbarItem) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SearchToolbarItemClass) Alloc() SearchToolbarItem {
	rv := objc.Send[SearchToolbarItem](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SearchToolbarItemClass) New() SearchToolbarItem {
	rv := objc.Send[SearchToolbarItem](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSearchToolbarItem creates and returns a new initialized instance.
func NewSearchToolbarItem() SearchToolbarItem {
	return SearchToolbarItemClass.New()
}

// Init initializes the instance.
func (s_ SearchToolbarItem) Init() SearchToolbarItem {
	rv := objc.Send[SearchToolbarItem](s_.ID(), selInit)
	return rv
}
