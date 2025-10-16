
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isContinuous] class.
var isContinuousClass _isContinuousClass

func init() {
	isContinuousClass = _isContinuousClass{objc.GetClass("isContinuous")}
}

type _isContinuousClass struct {
	objc.Class
}

// An interface definition for the [isContinuous] class.
type IisContinuous interface {
	ID() objc.ID
}

type isContinuous struct {
	id objc.ID
}

func isContinuousFrom(ptr unsafe.Pointer) isContinuous {
	return isContinuous{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isContinuous) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isContinuousClass) Alloc() isContinuous {
	rv := objc.Send[isContinuous](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isContinuousClass) New() isContinuous {
	rv := objc.Send[isContinuous](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisContinuous creates and returns a new initialized instance.
func NewisContinuous() isContinuous {
	return isContinuousClass.New()
}

// Init initializes the instance.
func (i_ isContinuous) Init() isContinuous {
	rv := objc.Send[isContinuous](i_.ID(), selInit)
	return rv
}
