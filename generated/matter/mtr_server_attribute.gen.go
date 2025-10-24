// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServerAttribute] class.
var (
	MTRServerAttributeClass     _MTRServerAttributeClass
	MTRServerAttributeClassOnce sync.Once
)

func getMTRServerAttributeClass() _MTRServerAttributeClass {
	MTRServerAttributeClassOnce.Do(func() {
		MTRServerAttributeClass = _MTRServerAttributeClass{objc.GetClass("MTRServerAttribute")}
	})
	return MTRServerAttributeClass
}

type _MTRServerAttributeClass struct {
	class objc.Class
}

// An interface definition for the [MTRServerAttribute] class.
type IMTRServerAttribute interface {
	objectivec.IObject
	// properties:
	AttributeID() objc.IObject /* cross-framework: NSNumber */
	SetAttributeID(value objc.IObject /* cross-framework: NSNumber */)
	IsWritable() bool
	SetIsWritable(value bool)
	RequiredReadPrivilege() MTRAccessControlEntryPrivilege
	SetRequiredReadPrivilege(value MTRAccessControlEntryPrivilege)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServerAttribute
type MTRServerAttribute struct {
	objectivec.Object
}

// MTRServerAttributeFrom constructs a [MTRServerAttribute] from an unsafe.Pointer.
func MTRServerAttributeFrom(ptr unsafe.Pointer) MTRServerAttribute {
	return MTRServerAttribute{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServerAttributeClass) Alloc() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServerAttributeClass) New() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServerAttribute) Init() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServerAttribute) Autorelease() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServerAttribute creates a new MTRServerAttribute instance.
func NewMTRServerAttribute() MTRServerAttribute {
	return getMTRServerAttributeClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserverattribute/attributeid
func (m_ MTRServerAttribute) AttributeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("attributeID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserverattribute/attributeid
func (m_ MTRServerAttribute) SetAttributeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeID:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserverattribute/iswritable
func (m_ MTRServerAttribute) IsWritable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isWritable"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserverattribute/iswritable
func (m_ MTRServerAttribute) SetIsWritable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsWritable:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserverattribute/requiredreadprivilege
func (m_ MTRServerAttribute) RequiredReadPrivilege() MTRAccessControlEntryPrivilege {
	rv := objc.Send[MTRAccessControlEntryPrivilege](m_.ID, objc.Sel("requiredReadPrivilege"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserverattribute/requiredreadprivilege
func (m_ MTRServerAttribute) SetRequiredReadPrivilege(value MTRAccessControlEntryPrivilege) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredReadPrivilege:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserverattribute/value
func (m_ MTRServerAttribute) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("value"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserverattribute/value
func (m_ MTRServerAttribute) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}
