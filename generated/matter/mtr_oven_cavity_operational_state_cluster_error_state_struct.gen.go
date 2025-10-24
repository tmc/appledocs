// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenCavityOperationalStateClusterErrorStateStruct] class.
var (
	MTROvenCavityOperationalStateClusterErrorStateStructClass     _MTROvenCavityOperationalStateClusterErrorStateStructClass
	MTROvenCavityOperationalStateClusterErrorStateStructClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterErrorStateStructClass() _MTROvenCavityOperationalStateClusterErrorStateStructClass {
	MTROvenCavityOperationalStateClusterErrorStateStructClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterErrorStateStructClass = _MTROvenCavityOperationalStateClusterErrorStateStructClass{objc.GetClass("MTROvenCavityOperationalStateClusterErrorStateStruct")}
	})
	return MTROvenCavityOperationalStateClusterErrorStateStructClass
}

type _MTROvenCavityOperationalStateClusterErrorStateStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenCavityOperationalStateClusterErrorStateStruct] class.
type IMTROvenCavityOperationalStateClusterErrorStateStruct interface {
	objectivec.IObject
	// properties:
	ErrorStateDetails() objc.IObject /* cross-framework: NSString */
	SetErrorStateDetails(value objc.IObject /* cross-framework: NSString */)
	ErrorStateID() objc.IObject /* cross-framework: NSNumber */
	SetErrorStateID(value objc.IObject /* cross-framework: NSNumber */)
	ErrorStateLabel() objc.IObject /* cross-framework: NSString */
	SetErrorStateLabel(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct
type MTROvenCavityOperationalStateClusterErrorStateStruct struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterErrorStateStructFrom constructs a [MTROvenCavityOperationalStateClusterErrorStateStruct] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterErrorStateStructFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterErrorStateStruct {
	return MTROvenCavityOperationalStateClusterErrorStateStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterErrorStateStructClass) Alloc() MTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenCavityOperationalStateClusterErrorStateStructClass) New() MTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) Init() MTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) Autorelease() MTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterErrorStateStruct creates a new MTROvenCavityOperationalStateClusterErrorStateStruct instance.
func NewMTROvenCavityOperationalStateClusterErrorStateStruct() MTROvenCavityOperationalStateClusterErrorStateStruct {
	return getMTROvenCavityOperationalStateClusterErrorStateStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct/errorStateDetails
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) ErrorStateDetails() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("errorStateDetails"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct/errorStateDetails
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) SetErrorStateDetails(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorStateDetails:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct/errorStateID
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) ErrorStateID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("errorStateID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct/errorStateID
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) SetErrorStateID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorStateID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct/errorStateLabel
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) ErrorStateLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("errorStateLabel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct/errorStateLabel
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) SetErrorStateLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorStateLabel:"), value)
}



