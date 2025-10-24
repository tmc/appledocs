// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenKeyAlgorithm */


/* debug [class_header]: Header for TKTokenKeyAlgorithm */
// The class instance for the [TKTokenKeyAlgorithm] class.
var (
	TKTokenKeyAlgorithmClass     _TKTokenKeyAlgorithmClass
	TKTokenKeyAlgorithmClassOnce sync.Once
)

func getTKTokenKeyAlgorithmClass() _TKTokenKeyAlgorithmClass {
	TKTokenKeyAlgorithmClassOnce.Do(func() {
		TKTokenKeyAlgorithmClass = _TKTokenKeyAlgorithmClass{objc.GetClass("TKTokenKeyAlgorithm")}
	})
	return TKTokenKeyAlgorithmClass
}

type _TKTokenKeyAlgorithmClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenKeyAlgorithm */
// An interface definition for the [TKTokenKeyAlgorithm] class.
type ITKTokenKeyAlgorithm interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenKeyAlgorithm */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenKeyAlgorithm */
	// methods:
	IsAlgorithm(algorithm unsafe.Pointer) bool
	SupportsAlgorithm(algorithm unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenKeyAlgorithm */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenKeyAlgorithmClass) Alloc() TKTokenKeyAlgorithm {
	rv := objc.Send[TKTokenKeyAlgorithm](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenKeyAlgorithmClass) New() TKTokenKeyAlgorithm {
	rv := objc.Send[TKTokenKeyAlgorithm](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenKeyAlgorithm) Init() TKTokenKeyAlgorithm {
	rv := objc.Send[TKTokenKeyAlgorithm](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenKeyAlgorithm) Autorelease() TKTokenKeyAlgorithm {
	rv := objc.Send[TKTokenKeyAlgorithm](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeyAlgorithm creates a new TKTokenKeyAlgorithm instance.
func NewTKTokenKeyAlgorithm() TKTokenKeyAlgorithm {
	return getTKTokenKeyAlgorithmClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenKeyAlgorithm */
// Cryptographic algorithms used by token keys.
//
// Typically, the supported algorithm for a token key can be represented by a value of the enumeration. However, tokens such as Smart Cards require that input data for operations take the format of a more specific algorithm. For example, a token may accept raw data to generate a cryptographic signature, but require that raw data to be formatted according to PKCS1 padding rules. To express such a requirement, a object defines a target algorithm and a set of other algorithms that were used. In the previous example, the target algorithm is and the algorithm is also reported as being used.


// Cryptographic algorithms used by token keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyAlgorithm
type TKTokenKeyAlgorithm struct {
	objectivec.Object
}

// TKTokenKeyAlgorithmFrom constructs a [TKTokenKeyAlgorithm] from an unsafe.Pointer.
//
// Cryptographic algorithms used by token keys.
func TKTokenKeyAlgorithmFrom(ptr unsafe.Pointer) TKTokenKeyAlgorithm {
	return TKTokenKeyAlgorithm{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenKeyAlgorithm *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenKeyAlgorithm */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenKeyAlgorithm */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenKeyAlgorithm */

// Returns whether the specified algorithm is the target operation algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyAlgorithm/isAlgorithm(_:)
func (t_ TKTokenKeyAlgorithm) IsAlgorithm(algorithm unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAlgorithm:"), algorithm)
	return rv
}/* debug [instance_methods/method]: IsAlgorithm */


// Whether the specified algorithm is the target operation algorithm, or one of the other algorithms used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyAlgorithm/supportsAlgorithm(_:)
func (t_ TKTokenKeyAlgorithm) SupportsAlgorithm(algorithm unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("supportsAlgorithm:"), algorithm)
	return rv
}/* debug [instance_methods/method]: SupportsAlgorithm */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenKeyAlgorithm */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenKeyAlgorithm */



