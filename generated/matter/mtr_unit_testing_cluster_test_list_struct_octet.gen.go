// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestListStructOctet] class.
var (
	MTRUnitTestingClusterTestListStructOctetClass     _MTRUnitTestingClusterTestListStructOctetClass
	MTRUnitTestingClusterTestListStructOctetClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListStructOctetClass() _MTRUnitTestingClusterTestListStructOctetClass {
	MTRUnitTestingClusterTestListStructOctetClassOnce.Do(func() {
		MTRUnitTestingClusterTestListStructOctetClass = _MTRUnitTestingClusterTestListStructOctetClass{objc.GetClass("MTRUnitTestingClusterTestListStructOctet")}
	})
	return MTRUnitTestingClusterTestListStructOctetClass
}

type _MTRUnitTestingClusterTestListStructOctetClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestListStructOctet] class.
type IMTRUnitTestingClusterTestListStructOctet interface {
	objectivec.IObject
	Member1() foundation.Number
	SetMember1(value foundation.INumber)
	Member2() foundation.Data
	SetMember2(value foundation.IData)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructOctet
type MTRUnitTestingClusterTestListStructOctet struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListStructOctetFrom constructs a [MTRUnitTestingClusterTestListStructOctet] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListStructOctetFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListStructOctet {
	return MTRUnitTestingClusterTestListStructOctet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListStructOctetClass) Alloc() MTRUnitTestingClusterTestListStructOctet {
	rv := objc.Send[MTRUnitTestingClusterTestListStructOctet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestListStructOctetClass) New() MTRUnitTestingClusterTestListStructOctet {
	rv := objc.Send[MTRUnitTestingClusterTestListStructOctet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListStructOctet) Init() MTRUnitTestingClusterTestListStructOctet {
	rv := objc.Send[MTRUnitTestingClusterTestListStructOctet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListStructOctet) Autorelease() MTRUnitTestingClusterTestListStructOctet {
	rv := objc.Send[MTRUnitTestingClusterTestListStructOctet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListStructOctet creates a new MTRUnitTestingClusterTestListStructOctet instance.
func NewMTRUnitTestingClusterTestListStructOctet() MTRUnitTestingClusterTestListStructOctet {
	return getMTRUnitTestingClusterTestListStructOctetClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructoctet/member1
func (m_ MTRUnitTestingClusterTestListStructOctet) Member1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("member1"))
	return rv
}


// SetMember1 sets the value of the member1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructoctet/member1
func (m_ MTRUnitTestingClusterTestListStructOctet) SetMember1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMember1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructoctet/member2
func (m_ MTRUnitTestingClusterTestListStructOctet) Member2() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("member2"))
	return rv
}


// SetMember2 sets the value of the member2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructoctet/member2
func (m_ MTRUnitTestingClusterTestListStructOctet) SetMember2(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMember2:"), value)
}



