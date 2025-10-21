// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CSRInfo] class.
var (
	CSRInfoClass     _CSRInfoClass
	CSRInfoClassOnce sync.Once
)

func getCSRInfoClass() _CSRInfoClass {
	CSRInfoClassOnce.Do(func() {
		CSRInfoClass = _CSRInfoClass{objc.GetClass("CSRInfo")}
	})
	return CSRInfoClass
}

type _CSRInfoClass struct {
	class objc.Class
}

// An interface definition for the [CSRInfo] class.
type ICSRInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo
type CSRInfo struct {
	objectivec.Object
}

// CSRInfoFrom constructs a [CSRInfo] from an unsafe.Pointer.
func CSRInfoFrom(ptr unsafe.Pointer) CSRInfo {
	return CSRInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CSRInfoClass) Alloc() CSRInfo {
	rv := objc.Send[CSRInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSRInfoClass) New() CSRInfo {
	rv := objc.Send[CSRInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSRInfo) Init() CSRInfo {
	rv := objc.Send[CSRInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSRInfo) Autorelease() CSRInfo {
	rv := objc.Send[CSRInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSRInfo creates a new CSRInfo instance.
func NewCSRInfo() CSRInfo {
	return getCSRInfoClass().New()
}




