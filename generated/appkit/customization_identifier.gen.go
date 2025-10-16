
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [customizationIdentifier] class.
var customizationIdentifierClass _customizationIdentifierClass

func init() {
	customizationIdentifierClass = _customizationIdentifierClass{objc.GetClass("customizationIdentifier")}
}

type _customizationIdentifierClass struct {
	objc.Class
}

// An interface definition for the [customizationIdentifier] class.
type IcustomizationIdentifier interface {
	ID() objc.ID
}

type customizationIdentifier struct {
	id objc.ID
}

func customizationIdentifierFrom(ptr unsafe.Pointer) customizationIdentifier {
	return customizationIdentifier{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ customizationIdentifier) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _customizationIdentifierClass) Alloc() customizationIdentifier {
	rv := objc.Send[customizationIdentifier](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _customizationIdentifierClass) New() customizationIdentifier {
	rv := objc.Send[customizationIdentifier](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcustomizationIdentifier creates and returns a new initialized instance.
func NewcustomizationIdentifier() customizationIdentifier {
	return customizationIdentifierClass.New()
}

// Init initializes the instance.
func (c_ customizationIdentifier) Init() customizationIdentifier {
	rv := objc.Send[customizationIdentifier](c_.ID(), selInit)
	return rv
}
