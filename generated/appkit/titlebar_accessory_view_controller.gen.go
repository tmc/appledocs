
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TitlebarAccessoryViewController] class.
var TitlebarAccessoryViewControllerClass _TitlebarAccessoryViewControllerClass

func init() {
	TitlebarAccessoryViewControllerClass = _TitlebarAccessoryViewControllerClass{objc.GetClass("NSTitlebarAccessoryViewController")}
}

type _TitlebarAccessoryViewControllerClass struct {
	objc.Class
}

// An interface definition for the [TitlebarAccessoryViewController] class.
type ITitlebarAccessoryViewController interface {
	ID() objc.ID
}

type TitlebarAccessoryViewController struct {
	id objc.ID
}

func TitlebarAccessoryViewControllerFrom(ptr unsafe.Pointer) TitlebarAccessoryViewController {
	return TitlebarAccessoryViewController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TitlebarAccessoryViewController) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TitlebarAccessoryViewControllerClass) Alloc() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TitlebarAccessoryViewControllerClass) New() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTitlebarAccessoryViewController creates and returns a new initialized instance.
func NewTitlebarAccessoryViewController() TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.New()
}

// Init initializes the instance.
func (t_ TitlebarAccessoryViewController) Init() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](t_.ID(), selInit)
	return rv
}
