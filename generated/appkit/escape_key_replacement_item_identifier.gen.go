
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [escapeKeyReplacementItemIdentifier] class.
var escapeKeyReplacementItemIdentifierClass _escapeKeyReplacementItemIdentifierClass

func init() {
	escapeKeyReplacementItemIdentifierClass = _escapeKeyReplacementItemIdentifierClass{objc.GetClass("escapeKeyReplacementItemIdentifier")}
}

type _escapeKeyReplacementItemIdentifierClass struct {
	objc.Class
}

// An interface definition for the [escapeKeyReplacementItemIdentifier] class.
type IescapeKeyReplacementItemIdentifier interface {
	ID() objc.ID
}

type escapeKeyReplacementItemIdentifier struct {
	id objc.ID
}

func escapeKeyReplacementItemIdentifierFrom(ptr unsafe.Pointer) escapeKeyReplacementItemIdentifier {
	return escapeKeyReplacementItemIdentifier{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ escapeKeyReplacementItemIdentifier) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _escapeKeyReplacementItemIdentifierClass) Alloc() escapeKeyReplacementItemIdentifier {
	rv := objc.Send[escapeKeyReplacementItemIdentifier](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _escapeKeyReplacementItemIdentifierClass) New() escapeKeyReplacementItemIdentifier {
	rv := objc.Send[escapeKeyReplacementItemIdentifier](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewescapeKeyReplacementItemIdentifier creates and returns a new initialized instance.
func NewescapeKeyReplacementItemIdentifier() escapeKeyReplacementItemIdentifier {
	return escapeKeyReplacementItemIdentifierClass.New()
}

// Init initializes the instance.
func (e_ escapeKeyReplacementItemIdentifier) Init() escapeKeyReplacementItemIdentifier {
	rv := objc.Send[escapeKeyReplacementItemIdentifier](e_.ID(), selInit)
	return rv
}
