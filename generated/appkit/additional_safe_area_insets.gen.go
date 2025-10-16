
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [additionalSafeAreaInsets] class.
var additionalSafeAreaInsetsClass _additionalSafeAreaInsetsClass

func init() {
	additionalSafeAreaInsetsClass = _additionalSafeAreaInsetsClass{objc.GetClass("additionalSafeAreaInsets")}
}

type _additionalSafeAreaInsetsClass struct {
	objc.Class
}

// An interface definition for the [additionalSafeAreaInsets] class.
type IadditionalSafeAreaInsets interface {
	ID() objc.ID
}

type additionalSafeAreaInsets struct {
	id objc.ID
}

func additionalSafeAreaInsetsFrom(ptr unsafe.Pointer) additionalSafeAreaInsets {
	return additionalSafeAreaInsets{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ additionalSafeAreaInsets) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _additionalSafeAreaInsetsClass) Alloc() additionalSafeAreaInsets {
	rv := objc.Send[additionalSafeAreaInsets](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _additionalSafeAreaInsetsClass) New() additionalSafeAreaInsets {
	rv := objc.Send[additionalSafeAreaInsets](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewadditionalSafeAreaInsets creates and returns a new initialized instance.
func NewadditionalSafeAreaInsets() additionalSafeAreaInsets {
	return additionalSafeAreaInsetsClass.New()
}

// Init initializes the instance.
func (a_ additionalSafeAreaInsets) Init() additionalSafeAreaInsets {
	rv := objc.Send[additionalSafeAreaInsets](a_.ID(), selInit)
	return rv
}
