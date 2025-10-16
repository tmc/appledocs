
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [automaticModeOptions] class.
var automaticModeOptionsClass _automaticModeOptionsClass

func init() {
	automaticModeOptionsClass = _automaticModeOptionsClass{objc.GetClass("automaticModeOptions")}
}

type _automaticModeOptionsClass struct {
	objc.Class
}

// An interface definition for the [automaticModeOptions] class.
type IautomaticModeOptions interface {
	ID() objc.ID
}

type automaticModeOptions struct {
	id objc.ID
}

func automaticModeOptionsFrom(ptr unsafe.Pointer) automaticModeOptions {
	return automaticModeOptions{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ automaticModeOptions) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _automaticModeOptionsClass) Alloc() automaticModeOptions {
	rv := objc.Send[automaticModeOptions](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _automaticModeOptionsClass) New() automaticModeOptions {
	rv := objc.Send[automaticModeOptions](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautomaticModeOptions creates and returns a new initialized instance.
func NewautomaticModeOptions() automaticModeOptions {
	return automaticModeOptionsClass.New()
}

// Init initializes the instance.
func (a_ automaticModeOptions) Init() automaticModeOptions {
	rv := objc.Send[automaticModeOptions](a_.ID(), selInit)
	return rv
}
