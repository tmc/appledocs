
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [becomesKeyOnlyIfNeeded] class.
var becomesKeyOnlyIfNeededClass _becomesKeyOnlyIfNeededClass

func init() {
	becomesKeyOnlyIfNeededClass = _becomesKeyOnlyIfNeededClass{objc.GetClass("becomesKeyOnlyIfNeeded")}
}

type _becomesKeyOnlyIfNeededClass struct {
	objc.Class
}

// An interface definition for the [becomesKeyOnlyIfNeeded] class.
type IbecomesKeyOnlyIfNeeded interface {
	ID() objc.ID
}

type becomesKeyOnlyIfNeeded struct {
	id objc.ID
}

func becomesKeyOnlyIfNeededFrom(ptr unsafe.Pointer) becomesKeyOnlyIfNeeded {
	return becomesKeyOnlyIfNeeded{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ becomesKeyOnlyIfNeeded) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _becomesKeyOnlyIfNeededClass) Alloc() becomesKeyOnlyIfNeeded {
	rv := objc.Send[becomesKeyOnlyIfNeeded](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _becomesKeyOnlyIfNeededClass) New() becomesKeyOnlyIfNeeded {
	rv := objc.Send[becomesKeyOnlyIfNeeded](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbecomesKeyOnlyIfNeeded creates and returns a new initialized instance.
func NewbecomesKeyOnlyIfNeeded() becomesKeyOnlyIfNeeded {
	return becomesKeyOnlyIfNeededClass.New()
}

// Init initializes the instance.
func (b_ becomesKeyOnlyIfNeeded) Init() becomesKeyOnlyIfNeeded {
	rv := objc.Send[becomesKeyOnlyIfNeeded](b_.ID(), selInit)
	return rv
}
