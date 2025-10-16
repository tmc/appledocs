
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maxNumberOfRows] class.
var maxNumberOfRowsClass _maxNumberOfRowsClass

func init() {
	maxNumberOfRowsClass = _maxNumberOfRowsClass{objc.GetClass("maxNumberOfRows")}
}

type _maxNumberOfRowsClass struct {
	objc.Class
}

// An interface definition for the [maxNumberOfRows] class.
type ImaxNumberOfRows interface {
	ID() objc.ID
}

type maxNumberOfRows struct {
	id objc.ID
}

func maxNumberOfRowsFrom(ptr unsafe.Pointer) maxNumberOfRows {
	return maxNumberOfRows{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maxNumberOfRows) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maxNumberOfRowsClass) Alloc() maxNumberOfRows {
	rv := objc.Send[maxNumberOfRows](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maxNumberOfRowsClass) New() maxNumberOfRows {
	rv := objc.Send[maxNumberOfRows](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaxNumberOfRows creates and returns a new initialized instance.
func NewmaxNumberOfRows() maxNumberOfRows {
	return maxNumberOfRowsClass.New()
}

// Init initializes the instance.
func (m_ maxNumberOfRows) Init() maxNumberOfRows {
	rv := objc.Send[maxNumberOfRows](m_.ID(), selInit)
	return rv
}
