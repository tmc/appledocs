// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMediaInputClusterInputInfo] class.
var (
	MTRMediaInputClusterInputInfoClass     _MTRMediaInputClusterInputInfoClass
	MTRMediaInputClusterInputInfoClassOnce sync.Once
)

func getMTRMediaInputClusterInputInfoClass() _MTRMediaInputClusterInputInfoClass {
	MTRMediaInputClusterInputInfoClassOnce.Do(func() {
		MTRMediaInputClusterInputInfoClass = _MTRMediaInputClusterInputInfoClass{objc.GetClass("MTRMediaInputClusterInputInfo")}
	})
	return MTRMediaInputClusterInputInfoClass
}

type _MTRMediaInputClusterInputInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaInputClusterInputInfo] class.
type IMTRMediaInputClusterInputInfo interface {
	IMTRMediaInputClusterInputInfoStruct
	// properties:
	DescriptionString() objc.IObject /* cross-framework: NSString */
	SetDescriptionString(value objc.IObject /* cross-framework: NSString */)
	Index() objc.IObject /* cross-framework: NSNumber */
	SetIndex(value objc.IObject /* cross-framework: NSNumber */)
	InputType() objc.IObject /* cross-framework: NSNumber */
	SetInputType(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaInputClusterInputInfo
type MTRMediaInputClusterInputInfo struct {
	MTRMediaInputClusterInputInfoStruct
}

// MTRMediaInputClusterInputInfoFrom constructs a [MTRMediaInputClusterInputInfo] from an unsafe.Pointer.
func MTRMediaInputClusterInputInfoFrom(ptr unsafe.Pointer) MTRMediaInputClusterInputInfo {
	return MTRMediaInputClusterInputInfo{
		MTRMediaInputClusterInputInfoStruct: MTRMediaInputClusterInputInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaInputClusterInputInfoClass) Alloc() MTRMediaInputClusterInputInfo {
	rv := objc.Send[MTRMediaInputClusterInputInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaInputClusterInputInfoClass) New() MTRMediaInputClusterInputInfo {
	rv := objc.Send[MTRMediaInputClusterInputInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaInputClusterInputInfo) Init() MTRMediaInputClusterInputInfo {
	rv := objc.Send[MTRMediaInputClusterInputInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaInputClusterInputInfo) Autorelease() MTRMediaInputClusterInputInfo {
	rv := objc.Send[MTRMediaInputClusterInputInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaInputClusterInputInfo creates a new MTRMediaInputClusterInputInfo instance.
func NewMTRMediaInputClusterInputInfo() MTRMediaInputClusterInputInfo {
	return getMTRMediaInputClusterInputInfoClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfo/descriptionstring
func (m_ MTRMediaInputClusterInputInfo) DescriptionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("descriptionString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfo/descriptionstring
func (m_ MTRMediaInputClusterInputInfo) SetDescriptionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDescriptionString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfo/index
func (m_ MTRMediaInputClusterInputInfo) Index() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("index"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfo/index
func (m_ MTRMediaInputClusterInputInfo) SetIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfo/inputtype
func (m_ MTRMediaInputClusterInputInfo) InputType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("inputType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfo/inputtype
func (m_ MTRMediaInputClusterInputInfo) SetInputType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfo/name
func (m_ MTRMediaInputClusterInputInfo) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfo/name
func (m_ MTRMediaInputClusterInputInfo) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}



