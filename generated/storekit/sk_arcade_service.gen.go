// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKArcadeService */


/* debug [class_header]: Header for SKArcadeService */
// The class instance for the [ArcadeService] class.
var (
	ArcadeServiceClass     _ArcadeServiceClass
	ArcadeServiceClassOnce sync.Once
)

func getArcadeServiceClass() _ArcadeServiceClass {
	ArcadeServiceClassOnce.Do(func() {
		ArcadeServiceClass = _ArcadeServiceClass{objc.GetClass("SKArcadeService")}
	})
	return ArcadeServiceClass
}

type _ArcadeServiceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ArcadeService */
// An interface definition for the [ArcadeService] class.
type IArcadeService interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ArcadeService */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ArcadeService */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ArcadeService */
// Alloc allocates a new instance without initialization.
func (ac _ArcadeServiceClass) Alloc() ArcadeService {
	rv := objc.Send[ArcadeService](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ArcadeServiceClass) New() ArcadeService {
	rv := objc.Send[ArcadeService](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArcadeService) Init() ArcadeService {
	rv := objc.Send[ArcadeService](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArcadeService) Autorelease() ArcadeService {
	rv := objc.Send[ArcadeService](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArcadeService creates a new ArcadeService instance.
func NewArcadeService() ArcadeService {
	return getArcadeServiceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ArcadeService */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKArcadeService
type ArcadeService struct {
	objectivec.Object
}

// ArcadeServiceFrom constructs a [ArcadeService] from an unsafe.Pointer.
func ArcadeServiceFrom(ptr unsafe.Pointer) ArcadeService {
	return ArcadeService{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ArcadeService *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ArcadeService */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKArcadeService/arcadeSubscriptionStatus(withNonce:resultHandler:)
func (ac _ArcadeServiceClass) ArcadeSubscriptionStatusWithNonceResultHandler(nonce uint64, resultHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("arcadeSubscriptionStatusWithNonce:resultHandler:"), nonce, resultHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArcadeSubscriptionStatusWithNonceResultHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKArcadeService/registerArcadeAppWithRandom(fromLib:randomFromLibLength:resultHandler:)
func (ac _ArcadeServiceClass) RegisterArcadeAppWithRandomFromLibRandomFromLibLengthResultHandler(randomFromLib objc.IObject /* cross-framework: NSData */, randomFromLibLength uint32 /* not a class type */, resultHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("registerArcadeAppWithRandomFromLib:randomFromLibLength:resultHandler:"), randomFromLib, randomFromLibLength, resultHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterArcadeAppWithRandomFromLibRandomFromLibLengthResultHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKArcadeService/repairArcadeApp()
func (ac _ArcadeServiceClass) RepairArcadeApp() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("repairArcadeApp"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RepairArcadeApp) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ArcadeService */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ArcadeService */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ArcadeService */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKArcadeService */



