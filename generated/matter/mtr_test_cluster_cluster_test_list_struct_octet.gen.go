// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestListStructOctet] class.
var (
	MTRTestClusterClusterTestListStructOctetClass     _MTRTestClusterClusterTestListStructOctetClass
	MTRTestClusterClusterTestListStructOctetClassOnce sync.Once
)

func getMTRTestClusterClusterTestListStructOctetClass() _MTRTestClusterClusterTestListStructOctetClass {
	MTRTestClusterClusterTestListStructOctetClassOnce.Do(func() {
		MTRTestClusterClusterTestListStructOctetClass = _MTRTestClusterClusterTestListStructOctetClass{objc.GetClass("MTRTestClusterClusterTestListStructOctet")}
	})
	return MTRTestClusterClusterTestListStructOctetClass
}

type _MTRTestClusterClusterTestListStructOctetClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestListStructOctet] class.
type IMTRTestClusterClusterTestListStructOctet interface {
	IMTRUnitTestingClusterTestListStructOctet
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestListStructOctet
type MTRTestClusterClusterTestListStructOctet struct {
	MTRUnitTestingClusterTestListStructOctet
}

// MTRTestClusterClusterTestListStructOctetFrom constructs a [MTRTestClusterClusterTestListStructOctet] from an unsafe.Pointer.
func MTRTestClusterClusterTestListStructOctetFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestListStructOctet {
	return MTRTestClusterClusterTestListStructOctet{
		MTRUnitTestingClusterTestListStructOctet: MTRUnitTestingClusterTestListStructOctetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestListStructOctetClass) Alloc() MTRTestClusterClusterTestListStructOctet {
	rv := objc.Send[MTRTestClusterClusterTestListStructOctet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestListStructOctetClass) New() MTRTestClusterClusterTestListStructOctet {
	rv := objc.Send[MTRTestClusterClusterTestListStructOctet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestListStructOctet) Init() MTRTestClusterClusterTestListStructOctet {
	rv := objc.Send[MTRTestClusterClusterTestListStructOctet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestListStructOctet) Autorelease() MTRTestClusterClusterTestListStructOctet {
	rv := objc.Send[MTRTestClusterClusterTestListStructOctet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestListStructOctet creates a new MTRTestClusterClusterTestListStructOctet instance.
func NewMTRTestClusterClusterTestListStructOctet() MTRTestClusterClusterTestListStructOctet {
	return getMTRTestClusterClusterTestListStructOctetClass().New()
}




