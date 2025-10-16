
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectedItemIdentifier] class.
var selectedItemIdentifierClass _selectedItemIdentifierClass

func init() {
	selectedItemIdentifierClass = _selectedItemIdentifierClass{objc.GetClass("selectedItemIdentifier")}
}

type _selectedItemIdentifierClass struct {
	objc.Class
}

// An interface definition for the [selectedItemIdentifier] class.
type IselectedItemIdentifier interface {
	ID() objc.ID
}

type selectedItemIdentifier struct {
	id objc.ID
}

func selectedItemIdentifierFrom(ptr unsafe.Pointer) selectedItemIdentifier {
	return selectedItemIdentifier{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectedItemIdentifier) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectedItemIdentifierClass) Alloc() selectedItemIdentifier {
	rv := objc.Send[selectedItemIdentifier](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectedItemIdentifierClass) New() selectedItemIdentifier {
	rv := objc.Send[selectedItemIdentifier](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectedItemIdentifier creates and returns a new initialized instance.
func NewselectedItemIdentifier() selectedItemIdentifier {
	return selectedItemIdentifierClass.New()
}

// Init initializes the instance.
func (s_ selectedItemIdentifier) Init() selectedItemIdentifier {
	rv := objc.Send[selectedItemIdentifier](s_.ID(), selInit)
	return rv
}
