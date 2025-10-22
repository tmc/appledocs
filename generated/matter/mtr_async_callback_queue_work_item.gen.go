// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAsyncCallbackQueueWorkItem] class.
var (
	MTRAsyncCallbackQueueWorkItemClass     _MTRAsyncCallbackQueueWorkItemClass
	MTRAsyncCallbackQueueWorkItemClassOnce sync.Once
)

func getMTRAsyncCallbackQueueWorkItemClass() _MTRAsyncCallbackQueueWorkItemClass {
	MTRAsyncCallbackQueueWorkItemClassOnce.Do(func() {
		MTRAsyncCallbackQueueWorkItemClass = _MTRAsyncCallbackQueueWorkItemClass{objc.GetClass("MTRAsyncCallbackQueueWorkItem")}
	})
	return MTRAsyncCallbackQueueWorkItemClass
}

type _MTRAsyncCallbackQueueWorkItemClass struct {
	class objc.Class
}

// An interface definition for the [MTRAsyncCallbackQueueWorkItem] class.
type IMTRAsyncCallbackQueueWorkItem interface {
	objectivec.IObject
	CancelHandler() unsafe.Pointer
	SetCancelHandler(value unsafe.Pointer)
	ReadyHandler() unsafe.Pointer
	SetReadyHandler(value unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackQueueWorkItem
type MTRAsyncCallbackQueueWorkItem struct {
	objectivec.Object
}

// MTRAsyncCallbackQueueWorkItemFrom constructs a [MTRAsyncCallbackQueueWorkItem] from an unsafe.Pointer.
func MTRAsyncCallbackQueueWorkItemFrom(ptr unsafe.Pointer) MTRAsyncCallbackQueueWorkItem {
	return MTRAsyncCallbackQueueWorkItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAsyncCallbackQueueWorkItemClass) Alloc() MTRAsyncCallbackQueueWorkItem {
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAsyncCallbackQueueWorkItemClass) New() MTRAsyncCallbackQueueWorkItem {
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAsyncCallbackQueueWorkItem) Init() MTRAsyncCallbackQueueWorkItem {
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAsyncCallbackQueueWorkItem) Autorelease() MTRAsyncCallbackQueueWorkItem {
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAsyncCallbackQueueWorkItem creates a new MTRAsyncCallbackQueueWorkItem instance.
func NewMTRAsyncCallbackQueueWorkItem() MTRAsyncCallbackQueueWorkItem {
	return getMTRAsyncCallbackQueueWorkItemClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrasynccallbackqueueworkitem/cancelhandler
func (m_ MTRAsyncCallbackQueueWorkItem) CancelHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cancelHandler"))
	return rv
}


// SetCancelHandler sets the value of the cancelHandler property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrasynccallbackqueueworkitem/cancelhandler
func (m_ MTRAsyncCallbackQueueWorkItem) SetCancelHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCancelHandler:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrasynccallbackqueueworkitem/readyhandler
func (m_ MTRAsyncCallbackQueueWorkItem) ReadyHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readyHandler"))
	return rv
}


// SetReadyHandler sets the value of the readyHandler property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrasynccallbackqueueworkitem/readyhandler
func (m_ MTRAsyncCallbackQueueWorkItem) SetReadyHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReadyHandler:"), value)
}



