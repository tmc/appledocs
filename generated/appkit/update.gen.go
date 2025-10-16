
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [update] class.
var updateClass _updateClass

func init() {
	updateClass = _updateClass{objc.GetClass("update")}
}

type _updateClass struct {
	objc.Class
}

// An interface definition for the [update] class.
type Iupdate interface {
	ID() objc.ID
}

type update struct {
	id objc.ID
}

func updateFrom(ptr unsafe.Pointer) update {
	return update{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ update) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _updateClass) Alloc() update {
	rv := objc.Send[update](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _updateClass) New() update {
	rv := objc.Send[update](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newupdate creates and returns a new initialized instance.
func Newupdate() update {
	return updateClass.New()
}

// Init initializes the instance.
func (u_ update) Init() update {
	rv := objc.Send[update](u_.ID(), selInit)
	return rv
}
