
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [activeCompressionOptions] class.
var activeCompressionOptionsClass _activeCompressionOptionsClass

func init() {
	activeCompressionOptionsClass = _activeCompressionOptionsClass{objc.GetClass("activeCompressionOptions")}
}

type _activeCompressionOptionsClass struct {
	objc.Class
}

// An interface definition for the [activeCompressionOptions] class.
type IactiveCompressionOptions interface {
	ID() objc.ID
}

type activeCompressionOptions struct {
	id objc.ID
}

func activeCompressionOptionsFrom(ptr unsafe.Pointer) activeCompressionOptions {
	return activeCompressionOptions{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ activeCompressionOptions) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _activeCompressionOptionsClass) Alloc() activeCompressionOptions {
	rv := objc.Send[activeCompressionOptions](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _activeCompressionOptionsClass) New() activeCompressionOptions {
	rv := objc.Send[activeCompressionOptions](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewactiveCompressionOptions creates and returns a new initialized instance.
func NewactiveCompressionOptions() activeCompressionOptions {
	return activeCompressionOptionsClass.New()
}

// Init initializes the instance.
func (a_ activeCompressionOptions) Init() activeCompressionOptions {
	rv := objc.Send[activeCompressionOptions](a_.ID(), selInit)
	return rv
}
