
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [preservesContentDuringLiveResize] class.
var preservesContentDuringLiveResizeClass _preservesContentDuringLiveResizeClass

func init() {
	preservesContentDuringLiveResizeClass = _preservesContentDuringLiveResizeClass{objc.GetClass("preservesContentDuringLiveResize")}
}

type _preservesContentDuringLiveResizeClass struct {
	objc.Class
}

// An interface definition for the [preservesContentDuringLiveResize] class.
type IpreservesContentDuringLiveResize interface {
	ID() objc.ID
}

type preservesContentDuringLiveResize struct {
	id objc.ID
}

func preservesContentDuringLiveResizeFrom(ptr unsafe.Pointer) preservesContentDuringLiveResize {
	return preservesContentDuringLiveResize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ preservesContentDuringLiveResize) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _preservesContentDuringLiveResizeClass) Alloc() preservesContentDuringLiveResize {
	rv := objc.Send[preservesContentDuringLiveResize](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _preservesContentDuringLiveResizeClass) New() preservesContentDuringLiveResize {
	rv := objc.Send[preservesContentDuringLiveResize](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreservesContentDuringLiveResize creates and returns a new initialized instance.
func NewpreservesContentDuringLiveResize() preservesContentDuringLiveResize {
	return preservesContentDuringLiveResizeClass.New()
}

// Init initializes the instance.
func (p_ preservesContentDuringLiveResize) Init() preservesContentDuringLiveResize {
	rv := objc.Send[preservesContentDuringLiveResize](p_.ID(), selInit)
	return rv
}
