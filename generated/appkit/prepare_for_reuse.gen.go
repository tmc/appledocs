
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [prepareForReuse] class.
var prepareForReuseClass _prepareForReuseClass

func init() {
	prepareForReuseClass = _prepareForReuseClass{objc.GetClass("prepareForReuse")}
}

type _prepareForReuseClass struct {
	objc.Class
}

// An interface definition for the [prepareForReuse] class.
type IprepareForReuse interface {
	ID() objc.ID
}

type prepareForReuse struct {
	id objc.ID
}

func prepareForReuseFrom(ptr unsafe.Pointer) prepareForReuse {
	return prepareForReuse{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ prepareForReuse) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _prepareForReuseClass) Alloc() prepareForReuse {
	rv := objc.Send[prepareForReuse](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _prepareForReuseClass) New() prepareForReuse {
	rv := objc.Send[prepareForReuse](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewprepareForReuse creates and returns a new initialized instance.
func NewprepareForReuse() prepareForReuse {
	return prepareForReuseClass.New()
}

// Init initializes the instance.
func (p_ prepareForReuse) Init() prepareForReuse {
	rv := objc.Send[prepareForReuse](p_.ID(), selInit)
	return rv
}
