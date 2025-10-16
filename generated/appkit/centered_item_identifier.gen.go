
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [centeredItemIdentifier] class.
var centeredItemIdentifierClass _centeredItemIdentifierClass

func init() {
	centeredItemIdentifierClass = _centeredItemIdentifierClass{objc.GetClass("centeredItemIdentifier")}
}

type _centeredItemIdentifierClass struct {
	objc.Class
}

// An interface definition for the [centeredItemIdentifier] class.
type IcenteredItemIdentifier interface {
	ID() objc.ID
}

type centeredItemIdentifier struct {
	id objc.ID
}

func centeredItemIdentifierFrom(ptr unsafe.Pointer) centeredItemIdentifier {
	return centeredItemIdentifier{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ centeredItemIdentifier) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _centeredItemIdentifierClass) Alloc() centeredItemIdentifier {
	rv := objc.Send[centeredItemIdentifier](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _centeredItemIdentifierClass) New() centeredItemIdentifier {
	rv := objc.Send[centeredItemIdentifier](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcenteredItemIdentifier creates and returns a new initialized instance.
func NewcenteredItemIdentifier() centeredItemIdentifier {
	return centeredItemIdentifierClass.New()
}

// Init initializes the instance.
func (c_ centeredItemIdentifier) Init() centeredItemIdentifier {
	rv := objc.Send[centeredItemIdentifier](c_.ID(), selInit)
	return rv
}
