
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [worksWhenModal] class.
var worksWhenModalClass _worksWhenModalClass

func init() {
	worksWhenModalClass = _worksWhenModalClass{objc.GetClass("worksWhenModal")}
}

type _worksWhenModalClass struct {
	objc.Class
}

// An interface definition for the [worksWhenModal] class.
type IworksWhenModal interface {
	ID() objc.ID
}

type worksWhenModal struct {
	id objc.ID
}

func worksWhenModalFrom(ptr unsafe.Pointer) worksWhenModal {
	return worksWhenModal{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ worksWhenModal) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _worksWhenModalClass) Alloc() worksWhenModal {
	rv := objc.Send[worksWhenModal](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _worksWhenModalClass) New() worksWhenModal {
	rv := objc.Send[worksWhenModal](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewworksWhenModal creates and returns a new initialized instance.
func NewworksWhenModal() worksWhenModal {
	return worksWhenModalClass.New()
}

// Init initializes the instance.
func (w_ worksWhenModal) Init() worksWhenModal {
	rv := objc.Send[worksWhenModal](w_.ID(), selInit)
	return rv
}
