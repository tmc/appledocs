
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsExpansionToolTips] class.
var allowsExpansionToolTipsClass _allowsExpansionToolTipsClass

func init() {
	allowsExpansionToolTipsClass = _allowsExpansionToolTipsClass{objc.GetClass("allowsExpansionToolTips")}
}

type _allowsExpansionToolTipsClass struct {
	objc.Class
}

// An interface definition for the [allowsExpansionToolTips] class.
type IallowsExpansionToolTips interface {
	ID() objc.ID
}

type allowsExpansionToolTips struct {
	id objc.ID
}

func allowsExpansionToolTipsFrom(ptr unsafe.Pointer) allowsExpansionToolTips {
	return allowsExpansionToolTips{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsExpansionToolTips) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsExpansionToolTipsClass) Alloc() allowsExpansionToolTips {
	rv := objc.Send[allowsExpansionToolTips](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsExpansionToolTipsClass) New() allowsExpansionToolTips {
	rv := objc.Send[allowsExpansionToolTips](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsExpansionToolTips creates and returns a new initialized instance.
func NewallowsExpansionToolTips() allowsExpansionToolTips {
	return allowsExpansionToolTipsClass.New()
}

// Init initializes the instance.
func (a_ allowsExpansionToolTips) Init() allowsExpansionToolTips {
	rv := objc.Send[allowsExpansionToolTips](a_.ID(), selInit)
	return rv
}
