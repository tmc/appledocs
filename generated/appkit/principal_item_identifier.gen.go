
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [principalItemIdentifier] class.
var principalItemIdentifierClass _principalItemIdentifierClass

func init() {
	principalItemIdentifierClass = _principalItemIdentifierClass{objc.GetClass("principalItemIdentifier")}
}

type _principalItemIdentifierClass struct {
	objc.Class
}

// An interface definition for the [principalItemIdentifier] class.
type IprincipalItemIdentifier interface {
	ID() objc.ID
}

type principalItemIdentifier struct {
	id objc.ID
}

func principalItemIdentifierFrom(ptr unsafe.Pointer) principalItemIdentifier {
	return principalItemIdentifier{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ principalItemIdentifier) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _principalItemIdentifierClass) Alloc() principalItemIdentifier {
	rv := objc.Send[principalItemIdentifier](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _principalItemIdentifierClass) New() principalItemIdentifier {
	rv := objc.Send[principalItemIdentifier](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewprincipalItemIdentifier creates and returns a new initialized instance.
func NewprincipalItemIdentifier() principalItemIdentifier {
	return principalItemIdentifierClass.New()
}

// Init initializes the instance.
func (p_ principalItemIdentifier) Init() principalItemIdentifier {
	rv := objc.Send[principalItemIdentifier](p_.ID(), selInit)
	return rv
}
