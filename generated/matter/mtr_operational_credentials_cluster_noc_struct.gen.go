// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterNOCStruct] class.
var (
	MTROperationalCredentialsClusterNOCStructClass     _MTROperationalCredentialsClusterNOCStructClass
	MTROperationalCredentialsClusterNOCStructClassOnce sync.Once
)

func getMTROperationalCredentialsClusterNOCStructClass() _MTROperationalCredentialsClusterNOCStructClass {
	MTROperationalCredentialsClusterNOCStructClassOnce.Do(func() {
		MTROperationalCredentialsClusterNOCStructClass = _MTROperationalCredentialsClusterNOCStructClass{objc.GetClass("MTROperationalCredentialsClusterNOCStruct")}
	})
	return MTROperationalCredentialsClusterNOCStructClass
}

type _MTROperationalCredentialsClusterNOCStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterNOCStruct] class.
type IMTROperationalCredentialsClusterNOCStruct interface {
	objectivec.IObject
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	Icac() objc.IObject /* cross-framework: Data */
	SetIcac(value objc.IObject /* cross-framework: Data */)
	Noc() objc.IObject /* cross-framework: Data */
	SetNoc(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct
type MTROperationalCredentialsClusterNOCStruct struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterNOCStructFrom constructs a [MTROperationalCredentialsClusterNOCStruct] from an unsafe.Pointer.
func MTROperationalCredentialsClusterNOCStructFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterNOCStruct {
	return MTROperationalCredentialsClusterNOCStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterNOCStructClass) Alloc() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterNOCStructClass) New() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterNOCStruct) Init() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterNOCStruct) Autorelease() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterNOCStruct creates a new MTROperationalCredentialsClusterNOCStruct instance.
func NewMTROperationalCredentialsClusterNOCStruct() MTROperationalCredentialsClusterNOCStruct {
	return getMTROperationalCredentialsClusterNOCStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocstruct/fabricindex
func (m_ MTROperationalCredentialsClusterNOCStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocstruct/fabricindex
func (m_ MTROperationalCredentialsClusterNOCStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocstruct/icac
func (m_ MTROperationalCredentialsClusterNOCStruct) Icac() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("icac"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocstruct/icac
func (m_ MTROperationalCredentialsClusterNOCStruct) SetIcac(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcac:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocstruct/noc
func (m_ MTROperationalCredentialsClusterNOCStruct) Noc() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("noc"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocstruct/noc
func (m_ MTROperationalCredentialsClusterNOCStruct) SetNoc(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNoc:"), value)
}



