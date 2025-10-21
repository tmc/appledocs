// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKQueryCursor] class.
var (
	CKQueryCursorClass     _CKQueryCursorClass
	CKQueryCursorClassOnce sync.Once
)

func getCKQueryCursorClass() _CKQueryCursorClass {
	CKQueryCursorClassOnce.Do(func() {
		CKQueryCursorClass = _CKQueryCursorClass{objc.GetClass("CKQueryCursor")}
	})
	return CKQueryCursorClass
}

type _CKQueryCursorClass struct {
	class objc.Class
}

// An interface definition for the [CKQueryCursor] class.
type ICKQueryCursor interface {
	objectivec.IObject
}

// An object that marks the stopping point for a query and the starting point for retrieving the remaining results.
//
// You don’t create instances of this class yourself. When fetching records using a query operation, if the number of results exceeds the limit for the query, CloudKit provides a cursor. Use that cursor to create a new instance of and retrieve the next batch of results for the same query. For information about how to use a object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/Cursor-swift.class
type CKQueryCursor struct {
	objectivec.Object
}

// CKQueryCursorFrom constructs a [CKQueryCursor] from an unsafe.Pointer.
//
// An object that marks the stopping point for a query and the starting point for retrieving the remaining results.
func CKQueryCursorFrom(ptr unsafe.Pointer) CKQueryCursor {
	return CKQueryCursor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKQueryCursorClass) Alloc() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKQueryCursorClass) New() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKQueryCursor) Init() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKQueryCursor) Autorelease() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQueryCursor creates a new CKQueryCursor instance.
func NewCKQueryCursor() CKQueryCursor {
	return getCKQueryCursorClass().New()
}




