// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOperationalCredentials] class.
var (
	MTRClusterOperationalCredentialsClass     _MTRClusterOperationalCredentialsClass
	MTRClusterOperationalCredentialsClassOnce sync.Once
)

func getMTRClusterOperationalCredentialsClass() _MTRClusterOperationalCredentialsClass {
	MTRClusterOperationalCredentialsClassOnce.Do(func() {
		MTRClusterOperationalCredentialsClass = _MTRClusterOperationalCredentialsClass{objc.GetClass("MTRClusterOperationalCredentials")}
	})
	return MTRClusterOperationalCredentialsClass
}

type _MTRClusterOperationalCredentialsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOperationalCredentials] class.
type IMTRClusterOperationalCredentials interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOperationalCredentials
type MTRClusterOperationalCredentials struct {
	MTRGenericCluster
}

// MTRClusterOperationalCredentialsFrom constructs a [MTRClusterOperationalCredentials] from an unsafe.Pointer.
func MTRClusterOperationalCredentialsFrom(ptr unsafe.Pointer) MTRClusterOperationalCredentials {
	return MTRClusterOperationalCredentials{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOperationalCredentialsClass) Alloc() MTRClusterOperationalCredentials {
	rv := objc.Send[MTRClusterOperationalCredentials](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOperationalCredentialsClass) New() MTRClusterOperationalCredentials {
	rv := objc.Send[MTRClusterOperationalCredentials](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOperationalCredentials) Init() MTRClusterOperationalCredentials {
	rv := objc.Send[MTRClusterOperationalCredentials](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOperationalCredentials) Autorelease() MTRClusterOperationalCredentials {
	rv := objc.Send[MTRClusterOperationalCredentials](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOperationalCredentials creates a new MTRClusterOperationalCredentials instance.
func NewMTRClusterOperationalCredentials() MTRClusterOperationalCredentials {
	return getMTRClusterOperationalCredentialsClass().New()
}




