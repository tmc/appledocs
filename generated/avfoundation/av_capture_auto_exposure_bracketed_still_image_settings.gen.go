// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureAutoExposureBracketedStillImageSettings */


/* debug [class_header]: Header for AVCaptureAutoExposureBracketedStillImageSettings */
// The class instance for the [CaptureAutoExposureBracketedStillImageSettings] class.
var (
	CaptureAutoExposureBracketedStillImageSettingsClass     _CaptureAutoExposureBracketedStillImageSettingsClass
	CaptureAutoExposureBracketedStillImageSettingsClassOnce sync.Once
)

func getCaptureAutoExposureBracketedStillImageSettingsClass() _CaptureAutoExposureBracketedStillImageSettingsClass {
	CaptureAutoExposureBracketedStillImageSettingsClassOnce.Do(func() {
		CaptureAutoExposureBracketedStillImageSettingsClass = _CaptureAutoExposureBracketedStillImageSettingsClass{objc.GetClass("AVCaptureAutoExposureBracketedStillImageSettings")}
	})
	return CaptureAutoExposureBracketedStillImageSettingsClass
}

type _CaptureAutoExposureBracketedStillImageSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureAutoExposureBracketedStillImageSettings */
// An interface definition for the [CaptureAutoExposureBracketedStillImageSettings] class.
type ICaptureAutoExposureBracketedStillImageSettings interface {
	ICaptureBracketedStillImageSettings
	
/* debug [class_interface_properties]: Properties for CaptureAutoExposureBracketedStillImageSettings */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureAutoExposureBracketedStillImageSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureAutoExposureBracketedStillImageSettings */
// Alloc allocates a new instance without initialization.
func (cc _CaptureAutoExposureBracketedStillImageSettingsClass) Alloc() CaptureAutoExposureBracketedStillImageSettings {
	rv := objc.Send[CaptureAutoExposureBracketedStillImageSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureAutoExposureBracketedStillImageSettingsClass) New() CaptureAutoExposureBracketedStillImageSettings {
	rv := objc.Send[CaptureAutoExposureBracketedStillImageSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureAutoExposureBracketedStillImageSettings) Init() CaptureAutoExposureBracketedStillImageSettings {
	rv := objc.Send[CaptureAutoExposureBracketedStillImageSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureAutoExposureBracketedStillImageSettings) Autorelease() CaptureAutoExposureBracketedStillImageSettings {
	rv := objc.Send[CaptureAutoExposureBracketedStillImageSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureAutoExposureBracketedStillImageSettings creates a new CaptureAutoExposureBracketedStillImageSettings instance.
func NewCaptureAutoExposureBracketedStillImageSettings() CaptureAutoExposureBracketedStillImageSettings {
	return getCaptureAutoExposureBracketedStillImageSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureAutoExposureBracketedStillImageSettings */
// A configuration for defining bracketed photo captures in terms of bias relative to automatic exposure.
//
// An instance defines the exposure target bias setting that should be applied to one image in a bracket. An array of objects is passed to to specify the bracketing. The minimum and maximum exposure target bias are properties of the instance supplying data to an instance. If you wish to leave unchanged for this bracketed still image, you may pass the value .


// A configuration for defining bracketed photo captures in terms of bias relative to automatic exposure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAutoExposureBracketedStillImageSettings
type CaptureAutoExposureBracketedStillImageSettings struct {
	CaptureBracketedStillImageSettings
}

// CaptureAutoExposureBracketedStillImageSettingsFrom constructs a [CaptureAutoExposureBracketedStillImageSettings] from an unsafe.Pointer.
//
// A configuration for defining bracketed photo captures in terms of bias relative to automatic exposure.
func CaptureAutoExposureBracketedStillImageSettingsFrom(ptr unsafe.Pointer) CaptureAutoExposureBracketedStillImageSettings {
	return CaptureAutoExposureBracketedStillImageSettings{
		CaptureBracketedStillImageSettings: CaptureBracketedStillImageSettingsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureAutoExposureBracketedStillImageSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureAutoExposureBracketedStillImageSettings */

// Creates an using the specified exposure target bias.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAutoExposureBracketedStillImageSettings/autoExposureSettings(exposureTargetBias:)
func (cc _CaptureAutoExposureBracketedStillImageSettingsClass) AutoExposureSettingsWithExposureTargetBias(exposureTargetBias float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("autoExposureSettingsWithExposureTargetBias:"), exposureTargetBias)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AutoExposureSettingsWithExposureTargetBias) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureAutoExposureBracketedStillImageSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureAutoExposureBracketedStillImageSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureAutoExposureBracketedStillImageSettings */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureAutoExposureBracketedStillImageSettings */


