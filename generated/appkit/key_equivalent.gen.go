
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [keyEquivalent] class.
var keyEquivalentClass _keyEquivalentClass

func init() {
	keyEquivalentClass = _keyEquivalentClass{objc.GetClass("keyEquivalent")}
}

type _keyEquivalentClass struct {
	objc.Class
}

// An interface definition for the [keyEquivalent] class.
type IkeyEquivalent interface {
	ID() objc.ID
}

type keyEquivalent struct {
	id objc.ID
}

func keyEquivalentFrom(ptr unsafe.Pointer) keyEquivalent {
	return keyEquivalent{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (k_ keyEquivalent) ID() objc.ID {
	return k_.id
}

// Alloc allocates a new instance without initialization.
func (kc _keyEquivalentClass) Alloc() keyEquivalent {
	rv := objc.Send[keyEquivalent](objc.ID(kc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (kc _keyEquivalentClass) New() keyEquivalent {
	rv := objc.Send[keyEquivalent](objc.ID(kc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewkeyEquivalent creates and returns a new initialized instance.
func NewkeyEquivalent() keyEquivalent {
	return keyEquivalentClass.New()
}

// Init initializes the instance.
func (k_ keyEquivalent) Init() keyEquivalent {
	rv := objc.Send[keyEquivalent](k_.ID(), selInit)
	return rv
}
