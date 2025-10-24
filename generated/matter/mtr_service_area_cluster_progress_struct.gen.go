// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterProgressStruct] class.
var (
	MTRServiceAreaClusterProgressStructClass     _MTRServiceAreaClusterProgressStructClass
	MTRServiceAreaClusterProgressStructClassOnce sync.Once
)

func getMTRServiceAreaClusterProgressStructClass() _MTRServiceAreaClusterProgressStructClass {
	MTRServiceAreaClusterProgressStructClassOnce.Do(func() {
		MTRServiceAreaClusterProgressStructClass = _MTRServiceAreaClusterProgressStructClass{objc.GetClass("MTRServiceAreaClusterProgressStruct")}
	})
	return MTRServiceAreaClusterProgressStructClass
}

type _MTRServiceAreaClusterProgressStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterProgressStruct] class.
type IMTRServiceAreaClusterProgressStruct interface {
	objectivec.IObject
	// properties:
	AreaID() objc.IObject /* cross-framework: NSNumber */
	SetAreaID(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TotalOperationalTime() objc.IObject /* cross-framework: NSNumber */
	SetTotalOperationalTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct
type MTRServiceAreaClusterProgressStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterProgressStructFrom constructs a [MTRServiceAreaClusterProgressStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterProgressStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterProgressStruct {
	return MTRServiceAreaClusterProgressStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterProgressStructClass) Alloc() MTRServiceAreaClusterProgressStruct {
	rv := objc.Send[MTRServiceAreaClusterProgressStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterProgressStructClass) New() MTRServiceAreaClusterProgressStruct {
	rv := objc.Send[MTRServiceAreaClusterProgressStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterProgressStruct) Init() MTRServiceAreaClusterProgressStruct {
	rv := objc.Send[MTRServiceAreaClusterProgressStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterProgressStruct) Autorelease() MTRServiceAreaClusterProgressStruct {
	rv := objc.Send[MTRServiceAreaClusterProgressStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterProgressStruct creates a new MTRServiceAreaClusterProgressStruct instance.
func NewMTRServiceAreaClusterProgressStruct() MTRServiceAreaClusterProgressStruct {
	return getMTRServiceAreaClusterProgressStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct/areaID
func (m_ MTRServiceAreaClusterProgressStruct) AreaID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("areaID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct/areaID
func (m_ MTRServiceAreaClusterProgressStruct) SetAreaID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAreaID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct/status
func (m_ MTRServiceAreaClusterProgressStruct) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct/status
func (m_ MTRServiceAreaClusterProgressStruct) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct/totalOperationalTime
func (m_ MTRServiceAreaClusterProgressStruct) TotalOperationalTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("totalOperationalTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct/totalOperationalTime
func (m_ MTRServiceAreaClusterProgressStruct) SetTotalOperationalTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalOperationalTime:"), value)
}



