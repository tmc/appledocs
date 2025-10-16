
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextTable] class.
var TextTableClass _TextTableClass

func init() {
	TextTableClass = _TextTableClass{objc.GetClass("NSTextTable")}
}

type _TextTableClass struct {
	objc.Class
}

// An interface definition for the [TextTable] class.
type ITextTable interface {
	ID() objc.ID
}

type TextTable struct {
	id objc.ID
}

func TextTableFrom(ptr unsafe.Pointer) TextTable {
	return TextTable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextTable) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextTableClass) Alloc() TextTable {
	rv := objc.Send[TextTable](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextTableClass) New() TextTable {
	rv := objc.Send[TextTable](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextTable creates and returns a new initialized instance.
func NewTextTable() TextTable {
	return TextTableClass.New()
}

// Init initializes the instance.
func (t_ TextTable) Init() TextTable {
	rv := objc.Send[TextTable](t_.ID(), selInit)
	return rv
}
