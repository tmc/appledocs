// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trustType] class.
var (
	TrustTypeClass     _trustTypeClass
	TrustTypeClassOnce sync.Once
)

func gettrustTypeClass() _trustTypeClass {
	TrustTypeClassOnce.Do(func() {
		TrustTypeClass = _trustTypeClass{objc.GetClass("trustType")}
	})
	return TrustTypeClass
}

type _trustTypeClass struct {
	class objc.Class
}

// An interface definition for the [trustType] class.
type ItrustType interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustType-c.ivar
type trustType struct {
	objectivec.Object
}

// trustTypeFrom constructs a [trustType] from an unsafe.Pointer.
func trustTypeFrom(ptr unsafe.Pointer) trustType {
	return trustType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _trustTypeClass) Alloc() trustType {
	rv := objc.Send[trustType](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _trustTypeClass) New() trustType {
	rv := objc.Send[trustType](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustType) Init() trustType {
	rv := objc.Send[trustType](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustType) Autorelease() trustType {
	rv := objc.Send[trustType](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustType creates a new trustType instance.
func NewtrustType() trustType {
	return gettrustTypeClass().New()
}




