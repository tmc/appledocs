
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [invalidateHashMarks] class.
var invalidateHashMarksClass _invalidateHashMarksClass

func init() {
	invalidateHashMarksClass = _invalidateHashMarksClass{objc.GetClass("invalidateHashMarks")}
}

type _invalidateHashMarksClass struct {
	objc.Class
}

// An interface definition for the [invalidateHashMarks] class.
type IinvalidateHashMarks interface {
	ID() objc.ID
}

type invalidateHashMarks struct {
	id objc.ID
}

func invalidateHashMarksFrom(ptr unsafe.Pointer) invalidateHashMarks {
	return invalidateHashMarks{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ invalidateHashMarks) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _invalidateHashMarksClass) Alloc() invalidateHashMarks {
	rv := objc.Send[invalidateHashMarks](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _invalidateHashMarksClass) New() invalidateHashMarks {
	rv := objc.Send[invalidateHashMarks](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinvalidateHashMarks creates and returns a new initialized instance.
func NewinvalidateHashMarks() invalidateHashMarks {
	return invalidateHashMarksClass.New()
}

// Init initializes the instance.
func (i_ invalidateHashMarks) Init() invalidateHashMarks {
	rv := objc.Send[invalidateHashMarks](i_.ID(), selInit)
	return rv
}
