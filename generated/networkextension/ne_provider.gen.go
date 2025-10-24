// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEProvider */


/* debug [class_header]: Header for NEProvider */
// The class instance for the [NEProvider] class.
var (
	NEProviderClass     _NEProviderClass
	NEProviderClassOnce sync.Once
)

func getNEProviderClass() _NEProviderClass {
	NEProviderClassOnce.Do(func() {
		NEProviderClass = _NEProviderClass{objc.GetClass("NEProvider")}
	})
	return NEProviderClass
}

type _NEProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEProvider */
// An interface definition for the [NEProvider] class.
type INEProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEProvider */
	// properties:
	DefaultPath() INWPath
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEProvider */
	// methods:
	SleepWithCompletionHandler(completionHandler unsafe.Pointer)
	Wake()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEProvider */
// Alloc allocates a new instance without initialization.
func (nc _NEProviderClass) Alloc() NEProvider {
	rv := objc.Send[NEProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEProviderClass) New() NEProvider {
	rv := objc.Send[NEProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEProvider) Init() NEProvider {
	rv := objc.Send[NEProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEProvider) Autorelease() NEProvider {
	rv := objc.Send[NEProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProvider creates a new NEProvider instance.
func NewNEProvider() NEProvider {
	return getNEProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEProvider */
// An abstract base class for all NetworkExtension providers.
//
// See the documentation for the subclasses for details about how to create Network Extension Provider extensions. The class and its subclasses expose methods and properties that allow Network Extension Provider extensions to participate in and affect the network data path on iOS and macOS. For example, the method in allows Filter Data Provider extensions to make pass/block decisions on TCP connections as the connections are established on the system.


// An abstract base class for all NetworkExtension providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProvider
type NEProvider struct {
	objectivec.Object
}

// NEProviderFrom constructs a [NEProvider] from an unsafe.Pointer.
//
// An abstract base class for all NetworkExtension providers.
func NEProviderFrom(ptr unsafe.Pointer) NEProvider {
	return NEProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEProvider */

// Starts the Network Extension machinery from inside a System Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProvider/startSystemExtensionMode()
func (nc _NEProviderClass) StartSystemExtensionMode() {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("startSystemExtensionMode"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StartSystemExtensionMode) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEProvider */

// Handle a sleep event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProvider/sleep(completionHandler:)
func (n_ NEProvider) SleepWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("sleepWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: SleepWithCompletionHandler */


// Handle a wake event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProvider/wake()
func (n_ NEProvider) Wake() {
	objc.Send[objc.ID](n_.ID, objc.Sel("wake"))
}/* debug [instance_methods/method]: Wake */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEProvider */

// The current default network path used for connections created by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProvider/defaultPath
func (n_ NEProvider) DefaultPath() INWPath {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("defaultPath"))
	return rv
}/* debug [instance_properties/getter]: defaultPath */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEProvider */



