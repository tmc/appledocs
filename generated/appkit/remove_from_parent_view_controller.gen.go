
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [removeFromParentViewController] class.
var removeFromParentViewControllerClass _removeFromParentViewControllerClass

func init() {
	removeFromParentViewControllerClass = _removeFromParentViewControllerClass{objc.GetClass("removeFromParentViewController")}
}

type _removeFromParentViewControllerClass struct {
	objc.Class
}

// An interface definition for the [removeFromParentViewController] class.
type IremoveFromParentViewController interface {
	ID() objc.ID
}

type removeFromParentViewController struct {
	id objc.ID
}

func removeFromParentViewControllerFrom(ptr unsafe.Pointer) removeFromParentViewController {
	return removeFromParentViewController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ removeFromParentViewController) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _removeFromParentViewControllerClass) Alloc() removeFromParentViewController {
	rv := objc.Send[removeFromParentViewController](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _removeFromParentViewControllerClass) New() removeFromParentViewController {
	rv := objc.Send[removeFromParentViewController](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewremoveFromParentViewController creates and returns a new initialized instance.
func NewremoveFromParentViewController() removeFromParentViewController {
	return removeFromParentViewControllerClass.New()
}

// Init initializes the instance.
func (r_ removeFromParentViewController) Init() removeFromParentViewController {
	rv := objc.Send[removeFromParentViewController](r_.ID(), selInit)
	return rv
}
