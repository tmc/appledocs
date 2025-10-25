// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSampleBufferGeneratorBatch */


/* debug [class_header]: Header for AVSampleBufferGeneratorBatch */
// The class instance for the [SampleBufferGeneratorBatch] class.
var (
	SampleBufferGeneratorBatchClass     _SampleBufferGeneratorBatchClass
	SampleBufferGeneratorBatchClassOnce sync.Once
)

func getSampleBufferGeneratorBatchClass() _SampleBufferGeneratorBatchClass {
	SampleBufferGeneratorBatchClassOnce.Do(func() {
		SampleBufferGeneratorBatchClass = _SampleBufferGeneratorBatchClass{objc.GetClass("AVSampleBufferGeneratorBatch")}
	})
	return SampleBufferGeneratorBatchClass
}

type _SampleBufferGeneratorBatchClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SampleBufferGeneratorBatch */
// An interface definition for the [SampleBufferGeneratorBatch] class.
type ISampleBufferGeneratorBatch interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SampleBufferGeneratorBatch */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SampleBufferGeneratorBatch */
	// methods:
	Cancel()
	MakeDataReadyWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SampleBufferGeneratorBatch */
// Alloc allocates a new instance without initialization.
func (sc _SampleBufferGeneratorBatchClass) Alloc() SampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SampleBufferGeneratorBatchClass) New() SampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferGeneratorBatch) Init() SampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferGeneratorBatch) Autorelease() SampleBufferGeneratorBatch {
	rv := objc.Send[SampleBufferGeneratorBatch](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferGeneratorBatch creates a new SampleBufferGeneratorBatch instance.
func NewSampleBufferGeneratorBatch() SampleBufferGeneratorBatch {
	return getSampleBufferGeneratorBatchClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SampleBufferGeneratorBatch */
// An object that generates sample buffers in a batch.
//
// The benefit of batching is it aggregates adjacent I/O requests and overlaps them when possible for all sample buffers within the batch.


// An object that generates sample buffers in a batch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch
type SampleBufferGeneratorBatch struct {
	objectivec.Object
}

// SampleBufferGeneratorBatchFrom constructs a [SampleBufferGeneratorBatch] from an unsafe.Pointer.
//
// An object that generates sample buffers in a batch.
func SampleBufferGeneratorBatchFrom(ptr unsafe.Pointer) SampleBufferGeneratorBatch {
	return SampleBufferGeneratorBatch{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SampleBufferGeneratorBatch *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SampleBufferGeneratorBatch */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SampleBufferGeneratorBatch */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SampleBufferGeneratorBatch */

// Cancels any I/O for this batch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch/cancel()
func (s_ SampleBufferGeneratorBatch) Cancel() {
	objc.Send[objc.ID](s_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Loads sample data asynchronously for all sample buffers within a batch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch/makeDataReady(completionHandler:)
func (s_ SampleBufferGeneratorBatch) MakeDataReadyWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("makeDataReadyWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: MakeDataReadyWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SampleBufferGeneratorBatch */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSampleBufferGeneratorBatch */



