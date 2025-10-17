
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ATSTypesetter] class.
var ATSTypesetterClass _ATSTypesetterClass

func init() {
	ATSTypesetterClass = _ATSTypesetterClass{objc.GetClass("NSATSTypesetter")}
}

type _ATSTypesetterClass struct {
	objc.Class
}

// An interface definition for the [ATSTypesetter] class.
type IATSTypesetter interface {
	ID() objc.ID
}

type ATSTypesetter struct {
	id objc.ID
}

func ATSTypesetterFrom(ptr unsafe.Pointer) ATSTypesetter {
	return ATSTypesetter{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ ATSTypesetter) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _ATSTypesetterClass) Alloc() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _ATSTypesetterClass) New() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewATSTypesetter creates and returns a new initialized instance.
func NewATSTypesetter() ATSTypesetter {
	return ATSTypesetterClass.New()
}

// Init initializes the instance.
func (a_ ATSTypesetter) Init() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](a_.ID(), selInit)
	return rv
}
