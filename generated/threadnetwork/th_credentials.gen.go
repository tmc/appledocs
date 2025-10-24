// Code generated from Apple documentation for ThreadNetwork. DO NOT EDIT.

package threadnetwork

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class THCredentials */


/* debug [class_header]: Header for THCredentials */
// The class instance for the [THCredentials] class.
var (
	THCredentialsClass     _THCredentialsClass
	THCredentialsClassOnce sync.Once
)

func getTHCredentialsClass() _THCredentialsClass {
	THCredentialsClassOnce.Do(func() {
		THCredentialsClass = _THCredentialsClass{objc.GetClass("THCredentials")}
	})
	return THCredentialsClass
}

type _THCredentialsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for THCredentials */
// An interface definition for the [THCredentials] class.
type ITHCredentials interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for THCredentials */
	// properties:
	ActiveOperationalDataSet() objc.IObject /* cross-framework: NSData */
	BorderAgentID() objc.IObject /* cross-framework: NSData */
	Channel() uint8 /* not a class type */
	SetChannel(value uint8 /* not a class type */)
	CreationDate() objc.IObject /* cross-framework: NSDate */
	ExtendedPANID() objc.IObject /* cross-framework: NSData */
	LastModificationDate() objc.IObject /* cross-framework: NSDate */
	NetworkKey() objc.IObject /* cross-framework: NSData */
	NetworkName() objc.IObject /* cross-framework: NSString */
	PanID() objc.IObject /* cross-framework: NSData */
	PSKC() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for THCredentials */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for THCredentials */
// Alloc allocates a new instance without initialization.
func (tc _THCredentialsClass) Alloc() THCredentials {
	rv := objc.Send[THCredentials](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _THCredentialsClass) New() THCredentials {
	rv := objc.Send[THCredentials](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ THCredentials) Init() THCredentials {
	rv := objc.Send[THCredentials](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ THCredentials) Autorelease() THCredentials {
	rv := objc.Send[THCredentials](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTHCredentials creates a new THCredentials instance.
func NewTHCredentials() THCredentials {
	return getTHCredentialsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for THCredentials */
// A class that contains credentials for a Thread network.
//
// A Thread network defines parameters that all connected devices use. provides these parameters.


// A class that contains credentials for a Thread network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials
type THCredentials struct {
	objectivec.Object
}

// THCredentialsFrom constructs a [THCredentials] from an unsafe.Pointer.
//
// A class that contains credentials for a Thread network.
func THCredentialsFrom(ptr unsafe.Pointer) THCredentials {
	return THCredentials{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for THCredentials *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for THCredentials */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for THCredentials */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for THCredentials */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for THCredentials */

// The essential operational parameters for the Thread network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/activeOperationalDataSet
func (t_ THCredentials) ActiveOperationalDataSet() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("activeOperationalDataSet"))
	return rv
}/* debug [instance_properties/getter]: activeOperationalDataSet */


// The identifier of an active Thread network Border Agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/borderAgentID
func (t_ THCredentials) BorderAgentID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("borderAgentID"))
	return rv
}/* debug [instance_properties/getter]: borderAgentID */


// The Thread network radio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/channel
func (t_ THCredentials) Channel() uint8 /* not a class type */ {
	rv := objc.Send[uint8](t_.ID, objc.Sel("channel"))
	return rv
}/* debug [instance_properties/getter]: channel */


// The Thread network radio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/channel
func (t_ THCredentials) SetChannel(value uint8 /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChannel:"), value)
}/* debug [instance_properties/setter]: channel */


// The date and time that the framework stored the credential in the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/creationDate
func (t_ THCredentials) CreationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("creationDate"))
	return rv
}/* debug [instance_properties/getter]: creationDate */


// The Thread network extended PAN identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/extendedPANID
func (t_ THCredentials) ExtendedPANID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("extendedPANID"))
	return rv
}/* debug [instance_properties/getter]: extendedPANID */


// The date and time that the framework updated the credential in the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/lastModificationDate
func (t_ THCredentials) LastModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("lastModificationDate"))
	return rv
}/* debug [instance_properties/getter]: lastModificationDate */


// The Thread network key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/networkKey
func (t_ THCredentials) NetworkKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("networkKey"))
	return rv
}/* debug [instance_properties/getter]: networkKey */


// The Thread network name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/networkName
func (t_ THCredentials) NetworkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("networkName"))
	return rv
}/* debug [instance_properties/getter]: networkName */


// The Thread network PAN identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/panID
func (t_ THCredentials) PanID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("panID"))
	return rv
}/* debug [instance_properties/getter]: panID */


// The Thread network pre-shared key (PSKC) for the Commissioner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/pskc
func (t_ THCredentials) PSKC() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("PSKC"))
	return rv
}/* debug [instance_properties/getter]: PSKC */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class THCredentials */






