// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROptionalQRCodeInfo] class.
var (
	MTROptionalQRCodeInfoClass     _MTROptionalQRCodeInfoClass
	MTROptionalQRCodeInfoClassOnce sync.Once
)

func getMTROptionalQRCodeInfoClass() _MTROptionalQRCodeInfoClass {
	MTROptionalQRCodeInfoClassOnce.Do(func() {
		MTROptionalQRCodeInfoClass = _MTROptionalQRCodeInfoClass{objc.GetClass("MTROptionalQRCodeInfo")}
	})
	return MTROptionalQRCodeInfoClass
}

type _MTROptionalQRCodeInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTROptionalQRCodeInfo] class.
type IMTROptionalQRCodeInfo interface {
	objectivec.IObject
	// properties:
	InfoType() objc.IObject /* cross-framework: NSNumber */
	SetInfoType(value objc.IObject /* cross-framework: NSNumber */)
	IntegerValue() objc.IObject /* cross-framework: NSNumber */
	SetIntegerValue(value objc.IObject /* cross-framework: NSNumber */)
	StringValue() objc.IObject /* cross-framework: NSString */
	SetStringValue(value objc.IObject /* cross-framework: NSString */)
	Tag() objc.IObject /* cross-framework: NSNumber */
	SetTag(value objc.IObject /* cross-framework: NSNumber */)
	Type() MTROptionalQRCodeInfoType
	SetType(value MTROptionalQRCodeInfoType)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo
type MTROptionalQRCodeInfo struct {
	objectivec.Object
}

// MTROptionalQRCodeInfoFrom constructs a [MTROptionalQRCodeInfo] from an unsafe.Pointer.
func MTROptionalQRCodeInfoFrom(ptr unsafe.Pointer) MTROptionalQRCodeInfo {
	return MTROptionalQRCodeInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROptionalQRCodeInfoClass) Alloc() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROptionalQRCodeInfoClass) New() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROptionalQRCodeInfo) Init() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROptionalQRCodeInfo) Autorelease() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROptionalQRCodeInfo creates a new MTROptionalQRCodeInfo instance.
func NewMTROptionalQRCodeInfo() MTROptionalQRCodeInfo {
	return getMTROptionalQRCodeInfoClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/infotype
func (m_ MTROptionalQRCodeInfo) InfoType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("infoType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/infotype
func (m_ MTROptionalQRCodeInfo) SetInfoType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInfoType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/integervalue
func (m_ MTROptionalQRCodeInfo) IntegerValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("integerValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/integervalue
func (m_ MTROptionalQRCodeInfo) SetIntegerValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntegerValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/stringvalue
func (m_ MTROptionalQRCodeInfo) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("stringValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/stringvalue
func (m_ MTROptionalQRCodeInfo) SetStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStringValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/tag
func (m_ MTROptionalQRCodeInfo) Tag() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("tag"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/tag
func (m_ MTROptionalQRCodeInfo) SetTag(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTag:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/type
func (m_ MTROptionalQRCodeInfo) Type() MTROptionalQRCodeInfoType {
	rv := objc.Send[MTROptionalQRCodeInfoType](m_.ID, objc.Sel("type"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/type
func (m_ MTROptionalQRCodeInfo) SetType(value MTROptionalQRCodeInfoType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



