// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioSessionDataSourceDescription] class.
var (
	AudioSessionDataSourceDescriptionClass     _AudioSessionDataSourceDescriptionClass
	AudioSessionDataSourceDescriptionClassOnce sync.Once
)

func getAudioSessionDataSourceDescriptionClass() _AudioSessionDataSourceDescriptionClass {
	AudioSessionDataSourceDescriptionClassOnce.Do(func() {
		AudioSessionDataSourceDescriptionClass = _AudioSessionDataSourceDescriptionClass{objc.GetClass("AVAudioSessionDataSourceDescription")}
	})
	return AudioSessionDataSourceDescriptionClass
}

type _AudioSessionDataSourceDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AudioSessionDataSourceDescription] class.
type IAudioSessionDataSourceDescription interface {
	objectivec.IObject
	// properties:
	OutputDataSource() IAVAudioSessionDataSourceDescription
	SetOutputDataSource(value IAVAudioSessionDataSourceDescription)
	OutputDataSources() IAVAudioSessionDataSourceDescription
	SetOutputDataSources(value IAVAudioSessionDataSourceDescription)
	DataSourceID() objc.IObject /* cross-framework: NSNumber */
	SetDataSourceID(value objc.IObject /* cross-framework: NSNumber */)
	DataSourceName() objc.IObject /* cross-framework: NSString */
	SetDataSourceName(value objc.IObject /* cross-framework: NSString */)
	Location() objc.IObject /* cross-framework: Location */
	SetLocation(value objc.IObject /* cross-framework: Location */)
	Orientation() unsafe.Pointer
	SetOrientation(value unsafe.Pointer)
	PreferredPolarPattern() unsafe.Pointer
	SetPreferredPolarPattern(value unsafe.Pointer)
	SelectedPolarPattern() unsafe.Pointer
	SetSelectedPolarPattern(value unsafe.Pointer)
	DataSources() IAVAudioSessionDataSourceDescription
	SetDataSources(value IAVAudioSessionDataSourceDescription)
	// methods:
}

// An object that defines a data source for an audio input or output, giving information such as the source’s name, location, and orientation.
//
// You obtain data source descriptions from the shared object or the objects corresponding to its input and output ports. Only built-in microphone ports on certain devices support the location, orientation, and polar pattern properties. If a port doesn’t support these features, the value of its property is . This class is especially useful for differentiating between microphone configurations on devices having more than one built-in microphone. Such devices may also support signal processing features for spatial filtering, or , in which the system makes the device more sensitive to audio signals from a particular direction. See for more information.


// An object that defines a data source for an audio input or output, giving information such as the source’s name, location, and orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription
type AudioSessionDataSourceDescription struct {
	objectivec.Object
}

// AudioSessionDataSourceDescriptionFrom constructs a [AudioSessionDataSourceDescription] from an unsafe.Pointer.
//
// An object that defines a data source for an audio input or output, giving information such as the source’s name, location, and orientation.
func AudioSessionDataSourceDescriptionFrom(ptr unsafe.Pointer) AudioSessionDataSourceDescription {
	return AudioSessionDataSourceDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioSessionDataSourceDescriptionClass) Alloc() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioSessionDataSourceDescriptionClass) New() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionDataSourceDescription) Init() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionDataSourceDescription) Autorelease() AudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionDataSourceDescription creates a new AudioSessionDataSourceDescription instance.
func NewAudioSessionDataSourceDescription() AudioSessionDataSourceDescription {
	return getAudioSessionDataSourceDescriptionClass().New()
}



// The currently selected output data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasource
func (a_ AudioSessionDataSourceDescription) OutputDataSource() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("outputDataSource"))
	return rv
}


// The currently selected output data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasource
func (a_ AudioSessionDataSourceDescription) SetOutputDataSource(value IAVAudioSessionDataSourceDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputDataSource:"), value)
}


// An array of available output data sources for the current audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasources
func (a_ AudioSessionDataSourceDescription) OutputDataSources() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("outputDataSources"))
	return rv
}


// An array of available output data sources for the current audio route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasources
func (a_ AudioSessionDataSourceDescription) SetOutputDataSources(value IAVAudioSessionDataSourceDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputDataSources:"), value)
}


// The system-assigned identifier for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/datasourceid
func (a_ AudioSessionDataSourceDescription) DataSourceID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("dataSourceID"))
	return rv
}


// The system-assigned identifier for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/datasourceid
func (a_ AudioSessionDataSourceDescription) SetDataSourceID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataSourceID:"), value)
}


// A human-readable name for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/datasourcename
func (a_ AudioSessionDataSourceDescription) DataSourceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("dataSourceName"))
	return rv
}


// A human-readable name for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/datasourcename
func (a_ AudioSessionDataSourceDescription) SetDataSourceName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataSourceName:"), value)
}


// The location of the data source on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/location
func (a_ AudioSessionDataSourceDescription) Location() objc.IObject /* cross-framework: Location */ {
	rv := objc.Send[corelocation.Location](a_.ID, objc.Sel("location"))
	return rv
}


// The location of the data source on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/location
func (a_ AudioSessionDataSourceDescription) SetLocation(value objc.IObject /* cross-framework: Location */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocation:"), value)
}


// The orientation of the data source relative to the device’s natural orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/orientation
func (a_ AudioSessionDataSourceDescription) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation of the data source relative to the device’s natural orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/orientation
func (a_ AudioSessionDataSourceDescription) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOrientation:"), value)
}


// The preferred directivity configuration for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/preferredpolarpattern
func (a_ AudioSessionDataSourceDescription) PreferredPolarPattern() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredPolarPattern"))
	return rv
}


// The preferred directivity configuration for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/preferredpolarpattern
func (a_ AudioSessionDataSourceDescription) SetPreferredPolarPattern(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredPolarPattern:"), value)
}


// The data source’s active polar pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/selectedpolarpattern
func (a_ AudioSessionDataSourceDescription) SelectedPolarPattern() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("selectedPolarPattern"))
	return rv
}


// The data source’s active polar pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/selectedpolarpattern
func (a_ AudioSessionDataSourceDescription) SetSelectedPolarPattern(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedPolarPattern:"), value)
}


// The available data sources for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/datasources
func (a_ AudioSessionDataSourceDescription) DataSources() IAVAudioSessionDataSourceDescription {
	rv := objc.Send[AudioSessionDataSourceDescription](a_.ID, objc.Sel("dataSources"))
	return rv
}


// The available data sources for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/datasources
func (a_ AudioSessionDataSourceDescription) SetDataSources(value IAVAudioSessionDataSourceDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataSources:"), value)
}


