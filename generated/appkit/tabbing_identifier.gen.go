
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tabbingIdentifier] class.
var tabbingIdentifierClass _tabbingIdentifierClass

func init() {
	tabbingIdentifierClass = _tabbingIdentifierClass{objc.GetClass("tabbingIdentifier")}
}

type _tabbingIdentifierClass struct {
	objc.Class
}

// An interface definition for the [tabbingIdentifier] class.
type ItabbingIdentifier interface {
	ID() objc.ID
}

type tabbingIdentifier struct {
	id objc.ID
}

func tabbingIdentifierFrom(ptr unsafe.Pointer) tabbingIdentifier {
	return tabbingIdentifier{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tabbingIdentifier) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabbingIdentifierClass) Alloc() tabbingIdentifier {
	rv := objc.Send[tabbingIdentifier](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabbingIdentifierClass) New() tabbingIdentifier {
	rv := objc.Send[tabbingIdentifier](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtabbingIdentifier creates and returns a new initialized instance.
func NewtabbingIdentifier() tabbingIdentifier {
	return tabbingIdentifierClass.New()
}

// Init initializes the instance.
func (t_ tabbingIdentifier) Init() tabbingIdentifier {
	rv := objc.Send[tabbingIdentifier](t_.ID(), selInit)
	return rv
}
