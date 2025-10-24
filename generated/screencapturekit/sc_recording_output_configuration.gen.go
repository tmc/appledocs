// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCRecordingOutputConfiguration */


/* debug [class_header]: Header for SCRecordingOutputConfiguration */
// The class instance for the [RecordingOutputConfiguration] class.
var (
	RecordingOutputConfigurationClass     _RecordingOutputConfigurationClass
	RecordingOutputConfigurationClassOnce sync.Once
)

func getRecordingOutputConfigurationClass() _RecordingOutputConfigurationClass {
	RecordingOutputConfigurationClassOnce.Do(func() {
		RecordingOutputConfigurationClass = _RecordingOutputConfigurationClass{objc.GetClass("SCRecordingOutputConfiguration")}
	})
	return RecordingOutputConfigurationClass
}

type _RecordingOutputConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecordingOutputConfiguration */
// An interface definition for the [RecordingOutputConfiguration] class.
type IRecordingOutputConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RecordingOutputConfiguration */
	// properties:
	AvailableOutputFileTypes() []string
	AvailableVideoCodecTypes() []string
	OutputFileType() FileType /* not a class type */
	SetOutputFileType(value FileType /* not a class type */)
	OutputURL() objc.IObject /* cross-framework: NSURL */
	SetOutputURL(value objc.IObject /* cross-framework: NSURL */)
	VideoCodecType() VideoCodecType /* not a class type */
	SetVideoCodecType(value VideoCodecType /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecordingOutputConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecordingOutputConfiguration */
// Alloc allocates a new instance without initialization.
func (rc _RecordingOutputConfigurationClass) Alloc() RecordingOutputConfiguration {
	rv := objc.Send[RecordingOutputConfiguration](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecordingOutputConfigurationClass) New() RecordingOutputConfiguration {
	rv := objc.Send[RecordingOutputConfiguration](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecordingOutputConfiguration) Init() RecordingOutputConfiguration {
	rv := objc.Send[RecordingOutputConfiguration](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecordingOutputConfiguration) Autorelease() RecordingOutputConfiguration {
	rv := objc.Send[RecordingOutputConfiguration](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecordingOutputConfiguration creates a new RecordingOutputConfiguration instance.
func NewRecordingOutputConfiguration() RecordingOutputConfiguration {
	return getRecordingOutputConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecordingOutputConfiguration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration
type RecordingOutputConfiguration struct {
	objectivec.Object
}

// RecordingOutputConfigurationFrom constructs a [RecordingOutputConfiguration] from an unsafe.Pointer.
func RecordingOutputConfigurationFrom(ptr unsafe.Pointer) RecordingOutputConfiguration {
	return RecordingOutputConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecordingOutputConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecordingOutputConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecordingOutputConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecordingOutputConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecordingOutputConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/availableOutputFileTypes
func (r_ RecordingOutputConfiguration) AvailableOutputFileTypes() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("availableOutputFileTypes"))
	return rv
}/* debug [instance_properties/getter]: availableOutputFileTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/availableVideoCodecTypes
func (r_ RecordingOutputConfiguration) AvailableVideoCodecTypes() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("availableVideoCodecTypes"))
	return rv
}/* debug [instance_properties/getter]: availableVideoCodecTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/outputFileType
func (r_ RecordingOutputConfiguration) OutputFileType() FileType /* not a class type */ {
	rv := objc.Send[FileType](r_.ID, objc.Sel("outputFileType"))
	return rv
}/* debug [instance_properties/getter]: outputFileType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/outputFileType
func (r_ RecordingOutputConfiguration) SetOutputFileType(value FileType /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFileType:"), value)
}/* debug [instance_properties/setter]: outputFileType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/outputURL
func (r_ RecordingOutputConfiguration) OutputURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](r_.ID, objc.Sel("outputURL"))
	return rv
}/* debug [instance_properties/getter]: outputURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/outputURL
func (r_ RecordingOutputConfiguration) SetOutputURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputURL:"), value)
}/* debug [instance_properties/setter]: outputURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/videoCodecType
func (r_ RecordingOutputConfiguration) VideoCodecType() VideoCodecType /* not a class type */ {
	rv := objc.Send[VideoCodecType](r_.ID, objc.Sel("videoCodecType"))
	return rv
}/* debug [instance_properties/getter]: videoCodecType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/videoCodecType
func (r_ RecordingOutputConfiguration) SetVideoCodecType(value VideoCodecType /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVideoCodecType:"), value)
}/* debug [instance_properties/setter]: videoCodecType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCRecordingOutputConfiguration */





