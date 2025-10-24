// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [cursor] class.
var (
	CursorClass     _cursorClass
	CursorClassOnce sync.Once
)

func getcursorClass() _cursorClass {
	CursorClassOnce.Do(func() {
		CursorClass = _cursorClass{objc.GetClass("cursor")}
	})
	return CursorClass
}

type _cursorClass struct {
	class objc.Class
}

// An interface definition for the [cursor] class.
type Icursor interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/cursor
type cursor struct {
	objectivec.Object
}

// cursorFrom constructs a [cursor] from an unsafe.Pointer.
func cursorFrom(ptr unsafe.Pointer) cursor {
	return cursor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _cursorClass) Alloc() cursor {
	rv := objc.Send[cursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _cursorClass) New() cursor {
	rv := objc.Send[cursor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ cursor) Init() cursor {
	rv := objc.Send[cursor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ cursor) Autorelease() cursor {
	rv := objc.Send[cursor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// Newcursor creates a new cursor instance.
func Newcursor() cursor {
	return getcursorClass().New()
}




