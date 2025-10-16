
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [wantsNotificationForMarkedText] class.
var wantsNotificationForMarkedTextClass _wantsNotificationForMarkedTextClass

func init() {
	wantsNotificationForMarkedTextClass = _wantsNotificationForMarkedTextClass{objc.GetClass("wantsNotificationForMarkedText")}
}

type _wantsNotificationForMarkedTextClass struct {
	objc.Class
}

// An interface definition for the [wantsNotificationForMarkedText] class.
type IwantsNotificationForMarkedText interface {
	ID() objc.ID
}

type wantsNotificationForMarkedText struct {
	id objc.ID
}

func wantsNotificationForMarkedTextFrom(ptr unsafe.Pointer) wantsNotificationForMarkedText {
	return wantsNotificationForMarkedText{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ wantsNotificationForMarkedText) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _wantsNotificationForMarkedTextClass) Alloc() wantsNotificationForMarkedText {
	rv := objc.Send[wantsNotificationForMarkedText](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _wantsNotificationForMarkedTextClass) New() wantsNotificationForMarkedText {
	rv := objc.Send[wantsNotificationForMarkedText](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwantsNotificationForMarkedText creates and returns a new initialized instance.
func NewwantsNotificationForMarkedText() wantsNotificationForMarkedText {
	return wantsNotificationForMarkedTextClass.New()
}

// Init initializes the instance.
func (w_ wantsNotificationForMarkedText) Init() wantsNotificationForMarkedText {
	rv := objc.Send[wantsNotificationForMarkedText](w_.ID(), selInit)
	return rv
}
