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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/type
func (m_ MTROptionalQRCodeInfo) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/type
func (m_ MTROptionalQRCodeInfo) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/infotype
func (m_ MTROptionalQRCodeInfo) InfoType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("infoType"))
	return rv
}


// SetInfoType sets the value of the infoType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/infotype
func (m_ MTROptionalQRCodeInfo) SetInfoType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInfoType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/tag
func (m_ MTROptionalQRCodeInfo) Tag() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/tag
func (m_ MTROptionalQRCodeInfo) SetTag(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTag:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/integervalue
func (m_ MTROptionalQRCodeInfo) IntegerValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("integerValue"))
	return rv
}


// SetIntegerValue sets the value of the integerValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/integervalue
func (m_ MTROptionalQRCodeInfo) SetIntegerValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntegerValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/stringvalue
func (m_ MTROptionalQRCodeInfo) StringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("stringValue"))
	return rv
}


// SetStringValue sets the value of the stringValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroptionalqrcodeinfo/stringvalue
func (m_ MTROptionalQRCodeInfo) SetStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStringValue:"), objc.String(value))
}



