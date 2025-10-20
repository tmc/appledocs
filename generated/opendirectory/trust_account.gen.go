// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trustAccount] class.
var (
	TrustAccountClass     _trustAccountClass
	TrustAccountClassOnce sync.Once
)

func gettrustAccountClass() _trustAccountClass {
	TrustAccountClassOnce.Do(func() {
		TrustAccountClass = _trustAccountClass{objc.GetClass("trustAccount")}
	})
	return TrustAccountClass
}

type _trustAccountClass struct {
	class objc.Class
}

// An interface definition for the [trustAccount] class.
type ItrustAccount interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustAccount-c.ivar
type trustAccount struct {
	objectivec.Object
}

// trustAccountFrom constructs a [trustAccount] from an unsafe.Pointer.
func trustAccountFrom(ptr unsafe.Pointer) trustAccount {
	return trustAccount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _trustAccountClass) Alloc() trustAccount {
	rv := objc.Send[trustAccount](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _trustAccountClass) New() trustAccount {
	rv := objc.Send[trustAccount](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustAccount) Init() trustAccount {
	rv := objc.Send[trustAccount](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustAccount) Autorelease() trustAccount {
	rv := objc.Send[trustAccount](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustAccount creates a new trustAccount instance.
func NewtrustAccount() trustAccount {
	return gettrustAccountClass().New()
}




