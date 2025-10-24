// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox
import (
	"unsafe"

	"github.com/tmc/appledocs/generated/avfaudio"

	"github.com/tmc/appledocs/generated/coreaudiotypes"

	"github.com/tmc/appledocs/generated/objectivec"
)


// C struct types
// AUChannelInfo - The audio input and output channel capabilities for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUChannelInfo
type AUChannelInfo struct {
	InChannels unsafe.Pointer // The number of input channels.
	OutChannels unsafe.Pointer // The number of output channels.
}

// AUDependentParameter - An audio unit parameter whose value can change in response to a change in its parent metaparameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUDependentParameter
type AUDependentParameter struct {
	MParameterID AudioUnitParameterID
	MScope AudioUnitScope
}

// AudioBalanceFade - Describes audio left/right balance and front/back fade values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBalanceFade
type AudioBalanceFade struct {
	MBackFrontFade unsafe.Pointer // The audio front/back fade, where -1 represents full rear, 0 represents center, and +1 represents full front.
	MChannelLayout avfaudio.AudioChannelLayout // The size, in bytes, of the   parameter.
	MLeftRightBalance unsafe.Pointer // The audio left/right balance, where -1 represents full left, 0 represents center, and +1 represents full right.
	MType AudioBalanceFadeType // An AudioBalanceFadeType constant. max unity gain, or equal power.
}

// AudioBytePacketTranslation - A data structure used by the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBytePacketTranslation
type AudioBytePacketTranslation struct {
	MByte unsafe.Pointer // A byte number.
	MByteOffsetInPacket unsafe.Pointer // A byte offset in a packet.
	MFlags AudioBytePacketTranslationFlags // A translation flag value.
	MPacket unsafe.Pointer // A packet number.
}

// AudioCodecMagicCookieInfo - A structure holding magic cookie information needed by some codecs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecMagicCookieInfo
type AudioCodecMagicCookieInfo struct {
	MMagicCookie unsafe.Pointer // Generic constant pointer to the magic cookie.
	MMagicCookieSize unsafe.Pointer // The size of the magic cookie.
}

// AudioCodecPrimeInfo - A structure specifying the number of leading and trailing empty frames to be inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecPrimeInfo
type AudioCodecPrimeInfo struct {
	LeadingFrames unsafe.Pointer // An unsigned integer specifying the number of leading empty frames.
	TrailingFrames unsafe.Pointer // An unsigned integer specifying the number of trailing empty frames.
}

// AudioComponentDescription - Identifying information for an audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentDescription
type AudioComponentDescription struct {
	ComponentFlags unsafe.Pointer // Set this value to zero.
	ComponentFlagsMask unsafe.Pointer // Set this value to zero.
	ComponentManufacturer unsafe.Pointer // The unique vendor identifier, registered with Apple, for the audio component.
	ComponentSubType unsafe.Pointer // A 4-byte code that you can use to indicate the purpose of a component. For example, you could use   or   as a mnemonic indication that an audio unit is a low-pass filter.
	ComponentType unsafe.Pointer // A unique 4-byte code identifying the interface for the component.
}

// AudioComponentPlugInInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentPlugInInterface
type AudioComponentPlugInInterface struct {
	Close unsafe.Pointer
	Lookup unsafe.Pointer
	Open unsafe.Pointer
	Reserved unsafe.Pointer
}

// AudioConverterPrimeInfo - Specifies priming information for an audio converter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterPrimeInfo
type AudioConverterPrimeInfo struct {
	LeadingFrames unsafe.Pointer // The number of leading frames of input audio data required for the converter to perform high-quality conversion.
	TrailingFrames unsafe.Pointer // The number of trailing frames of input audio data required by the converter to perform high-quality conversion. Trailing frames follow, in time, the expected final input frame. Your application should be prepared to provide this number of additional input frames except when using the   value for the   property. If no additional frames are available in the input stream (because, for example, the desired end frame is at the end of an audio file), then the audio converter synthesizes a sufficient number of silent ( -valued) trailing frames.
}

// AudioFile_SMPTE_Time - A data structure for describing SMPTE (Society of Motion Picture and Television Engineers) time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_SMPTE_Time
type AudioFile_SMPTE_Time struct {
	MFrames unsafe.Pointer // The frames.
	MHours unsafe.Pointer // The hours.
	MMinutes unsafe.Pointer // The minutes.
	MSeconds unsafe.Pointer // The seconds.
	MSubFrameSampleOffset unsafe.Pointer // The sample offset within a frame.
}

// AudioFileFDFTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFDFTable
type AudioFileFDFTable struct {
	MComponentStorage unsafe.Pointer
	MCountUserDataFDF CountUserDataFDF
	MGetPropertyFDF GetPropertyFDF
	MGetPropertyInfoFDF GetPropertyInfoFDF
	MGetUserDataFDF GetUserDataFDF
	MGetUserDataSizeFDF GetUserDataSizeFDF
	MReadBytesFDF ReadBytesFDF
	MReadPacketsFDF ReadPacketsFDF
	MSetPropertyFDF SetPropertyFDF
	MSetUserDataFDF SetUserDataFDF
	MWriteBytesFDF WriteBytesFDF
	MWritePacketsFDF WritePacketsFDF
}

// AudioFileFDFTableExtended
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFDFTableExtended
type AudioFileFDFTableExtended struct {
	MComponentStorage unsafe.Pointer
	MCountUserDataFDF CountUserDataFDF
	MGetPropertyFDF GetPropertyFDF
	MGetPropertyInfoFDF GetPropertyInfoFDF
	MGetUserDataFDF GetUserDataFDF
	MGetUserDataSizeFDF GetUserDataSizeFDF
	MReadBytesFDF ReadBytesFDF
	MReadPacketDataFDF ReadPacketDataFDF
	MReadPacketsFDF ReadPacketsFDF
	MSetPropertyFDF SetPropertyFDF
	MSetUserDataFDF SetUserDataFDF
	MWriteBytesFDF WriteBytesFDF
	MWritePacketsFDF WritePacketsFDF
}

// AudioFileMarker - Annotates a position in an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileMarker
type AudioFileMarker struct {
	MChannel unsafe.Pointer // The channel number referred to by the marker. Set to   if the marker applies to all channels.
	MFramePosition unsafe.Pointer // The frame in the file, counting from the start of the audio data.
	MMarkerID unsafe.Pointer // A unique ID for the marker.
	MName StringRef // The name of the marker.
	MReserved unsafe.Pointer // A reserved field. Set to  .
	MSMPTETime AudioFile_SMPTE_Time // The SMPTE time for this marker.
	MType unsafe.Pointer // The marker type.
}

// AudioFileMarkerList - A list of markers associated with an audio file, including their SMPTE time type, the number of markers, and the markers themselves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileMarkerList
type AudioFileMarkerList struct {
	MMarkers AudioFileMarker // An array of   elements, each of which is an audio file marker.
	MNumberMarkers unsafe.Pointer // The number of markers in the list.
	MSMPTE_TimeType unsafe.Pointer // The SMPTE time type of the whole list of markers in an audio file.
}

// AudioFilePacketTableInfo - Contains information about the number of valid frames in a file and where they begin and end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFilePacketTableInfo
type AudioFilePacketTableInfo struct {
	MNumberValidFrames unsafe.Pointer // The number of valid frames in the file.
	MPrimingFrames unsafe.Pointer // The number of invalid frames at the beginning of the file.
	MRemainderFrames unsafe.Pointer // The number of invalid frames at the end of the file.
}

// AudioFileRegion - An audio file region specifies a segment of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegion
type AudioFileRegion struct {
	MFlags AudioFileRegionFlags // Audio File Services region flags.
	MMarkers AudioFileMarker // An array of   elements describing where the data in the region starts.
	MName StringRef // The name of the region.
	MNumberMarkers unsafe.Pointer // The number of markers in the array specified in the    parameter.
	MRegionID unsafe.Pointer // A unique ID associated with the audio file region.
}

// AudioFileRegionList - A list of the audio file regions in a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegionList
type AudioFileRegionList struct {
	MNumberRegions unsafe.Pointer // The number of regions in the list specified in the    parameter.
	MRegions AudioFileRegion // A variable length array of audio file regions.
	MSMPTE_TimeType unsafe.Pointer // The SMPTE timing scheme used in the file. See Core Audio’s   header file for the values used here. For more information, see  .
}

// AudioFileTypeAndFormatID - A specifier for the constant
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileTypeAndFormatID
type AudioFileTypeAndFormatID struct {
	MFileType AudioFileTypeID // A four-character code for the file type.
	MFormatID unsafe.Pointer // A four-character code for the format ID such as  ,  , and so forth. (See the   header file for declarations.)
}

// AudioFormatInfo - A structure that specifies an audio format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatInfo
type AudioFormatInfo struct {
	MASBD coreaudiotypes.AudioStreamBasicDescription // An   structure.
	MMagicCookie unsafe.Pointer // A pointer to the decompression information for the data described in the   parameter.
	MMagicCookieSize unsafe.Pointer // The size, in bytes, of the   parameter.
}

// AudioFramePacketTranslation - A structure that specifies frame and packet translations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFramePacketTranslation
type AudioFramePacketTranslation struct {
	MFrame unsafe.Pointer // A frame number.
	MFrameOffsetInPacket unsafe.Pointer // A frame offset in a packet.
	MPacket unsafe.Pointer // A packet number.
}

// AudioIndependentPacketTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioIndependentPacketTranslation
type AudioIndependentPacketTranslation struct {
	MIndependentlyDecodablePacket unsafe.Pointer
	MPacket unsafe.Pointer
}

// AudioOutputUnitMIDICallbacks
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitMIDICallbacks
type AudioOutputUnitMIDICallbacks struct {
	MIDIEventProc unsafe.Pointer
	MIDISysExProc unsafe.Pointer
	UserData unsafe.Pointer
}

// AudioOutputUnitStartAtTimeParams - A timestamp for scheduled starting of an I/O audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStartAtTimeParams
type AudioOutputUnitStartAtTimeParams struct {
	MFlags unsafe.Pointer
	MTimestamp coreaudiotypes.AudioTimeStamp
}

// AudioPacketDependencyInfoTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketDependencyInfoTranslation
type AudioPacketDependencyInfoTranslation struct {
	MIsIndependentlyDecodable unsafe.Pointer
	MNumberPrerollPackets unsafe.Pointer
	MPacket unsafe.Pointer
}

// AudioPacketRangeByteCountTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketRangeByteCountTranslation
type AudioPacketRangeByteCountTranslation struct {
	MByteCountUpperBound unsafe.Pointer
	MPacket unsafe.Pointer
	MPacketCount unsafe.Pointer
}

// AudioPacketRollDistanceTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketRollDistanceTranslation
type AudioPacketRollDistanceTranslation struct {
	MPacket unsafe.Pointer
	MRollDistance unsafe.Pointer
}

// AudioPanningInfo - Audio panning information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPanningInfo
type AudioPanningInfo struct {
	MCoordinateFlags unsafe.Pointer // For the available coordinate flags, see Channel Coordinate Flags.
	MCoordinates unsafe.Pointer // For the available coordinate index constants, see Channel Coordinate Index Constants.
	MGainScale unsafe.Pointer // A multiplier for audio panning values, typically representing a volume value in the range from 0 to 1. A value of 1 results in audio panning at unity gain. A value of 0 silences all channels.
	MOutputChannelMap avfaudio.AudioChannelLayout // The channel map used to determine channel volumes for the audio panning.
	MPanningMode AudioPanningMode // The mode to use for panning.
}

// AudioQueueBuffer - Defines an audio queue buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueBuffer
type AudioQueueBuffer struct {
	MAudioData unsafe.Pointer // The audio data owned the audio queue buffer. The buffer address cannot be changed.
	MAudioDataBytesCapacity unsafe.Pointer // The size of the audio queue buffer, in bytes. This size is set when a buffer is allocated and cannot be changed.
	MAudioDataByteSize unsafe.Pointer // The number of bytes of valid audio data in the audio queue buffer’s   field, initially set to  . Your callback must set this value for a playback audio queue; for recording, the recording audio queue sets the value.
	MPacketDescriptionCapacity unsafe.Pointer // The maximum number of packet descriptions that can be stored in the   field.
	MPacketDescriptionCount unsafe.Pointer // The number of valid packet descriptions in the buffer. You set this value when providing buffers for playback. The audio queue sets this value when returning buffers from a recording queue.
	MPacketDescriptions coreaudiotypes.AudioStreamPacketDescription // An array of   structures for the buffer.
	MUserData unsafe.Pointer // The custom data structure you specify, for use by your callback function, when creating a recording or playback audio queue.
}

// AudioQueueChannelAssignment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueChannelAssignment
type AudioQueueChannelAssignment struct {
	MChannelNumber unsafe.Pointer
	MDeviceUID StringRef
}

// AudioQueueLevelMeterState - Specifies the current level metering information for one channel of an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueLevelMeterState
type AudioQueueLevelMeterState struct {
	MAveragePower unsafe.Pointer // The audio channel’s average RMS power.
	MPeakPower unsafe.Pointer // The audio channel’s peak RMS power.
}

// AudioQueueParameterEvent - Specifies an audio queue parameter and associated value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueParameterEvent
type AudioQueueParameterEvent struct {
	MID AudioQueueParameterID // The parameter.
	MValue AudioQueueParameterValue // The value of the specified parameter.
}

// AudioUnitCocoaViewInfo - The name and number of custom Cocoa views for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitCocoaViewInfo
type AudioUnitCocoaViewInfo struct {
	MCocoaAUViewBundleLocation URLRef
	MCocoaAUViewClass StringRef
}

// AudioUnitConnection - An audio unit source-to-destination connection specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitConnection
type AudioUnitConnection struct {
	DestInputNumber unsafe.Pointer // The destination audio unit’s input element to be used in the connection.
	SourceAudioUnit AudioUnit // The audio unit that is serves as the source in the connection.
	SourceOutputNumber unsafe.Pointer // The source audio unit’s output element to be used in the connection.
}

// AudioUnitEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEvent
type AudioUnitEvent struct {
	MArgument unsafe.Pointer
	MParameter AudioUnitParameter
	MProperty AudioUnitProperty
	MEventType AudioUnitEventType
}

// AudioUnitExternalBuffer - Allows an audio unit host application to tell an audio unit to use a specified buffer for its input callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExternalBuffer
type AudioUnitExternalBuffer struct {
	Buffer unsafe.Pointer
	Size unsafe.Pointer
}

// AudioUnitFrequencyResponseBin - An audio unit’s audio level at a particular frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitFrequencyResponseBin
type AudioUnitFrequencyResponseBin struct {
	MFrequency unsafe.Pointer
	MMagnitude unsafe.Pointer
}

// AudioUnitMeterClipping - Audio clipping that has occurred in a mixer unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitMeterClipping
type AudioUnitMeterClipping struct {
	PeakValueSinceLastCall unsafe.Pointer // The maximum value seen on the channel since the last time the property was retrieved.
	SawInfinity unsafe.Pointer //  if there was an infinite value on this channel since the last time the property was retrieved.
	SawNotANumber unsafe.Pointer //  if there was a floating point “not a number” value on this channel since the last time the property was retrieved.
}

// AudioUnitMIDIControlMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitMIDIControlMapping
type AudioUnitMIDIControlMapping struct {
	Element AudioUnitElement
	MidiControl unsafe.Pointer
	MidiNRPN unsafe.Pointer
	Parameter AudioUnitParameterID
	Scope unsafe.Pointer
}

// AudioUnitNodeConnection - A connection between two node objects in an audio processing graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitNodeConnection
type AudioUnitNodeConnection struct {
	DestInputNumber unsafe.Pointer
	DestNode Node
	SourceNode Node
	SourceOutputNumber unsafe.Pointer
}

// AudioUnitOtherPluginDesc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitOtherPluginDesc
type AudioUnitOtherPluginDesc struct {
	Format unsafe.Pointer
	Plugin coreaudiotypes.AudioClassDescription
}

// AudioUnitParameter - An adjustable audio unit attribute such as volume, pitch, or filter cutoff frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameter
type AudioUnitParameter struct {
	MAudioUnit AudioUnit // The audio unit instance that the parameter applies to.
	MElement AudioUnitElement // The audio unit element for the parameter.
	MParameterID AudioUnitParameterID // The audio unit parameter identifier.
	MScope AudioUnitScope // The audio unit scope for the parameter.
}

// AudioUnitParameterEvent - A scheduled change to an audio unit parameter’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterEvent
type AudioUnitParameterEvent struct {
	Element AudioUnitElement // The element for this parameter event.
	EventType ParameterEventType // The type for this parameter event.
	EventValues unsafe.Pointer // The values for this parameter event.
	Immediate unsafe.Pointer
	BufferOffset unsafe.Pointer
	Value AudioUnitParameterValue
	Ramp unsafe.Pointer
	DurationInFrames unsafe.Pointer
	EndValue AudioUnitParameterValue
	StartBufferOffset unsafe.Pointer
	StartValue AudioUnitParameterValue
	Parameter AudioUnitParameterID // An identifier for this parameter event.
	Scope AudioUnitScope // The scope for this parameter event.
}

// AudioUnitParameterHistoryInfo - The suggested update rate and history duration for parameters which have the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterHistoryInfo
type AudioUnitParameterHistoryInfo struct {
	HistoryDurationInSeconds unsafe.Pointer
	UpdatesPerSecond unsafe.Pointer
}

// AudioUnitParameterInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterInfo
type AudioUnitParameterInfo struct {
	CfNameString StringRef // Only valid if   is set.
	ClumpID unsafe.Pointer // Only valid if   is set.
	DefaultValue AudioUnitParameterValue
	Flags AudioUnitParameterOptions // The host should check for this flag and, if present, release the parameter name when it is finished with it.
	MaxValue AudioUnitParameterValue
	MinValue AudioUnitParameterValue
	Name unsafe.Pointer // Must be set to  .
	Unit AudioUnitParameterUnit // If the   field contains a value not in the   enumeration, then assume the unit type is  .
	UnitName StringRef // If   is set, this field must contain a valid   object. Only valid if   is set.
}

// AudioUnitParameterNameInfo - A short version of the name for an audio unit parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterNameInfo
type AudioUnitParameterNameInfo struct {
	InDesiredLength unsafe.Pointer // When setting an audio unit property that uses this data structure for its value, the maximum length that you are specifying for the audio unit parameter name.
	InID AudioUnitParameterID // The identifier for the audio unit parameter.
	OutName StringRef // When getting an audio unit property that uses this data structure for its value, the short version of the parameter name provided by the audio unit. The host application then owns the string and is responsible for releasing it.
}

// AudioUnitParameterStringFromValue - A string representation of a parameter’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterStringFromValue
type AudioUnitParameterStringFromValue struct {
	InParamID AudioUnitParameterID
	InValue AudioUnitParameterValue
	OutString StringRef
}

// AudioUnitParameterValueFromString - A parameter’s value based on a string representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueFromString
type AudioUnitParameterValueFromString struct {
	InParamID AudioUnitParameterID
	InString StringRef
	OutValue AudioUnitParameterValue
}

// AudioUnitParameterValueName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueName
type AudioUnitParameterValueName struct {
	InParamID AudioUnitParameterID
	InValue unsafe.Pointer
	OutName StringRef
}

// AudioUnitParameterValueTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueTranslation
type AudioUnitParameterValueTranslation struct {
	AuParamID AudioUnitParameterID
	AuValue AudioUnitParameterValue
	OtherDesc AudioUnitOtherPluginDesc
	OtherParamID unsafe.Pointer
	OtherValue unsafe.Pointer
}

// AudioUnitPresetMAS_SettingData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPresetMAS_SettingData
type AudioUnitPresetMAS_SettingData struct {
	Data unsafe.Pointer
	DataLen unsafe.Pointer
	IsStockSetting unsafe.Pointer
	SettingID unsafe.Pointer
}

// AudioUnitPresetMAS_Settings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPresetMAS_Settings
type AudioUnitPresetMAS_Settings struct {
	EffectID unsafe.Pointer
	ManufacturerID unsafe.Pointer
	NumberOfSettings unsafe.Pointer
	Settings AudioUnitPresetMAS_SettingData
	SettingsVersion unsafe.Pointer
	VariantID unsafe.Pointer
}

// AudioUnitProperty - A key-value pair that declares an attribute or behavior for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProperty
type AudioUnitProperty struct {
	MAudioUnit AudioUnit // The audio unit instance that the parameter applies to.
	MElement AudioUnitElement // The audio unit element for the property.
	MPropertyID AudioUnitPropertyID // The audio unit property identifier.
	MScope AudioUnitScope // The audio unit scope for the property.
}

// AudioUnitRenderContext - A structure that contains thread context information for a real-time rendering operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderContext
type AudioUnitRenderContext struct {
	Reserved uint32 // System-specific information.
	Workgroup unsafe.Pointer // The workgroup that manages the rendering threads of the audio unit.
}

// AUDistanceAttenuationData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUDistanceAttenuationData
type AUDistanceAttenuationData struct {
	InNumberOfPairs unsafe.Pointer
	Pairs unsafe.Pointer
	InDistance unsafe.Pointer
	OutGain unsafe.Pointer
}

// AUHostIdentifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUHostIdentifier
type AUHostIdentifier struct {
	HostName StringRef
	HostVersion NumVersion
}

// AUHostVersionIdentifier - The name and version of an audio unit’s host application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUHostVersionIdentifier
type AUHostVersionIdentifier struct {
	HostName StringRef
	HostVersion unsafe.Pointer
}

// AUInputSamplesInOutputCallbackStruct - The callback function and custom data for providing input-to-output sample mapping for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUInputSamplesInOutputCallbackStruct
type AUInputSamplesInOutputCallbackStruct struct {
	InputToOutputCallback InputSamplesInOutputCallback // The callback function that provides input-to-output sample mapping for an audio unit.
	UserData unsafe.Pointer // Custom data for input-to-output sample mapping for an audio unit.
}

// AUMIDIEvent - A structure that describes a scheduled MIDI event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIEvent
type AUMIDIEvent struct {
	Cable uint8 // The virtual cable number.
	Data uint8 // The bytes of the MIDI event. Running status is not used.
	EventSampleTime EventSampleTime // The sample time at which the event is scheduled to occur.
	EventType RenderEventType // The type of render event. Must be   or  .
	Length uint16 // The number of valid MIDI bytes in the data field. For most MIDI events this value is usually  ,  , or  , but it can be longer for system-exclusive events.
	Next unsafe.Pointer // The next event in a linked list of events.
	Reserved uint8 // Reserved field. Must be  .
}

// AUMIDIEventList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIEventList
type AUMIDIEventList struct {
	Cable uint8
	EventList MIDIEventList
	EventSampleTime EventSampleTime
	EventType RenderEventType
	Next unsafe.Pointer
	Reserved uint8
}

// AUMIDIOutputCallbackStruct - The callback function and custom data for an audio unit that provides MIDI output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIOutputCallbackStruct
type AUMIDIOutputCallbackStruct struct {
	MidiOutputCallback MIDIOutputCallback // The callback function for an audio unit that provides MIDI output.
	UserData unsafe.Pointer // Custom data for an audio unit that provides MIDI output.
}

// AUNodeInteraction - Describes the interaction between two node objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNodeInteraction
type AUNodeInteraction struct {
	NodeInteraction unsafe.Pointer // A union providing information about a node interaction.
	Connection NodeConnection
	InputCallback NodeRenderCallback
	NodeInteractionType unsafe.Pointer // The interaction type.
}

// AUNodeRenderCallback - A callback used to provide input to an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNodeRenderCallback
type AUNodeRenderCallback struct {
	Cback RenderCallbackStruct
	DestInputNumber AudioUnitElement
	DestNode Node
}

// AUNumVersion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNumVersion
type AUNumVersion struct {
	MajorRev unsafe.Pointer
	MinorAndBugRev unsafe.Pointer
	NonRelRev unsafe.Pointer
	Stage unsafe.Pointer
}

// AUParameterAutomationEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationEvent
type AUParameterAutomationEvent struct {
	Address ParameterAddress
	EventType ParameterAutomationEventType
	HostTime uint64
	Reserved uint64
	Value Value
}

// AUParameterEvent - A structure that describes a scheduled parameter event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterEvent
type AUParameterEvent struct {
	EventSampleTime EventSampleTime // The sample time at which the event is scheduled to occur.
	EventType RenderEventType // The type of render event. Must be   or  .
	Next unsafe.Pointer // The next event in a linked list of events.
	ParameterAddress ParameterAddress // The parameter to change.
	RampDurationSampleFrames AudioFrameCount // The ramp duration, in sample frames. Must be   for a non-ramped event; otherwise, must be greater than   for a ramped event.
	Reserved uint8 // Reserved field. Must be  .
	Value Value // For a non-ramped event, this is the new parameter value. For a ramped event, this is the parameter value at the end of the ramp.
}

// AUParameterMIDIMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterMIDIMapping
type AUParameterMIDIMapping struct {
	MData1 unsafe.Pointer
	MElement AudioUnitElement
	MFlags ParameterMIDIMappingFlags
	MParameterID AudioUnitParameterID
	MScope AudioUnitScope
	MStatus unsafe.Pointer
	MSubRangeMax AudioUnitParameterValue
	MSubRangeMin AudioUnitParameterValue
	Reserved1 unsafe.Pointer
	Reserved2 unsafe.Pointer
	Reserved3 unsafe.Pointer
}

// AUPreset - Used to set factory presets for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUPreset
type AUPreset struct {
	PresetName StringRef // If a factory preset, the name of the specified factory preset.
	PresetNumber unsafe.Pointer // If less than  , then the preset is a user preset. If greater than or equal to  , then this field is used to select a factory preset.
}

// AUPresetEvent - Describes an audio unit preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUPresetEvent
type AUPresetEvent struct {
	Element AudioUnitElement
	Preset PropertyListRef
	Scope AudioUnitScope
}

// AURecordedParameterEvent - An event recording the changing of a parameter at a particular host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURecordedParameterEvent
type AURecordedParameterEvent struct {
	Address ParameterAddress // The address of the parameter whose value changed.
	HostTime uint64 // The host time at which the event occurred.
	Value Value // The value of the parameter at the given time.
}

// AURenderCallbackStruct - Used for registering an input callback function with an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderCallbackStruct
type AURenderCallbackStruct struct {
	InputProc RenderCallback
	InputProcRefCon unsafe.Pointer
}

// AURenderEventHeader - The common header for a render event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderEventHeader
type AURenderEventHeader struct {
	EventSampleTime EventSampleTime // The sample time at which the event is scheduled to occur.
	EventType RenderEventType // The type of the event.
	Next unsafe.Pointer // The next event in a linked list of events.
	Reserved uint8 // Reserved field. Must be  .
}

// AUSamplerBankPresetData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSamplerBankPresetData
type AUSamplerBankPresetData struct {
	BankLSB unsafe.Pointer
	BankMSB unsafe.Pointer
	BankURL URLRef
	PresetID unsafe.Pointer
	Reserved unsafe.Pointer
}

// AUSamplerInstrumentData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSamplerInstrumentData
type AUSamplerInstrumentData struct {
	BankLSB unsafe.Pointer
	BankMSB unsafe.Pointer
	FileURL URLRef
	InstrumentType unsafe.Pointer
	PresetID unsafe.Pointer
}

// AUVoiceIOOtherAudioDuckingConfiguration - A structure that you use to configure ducking of other non-voice audio in a voice chat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOOtherAudioDuckingConfiguration
type AUVoiceIOOtherAudioDuckingConfiguration struct {
	MDuckingLevel VoiceIOOtherAudioDuckingLevel // The ducking level of other non-voice audio.
	MEnableAdvancedDucking unsafe.Pointer // A Boolean value that specifies whether to enable advanced ducking.
}

// CABarBeatTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CABarBeatTime
type CABarBeatTime struct {
	Bar unsafe.Pointer
	Beat unsafe.Pointer
	Reserved unsafe.Pointer
	Subbeat unsafe.Pointer
	SubbeatDivisor unsafe.Pointer
}

// CAClockTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTime
type CAClockTime struct {
	Format ClockTimeFormat
	Reserved unsafe.Pointer
	Time unsafe.Pointer
	Beats ClockBeats
	HostTime unsafe.Pointer
	Samples ClockSamples
	Seconds ClockSeconds
	Smpte objectivec.IObject
}

// CAF_SMPTE_Time
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAF_SMPTE_Time
type CAF_SMPTE_Time struct {
	MFrames unsafe.Pointer
	MHours unsafe.Pointer
	MMinutes unsafe.Pointer
	MSeconds unsafe.Pointer
	MSubFrameSampleOffset unsafe.Pointer
}

// CAF_UUID_ChunkHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAF_UUID_ChunkHeader
type CAF_UUID_ChunkHeader struct {
	MHeader FChunkHeader
	MUUID unsafe.Pointer
}

// CAFAudioDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFAudioDescription
type CAFAudioDescription struct {
	MBitsPerChannel unsafe.Pointer
	MBytesPerPacket unsafe.Pointer
	MChannelsPerFrame unsafe.Pointer
	MFormatFlags FFormatFlags
	MFormatID unsafe.Pointer
	MFramesPerPacket unsafe.Pointer
	MSampleRate unsafe.Pointer
}

// CAFAudioFormatListItem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFAudioFormatListItem
type CAFAudioFormatListItem struct {
	MChannelLayoutTag unsafe.Pointer
	MFormat FAudioDescription
}

// CAFChunkHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFChunkHeader
type CAFChunkHeader struct {
	MChunkSize unsafe.Pointer
	MChunkType unsafe.Pointer
}

// CAFDataChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFDataChunk
type CAFDataChunk struct {
	MData unsafe.Pointer
	MEditCount unsafe.Pointer
}

// CAFFileHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFFileHeader
type CAFFileHeader struct {
	MFileFlags unsafe.Pointer
	MFileType unsafe.Pointer
	MFileVersion unsafe.Pointer
}

// CAFInfoStrings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFInfoStrings
type CAFInfoStrings struct {
	MNumEntries unsafe.Pointer
}

// CAFInstrumentChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFInstrumentChunk
type CAFInstrumentChunk struct {
	MBaseNote unsafe.Pointer
	MdBGain unsafe.Pointer
	MInstrumentID unsafe.Pointer
	MMIDIHighNote unsafe.Pointer
	MMIDIHighVelocity unsafe.Pointer
	MMIDILowNote unsafe.Pointer
	MMIDILowVelocity unsafe.Pointer
	MReleaseRegionID unsafe.Pointer
	MStartRegionID unsafe.Pointer
	MSustainRegionID unsafe.Pointer
}

// CAFMarker
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFMarker
type CAFMarker struct {
	MChannel unsafe.Pointer
	MFramePosition unsafe.Pointer
	MMarkerID unsafe.Pointer
	MSMPTETime F_SMPTE_Time
	MType unsafe.Pointer
}

// CAFMarkerChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFMarkerChunk
type CAFMarkerChunk struct {
	MMarkers FMarker
	MNumberMarkers unsafe.Pointer
	MSMPTE_TimeType unsafe.Pointer
}

// CAFOverviewChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFOverviewChunk
type CAFOverviewChunk struct {
	MData FOverviewSample
	MEditCount unsafe.Pointer
	MNumFramesPerOVWSample unsafe.Pointer
}

// CAFOverviewSample
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFOverviewSample
type CAFOverviewSample struct {
	MMaxValue unsafe.Pointer
	MMinValue unsafe.Pointer
}

// CAFPacketTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPacketTableHeader
type CAFPacketTableHeader struct {
	MNumberPackets unsafe.Pointer
	MNumberValidFrames unsafe.Pointer
	MPacketDescriptions unsafe.Pointer
	MPrimingFrames unsafe.Pointer
	MRemainderFrames unsafe.Pointer
}

// CAFPeakChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPeakChunk
type CAFPeakChunk struct {
	MEditCount unsafe.Pointer
	MPeaks FPositionPeak
}

// CAFPositionPeak
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPositionPeak
type CAFPositionPeak struct {
	MFrameNumber unsafe.Pointer
	MValue unsafe.Pointer
}

// CAFRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegion
type CAFRegion struct {
	MFlags FRegionFlags
	MMarkers FMarker
	MNumberMarkers unsafe.Pointer
	MRegionID unsafe.Pointer
}

// CAFRegionChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegionChunk
type CAFRegionChunk struct {
	MNumberRegions unsafe.Pointer
	MRegions FRegion
	MSMPTE_TimeType unsafe.Pointer
}

// CAFStringID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFStringID
type CAFStringID struct {
	MStringID unsafe.Pointer
	MStringStartByteOffset unsafe.Pointer
}

// CAFStrings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFStrings
type CAFStrings struct {
	MNumEntries unsafe.Pointer
	MStringsIDs FStringID
}

// CAFUMIDChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFUMIDChunk
type CAFUMIDChunk struct {
	MBytes unsafe.Pointer
}

// CAMeterTrackEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAMeterTrackEntry
type CAMeterTrackEntry struct {
	Beats ClockBeats
	MeterDenom unsafe.Pointer
	MeterNumer unsafe.Pointer
}

// CATempoMapEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CATempoMapEntry
type CATempoMapEntry struct {
	Beats ClockBeats
	TempoBPM ClockTempo
}

// ExtendedAudioFormatInfo - A specifier for the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedAudioFormatInfo
type ExtendedAudioFormatInfo struct {
	MASBD coreaudiotypes.AudioStreamBasicDescription // A format specification for an audio stream.
	MClassDescription coreaudiotypes.AudioClassDescription // A structure that describes an audio codec.
	MMagicCookie unsafe.Pointer // Decompression information for the audio data format specified in the   field.
	MMagicCookieSize unsafe.Pointer // The size, in bytes, of the   field.
}

// ExtendedControlEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedControlEvent
type ExtendedControlEvent struct {
	ControlID AudioUnitParameterID
	GroupID MusicDeviceGroupID
	Value AudioUnitParameterValue
}

// ExtendedNoteOnEvent - Describes a note-on event with extended parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedNoteOnEvent
type ExtendedNoteOnEvent struct {
	Duration unsafe.Pointer
	ExtendedParams MusicDeviceNoteParams
	GroupID MusicDeviceGroupID
	InstrumentID MusicDeviceInstrumentID
}

// ExtendedTempoEvent - Describes a music track tempo in beats-per-minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedTempoEvent
type ExtendedTempoEvent struct {
	Bpm unsafe.Pointer // The number of beats-per-minute.
}

// HostCallbackInfo - The time- and transport-related callback functions for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallbackInfo
type HostCallbackInfo struct {
	BeatAndTempoProc HostCallback_GetBeatAndTempo // Your callback function that provides beat and tempo information to an audio unit. May be  .
	HostUserData unsafe.Pointer // Custom data specified by your application. May be  .
	MusicalTimeLocationProc HostCallback_GetMusicalTimeLocation // Your callback function that provides musical timeline information to an audio unit. May be  .
	TransportStateProc HostCallback_GetTransportState // Your callback function that provides audio transport state information ( ,  , and so on) to an audio unit. May be  .
	TransportStateProc2 HostCallback_GetTransportState2
}

// MIDIChannelMessage - Describes a MIDI channel message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIChannelMessage
type MIDIChannelMessage struct {
	Data1 unsafe.Pointer
	Data2 unsafe.Pointer
	Reserved unsafe.Pointer
	Status unsafe.Pointer // Data specific to the channel message.
}

// MIDIMetaEvent - Describes a MIDI metaevent such as lyric text, time signature, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIMetaEvent
type MIDIMetaEvent struct {
	Data unsafe.Pointer
	DataLength unsafe.Pointer
	MetaEventType unsafe.Pointer // An integer that designates one of the types of MIDI metaevents.
	Unused1 unsafe.Pointer
	Unused2 unsafe.Pointer
	Unused3 unsafe.Pointer
}

// MIDINoteMessage - Describes a MIDI note.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDINoteMessage
type MIDINoteMessage struct {
	Channel unsafe.Pointer // The MIDI channel to play the note on.
	Duration unsafe.Pointer // The duration for the note.
	Note unsafe.Pointer // The note to play.
	ReleaseVelocity unsafe.Pointer // The key-release velocity for the note. Use 0 if you don’t want to specify a particular value.
	Velocity unsafe.Pointer // The key-press velocity for the note.
}

// MIDIRawData - Describes a MIDI system-exclusive (SysEx) message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIRawData
type MIDIRawData struct {
	Data unsafe.Pointer
	Length unsafe.Pointer
}

// MixerDistanceParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MixerDistanceParams
type MixerDistanceParams struct {
	MMaxAttenuation unsafe.Pointer
	MMaxDistance unsafe.Pointer
	MReferenceDistance unsafe.Pointer
}

// MusicDeviceNoteParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceNoteParams
type MusicDeviceNoteParams struct {
	ArgCount unsafe.Pointer
	MControls NoteParamsControlValue
	MPitch unsafe.Pointer
	MVelocity unsafe.Pointer
}

// MusicDeviceStdNoteParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStdNoteParams
type MusicDeviceStdNoteParams struct {
	ArgCount unsafe.Pointer
	MPitch unsafe.Pointer
	MVelocity unsafe.Pointer
}

// MusicEventUserData - Describes a user-defined event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventUserData
type MusicEventUserData struct {
	Data unsafe.Pointer // User-defined data.
	Length unsafe.Pointer // The size, in bytes, of the user data.
}

// MusicTrackLoopInfo - Supports control of the looping behavior of a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackLoopInfo
type MusicTrackLoopInfo struct {
	LoopDuration MusicTimeStamp // The point in a music track, measured in beats from the end of the music track, at which to begin playback during looped playback.
	NumberOfLoops unsafe.Pointer // The number of times to play the designated portion of the music track.
}

// NoteParamsControlValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NoteParamsControlValue
type NoteParamsControlValue struct {
	MID AudioUnitParameterID
	MValue AudioUnitParameterValue
}

// ParameterEvent - Describes an audio unit parameter automation event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ParameterEvent
type ParameterEvent struct {
	Element AudioUnitElement
	ParameterID AudioUnitParameterID
	Scope AudioUnitScope
	Value AudioUnitParameterValue
}

// ScheduledAudioFileRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioFileRegion
type ScheduledAudioFileRegion struct {
	MAudioFile unsafe.Pointer // Must be a valid and already-open audio file object (of type  ), as declared in  .
	MCompletionProc ScheduledAudioFileRegionCompletionProc // may be 
	MCompletionProcUserData unsafe.Pointer
	MFramesToPlay unsafe.Pointer // The number of frames to play.
	MLoopCount unsafe.Pointer //  = do not loop
	MStartFrame unsafe.Pointer // The frame offset into the file.
	MTimeStamp coreaudiotypes.AudioTimeStamp
}

// ScheduledAudioSlice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioSlice
type ScheduledAudioSlice struct {
	MBufferList coreaudiotypes.AudioBufferList
	MCompletionProc ScheduledAudioSliceCompletionProc
	MCompletionProcUserData unsafe.Pointer
	MFlags ScheduledAudioSliceFlags
	MNumberFrames unsafe.Pointer
	MReserved unsafe.Pointer
	MReserved2 unsafe.Pointer
	MTimeStamp coreaudiotypes.AudioTimeStamp
}





