
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [defaultItemIdentifiers] class.
var defaultItemIdentifiersClass _defaultItemIdentifiersClass

func init() {
	defaultItemIdentifiersClass = _defaultItemIdentifiersClass{objc.GetClass("defaultItemIdentifiers")}
}

type _defaultItemIdentifiersClass struct {
	objc.Class
}

// An interface definition for the [defaultItemIdentifiers] class.
type IdefaultItemIdentifiers interface {
	ID() objc.ID
}

type defaultItemIdentifiers struct {
	id objc.ID
}

func defaultItemIdentifiersFrom(ptr unsafe.Pointer) defaultItemIdentifiers {
	return defaultItemIdentifiers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ defaultItemIdentifiers) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _defaultItemIdentifiersClass) Alloc() defaultItemIdentifiers {
	rv := objc.Send[defaultItemIdentifiers](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _defaultItemIdentifiersClass) New() defaultItemIdentifiers {
	rv := objc.Send[defaultItemIdentifiers](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdefaultItemIdentifiers creates and returns a new initialized instance.
func NewdefaultItemIdentifiers() defaultItemIdentifiers {
	return defaultItemIdentifiersClass.New()
}

// Init initializes the instance.
func (d_ defaultItemIdentifiers) Init() defaultItemIdentifiers {
	rv := objc.Send[defaultItemIdentifiers](d_.ID(), selInit)
	return rv
}
