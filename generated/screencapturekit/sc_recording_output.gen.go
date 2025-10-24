// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCRecordingOutput */


/* debug [class_header]: Header for SCRecordingOutput */
// The class instance for the [RecordingOutput] class.
var (
	RecordingOutputClass     _RecordingOutputClass
	RecordingOutputClassOnce sync.Once
)

func getRecordingOutputClass() _RecordingOutputClass {
	RecordingOutputClassOnce.Do(func() {
		RecordingOutputClass = _RecordingOutputClass{objc.GetClass("SCRecordingOutput")}
	})
	return RecordingOutputClass
}

type _RecordingOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecordingOutput */
// An interface definition for the [RecordingOutput] class.
type IRecordingOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RecordingOutput */
	// properties:
	RecordedDuration() objc.IObject /* cross-framework: Time */
	RecordedFileSize() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecordingOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecordingOutput */
// Alloc allocates a new instance without initialization.
func (rc _RecordingOutputClass) Alloc() RecordingOutput {
	rv := objc.Send[RecordingOutput](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecordingOutputClass) New() RecordingOutput {
	rv := objc.Send[RecordingOutput](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecordingOutput) Init() RecordingOutput {
	rv := objc.Send[RecordingOutput](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecordingOutput) Autorelease() RecordingOutput {
	rv := objc.Send[RecordingOutput](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecordingOutput creates a new RecordingOutput instance.
func NewRecordingOutput() RecordingOutput {
	return getRecordingOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecordingOutput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput
type RecordingOutput struct {
	objectivec.Object
}

// RecordingOutputFrom constructs a [RecordingOutput] from an unsafe.Pointer.
func RecordingOutputFrom(ptr unsafe.Pointer) RecordingOutput {
	return RecordingOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecordingOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput/init(configuration:delegate:)
func NewRecordingOutputWithConfigurationDelegate(recordingOutputConfiguration ISCRecordingOutputConfiguration, delegate unsafe.Pointer) RecordingOutput {
	instance := getRecordingOutputClass().Alloc()
	rv := objc.Send[RecordingOutput](instance.ID, objc.Sel("initWithConfiguration:delegate:"), recordingOutputConfiguration, delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRecordingOutputWithConfigurationDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecordingOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecordingOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecordingOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecordingOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput/recordedDuration
func (r_ RecordingOutput) RecordedDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](r_.ID, objc.Sel("recordedDuration"))
	return rv
}/* debug [instance_properties/getter]: recordedDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput/recordedFileSize
func (r_ RecordingOutput) RecordedFileSize() int {
	rv := objc.Send[int](r_.ID, objc.Sel("recordedFileSize"))
	return rv
}/* debug [instance_properties/getter]: recordedFileSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCRecordingOutput */


