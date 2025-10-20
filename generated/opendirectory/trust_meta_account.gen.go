// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trustMetaAccount] class.
var (
	TrustMetaAccountClass     _trustMetaAccountClass
	TrustMetaAccountClassOnce sync.Once
)

func gettrustMetaAccountClass() _trustMetaAccountClass {
	TrustMetaAccountClassOnce.Do(func() {
		TrustMetaAccountClass = _trustMetaAccountClass{objc.GetClass("trustMetaAccount")}
	})
	return TrustMetaAccountClass
}

type _trustMetaAccountClass struct {
	class objc.Class
}

// An interface definition for the [trustMetaAccount] class.
type ItrustMetaAccount interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustMetaAccount-c.ivar
type trustMetaAccount struct {
	objectivec.Object
}

// trustMetaAccountFrom constructs a [trustMetaAccount] from an unsafe.Pointer.
func trustMetaAccountFrom(ptr unsafe.Pointer) trustMetaAccount {
	return trustMetaAccount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _trustMetaAccountClass) Alloc() trustMetaAccount {
	rv := objc.Send[trustMetaAccount](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _trustMetaAccountClass) New() trustMetaAccount {
	rv := objc.Send[trustMetaAccount](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustMetaAccount) Init() trustMetaAccount {
	rv := objc.Send[trustMetaAccount](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustMetaAccount) Autorelease() trustMetaAccount {
	rv := objc.Send[trustMetaAccount](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustMetaAccount creates a new trustMetaAccount instance.
func NewtrustMetaAccount() trustMetaAccount {
	return gettrustMetaAccountClass().New()
}




