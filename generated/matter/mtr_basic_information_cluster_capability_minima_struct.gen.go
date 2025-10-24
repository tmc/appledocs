// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBasicInformationClusterCapabilityMinimaStruct] class.
var (
	MTRBasicInformationClusterCapabilityMinimaStructClass     _MTRBasicInformationClusterCapabilityMinimaStructClass
	MTRBasicInformationClusterCapabilityMinimaStructClassOnce sync.Once
)

func getMTRBasicInformationClusterCapabilityMinimaStructClass() _MTRBasicInformationClusterCapabilityMinimaStructClass {
	MTRBasicInformationClusterCapabilityMinimaStructClassOnce.Do(func() {
		MTRBasicInformationClusterCapabilityMinimaStructClass = _MTRBasicInformationClusterCapabilityMinimaStructClass{objc.GetClass("MTRBasicInformationClusterCapabilityMinimaStruct")}
	})
	return MTRBasicInformationClusterCapabilityMinimaStructClass
}

type _MTRBasicInformationClusterCapabilityMinimaStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicInformationClusterCapabilityMinimaStruct] class.
type IMTRBasicInformationClusterCapabilityMinimaStruct interface {
	objectivec.IObject
	// properties:
	CaseSessionsPerFabric() objc.IObject /* cross-framework: NSNumber */
	SetCaseSessionsPerFabric(value objc.IObject /* cross-framework: NSNumber */)
	SubscriptionsPerFabric() objc.IObject /* cross-framework: NSNumber */
	SetSubscriptionsPerFabric(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicInformationClusterCapabilityMinimaStruct
type MTRBasicInformationClusterCapabilityMinimaStruct struct {
	objectivec.Object
}

// MTRBasicInformationClusterCapabilityMinimaStructFrom constructs a [MTRBasicInformationClusterCapabilityMinimaStruct] from an unsafe.Pointer.
func MTRBasicInformationClusterCapabilityMinimaStructFrom(ptr unsafe.Pointer) MTRBasicInformationClusterCapabilityMinimaStruct {
	return MTRBasicInformationClusterCapabilityMinimaStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicInformationClusterCapabilityMinimaStructClass) Alloc() MTRBasicInformationClusterCapabilityMinimaStruct {
	rv := objc.Send[MTRBasicInformationClusterCapabilityMinimaStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicInformationClusterCapabilityMinimaStructClass) New() MTRBasicInformationClusterCapabilityMinimaStruct {
	rv := objc.Send[MTRBasicInformationClusterCapabilityMinimaStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicInformationClusterCapabilityMinimaStruct) Init() MTRBasicInformationClusterCapabilityMinimaStruct {
	rv := objc.Send[MTRBasicInformationClusterCapabilityMinimaStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicInformationClusterCapabilityMinimaStruct) Autorelease() MTRBasicInformationClusterCapabilityMinimaStruct {
	rv := objc.Send[MTRBasicInformationClusterCapabilityMinimaStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicInformationClusterCapabilityMinimaStruct creates a new MTRBasicInformationClusterCapabilityMinimaStruct instance.
func NewMTRBasicInformationClusterCapabilityMinimaStruct() MTRBasicInformationClusterCapabilityMinimaStruct {
	return getMTRBasicInformationClusterCapabilityMinimaStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicinformationclustercapabilityminimastruct/casesessionsperfabric
func (m_ MTRBasicInformationClusterCapabilityMinimaStruct) CaseSessionsPerFabric() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("caseSessionsPerFabric"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicinformationclustercapabilityminimastruct/casesessionsperfabric
func (m_ MTRBasicInformationClusterCapabilityMinimaStruct) SetCaseSessionsPerFabric(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseSessionsPerFabric:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicinformationclustercapabilityminimastruct/subscriptionsperfabric
func (m_ MTRBasicInformationClusterCapabilityMinimaStruct) SubscriptionsPerFabric() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("subscriptionsPerFabric"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicinformationclustercapabilityminimastruct/subscriptionsperfabric
func (m_ MTRBasicInformationClusterCapabilityMinimaStruct) SetSubscriptionsPerFabric(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubscriptionsPerFabric:"), value)
}



