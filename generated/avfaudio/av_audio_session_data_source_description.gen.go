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
	SetPreferredPolarPatternError(pattern IAudioSessionPolarPattern, outError unsafe.Pointer) bool
	PreferredPolarPattern() AudioSessionPolarPattern
	SelectedPolarPattern() AudioSessionPolarPattern
	SupportedPolarPatterns() []string
	OutputDataSource() AVAudioSessionDataSourceDescription
	SetOutputDataSource(value IAVAudioSessionDataSourceDescription)
	OutputDataSources() AVAudioSessionDataSourceDescription
	SetOutputDataSources(value IAVAudioSessionDataSourceDescription)
	DataSourceID() foundation.Number
	SetDataSourceID(value foundation.INumber)
	DataSourceName() string
	SetDataSourceName(value string)
	Location() corelocation.Location
	SetLocation(value corelocation.ILocation)
	Orientation() unsafe.Pointer
	SetOrientation(value unsafe.Pointer)
	DataSources() AVAudioSessionDataSourceDescription
	SetDataSources(value IAVAudioSessionDataSourceDescription)
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



// Selects the preferred directivity configuration for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/setPreferredPolarPattern(_:)
func (a_ AudioSessionDataSourceDescription) SetPreferredPolarPatternError(pattern IAudioSessionPolarPattern, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setPreferredPolarPattern:error:"), pattern, outError)
	return rv
}


// The preferred directivity configuration for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/preferredPolarPattern
func (a_ AudioSessionDataSourceDescription) PreferredPolarPattern() AudioSessionPolarPattern {
	rv := objc.Send[AudioSessionPolarPattern](a_.ID, objc.Sel("preferredPolarPattern"))
	return rv
}


// The data source’s active polar pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/selectedPolarPattern
func (a_ AudioSessionDataSourceDescription) SelectedPolarPattern() AudioSessionPolarPattern {
	rv := objc.Send[AudioSessionPolarPattern](a_.ID, objc.Sel("selectedPolarPattern"))
	return rv
}


// The set of directivity configurations supported by the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionDataSourceDescription/supportedPolarPatterns
func (a_ AudioSessionDataSourceDescription) SupportedPolarPatterns() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("supportedPolarPatterns"))
	return rv
}


// The currently selected output data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/outputdatasource
func (a_ AudioSessionDataSourceDescription) OutputDataSource() AVAudioSessionDataSourceDescription {
	rv := objc.Send[AVAudioSessionDataSourceDescription](a_.ID, objc.Sel("outputDataSource"))
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
func (a_ AudioSessionDataSourceDescription) OutputDataSources() AVAudioSessionDataSourceDescription {
	rv := objc.Send[AVAudioSessionDataSourceDescription](a_.ID, objc.Sel("outputDataSources"))
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
func (a_ AudioSessionDataSourceDescription) DataSourceID() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("dataSourceID"))
	return rv
}


// The system-assigned identifier for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/datasourceid
func (a_ AudioSessionDataSourceDescription) SetDataSourceID(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataSourceID:"), value)
}


// A human-readable name for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/datasourcename
func (a_ AudioSessionDataSourceDescription) DataSourceName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("dataSourceName"))
	return rv
}


// A human-readable name for the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/datasourcename
func (a_ AudioSessionDataSourceDescription) SetDataSourceName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataSourceName:"), objc.String(value))
}


// The location of the data source on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/location
func (a_ AudioSessionDataSourceDescription) Location() corelocation.Location {
	rv := objc.Send[corelocation.Location](a_.ID, objc.Sel("location"))
	return rv
}


// The location of the data source on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessiondatasourcedescription/location
func (a_ AudioSessionDataSourceDescription) SetLocation(value corelocation.ILocation) {
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


// The available data sources for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/datasources
func (a_ AudioSessionDataSourceDescription) DataSources() AVAudioSessionDataSourceDescription {
	rv := objc.Send[AVAudioSessionDataSourceDescription](a_.ID, objc.Sel("dataSources"))
	return rv
}


// The available data sources for the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessionportdescription/datasources
func (a_ AudioSessionDataSourceDescription) SetDataSources(value IAVAudioSessionDataSourceDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataSources:"), value)
}



