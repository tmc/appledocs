
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [identifier] class.
var identifierClass _identifierClass

func init() {
	identifierClass = _identifierClass{objc.GetClass("identifier")}
}

type _identifierClass struct {
	objc.Class
}

// An interface definition for the [identifier] class.
type Iidentifier interface {
	ID() objc.ID
}

type identifier struct {
	id objc.ID
}

func identifierFrom(ptr unsafe.Pointer) identifier {
	return identifier{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ identifier) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _identifierClass) Alloc() identifier {
	rv := objc.Send[identifier](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _identifierClass) New() identifier {
	rv := objc.Send[identifier](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newidentifier creates and returns a new initialized instance.
func Newidentifier() identifier {
	return identifierClass.New()
}

// Init initializes the instance.
func (i_ identifier) Init() identifier {
	rv := objc.Send[identifier](i_.ID(), selInit)
	return rv
}
