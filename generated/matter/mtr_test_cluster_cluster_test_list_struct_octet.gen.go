// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Member1() foundation.Number
	SetMember1(value foundation.INumber)
	Member2() foundation.Data
	SetMember2(value foundation.IData)
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructoctet/member1
func (m_ MTRTestClusterClusterTestListStructOctet) Member1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("member1"))
	return rv
}


// SetMember1 sets the value of the member1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructoctet/member1
func (m_ MTRTestClusterClusterTestListStructOctet) SetMember1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMember1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructoctet/member2
func (m_ MTRTestClusterClusterTestListStructOctet) Member2() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("member2"))
	return rv
}


// SetMember2 sets the value of the member2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructoctet/member2
func (m_ MTRTestClusterClusterTestListStructOctet) SetMember2(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMember2:"), value)
}



