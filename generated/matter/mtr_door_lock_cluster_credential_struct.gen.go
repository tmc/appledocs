// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterCredentialStruct */


/* debug [class_header]: Header for MTRDoorLockClusterCredentialStruct */
// The class instance for the [MTRDoorLockClusterCredentialStruct] class.
var (
	MTRDoorLockClusterCredentialStructClass     _MTRDoorLockClusterCredentialStructClass
	MTRDoorLockClusterCredentialStructClassOnce sync.Once
)

func getMTRDoorLockClusterCredentialStructClass() _MTRDoorLockClusterCredentialStructClass {
	MTRDoorLockClusterCredentialStructClassOnce.Do(func() {
		MTRDoorLockClusterCredentialStructClass = _MTRDoorLockClusterCredentialStructClass{objc.GetClass("MTRDoorLockClusterCredentialStruct")}
	})
	return MTRDoorLockClusterCredentialStructClass
}

type _MTRDoorLockClusterCredentialStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterCredentialStruct */
// An interface definition for the [MTRDoorLockClusterCredentialStruct] class.
type IMTRDoorLockClusterCredentialStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterCredentialStruct */
	// properties:
	CredentialIndex() objc.IObject /* cross-framework: NSNumber */
	SetCredentialIndex(value objc.IObject /* cross-framework: NSNumber */)
	CredentialType() objc.IObject /* cross-framework: NSNumber */
	SetCredentialType(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterCredentialStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterCredentialStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterCredentialStructClass) Alloc() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterCredentialStructClass) New() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterCredentialStruct) Init() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterCredentialStruct) Autorelease() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterCredentialStruct creates a new MTRDoorLockClusterCredentialStruct instance.
func NewMTRDoorLockClusterCredentialStruct() MTRDoorLockClusterCredentialStruct {
	return getMTRDoorLockClusterCredentialStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterCredentialStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterCredentialStruct
type MTRDoorLockClusterCredentialStruct struct {
	objectivec.Object
}

// MTRDoorLockClusterCredentialStructFrom constructs a [MTRDoorLockClusterCredentialStruct] from an unsafe.Pointer.
func MTRDoorLockClusterCredentialStructFrom(ptr unsafe.Pointer) MTRDoorLockClusterCredentialStruct {
	return MTRDoorLockClusterCredentialStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterCredentialStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterCredentialStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterCredentialStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterCredentialStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterCredentialStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterCredentialStruct/credentialIndex
func (m_ MTRDoorLockClusterCredentialStruct) CredentialIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("credentialIndex"))
	return rv
}/* debug [instance_properties/getter]: credentialIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterCredentialStruct/credentialIndex
func (m_ MTRDoorLockClusterCredentialStruct) SetCredentialIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialIndex:"), value)
}/* debug [instance_properties/setter]: credentialIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterCredentialStruct/credentialType
func (m_ MTRDoorLockClusterCredentialStruct) CredentialType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("credentialType"))
	return rv
}/* debug [instance_properties/getter]: credentialType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterCredentialStruct/credentialType
func (m_ MTRDoorLockClusterCredentialStruct) SetCredentialType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialType:"), value)
}/* debug [instance_properties/setter]: credentialType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterCredentialStruct */



