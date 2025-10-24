// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaInputClusterInputInfoStruct] class.
var (
	MTRMediaInputClusterInputInfoStructClass     _MTRMediaInputClusterInputInfoStructClass
	MTRMediaInputClusterInputInfoStructClassOnce sync.Once
)

func getMTRMediaInputClusterInputInfoStructClass() _MTRMediaInputClusterInputInfoStructClass {
	MTRMediaInputClusterInputInfoStructClassOnce.Do(func() {
		MTRMediaInputClusterInputInfoStructClass = _MTRMediaInputClusterInputInfoStructClass{objc.GetClass("MTRMediaInputClusterInputInfoStruct")}
	})
	return MTRMediaInputClusterInputInfoStructClass
}

type _MTRMediaInputClusterInputInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaInputClusterInputInfoStruct] class.
type IMTRMediaInputClusterInputInfoStruct interface {
	objectivec.IObject
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaInputClusterInputInfoStruct
type MTRMediaInputClusterInputInfoStruct struct {
	objectivec.Object
}

// MTRMediaInputClusterInputInfoStructFrom constructs a [MTRMediaInputClusterInputInfoStruct] from an unsafe.Pointer.
func MTRMediaInputClusterInputInfoStructFrom(ptr unsafe.Pointer) MTRMediaInputClusterInputInfoStruct {
	return MTRMediaInputClusterInputInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaInputClusterInputInfoStructClass) Alloc() MTRMediaInputClusterInputInfoStruct {
	rv := objc.Send[MTRMediaInputClusterInputInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaInputClusterInputInfoStructClass) New() MTRMediaInputClusterInputInfoStruct {
	rv := objc.Send[MTRMediaInputClusterInputInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaInputClusterInputInfoStruct) Init() MTRMediaInputClusterInputInfoStruct {
	rv := objc.Send[MTRMediaInputClusterInputInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaInputClusterInputInfoStruct) Autorelease() MTRMediaInputClusterInputInfoStruct {
	rv := objc.Send[MTRMediaInputClusterInputInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaInputClusterInputInfoStruct creates a new MTRMediaInputClusterInputInfoStruct instance.
func NewMTRMediaInputClusterInputInfoStruct() MTRMediaInputClusterInputInfoStruct {
	return getMTRMediaInputClusterInputInfoStructClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfostruct/descriptionstring
func (m_ MTRMediaInputClusterInputInfoStruct) DescriptionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("descriptionString"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfostruct/descriptionstring
func (m_ MTRMediaInputClusterInputInfoStruct) SetDescriptionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDescriptionString:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfostruct/index
func (m_ MTRMediaInputClusterInputInfoStruct) Index() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("index"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfostruct/index
func (m_ MTRMediaInputClusterInputInfoStruct) SetIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfostruct/inputtype
func (m_ MTRMediaInputClusterInputInfoStruct) InputType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("inputType"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfostruct/inputtype
func (m_ MTRMediaInputClusterInputInfoStruct) SetInputType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputType:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfostruct/name
func (m_ MTRMediaInputClusterInputInfoStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterinputinfostruct/name
func (m_ MTRMediaInputClusterInputInfoStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}
