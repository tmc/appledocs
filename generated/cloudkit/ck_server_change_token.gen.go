// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKServerChangeToken] class.
var (
	CKServerChangeTokenClass     _CKServerChangeTokenClass
	CKServerChangeTokenClassOnce sync.Once
)

func getCKServerChangeTokenClass() _CKServerChangeTokenClass {
	CKServerChangeTokenClassOnce.Do(func() {
		CKServerChangeTokenClass = _CKServerChangeTokenClass{objc.GetClass("CKServerChangeToken")}
	})
	return CKServerChangeTokenClass
}

type _CKServerChangeTokenClass struct {
	class objc.Class
}

// An interface definition for the [CKServerChangeToken] class.
type ICKServerChangeToken interface {
	objectivec.IObject
}

// An opaque token that represents a specific point in a database’s history.
//
// CloudKit uses server change tokens to record significant events in a database’s history, such as record creation, modification, and deletion. Using change tokens helps reduce the cost of a fetch operation — both the time to execute the fetch and the overall number of records it returns. You don’t create change tokens. Instead, and provide them during their execution and when they complete. Cache each token as you receive it, overwriting any previous token for the database or record zone you’re fetching from. Then, pass the cached token with your next fetch and CloudKit returns only the changes that occur after that point. Don’t infer any behavior or order from a token’s contents. The change tokens that provides aren’t compatible with and vice versa, so segregate them in your cache. Change tokens conform to and are safe to cache on-disk, as the following example shows:
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKServerChangeToken
type CKServerChangeToken struct {
	objectivec.Object
}

// CKServerChangeTokenFrom constructs a [CKServerChangeToken] from an unsafe.Pointer.
//
// An opaque token that represents a specific point in a database’s history.
func CKServerChangeTokenFrom(ptr unsafe.Pointer) CKServerChangeToken {
	return CKServerChangeToken{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKServerChangeTokenClass) Alloc() CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKServerChangeTokenClass) New() CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKServerChangeToken) Init() CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKServerChangeToken) Autorelease() CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKServerChangeToken creates a new CKServerChangeToken instance.
func NewCKServerChangeToken() CKServerChangeToken {
	return getCKServerChangeTokenClass().New()
}




