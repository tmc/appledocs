
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [disableScreenUpdatesUntilFlush] class.
var disableScreenUpdatesUntilFlushClass _disableScreenUpdatesUntilFlushClass

func init() {
	disableScreenUpdatesUntilFlushClass = _disableScreenUpdatesUntilFlushClass{objc.GetClass("disableScreenUpdatesUntilFlush")}
}

type _disableScreenUpdatesUntilFlushClass struct {
	objc.Class
}

// An interface definition for the [disableScreenUpdatesUntilFlush] class.
type IdisableScreenUpdatesUntilFlush interface {
	ID() objc.ID
}

type disableScreenUpdatesUntilFlush struct {
	id objc.ID
}

func disableScreenUpdatesUntilFlushFrom(ptr unsafe.Pointer) disableScreenUpdatesUntilFlush {
	return disableScreenUpdatesUntilFlush{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ disableScreenUpdatesUntilFlush) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _disableScreenUpdatesUntilFlushClass) Alloc() disableScreenUpdatesUntilFlush {
	rv := objc.Send[disableScreenUpdatesUntilFlush](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _disableScreenUpdatesUntilFlushClass) New() disableScreenUpdatesUntilFlush {
	rv := objc.Send[disableScreenUpdatesUntilFlush](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisableScreenUpdatesUntilFlush creates and returns a new initialized instance.
func NewdisableScreenUpdatesUntilFlush() disableScreenUpdatesUntilFlush {
	return disableScreenUpdatesUntilFlushClass.New()
}

// Init initializes the instance.
func (d_ disableScreenUpdatesUntilFlush) Init() disableScreenUpdatesUntilFlush {
	rv := objc.Send[disableScreenUpdatesUntilFlush](d_.ID(), selInit)
	return rv
}
