// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDistinguishedNameInfo */


/* debug [class_header]: Header for MTRDistinguishedNameInfo */
// The class instance for the [MTRDistinguishedNameInfo] class.
var (
	MTRDistinguishedNameInfoClass     _MTRDistinguishedNameInfoClass
	MTRDistinguishedNameInfoClassOnce sync.Once
)

func getMTRDistinguishedNameInfoClass() _MTRDistinguishedNameInfoClass {
	MTRDistinguishedNameInfoClassOnce.Do(func() {
		MTRDistinguishedNameInfoClass = _MTRDistinguishedNameInfoClass{objc.GetClass("MTRDistinguishedNameInfo")}
	})
	return MTRDistinguishedNameInfoClass
}

type _MTRDistinguishedNameInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDistinguishedNameInfo */
// An interface definition for the [MTRDistinguishedNameInfo] class.
type IMTRDistinguishedNameInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDistinguishedNameInfo */
	// properties:
	CaseAuthenticatedTags() unsafe.Pointer
	FabricID() objc.IObject /* cross-framework: NSNumber */
	IntermediateCACertificateID() objc.IObject /* cross-framework: NSNumber */
	NodeID() objc.IObject /* cross-framework: NSNumber */
	RootCACertificateID() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDistinguishedNameInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDistinguishedNameInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTRDistinguishedNameInfoClass) Alloc() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDistinguishedNameInfoClass) New() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDistinguishedNameInfo) Init() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDistinguishedNameInfo) Autorelease() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDistinguishedNameInfo creates a new MTRDistinguishedNameInfo instance.
func NewMTRDistinguishedNameInfo() MTRDistinguishedNameInfo {
	return getMTRDistinguishedNameInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDistinguishedNameInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDistinguishedNameInfo
type MTRDistinguishedNameInfo struct {
	objectivec.Object
}

// MTRDistinguishedNameInfoFrom constructs a [MTRDistinguishedNameInfo] from an unsafe.Pointer.
func MTRDistinguishedNameInfoFrom(ptr unsafe.Pointer) MTRDistinguishedNameInfo {
	return MTRDistinguishedNameInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDistinguishedNameInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDistinguishedNameInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDistinguishedNameInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDistinguishedNameInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDistinguishedNameInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDistinguishedNameInfo/caseAuthenticatedTags
func (m_ MTRDistinguishedNameInfo) CaseAuthenticatedTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("caseAuthenticatedTags"))
	return rv
}/* debug [instance_properties/getter]: caseAuthenticatedTags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDistinguishedNameInfo/fabricID
func (m_ MTRDistinguishedNameInfo) FabricID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricID"))
	return rv
}/* debug [instance_properties/getter]: fabricID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDistinguishedNameInfo/intermediateCACertificateID
func (m_ MTRDistinguishedNameInfo) IntermediateCACertificateID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("intermediateCACertificateID"))
	return rv
}/* debug [instance_properties/getter]: intermediateCACertificateID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDistinguishedNameInfo/nodeID
func (m_ MTRDistinguishedNameInfo) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}/* debug [instance_properties/getter]: nodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDistinguishedNameInfo/rootCACertificateID
func (m_ MTRDistinguishedNameInfo) RootCACertificateID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rootCACertificateID"))
	return rv
}/* debug [instance_properties/getter]: rootCACertificateID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDistinguishedNameInfo */



