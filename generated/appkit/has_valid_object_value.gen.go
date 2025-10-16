
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasValidObjectValue] class.
var hasValidObjectValueClass _hasValidObjectValueClass

func init() {
	hasValidObjectValueClass = _hasValidObjectValueClass{objc.GetClass("hasValidObjectValue")}
}

type _hasValidObjectValueClass struct {
	objc.Class
}

// An interface definition for the [hasValidObjectValue] class.
type IhasValidObjectValue interface {
	ID() objc.ID
}

type hasValidObjectValue struct {
	id objc.ID
}

func hasValidObjectValueFrom(ptr unsafe.Pointer) hasValidObjectValue {
	return hasValidObjectValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasValidObjectValue) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasValidObjectValueClass) Alloc() hasValidObjectValue {
	rv := objc.Send[hasValidObjectValue](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasValidObjectValueClass) New() hasValidObjectValue {
	rv := objc.Send[hasValidObjectValue](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasValidObjectValue creates and returns a new initialized instance.
func NewhasValidObjectValue() hasValidObjectValue {
	return hasValidObjectValueClass.New()
}

// Init initializes the instance.
func (h_ hasValidObjectValue) Init() hasValidObjectValue {
	rv := objc.Send[hasValidObjectValue](h_.ID(), selInit)
	return rv
}
