// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCTensorData */


/* debug [class_header]: Header for MLCTensorData */
// The class instance for the [CTensorData] class.
var (
	CTensorDataClass     _CTensorDataClass
	CTensorDataClassOnce sync.Once
)

func getCTensorDataClass() _CTensorDataClass {
	CTensorDataClassOnce.Do(func() {
		CTensorDataClass = _CTensorDataClass{objc.GetClass("MLCTensorData")}
	})
	return CTensorDataClass
}

type _CTensorDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CTensorData */
// An interface definition for the [CTensorData] class.
type ICTensorData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CTensorData */
	// properties:
	Bytes() unsafe.Pointer
	Length() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CTensorData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CTensorData */
// Alloc allocates a new instance without initialization.
func (cc _CTensorDataClass) Alloc() CTensorData {
	rv := objc.Send[CTensorData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CTensorDataClass) New() CTensorData {
	rv := objc.Send[CTensorData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensorData) Init() CTensorData {
	rv := objc.Send[CTensorData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensorData) Autorelease() CTensorData {
	rv := objc.Send[CTensorData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensorData creates a new CTensorData instance.
func NewCTensorData() CTensorData {
	return getCTensorDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CTensorData */
// An encapsulation of the memory that tensor data uses.


// An encapsulation of the memory that tensor data uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData
type CTensorData struct {
	objectivec.Object
}

// CTensorDataFrom constructs a [CTensorData] from an unsafe.Pointer.
//
// An encapsulation of the memory that tensor data uses.
func CTensorDataFrom(ptr unsafe.Pointer) CTensorData {
	return CTensorData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CTensorData */

// Creates a tensor data instance with the buffer of data and length of bytes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData/init(bytesNoCopy:length:)
func NewCTensorDataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) CTensorData {
	rv := objc.Send[CTensorData](objc.ID(getCTensorDataClass().class), objc.Sel("dataWithBytesNoCopy:length:"), bytes, length)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorDataWithBytesNoCopyLength */


// Creates a tensor data instance with a data buffer, byte length, and custom deallocator closure you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData/init(bytesNoCopy:length:deallocator:)
func NewCTensorDataWithBytesNoCopyLengthDeallocator(bytes unsafe.Pointer, length uint, deallocator unsafe.Pointer) CTensorData {
	rv := objc.Send[CTensorData](objc.ID(getCTensorDataClass().class), objc.Sel("dataWithBytesNoCopy:length:deallocator:"), bytes, length, deallocator)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorDataWithBytesNoCopyLengthDeallocator */


// Creates a tensor data instance with the buffer of immutable data and length of bytes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData/init(immutableBytesNoCopy:length:)
func NewCTensorDataWithImmutableBytesNoCopyLength(bytes unsafe.Pointer, length uint) CTensorData {
	rv := objc.Send[CTensorData](objc.ID(getCTensorDataClass().class), objc.Sel("dataWithImmutableBytesNoCopy:length:"), bytes, length)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorDataWithImmutableBytesNoCopyLength */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CTensorData */

// Creates a tensor data instance with the buffer of data and length of bytes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData/init(bytesNoCopy:length:)
func (cc _CTensorDataClass) DataWithBytesNoCopyLength(bytes unsafe.Pointer, length uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("dataWithBytesNoCopy:length:"), bytes, length)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithBytesNoCopyLength) */


// Creates a tensor data instance with a data buffer, byte length, and custom deallocator closure you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData/init(bytesNoCopy:length:deallocator:)
func (cc _CTensorDataClass) DataWithBytesNoCopyLengthDeallocator(bytes unsafe.Pointer, length uint, deallocator unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("dataWithBytesNoCopy:length:deallocator:"), bytes, length, deallocator)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithBytesNoCopyLengthDeallocator) */


// Creates a tensor data instance with the buffer of immutable data and length of bytes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData/init(immutableBytesNoCopy:length:)
func (cc _CTensorDataClass) DataWithImmutableBytesNoCopyLength(bytes unsafe.Pointer, length uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("dataWithImmutableBytesNoCopy:length:"), bytes, length)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithImmutableBytesNoCopyLength) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CTensorData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CTensorData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CTensorData */

// A buffer that conains data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData/bytes
func (c_ CTensorData) Bytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("bytes"))
	return rv
}/* debug [instance_properties/getter]: bytes */


// The number of bytes you choose to hold for this tensor data instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData/length
func (c_ CTensorData) Length() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCTensorData */


