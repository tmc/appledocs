
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [usesSingleLineMode] class.
var usesSingleLineModeClass _usesSingleLineModeClass

func init() {
	usesSingleLineModeClass = _usesSingleLineModeClass{objc.GetClass("usesSingleLineMode")}
}

type _usesSingleLineModeClass struct {
	objc.Class
}

// An interface definition for the [usesSingleLineMode] class.
type IusesSingleLineMode interface {
	ID() objc.ID
}

type usesSingleLineMode struct {
	id objc.ID
}

func usesSingleLineModeFrom(ptr unsafe.Pointer) usesSingleLineMode {
	return usesSingleLineMode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ usesSingleLineMode) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _usesSingleLineModeClass) Alloc() usesSingleLineMode {
	rv := objc.Send[usesSingleLineMode](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _usesSingleLineModeClass) New() usesSingleLineMode {
	rv := objc.Send[usesSingleLineMode](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewusesSingleLineMode creates and returns a new initialized instance.
func NewusesSingleLineMode() usesSingleLineMode {
	return usesSingleLineModeClass.New()
}

// Init initializes the instance.
func (u_ usesSingleLineMode) Init() usesSingleLineMode {
	rv := objc.Send[usesSingleLineMode](u_.ID(), selInit)
	return rv
}
