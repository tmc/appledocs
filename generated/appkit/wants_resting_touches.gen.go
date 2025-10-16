
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [wantsRestingTouches] class.
var wantsRestingTouchesClass _wantsRestingTouchesClass

func init() {
	wantsRestingTouchesClass = _wantsRestingTouchesClass{objc.GetClass("wantsRestingTouches")}
}

type _wantsRestingTouchesClass struct {
	objc.Class
}

// An interface definition for the [wantsRestingTouches] class.
type IwantsRestingTouches interface {
	ID() objc.ID
}

type wantsRestingTouches struct {
	id objc.ID
}

func wantsRestingTouchesFrom(ptr unsafe.Pointer) wantsRestingTouches {
	return wantsRestingTouches{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ wantsRestingTouches) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _wantsRestingTouchesClass) Alloc() wantsRestingTouches {
	rv := objc.Send[wantsRestingTouches](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _wantsRestingTouchesClass) New() wantsRestingTouches {
	rv := objc.Send[wantsRestingTouches](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwantsRestingTouches creates and returns a new initialized instance.
func NewwantsRestingTouches() wantsRestingTouches {
	return wantsRestingTouchesClass.New()
}

// Init initializes the instance.
func (w_ wantsRestingTouches) Init() wantsRestingTouches {
	rv := objc.Send[wantsRestingTouches](w_.ID(), selInit)
	return rv
}
