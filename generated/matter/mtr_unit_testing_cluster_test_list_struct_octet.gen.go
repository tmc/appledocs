// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestListStructOctet */


/* debug [class_header]: Header for MTRUnitTestingClusterTestListStructOctet */
// The class instance for the [MTRUnitTestingClusterTestListStructOctet] class.
var (
	MTRUnitTestingClusterTestListStructOctetClass     _MTRUnitTestingClusterTestListStructOctetClass
	MTRUnitTestingClusterTestListStructOctetClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListStructOctetClass() _MTRUnitTestingClusterTestListStructOctetClass {
	MTRUnitTestingClusterTestListStructOctetClassOnce.Do(func() {
		MTRUnitTestingClusterTestListStructOctetClass = _MTRUnitTestingClusterTestListStructOctetClass{objc.GetClass("MTRUnitTestingClusterTestListStructOctet")}
	})
	return MTRUnitTestingClusterTestListStructOctetClass
}

type _MTRUnitTestingClusterTestListStructOctetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestListStructOctet */
// An interface definition for the [MTRUnitTestingClusterTestListStructOctet] class.
type IMTRUnitTestingClusterTestListStructOctet interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestListStructOctet */
	// properties:
	Member1() objc.IObject /* cross-framework: NSNumber */
	SetMember1(value objc.IObject /* cross-framework: NSNumber */)
	Member2() objc.IObject /* cross-framework: NSData */
	SetMember2(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestListStructOctet */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestListStructOctet */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListStructOctetClass) Alloc() MTRUnitTestingClusterTestListStructOctet {
	rv := objc.Send[MTRUnitTestingClusterTestListStructOctet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestListStructOctetClass) New() MTRUnitTestingClusterTestListStructOctet {
	rv := objc.Send[MTRUnitTestingClusterTestListStructOctet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListStructOctet) Init() MTRUnitTestingClusterTestListStructOctet {
	rv := objc.Send[MTRUnitTestingClusterTestListStructOctet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListStructOctet) Autorelease() MTRUnitTestingClusterTestListStructOctet {
	rv := objc.Send[MTRUnitTestingClusterTestListStructOctet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListStructOctet creates a new MTRUnitTestingClusterTestListStructOctet instance.
func NewMTRUnitTestingClusterTestListStructOctet() MTRUnitTestingClusterTestListStructOctet {
	return getMTRUnitTestingClusterTestListStructOctetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestListStructOctet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructOctet
type MTRUnitTestingClusterTestListStructOctet struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListStructOctetFrom constructs a [MTRUnitTestingClusterTestListStructOctet] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListStructOctetFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListStructOctet {
	return MTRUnitTestingClusterTestListStructOctet{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestListStructOctet *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestListStructOctet */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestListStructOctet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestListStructOctet */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestListStructOctet */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructOctet/member1
func (m_ MTRUnitTestingClusterTestListStructOctet) Member1() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("member1"))
	return rv
}/* debug [instance_properties/getter]: member1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructOctet/member1
func (m_ MTRUnitTestingClusterTestListStructOctet) SetMember1(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMember1:"), value)
}/* debug [instance_properties/setter]: member1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructOctet/member2
func (m_ MTRUnitTestingClusterTestListStructOctet) Member2() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("member2"))
	return rv
}/* debug [instance_properties/getter]: member2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructOctet/member2
func (m_ MTRUnitTestingClusterTestListStructOctet) SetMember2(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMember2:"), value)
}/* debug [instance_properties/setter]: member2 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestListStructOctet */



