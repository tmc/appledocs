// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRControllerFactoryParams */


/* debug [class_header]: Header for MTRControllerFactoryParams */
// The class instance for the [MTRControllerFactoryParams] class.
var (
	MTRControllerFactoryParamsClass     _MTRControllerFactoryParamsClass
	MTRControllerFactoryParamsClassOnce sync.Once
)

func getMTRControllerFactoryParamsClass() _MTRControllerFactoryParamsClass {
	MTRControllerFactoryParamsClassOnce.Do(func() {
		MTRControllerFactoryParamsClass = _MTRControllerFactoryParamsClass{objc.GetClass("MTRControllerFactoryParams")}
	})
	return MTRControllerFactoryParamsClass
}

type _MTRControllerFactoryParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRControllerFactoryParams */
// An interface definition for the [MTRControllerFactoryParams] class.
type IMTRControllerFactoryParams interface {
	IMTRDeviceControllerFactoryParams
	
/* debug [class_interface_properties]: Properties for MTRControllerFactoryParams */
	// properties:
	CdCerts() []foundation.Data
	SetCdCerts(value []foundation.Data)
	PaaCerts() []foundation.Data
	SetPaaCerts(value []foundation.Data)
	StartServer() bool
	SetStartServer(value bool)
	StorageDelegate() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRControllerFactoryParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRControllerFactoryParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRControllerFactoryParamsClass) Alloc() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRControllerFactoryParamsClass) New() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRControllerFactoryParams) Init() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRControllerFactoryParams) Autorelease() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRControllerFactoryParams creates a new MTRControllerFactoryParams instance.
func NewMTRControllerFactoryParams() MTRControllerFactoryParams {
	return getMTRControllerFactoryParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRControllerFactoryParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams
type MTRControllerFactoryParams struct {
	MTRDeviceControllerFactoryParams
}

// MTRControllerFactoryParamsFrom constructs a [MTRControllerFactoryParams] from an unsafe.Pointer.
func MTRControllerFactoryParamsFrom(ptr unsafe.Pointer) MTRControllerFactoryParams {
	return MTRControllerFactoryParams{
		MTRDeviceControllerFactoryParams: MTRDeviceControllerFactoryParamsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRControllerFactoryParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRControllerFactoryParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRControllerFactoryParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRControllerFactoryParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRControllerFactoryParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams/cdCerts
func (m_ MTRControllerFactoryParams) CdCerts() []foundation.Data {
	rv := objc.Send[[]foundation.Data](m_.ID, objc.Sel("cdCerts"))
	return rv
}/* debug [instance_properties/getter]: cdCerts */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams/cdCerts
func (m_ MTRControllerFactoryParams) SetCdCerts(value []foundation.Data) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setCdCerts:"), nsArray)
}/* debug [instance_properties/setter]: cdCerts */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams/paaCerts
func (m_ MTRControllerFactoryParams) PaaCerts() []foundation.Data {
	rv := objc.Send[[]foundation.Data](m_.ID, objc.Sel("paaCerts"))
	return rv
}/* debug [instance_properties/getter]: paaCerts */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams/paaCerts
func (m_ MTRControllerFactoryParams) SetPaaCerts(value []foundation.Data) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setPaaCerts:"), nsArray)
}/* debug [instance_properties/setter]: paaCerts */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams/startServer
func (m_ MTRControllerFactoryParams) StartServer() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("startServer"))
	return rv
}/* debug [instance_properties/getter]: startServer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams/startServer
func (m_ MTRControllerFactoryParams) SetStartServer(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartServer:"), value)
}/* debug [instance_properties/setter]: startServer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactoryParams/storageDelegate
func (m_ MTRControllerFactoryParams) StorageDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("storageDelegate"))
	return rv
}/* debug [instance_properties/getter]: storageDelegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRControllerFactoryParams */



