// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cursor] class.
var (
	CursorClass     _CursorClass
	CursorClassOnce sync.Once
)

func getCursorClass() _CursorClass {
	CursorClassOnce.Do(func() {
		CursorClass = _CursorClass{objc.GetClass("NSCursor")}
	})
	return CursorClass
}

type _CursorClass struct {
	class objc.Class
}

// An interface definition for the [Cursor] class.
type ICursor interface {
	objectivec.IObject
	Set()
}

// A pointer (also called a cursor).
//
// The following table shows and describes the system cursors, and indicates the class method for obtaining them: In macOS 10.3 and later, cursor size is no longer limited to 16 by 16 pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor
type Cursor struct {
	objectivec.Object
}

// CursorFrom constructs a [Cursor] from an unsafe.Pointer.
//
// A pointer (also called a cursor).
func CursorFrom(ptr unsafe.Pointer) Cursor {
	return Cursor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CursorClass) Alloc() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CursorClass) New() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Cursor) Init() Cursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Cursor) Autorelease() Cursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCursor creates a new Cursor instance.
func NewCursor() Cursor {
	return getCursorClass().New()
}

// Returns a cursor indicating that the current operation will result in a link action.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (cc _CursorClass) DragLinkCursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("dragLinkCursor"))
	return rv
}

// Returns the zoom-out cursor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomOut
func (cc _CursorClass) ZoomOutCursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("zoomOutCursor"))
	return rv
}

// Makes the receiver the current cursor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/set()
func (c_ Cursor) Set() {
	objc.Send[objc.ID](c_.ID, objc.Sel("set"))
}

// Returns a cursor indicating that the current operation will result in a link action.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (c_ Cursor) DragLinkCursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dragLinkCursor"))
	return rv
}

// Returns the zoom-out cursor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomOut
func (c_ Cursor) ZoomOutCursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("zoomOutCursor"))
	return rv
}
