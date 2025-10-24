// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDataTypeAtomicAttributeStatusStruct] class.
var (
	MTRDataTypeAtomicAttributeStatusStructClass     _MTRDataTypeAtomicAttributeStatusStructClass
	MTRDataTypeAtomicAttributeStatusStructClassOnce sync.Once
)

func getMTRDataTypeAtomicAttributeStatusStructClass() _MTRDataTypeAtomicAttributeStatusStructClass {
	MTRDataTypeAtomicAttributeStatusStructClassOnce.Do(func() {
		MTRDataTypeAtomicAttributeStatusStructClass = _MTRDataTypeAtomicAttributeStatusStructClass{objc.GetClass("MTRDataTypeAtomicAttributeStatusStruct")}
	})
	return MTRDataTypeAtomicAttributeStatusStructClass
}

type _MTRDataTypeAtomicAttributeStatusStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDataTypeAtomicAttributeStatusStruct] class.
type IMTRDataTypeAtomicAttributeStatusStruct interface {
	objectivec.IObject
	// properties:
	AttributeID() objc.IObject /* cross-framework: NSNumber */
	SetAttributeID(value objc.IObject /* cross-framework: NSNumber */)
	StatusCode() objc.IObject /* cross-framework: NSNumber */
	SetStatusCode(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeAtomicAttributeStatusStruct
type MTRDataTypeAtomicAttributeStatusStruct struct {
	objectivec.Object
}

// MTRDataTypeAtomicAttributeStatusStructFrom constructs a [MTRDataTypeAtomicAttributeStatusStruct] from an unsafe.Pointer.
func MTRDataTypeAtomicAttributeStatusStructFrom(ptr unsafe.Pointer) MTRDataTypeAtomicAttributeStatusStruct {
	return MTRDataTypeAtomicAttributeStatusStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDataTypeAtomicAttributeStatusStructClass) Alloc() MTRDataTypeAtomicAttributeStatusStruct {
	rv := objc.Send[MTRDataTypeAtomicAttributeStatusStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDataTypeAtomicAttributeStatusStructClass) New() MTRDataTypeAtomicAttributeStatusStruct {
	rv := objc.Send[MTRDataTypeAtomicAttributeStatusStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDataTypeAtomicAttributeStatusStruct) Init() MTRDataTypeAtomicAttributeStatusStruct {
	rv := objc.Send[MTRDataTypeAtomicAttributeStatusStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDataTypeAtomicAttributeStatusStruct) Autorelease() MTRDataTypeAtomicAttributeStatusStruct {
	rv := objc.Send[MTRDataTypeAtomicAttributeStatusStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDataTypeAtomicAttributeStatusStruct creates a new MTRDataTypeAtomicAttributeStatusStruct instance.
func NewMTRDataTypeAtomicAttributeStatusStruct() MTRDataTypeAtomicAttributeStatusStruct {
	return getMTRDataTypeAtomicAttributeStatusStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeAtomicAttributeStatusStruct/attributeID
func (m_ MTRDataTypeAtomicAttributeStatusStruct) AttributeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("attributeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeAtomicAttributeStatusStruct/attributeID
func (m_ MTRDataTypeAtomicAttributeStatusStruct) SetAttributeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeAtomicAttributeStatusStruct/statusCode
func (m_ MTRDataTypeAtomicAttributeStatusStruct) StatusCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("statusCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeAtomicAttributeStatusStruct/statusCode
func (m_ MTRDataTypeAtomicAttributeStatusStruct) SetStatusCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}



