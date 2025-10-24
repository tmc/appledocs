// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDataTypeAtomicAttributeStatusStruct */


/* debug [class_header]: Header for MTRDataTypeAtomicAttributeStatusStruct */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDataTypeAtomicAttributeStatusStruct */
// An interface definition for the [MTRDataTypeAtomicAttributeStatusStruct] class.
type IMTRDataTypeAtomicAttributeStatusStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDataTypeAtomicAttributeStatusStruct */
	// properties:
	AttributeID() objc.IObject /* cross-framework: NSNumber */
	SetAttributeID(value objc.IObject /* cross-framework: NSNumber */)
	StatusCode() objc.IObject /* cross-framework: NSNumber */
	SetStatusCode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDataTypeAtomicAttributeStatusStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDataTypeAtomicAttributeStatusStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDataTypeAtomicAttributeStatusStructClass) Alloc() MTRDataTypeAtomicAttributeStatusStruct {
	rv := objc.Send[MTRDataTypeAtomicAttributeStatusStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDataTypeAtomicAttributeStatusStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeAtomicAttributeStatusStruct
type MTRDataTypeAtomicAttributeStatusStruct struct {
	objectivec.Object
}

// MTRDataTypeAtomicAttributeStatusStructFrom constructs a [MTRDataTypeAtomicAttributeStatusStruct] from an unsafe.Pointer.
func MTRDataTypeAtomicAttributeStatusStructFrom(ptr unsafe.Pointer) MTRDataTypeAtomicAttributeStatusStruct {
	return MTRDataTypeAtomicAttributeStatusStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDataTypeAtomicAttributeStatusStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDataTypeAtomicAttributeStatusStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDataTypeAtomicAttributeStatusStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDataTypeAtomicAttributeStatusStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDataTypeAtomicAttributeStatusStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeAtomicAttributeStatusStruct/attributeID
func (m_ MTRDataTypeAtomicAttributeStatusStruct) AttributeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("attributeID"))
	return rv
}/* debug [instance_properties/getter]: attributeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeAtomicAttributeStatusStruct/attributeID
func (m_ MTRDataTypeAtomicAttributeStatusStruct) SetAttributeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeID:"), value)
}/* debug [instance_properties/setter]: attributeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdatatypeatomicattributestatusstruct/statuscode
func (m_ MTRDataTypeAtomicAttributeStatusStruct) StatusCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("statusCode"))
	return rv
}/* debug [instance_properties/getter]: statusCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdatatypeatomicattributestatusstruct/statuscode
func (m_ MTRDataTypeAtomicAttributeStatusStruct) SetStatusCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}/* debug [instance_properties/setter]: statusCode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDataTypeAtomicAttributeStatusStruct */



