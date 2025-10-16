
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [orderFrontRegardless] class.
var orderFrontRegardlessClass _orderFrontRegardlessClass

func init() {
	orderFrontRegardlessClass = _orderFrontRegardlessClass{objc.GetClass("orderFrontRegardless")}
}

type _orderFrontRegardlessClass struct {
	objc.Class
}

// An interface definition for the [orderFrontRegardless] class.
type IorderFrontRegardless interface {
	ID() objc.ID
}

type orderFrontRegardless struct {
	id objc.ID
}

func orderFrontRegardlessFrom(ptr unsafe.Pointer) orderFrontRegardless {
	return orderFrontRegardless{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ orderFrontRegardless) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _orderFrontRegardlessClass) Alloc() orderFrontRegardless {
	rv := objc.Send[orderFrontRegardless](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _orderFrontRegardlessClass) New() orderFrontRegardless {
	rv := objc.Send[orderFrontRegardless](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NeworderFrontRegardless creates and returns a new initialized instance.
func NeworderFrontRegardless() orderFrontRegardless {
	return orderFrontRegardlessClass.New()
}

// Init initializes the instance.
func (o_ orderFrontRegardless) Init() orderFrontRegardless {
	rv := objc.Send[orderFrontRegardless](o_.ID(), selInit)
	return rv
}
