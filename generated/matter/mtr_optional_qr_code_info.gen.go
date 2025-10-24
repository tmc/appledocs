// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROptionalQRCodeInfo */


/* debug [class_header]: Header for MTROptionalQRCodeInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROptionalQRCodeInfo */
// An interface definition for the [MTROptionalQRCodeInfo] class.
type IMTROptionalQRCodeInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROptionalQRCodeInfo */
	// properties:
	InfoType() objc.IObject /* cross-framework: NSNumber */
	SetInfoType(value objc.IObject /* cross-framework: NSNumber */)
	IntegerValue() objc.IObject /* cross-framework: NSNumber */
	StringValue() objc.IObject /* cross-framework: NSString */
	Tag() objc.IObject /* cross-framework: NSNumber */
	Type() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROptionalQRCodeInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROptionalQRCodeInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTROptionalQRCodeInfoClass) Alloc() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROptionalQRCodeInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo
type MTROptionalQRCodeInfo struct {
	objectivec.Object
}

// MTROptionalQRCodeInfoFrom constructs a [MTROptionalQRCodeInfo] from an unsafe.Pointer.
func MTROptionalQRCodeInfoFrom(ptr unsafe.Pointer) MTROptionalQRCodeInfo {
	return MTROptionalQRCodeInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROptionalQRCodeInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo/init(tag:int32Value:)
func NewMTROptionalQRCodeInfoWithTagInt32Value(tag objc.IObject /* cross-framework: NSNumber */, value int32 /* not a class type */) MTROptionalQRCodeInfo {
	instance := getMTROptionalQRCodeInfoClass().Alloc()
	rv := objc.Send[MTROptionalQRCodeInfo](instance.ID, objc.Sel("initWithTag:int32Value:"), tag, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROptionalQRCodeInfoWithTagInt32Value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo/init(tag:stringValue:)
func NewMTROptionalQRCodeInfoWithTagStringValue(tag objc.IObject /* cross-framework: NSNumber */, value objc.IObject /* cross-framework: NSString */) MTROptionalQRCodeInfo {
	instance := getMTROptionalQRCodeInfoClass().Alloc()
	rv := objc.Send[MTROptionalQRCodeInfo](instance.ID, objc.Sel("initWithTag:stringValue:"), tag, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROptionalQRCodeInfoWithTagStringValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROptionalQRCodeInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROptionalQRCodeInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROptionalQRCodeInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROptionalQRCodeInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo/infoType
func (m_ MTROptionalQRCodeInfo) InfoType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("infoType"))
	return rv
}/* debug [instance_properties/getter]: infoType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo/infoType
func (m_ MTROptionalQRCodeInfo) SetInfoType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInfoType:"), value)
}/* debug [instance_properties/setter]: infoType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo/integerValue
func (m_ MTROptionalQRCodeInfo) IntegerValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("integerValue"))
	return rv
}/* debug [instance_properties/getter]: integerValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo/stringValue
func (m_ MTROptionalQRCodeInfo) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("stringValue"))
	return rv
}/* debug [instance_properties/getter]: stringValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo/tag
func (m_ MTROptionalQRCodeInfo) Tag() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo/type
func (m_ MTROptionalQRCodeInfo) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROptionalQRCodeInfo */


