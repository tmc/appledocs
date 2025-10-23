// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKHealthConceptIdentifier] class.
var (
	HKHealthConceptIdentifierClass     _HKHealthConceptIdentifierClass
	HKHealthConceptIdentifierClassOnce sync.Once
)

func getHKHealthConceptIdentifierClass() _HKHealthConceptIdentifierClass {
	HKHealthConceptIdentifierClassOnce.Do(func() {
		HKHealthConceptIdentifierClass = _HKHealthConceptIdentifierClass{objc.GetClass("HKHealthConceptIdentifier")}
	})
	return HKHealthConceptIdentifierClass
}

type _HKHealthConceptIdentifierClass struct {
	class objc.Class
}

// An interface definition for the [HKHealthConceptIdentifier] class.
type IHKHealthConceptIdentifier interface {
	objectivec.IObject
	// properties:
	Domain() HKHealthConceptDomain
	SetDomain(value HKHealthConceptDomain)
	// methods:
}

// A unique identifier for a specific health concept within a domain.
//
// Each identifier points to one concept inside a domain. For example, within the medication domain, one identifier might represent ibuprofen while another represents insulin.


// A unique identifier for a specific health concept within a domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthConceptIdentifier
type HKHealthConceptIdentifier struct {
	objectivec.Object
}

// HKHealthConceptIdentifierFrom constructs a [HKHealthConceptIdentifier] from an unsafe.Pointer.
//
// A unique identifier for a specific health concept within a domain.
func HKHealthConceptIdentifierFrom(ptr unsafe.Pointer) HKHealthConceptIdentifier {
	return HKHealthConceptIdentifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKHealthConceptIdentifierClass) Alloc() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKHealthConceptIdentifierClass) New() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKHealthConceptIdentifier) Init() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKHealthConceptIdentifier) Autorelease() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKHealthConceptIdentifier creates a new HKHealthConceptIdentifier instance.
func NewHKHealthConceptIdentifier() HKHealthConceptIdentifier {
	return getHKHealthConceptIdentifierClass().New()
}



// The domain this identifier belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkhealthconceptidentifier/domain
func (h_ HKHealthConceptIdentifier) Domain() HKHealthConceptDomain {
	rv := objc.Send[HKHealthConceptDomain](h_.ID, objc.Sel("domain"))
	return rv
}


// The domain this identifier belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkhealthconceptidentifier/domain
func (h_ HKHealthConceptIdentifier) SetDomain(value HKHealthConceptDomain) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDomain:"), value)
}



