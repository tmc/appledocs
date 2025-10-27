// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4CommitOptions] class.
var (
	MTL4CommitOptionsClass     _MTL4CommitOptionsClass
	MTL4CommitOptionsClassOnce sync.Once
)

func getMTL4CommitOptionsClass() _MTL4CommitOptionsClass {
	MTL4CommitOptionsClassOnce.Do(func() {
		MTL4CommitOptionsClass = _MTL4CommitOptionsClass{objc.GetClass("MTL4CommitOptions")}
	})
	return MTL4CommitOptionsClass
}

type _MTL4CommitOptionsClass struct {
	class objc.Class
}





// An interface definition for the [MTL4CommitOptions] class.
type IMTL4CommitOptions interface {
	objectivec.IObject
	

	// properties:
	MTL4CommandQueueErrorDomain() foundation.foundation.INSString


	

	// methods:
	AddFeedbackHandler(block objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4CommitOptionsClass) Alloc() MTL4CommitOptions {
	rv := objc.Send[MTL4CommitOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4CommitOptionsClass) New() MTL4CommitOptions {
	rv := objc.Send[MTL4CommitOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CommitOptions) Init() MTL4CommitOptions {
	rv := objc.Send[MTL4CommitOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CommitOptions) Autorelease() MTL4CommitOptions {
	rv := objc.Send[MTL4CommitOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommitOptions creates a new MTL4CommitOptions instance.
func NewMTL4CommitOptions() MTL4CommitOptions {
	return getMTL4CommitOptionsClass().New()
}





// Represents options to configure a commit operation on a command queue.
//
// You pass these options as a parameter when you call . Note Instances of this class are not thread-safe. If your app modifies a shared commit options instance from multiple threads simultaneously, you are responsible for providing external synchronization.


// Represents options to configure a commit operation on a command queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommitOptions
type MTL4CommitOptions struct {
	objectivec.Object
}

// MTL4CommitOptionsFrom constructs a [MTL4CommitOptions] from an unsafe.Pointer.
//
// Represents options to configure a commit operation on a command queue.
func MTL4CommitOptionsFrom(ptr unsafe.Pointer) MTL4CommitOptions {
	return MTL4CommitOptions{objectivec.Object{objc.ID(ptr)}}
}




















// Registers a commit feedback handler that Metal calls with feedback data when available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommitOptions/addFeedbackHandler(_:)
func (m_ MTL4CommitOptions) AddFeedbackHandler(block objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addFeedbackHandler:"), block)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4CommitOptions) MTL4CommandQueueErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}








