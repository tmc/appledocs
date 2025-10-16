
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maxNumberOfColumns] class.
var maxNumberOfColumnsClass _maxNumberOfColumnsClass

func init() {
	maxNumberOfColumnsClass = _maxNumberOfColumnsClass{objc.GetClass("maxNumberOfColumns")}
}

type _maxNumberOfColumnsClass struct {
	objc.Class
}

// An interface definition for the [maxNumberOfColumns] class.
type ImaxNumberOfColumns interface {
	ID() objc.ID
}

type maxNumberOfColumns struct {
	id objc.ID
}

func maxNumberOfColumnsFrom(ptr unsafe.Pointer) maxNumberOfColumns {
	return maxNumberOfColumns{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maxNumberOfColumns) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maxNumberOfColumnsClass) Alloc() maxNumberOfColumns {
	rv := objc.Send[maxNumberOfColumns](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maxNumberOfColumnsClass) New() maxNumberOfColumns {
	rv := objc.Send[maxNumberOfColumns](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaxNumberOfColumns creates and returns a new initialized instance.
func NewmaxNumberOfColumns() maxNumberOfColumns {
	return maxNumberOfColumnsClass.New()
}

// Init initializes the instance.
func (m_ maxNumberOfColumns) Init() maxNumberOfColumns {
	rv := objc.Send[maxNumberOfColumns](m_.ID(), selInit)
	return rv
}
