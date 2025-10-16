
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [shouldInvalidateLayoutForSelectionChange] class.
var shouldInvalidateLayoutForSelectionChangeClass _shouldInvalidateLayoutForSelectionChangeClass

func init() {
	shouldInvalidateLayoutForSelectionChangeClass = _shouldInvalidateLayoutForSelectionChangeClass{objc.GetClass("shouldInvalidateLayoutForSelectionChange")}
}

type _shouldInvalidateLayoutForSelectionChangeClass struct {
	objc.Class
}

// An interface definition for the [shouldInvalidateLayoutForSelectionChange] class.
type IshouldInvalidateLayoutForSelectionChange interface {
	ID() objc.ID
}

type shouldInvalidateLayoutForSelectionChange struct {
	id objc.ID
}

func shouldInvalidateLayoutForSelectionChangeFrom(ptr unsafe.Pointer) shouldInvalidateLayoutForSelectionChange {
	return shouldInvalidateLayoutForSelectionChange{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ shouldInvalidateLayoutForSelectionChange) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _shouldInvalidateLayoutForSelectionChangeClass) Alloc() shouldInvalidateLayoutForSelectionChange {
	rv := objc.Send[shouldInvalidateLayoutForSelectionChange](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _shouldInvalidateLayoutForSelectionChangeClass) New() shouldInvalidateLayoutForSelectionChange {
	rv := objc.Send[shouldInvalidateLayoutForSelectionChange](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshouldInvalidateLayoutForSelectionChange creates and returns a new initialized instance.
func NewshouldInvalidateLayoutForSelectionChange() shouldInvalidateLayoutForSelectionChange {
	return shouldInvalidateLayoutForSelectionChangeClass.New()
}

// Init initializes the instance.
func (s_ shouldInvalidateLayoutForSelectionChange) Init() shouldInvalidateLayoutForSelectionChange {
	rv := objc.Send[shouldInvalidateLayoutForSelectionChange](s_.ID(), selInit)
	return rv
}
