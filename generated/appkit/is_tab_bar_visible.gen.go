
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isTabBarVisible] class.
var isTabBarVisibleClass _isTabBarVisibleClass

func init() {
	isTabBarVisibleClass = _isTabBarVisibleClass{objc.GetClass("isTabBarVisible")}
}

type _isTabBarVisibleClass struct {
	objc.Class
}

// An interface definition for the [isTabBarVisible] class.
type IisTabBarVisible interface {
	ID() objc.ID
}

type isTabBarVisible struct {
	id objc.ID
}

func isTabBarVisibleFrom(ptr unsafe.Pointer) isTabBarVisible {
	return isTabBarVisible{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isTabBarVisible) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isTabBarVisibleClass) Alloc() isTabBarVisible {
	rv := objc.Send[isTabBarVisible](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isTabBarVisibleClass) New() isTabBarVisible {
	rv := objc.Send[isTabBarVisible](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisTabBarVisible creates and returns a new initialized instance.
func NewisTabBarVisible() isTabBarVisible {
	return isTabBarVisibleClass.New()
}

// Init initializes the instance.
func (i_ isTabBarVisible) Init() isTabBarVisible {
	rv := objc.Send[isTabBarVisible](i_.ID(), selInit)
	return rv
}
