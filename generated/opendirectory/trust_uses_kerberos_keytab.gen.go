// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trustUsesKerberosKeytab] class.
var (
	TrustUsesKerberosKeytabClass     _trustUsesKerberosKeytabClass
	TrustUsesKerberosKeytabClassOnce sync.Once
)

func gettrustUsesKerberosKeytabClass() _trustUsesKerberosKeytabClass {
	TrustUsesKerberosKeytabClassOnce.Do(func() {
		TrustUsesKerberosKeytabClass = _trustUsesKerberosKeytabClass{objc.GetClass("trustUsesKerberosKeytab")}
	})
	return TrustUsesKerberosKeytabClass
}

type _trustUsesKerberosKeytabClass struct {
	class objc.Class
}

// An interface definition for the [trustUsesKerberosKeytab] class.
type ItrustUsesKerberosKeytab interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesKerberosKeytab-c.ivar
type trustUsesKerberosKeytab struct {
	objectivec.Object
}

// trustUsesKerberosKeytabFrom constructs a [trustUsesKerberosKeytab] from an unsafe.Pointer.
func trustUsesKerberosKeytabFrom(ptr unsafe.Pointer) trustUsesKerberosKeytab {
	return trustUsesKerberosKeytab{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _trustUsesKerberosKeytabClass) Alloc() trustUsesKerberosKeytab {
	rv := objc.Send[trustUsesKerberosKeytab](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _trustUsesKerberosKeytabClass) New() trustUsesKerberosKeytab {
	rv := objc.Send[trustUsesKerberosKeytab](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustUsesKerberosKeytab) Init() trustUsesKerberosKeytab {
	rv := objc.Send[trustUsesKerberosKeytab](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustUsesKerberosKeytab) Autorelease() trustUsesKerberosKeytab {
	rv := objc.Send[trustUsesKerberosKeytab](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustUsesKerberosKeytab creates a new trustUsesKerberosKeytab instance.
func NewtrustUsesKerberosKeytab() trustUsesKerberosKeytab {
	return gettrustUsesKerberosKeytabClass().New()
}




