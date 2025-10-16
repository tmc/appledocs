
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [interfaceStyle] class.
var interfaceStyleClass _interfaceStyleClass

func init() {
	interfaceStyleClass = _interfaceStyleClass{objc.GetClass("interfaceStyle")}
}

type _interfaceStyleClass struct {
	objc.Class
}

// An interface definition for the [interfaceStyle] class.
type IinterfaceStyle interface {
	ID() objc.ID
}

type interfaceStyle struct {
	id objc.ID
}

func interfaceStyleFrom(ptr unsafe.Pointer) interfaceStyle {
	return interfaceStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ interfaceStyle) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _interfaceStyleClass) Alloc() interfaceStyle {
	rv := objc.Send[interfaceStyle](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _interfaceStyleClass) New() interfaceStyle {
	rv := objc.Send[interfaceStyle](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinterfaceStyle creates and returns a new initialized instance.
func NewinterfaceStyle() interfaceStyle {
	return interfaceStyleClass.New()
}

// Init initializes the instance.
func (i_ interfaceStyle) Init() interfaceStyle {
	rv := objc.Send[interfaceStyle](i_.ID(), selInit)
	return rv
}
