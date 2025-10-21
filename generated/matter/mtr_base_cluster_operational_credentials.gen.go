// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOperationalCredentials] class.
var (
	MTRBaseClusterOperationalCredentialsClass     _MTRBaseClusterOperationalCredentialsClass
	MTRBaseClusterOperationalCredentialsClassOnce sync.Once
)

func getMTRBaseClusterOperationalCredentialsClass() _MTRBaseClusterOperationalCredentialsClass {
	MTRBaseClusterOperationalCredentialsClassOnce.Do(func() {
		MTRBaseClusterOperationalCredentialsClass = _MTRBaseClusterOperationalCredentialsClass{objc.GetClass("MTRBaseClusterOperationalCredentials")}
	})
	return MTRBaseClusterOperationalCredentialsClass
}

type _MTRBaseClusterOperationalCredentialsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOperationalCredentials] class.
type IMTRBaseClusterOperationalCredentials interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOperationalCredentials
type MTRBaseClusterOperationalCredentials struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOperationalCredentialsFrom constructs a [MTRBaseClusterOperationalCredentials] from an unsafe.Pointer.
func MTRBaseClusterOperationalCredentialsFrom(ptr unsafe.Pointer) MTRBaseClusterOperationalCredentials {
	return MTRBaseClusterOperationalCredentials{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOperationalCredentialsClass) Alloc() MTRBaseClusterOperationalCredentials {
	rv := objc.Send[MTRBaseClusterOperationalCredentials](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOperationalCredentialsClass) New() MTRBaseClusterOperationalCredentials {
	rv := objc.Send[MTRBaseClusterOperationalCredentials](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOperationalCredentials) Init() MTRBaseClusterOperationalCredentials {
	rv := objc.Send[MTRBaseClusterOperationalCredentials](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOperationalCredentials) Autorelease() MTRBaseClusterOperationalCredentials {
	rv := objc.Send[MTRBaseClusterOperationalCredentials](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOperationalCredentials creates a new MTRBaseClusterOperationalCredentials instance.
func NewMTRBaseClusterOperationalCredentials() MTRBaseClusterOperationalCredentials {
	return getMTRBaseClusterOperationalCredentialsClass().New()
}




