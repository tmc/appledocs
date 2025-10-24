// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASENumericPair] class.
var (
	PHASENumericPairClass     _PHASENumericPairClass
	PHASENumericPairClassOnce sync.Once
)

func getPHASENumericPairClass() _PHASENumericPairClass {
	PHASENumericPairClassOnce.Do(func() {
		PHASENumericPairClass = _PHASENumericPairClass{objc.GetClass("PHASENumericPair")}
	})
	return PHASENumericPairClass
}

type _PHASENumericPairClass struct {
	class objc.Class
}

// An interface definition for the [PHASENumericPair] class.
type IPHASENumericPair interface {
	objectivec.IObject
	// properties:
	First() float64
	SetFirst(value float64)
	Second() float64
	SetSecond(value float64)
	Domain() IPHASENumericPair
	SetDomain(value IPHASENumericPair)
	Range() IPHASENumericPair
	SetRange(value IPHASENumericPair)
	// methods:
}

// An ordered pair that defines a bounding box for an envelope.
//
// A object uses this class to bound the value of its and .


// An ordered pair that defines a bounding box for an envelope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumericPair
type PHASENumericPair struct {
	objectivec.Object
}

// PHASENumericPairFrom constructs a [PHASENumericPair] from an unsafe.Pointer.
//
// An ordered pair that defines a bounding box for an envelope.
func PHASENumericPairFrom(ptr unsafe.Pointer) PHASENumericPair {
	return PHASENumericPair{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASENumericPairClass) Alloc() PHASENumericPair {
	rv := objc.Send[PHASENumericPair](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASENumericPairClass) New() PHASENumericPair {
	rv := objc.Send[PHASENumericPair](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASENumericPair) Init() PHASENumericPair {
	rv := objc.Send[PHASENumericPair](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASENumericPair) Autorelease() PHASENumericPair {
	rv := objc.Send[PHASENumericPair](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASENumericPair creates a new PHASENumericPair instance.
func NewPHASENumericPair() PHASENumericPair {
	return getPHASENumericPairClass().New()
}



// Creates a pair of numbers with the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumericPair/init(firstValue:secondValue:)
func NewPHASENumericPairWithFirstValueSecondValue(first float64, second float64) PHASENumericPair {
	instance := getPHASENumericPairClass().Alloc()
	rv := objc.Send[PHASENumericPair](instance.ID, objc.Sel("initWithFirstValue:secondValue:"), first, second)
	rv.Autorelease()
	return rv
}



// The first value in the pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumericPair/first
func (p_ PHASENumericPair) First() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("first"))
	return rv
}


// The first value in the pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumericPair/first
func (p_ PHASENumericPair) SetFirst(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFirst:"), value)
}


// The second value in the pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumericPair/second
func (p_ PHASENumericPair) Second() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("second"))
	return rv
}


// The second value in the pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumericPair/second
func (p_ PHASENumericPair) SetSecond(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSecond:"), value)
}


// The range of the envelope’s possible input values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseenvelope/domain
func (p_ PHASENumericPair) Domain() IPHASENumericPair {
	rv := objc.Send[PHASENumericPair](p_.ID, objc.Sel("domain"))
	return rv
}


// The range of the envelope’s possible input values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseenvelope/domain
func (p_ PHASENumericPair) SetDomain(value IPHASENumericPair) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDomain:"), value)
}


// The bounds of the output value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseenvelope/range
func (p_ PHASENumericPair) Range() IPHASENumericPair {
	rv := objc.Send[PHASENumericPair](p_.ID, objc.Sel("range"))
	return rv
}


// The bounds of the output value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseenvelope/range
func (p_ PHASENumericPair) SetRange(value IPHASENumericPair) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRange:"), value)
}


