// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MTRControllerFactoryParams] class.
type IMTRControllerFactoryParams interface {
	IMTRDeviceControllerFactoryParams
	// properties:
	CdCerts() objc.IObject /* cross-framework: Data */
	SetCdCerts(value objc.IObject /* cross-framework: Data */)
	PaaCerts() objc.IObject /* cross-framework: Data */
	SetPaaCerts(value objc.IObject /* cross-framework: Data */)
	StartServer() bool
	SetStartServer(value bool)
	StorageDelegate() unsafe.Pointer
	SetStorageDelegate(value unsafe.Pointer)
	// methods:
}



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

// Alloc allocates a new instance without initialization.
func (mc _MTRControllerFactoryParamsClass) Alloc() MTRControllerFactoryParams {
	rv := objc.Send[MTRControllerFactoryParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactoryparams/cdcerts
func (m_ MTRControllerFactoryParams) CdCerts() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("cdCerts"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactoryparams/cdcerts
func (m_ MTRControllerFactoryParams) SetCdCerts(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCdCerts:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactoryparams/paacerts
func (m_ MTRControllerFactoryParams) PaaCerts() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("paaCerts"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactoryparams/paacerts
func (m_ MTRControllerFactoryParams) SetPaaCerts(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPaaCerts:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactoryparams/startserver
func (m_ MTRControllerFactoryParams) StartServer() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("startServer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactoryparams/startserver
func (m_ MTRControllerFactoryParams) SetStartServer(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartServer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactoryparams/storagedelegate
func (m_ MTRControllerFactoryParams) StorageDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("storageDelegate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontrollerfactoryparams/storagedelegate
func (m_ MTRControllerFactoryParams) SetStorageDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStorageDelegate:"), value)
}



