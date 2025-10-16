
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowedTouchTypes] class.
var allowedTouchTypesClass _allowedTouchTypesClass

func init() {
	allowedTouchTypesClass = _allowedTouchTypesClass{objc.GetClass("allowedTouchTypes")}
}

type _allowedTouchTypesClass struct {
	objc.Class
}

// An interface definition for the [allowedTouchTypes] class.
type IallowedTouchTypes interface {
	ID() objc.ID
}

type allowedTouchTypes struct {
	id objc.ID
}

func allowedTouchTypesFrom(ptr unsafe.Pointer) allowedTouchTypes {
	return allowedTouchTypes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowedTouchTypes) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowedTouchTypesClass) Alloc() allowedTouchTypes {
	rv := objc.Send[allowedTouchTypes](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowedTouchTypesClass) New() allowedTouchTypes {
	rv := objc.Send[allowedTouchTypes](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowedTouchTypes creates and returns a new initialized instance.
func NewallowedTouchTypes() allowedTouchTypes {
	return allowedTouchTypesClass.New()
}

// Init initializes the instance.
func (a_ allowedTouchTypes) Init() allowedTouchTypes {
	rv := objc.Send[allowedTouchTypes](a_.ID(), selInit)
	return rv
}
