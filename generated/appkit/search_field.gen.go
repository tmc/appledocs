
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SearchField] class.
var SearchFieldClass _SearchFieldClass

func init() {
	SearchFieldClass = _SearchFieldClass{objc.GetClass("NSSearchField")}
}

type _SearchFieldClass struct {
	objc.Class
}

// An interface definition for the [SearchField] class.
type ISearchField interface {
	ID() objc.ID
}

type SearchField struct {
	id objc.ID
}

func SearchFieldFrom(ptr unsafe.Pointer) SearchField {
	return SearchField{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SearchField) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SearchFieldClass) Alloc() SearchField {
	rv := objc.Send[SearchField](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SearchFieldClass) New() SearchField {
	rv := objc.Send[SearchField](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSearchField creates and returns a new initialized instance.
func NewSearchField() SearchField {
	return SearchFieldClass.New()
}

// Init initializes the instance.
func (s_ SearchField) Init() SearchField {
	rv := objc.Send[SearchField](s_.ID(), selInit)
	return rv
}
