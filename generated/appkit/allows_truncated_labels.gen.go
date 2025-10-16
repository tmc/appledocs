
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsTruncatedLabels] class.
var allowsTruncatedLabelsClass _allowsTruncatedLabelsClass

func init() {
	allowsTruncatedLabelsClass = _allowsTruncatedLabelsClass{objc.GetClass("allowsTruncatedLabels")}
}

type _allowsTruncatedLabelsClass struct {
	objc.Class
}

// An interface definition for the [allowsTruncatedLabels] class.
type IallowsTruncatedLabels interface {
	ID() objc.ID
}

type allowsTruncatedLabels struct {
	id objc.ID
}

func allowsTruncatedLabelsFrom(ptr unsafe.Pointer) allowsTruncatedLabels {
	return allowsTruncatedLabels{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsTruncatedLabels) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsTruncatedLabelsClass) Alloc() allowsTruncatedLabels {
	rv := objc.Send[allowsTruncatedLabels](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsTruncatedLabelsClass) New() allowsTruncatedLabels {
	rv := objc.Send[allowsTruncatedLabels](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsTruncatedLabels creates and returns a new initialized instance.
func NewallowsTruncatedLabels() allowsTruncatedLabels {
	return allowsTruncatedLabelsClass.New()
}

// Init initializes the instance.
func (a_ allowsTruncatedLabels) Init() allowsTruncatedLabels {
	rv := objc.Send[allowsTruncatedLabels](a_.ID(), selInit)
	return rv
}
