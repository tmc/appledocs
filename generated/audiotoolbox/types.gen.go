// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox
import (
	"unsafe"
)


// C struct types
// AUChannelInfo - The audio input and output channel capabilities for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUChannelInfo
type AUChannelInfo struct {
	InChannels unsafe.Pointer // The number of input channels.
	OutChannels unsafe.Pointer // The number of output channels.
}// AUDependentParameter - An audio unit parameter whose value can change in response to a change in its parent metaparameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUDependentParameter
type AUDependentParameter struct {
	MParameterID AudioUnitParameterID
	MScope AudioUnitScope
}// AUDistanceAttenuationData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUDistanceAttenuationData
type AUDistanceAttenuationData struct {
	InNumberOfPairs unsafe.Pointer
	Pairs unsafe.Pointer
}// AUHostIdentifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUHostIdentifier
type AUHostIdentifier struct {
	HostName unsafe.Pointer
	HostVersion unsafe.Pointer
}// AUHostVersionIdentifier - The name and version of an audio unit’s host application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUHostVersionIdentifier
type AUHostVersionIdentifier struct {
	HostVersion unsafe.Pointer
}// AUInputSamplesInOutputCallbackStruct - The callback function and custom data for providing input-to-output sample mapping for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUInputSamplesInOutputCallbackStruct
type AUInputSamplesInOutputCallbackStruct struct {
	UserData unsafe.Pointer // Custom data for input-to-output sample mapping for an audio unit.
}// AUMIDIEvent - A structure that describes a scheduled MIDI event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIEvent
type AUMIDIEvent struct {
	Next unsafe.Pointer // The next event in a linked list of events.
}// AUMIDIEventList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIEventList
type AUMIDIEventList struct {
	Cable uint8
	EventList unsafe.Pointer
	EventSampleTime EventSampleTime
	EventType unsafe.Pointer
	Next unsafe.Pointer
	Reserved uint8
}// AUMIDIOutputCallbackStruct - The callback function and custom data for an audio unit that provides MIDI output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIOutputCallbackStruct
type AUMIDIOutputCallbackStruct struct {
	UserData unsafe.Pointer // Custom data for an audio unit that provides MIDI output.
}// AUNodeInteraction - Describes the interaction between two node objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNodeInteraction
type AUNodeInteraction struct {
}// AUNumVersion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNumVersion
type AUNumVersion struct {
	MajorRev unsafe.Pointer
}// AUParameterAutomationEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationEvent
type AUParameterAutomationEvent struct {
	Address ParameterAddress
}// AUParameterEvent - A structure that describes a scheduled parameter event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterEvent
type AUParameterEvent struct {
	ParameterAddress ParameterAddress // The parameter to change.
	RampDurationSampleFrames AudioFrameCount // The ramp duration, in sample frames. Must be   for a non-ramped event; otherwise, must be greater than   for a ramped event.
}// AUParameterMIDIMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterMIDIMapping
type AUParameterMIDIMapping struct {
	MScope AudioUnitScope
	MSubRangeMax AudioUnitParameterValue
	MSubRangeMin AudioUnitParameterValue
	Reserved2 unsafe.Pointer
}// AUPreset - Used to set factory presets for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUPreset
type AUPreset struct {
}// AUPresetEvent - Describes an audio unit preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUPresetEvent
type AUPresetEvent struct {
	Element AudioUnitElement
}// AURecordedParameterEvent - An event recording the changing of a parameter at a particular host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURecordedParameterEvent
type AURecordedParameterEvent struct {
	Address ParameterAddress // The address of the parameter whose value changed.
	Value Value // The value of the parameter at the given time.
}// AURenderCallbackStruct - Used for registering an input callback function with an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderCallbackStruct
type AURenderCallbackStruct struct {
	InputProcRefCon unsafe.Pointer
}// AURenderEventHeader - The common header for a render event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderEventHeader
type AURenderEventHeader struct {
	Reserved uint8 // Reserved field. Must be  .
}// AUSamplerBankPresetData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSamplerBankPresetData
type AUSamplerBankPresetData struct {
	BankLSB unsafe.Pointer
	BankURL unsafe.Pointer
}// AUSamplerInstrumentData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSamplerInstrumentData
type AUSamplerInstrumentData struct {
	FileURL unsafe.Pointer
	PresetID unsafe.Pointer
}// AUVoiceIOOtherAudioDuckingConfiguration - A structure that you use to configure ducking of other non-voice audio in a voice chat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOOtherAudioDuckingConfiguration
type AUVoiceIOOtherAudioDuckingConfiguration struct {
	MDuckingLevel unsafe.Pointer // The ducking level of other non-voice audio.
	MEnableAdvancedDucking unsafe.Pointer // A Boolean value that specifies whether to enable advanced ducking.
}// AudioBalanceFade - Describes audio left/right balance and front/back fade values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBalanceFade
type AudioBalanceFade struct {
	MChannelLayout unsafe.Pointer // The size, in bytes, of the   parameter.
}// AudioBytePacketTranslation - A data structure used by the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBytePacketTranslation
type AudioBytePacketTranslation struct {
	MByteOffsetInPacket unsafe.Pointer // A byte offset in a packet.
}// AudioCodecMagicCookieInfo - A structure holding magic cookie information needed by some codecs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecMagicCookieInfo
type AudioCodecMagicCookieInfo struct {
	MMagicCookie unsafe.Pointer // Generic constant pointer to the magic cookie.
	MMagicCookieSize unsafe.Pointer // The size of the magic cookie.
}// AudioCodecPrimeInfo - A structure specifying the number of leading and trailing empty frames to be inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecPrimeInfo
type AudioCodecPrimeInfo struct {
}// AudioComponentDescription - Identifying information for an audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentDescription
type AudioComponentDescription struct {
	ComponentFlags unsafe.Pointer // Set this value to zero.
	ComponentFlagsMask unsafe.Pointer // Set this value to zero.
	ComponentManufacturer unsafe.Pointer // The unique vendor identifier, registered with Apple, for the audio component.
	ComponentSubType unsafe.Pointer // A 4-byte code that you can use to indicate the purpose of a component. For example, you could use   or   as a mnemonic indication that an audio unit is a low-pass filter.
	ComponentType unsafe.Pointer // A unique 4-byte code identifying the interface for the component.
}// AudioComponentPlugInInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentPlugInInterface
type AudioComponentPlugInInterface struct {
	Close unsafe.Pointer
	Lookup unsafe.Pointer
	Open unsafe.Pointer
	Reserved unsafe.Pointer
}// AudioConverterPrimeInfo - Specifies priming information for an audio converter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterPrimeInfo
type AudioConverterPrimeInfo struct {
}// AudioFileFDFTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFDFTable
type AudioFileFDFTable struct {
	MGetUserDataSizeFDF GetUserDataSizeFDF
}// AudioFileFDFTableExtended
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFDFTableExtended
type AudioFileFDFTableExtended struct {
	MReadPacketsFDF ReadPacketsFDF
}// AudioFileMarker - Annotates a position in an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileMarker
type AudioFileMarker struct {
	MReserved unsafe.Pointer // A reserved field. Set to  .
	MSMPTETime unsafe.Pointer // The SMPTE time for this marker.
}// AudioFileMarkerList - A list of markers associated with an audio file, including their SMPTE time type, the number of markers, and the markers themselves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileMarkerList
type AudioFileMarkerList struct {
	MMarkers unsafe.Pointer // An array of   elements, each of which is an audio file marker.
	MNumberMarkers unsafe.Pointer // The number of markers in the list.
	MSMPTE_TimeType unsafe.Pointer // The SMPTE time type of the whole list of markers in an audio file.
}// AudioFilePacketTableInfo - Contains information about the number of valid frames in a file and where they begin and end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFilePacketTableInfo
type AudioFilePacketTableInfo struct {
	MPrimingFrames unsafe.Pointer // The number of invalid frames at the beginning of the file.
}// AudioFileRegion - An audio file region specifies a segment of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegion
type AudioFileRegion struct {
	MFlags unsafe.Pointer // Audio File Services region flags.
	MMarkers unsafe.Pointer // An array of   elements describing where the data in the region starts.
	MName unsafe.Pointer // The name of the region.
	MNumberMarkers unsafe.Pointer // The number of markers in the array specified in the    parameter.
	MRegionID unsafe.Pointer // A unique ID associated with the audio file region.
}// AudioFileRegionList - A list of the audio file regions in a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegionList
type AudioFileRegionList struct {
	MRegions unsafe.Pointer // A variable length array of audio file regions.
}// AudioFileTypeAndFormatID - A specifier for the constant
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileTypeAndFormatID
type AudioFileTypeAndFormatID struct {
	MFileType AudioFileTypeID // A four-character code for the file type.
	MFormatID unsafe.Pointer // A four-character code for the format ID such as  ,  , and so forth. (See the   header file for declarations.)
}// AudioFile_SMPTE_Time - A data structure for describing SMPTE (Society of Motion Picture and Television Engineers) time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_SMPTE_Time
type AudioFile_SMPTE_Time struct {
	MFrames unsafe.Pointer // The frames.
	MHours unsafe.Pointer // The hours.
}// AudioFormatInfo - A structure that specifies an audio format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatInfo
type AudioFormatInfo struct {
	MASBD unsafe.Pointer // An   structure.
	MMagicCookie unsafe.Pointer // A pointer to the decompression information for the data described in the   parameter.
	MMagicCookieSize unsafe.Pointer // The size, in bytes, of the   parameter.
}// AudioFramePacketTranslation - A structure that specifies frame and packet translations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFramePacketTranslation
type AudioFramePacketTranslation struct {
	MFrame unsafe.Pointer // A frame number.
	MFrameOffsetInPacket unsafe.Pointer // A frame offset in a packet.
	MPacket unsafe.Pointer // A packet number.
}// AudioIndependentPacketTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioIndependentPacketTranslation
type AudioIndependentPacketTranslation struct {
	MIndependentlyDecodablePacket unsafe.Pointer
	MPacket unsafe.Pointer
}// AudioOutputUnitMIDICallbacks
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitMIDICallbacks
type AudioOutputUnitMIDICallbacks struct {
	MIDIEventProc unsafe.Pointer
	MIDISysExProc unsafe.Pointer
	UserData unsafe.Pointer
}// AudioOutputUnitStartAtTimeParams - A timestamp for scheduled starting of an I/O audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStartAtTimeParams
type AudioOutputUnitStartAtTimeParams struct {
	MFlags unsafe.Pointer
	MTimestamp unsafe.Pointer
}// AudioPacketDependencyInfoTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketDependencyInfoTranslation
type AudioPacketDependencyInfoTranslation struct {
}// AudioPacketRangeByteCountTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketRangeByteCountTranslation
type AudioPacketRangeByteCountTranslation struct {
	MPacketCount unsafe.Pointer
}// AudioPacketRollDistanceTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketRollDistanceTranslation
type AudioPacketRollDistanceTranslation struct {
	MPacket unsafe.Pointer
	MRollDistance unsafe.Pointer
}// AudioPanningInfo - Audio panning information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPanningInfo
type AudioPanningInfo struct {
	MCoordinateFlags unsafe.Pointer // For the available coordinate flags, see Channel Coordinate Flags.
	MCoordinates unsafe.Pointer // For the available coordinate index constants, see Channel Coordinate Index Constants.
	MGainScale unsafe.Pointer // A multiplier for audio panning values, typically representing a volume value in the range from 0 to 1. A value of 1 results in audio panning at unity gain. A value of 0 silences all channels.
	MOutputChannelMap unsafe.Pointer // The channel map used to determine channel volumes for the audio panning.
	MPanningMode unsafe.Pointer // The mode to use for panning.
}// AudioQueueBuffer - Defines an audio queue buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueBuffer
type AudioQueueBuffer struct {
	MPacketDescriptions unsafe.Pointer // An array of   structures for the buffer.
}// AudioQueueChannelAssignment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueChannelAssignment
type AudioQueueChannelAssignment struct {
	MChannelNumber unsafe.Pointer
}// AudioQueueLevelMeterState - Specifies the current level metering information for one channel of an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueLevelMeterState
type AudioQueueLevelMeterState struct {
}// AudioQueueParameterEvent - Specifies an audio queue parameter and associated value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueParameterEvent
type AudioQueueParameterEvent struct {
	MID AudioQueueParameterID // The parameter.
	MValue AudioQueueParameterValue // The value of the specified parameter.
}// AudioUnitCocoaViewInfo - The name and number of custom Cocoa views for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitCocoaViewInfo
type AudioUnitCocoaViewInfo struct {
	MCocoaAUViewBundleLocation unsafe.Pointer
	MCocoaAUViewClass unsafe.Pointer
}// AudioUnitConnection - An audio unit source-to-destination connection specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitConnection
type AudioUnitConnection struct {
	DestInputNumber unsafe.Pointer // The destination audio unit’s input element to be used in the connection.
	SourceAudioUnit AudioUnit // The audio unit that is serves as the source in the connection.
	SourceOutputNumber unsafe.Pointer // The source audio unit’s output element to be used in the connection.
}// AudioUnitEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEvent
type AudioUnitEvent struct {
	MArgument unsafe.Pointer
	MEventType unsafe.Pointer
}// AudioUnitExternalBuffer - Allows an audio unit host application to tell an audio unit to use a specified buffer for its input callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExternalBuffer
type AudioUnitExternalBuffer struct {
	Buffer unsafe.Pointer
	Size unsafe.Pointer
}// AudioUnitFrequencyResponseBin - An audio unit’s audio level at a particular frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitFrequencyResponseBin
type AudioUnitFrequencyResponseBin struct {
}// AudioUnitMIDIControlMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitMIDIControlMapping
type AudioUnitMIDIControlMapping struct {
	Element AudioUnitElement
	Parameter AudioUnitParameterID
}// AudioUnitMeterClipping - Audio clipping that has occurred in a mixer unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitMeterClipping
type AudioUnitMeterClipping struct {
	SawInfinity unsafe.Pointer //  if there was an infinite value on this channel since the last time the property was retrieved.
	SawNotANumber unsafe.Pointer //  if there was a floating point “not a number” value on this channel since the last time the property was retrieved.
}// AudioUnitNodeConnection - A connection between two node objects in an audio processing graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitNodeConnection
type AudioUnitNodeConnection struct {
}// AudioUnitOtherPluginDesc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitOtherPluginDesc
type AudioUnitOtherPluginDesc struct {
	Format unsafe.Pointer
}// AudioUnitParameter - An adjustable audio unit attribute such as volume, pitch, or filter cutoff frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameter
type AudioUnitParameter struct {
	MElement AudioUnitElement // The audio unit element for the parameter.
}// AudioUnitParameterEvent - A scheduled change to an audio unit parameter’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterEvent
type AudioUnitParameterEvent struct {
}// AudioUnitParameterHistoryInfo - The suggested update rate and history duration for parameters which have the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterHistoryInfo
type AudioUnitParameterHistoryInfo struct {
}// AudioUnitParameterInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterInfo
type AudioUnitParameterInfo struct {
	Name unsafe.Pointer // Must be set to  .
}// AudioUnitParameterNameInfo - A short version of the name for an audio unit parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterNameInfo
type AudioUnitParameterNameInfo struct {
	InDesiredLength unsafe.Pointer // When setting an audio unit property that uses this data structure for its value, the maximum length that you are specifying for the audio unit parameter name.
}// AudioUnitParameterStringFromValue - A string representation of a parameter’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterStringFromValue
type AudioUnitParameterStringFromValue struct {
}// AudioUnitParameterValueFromString - A parameter’s value based on a string representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueFromString
type AudioUnitParameterValueFromString struct {
	InParamID AudioUnitParameterID
	InString unsafe.Pointer
	OutValue AudioUnitParameterValue
}// AudioUnitParameterValueName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueName
type AudioUnitParameterValueName struct {
	InParamID AudioUnitParameterID
	InValue unsafe.Pointer
	OutName unsafe.Pointer
}// AudioUnitParameterValueTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueTranslation
type AudioUnitParameterValueTranslation struct {
	AuParamID AudioUnitParameterID
	AuValue AudioUnitParameterValue
	OtherDesc unsafe.Pointer
	OtherValue unsafe.Pointer
}// AudioUnitPresetMAS_SettingData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPresetMAS_SettingData
type AudioUnitPresetMAS_SettingData struct {
	Data unsafe.Pointer
	DataLen unsafe.Pointer
	IsStockSetting unsafe.Pointer
	SettingID unsafe.Pointer
}// AudioUnitPresetMAS_Settings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPresetMAS_Settings
type AudioUnitPresetMAS_Settings struct {
	Settings unsafe.Pointer
}// AudioUnitProperty - A key-value pair that declares an attribute or behavior for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProperty
type AudioUnitProperty struct {
	MPropertyID AudioUnitPropertyID // The audio unit property identifier.
}// AudioUnitRenderContext - A structure that contains thread context information for a real-time rendering operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderContext
type AudioUnitRenderContext struct {
	Reserved uint32 // System-specific information.
	Workgroup unsafe.Pointer // The workgroup that manages the rendering threads of the audio unit.
}// CABarBeatTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CABarBeatTime
type CABarBeatTime struct {
	Bar unsafe.Pointer
	Beat unsafe.Pointer
	Reserved unsafe.Pointer
	Subbeat unsafe.Pointer
	SubbeatDivisor unsafe.Pointer
}// CAClockTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTime
type CAClockTime struct {
	Format unsafe.Pointer
	Reserved unsafe.Pointer
	Time unsafe.Pointer
}// CAFAudioDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFAudioDescription
type CAFAudioDescription struct {
	MBitsPerChannel unsafe.Pointer
	MBytesPerPacket unsafe.Pointer
	MChannelsPerFrame unsafe.Pointer
	MFormatFlags unsafe.Pointer
	MFormatID unsafe.Pointer
	MFramesPerPacket unsafe.Pointer
	MSampleRate unsafe.Pointer
}// CAFAudioFormatListItem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFAudioFormatListItem
type CAFAudioFormatListItem struct {
	MChannelLayoutTag unsafe.Pointer
	MFormat unsafe.Pointer
}// CAFChunkHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFChunkHeader
type CAFChunkHeader struct {
	MChunkSize unsafe.Pointer
	MChunkType unsafe.Pointer
}// CAFDataChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFDataChunk
type CAFDataChunk struct {
	MData unsafe.Pointer
	MEditCount unsafe.Pointer
}// CAFFileHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFFileHeader
type CAFFileHeader struct {
	MFileFlags unsafe.Pointer
	MFileType unsafe.Pointer
	MFileVersion unsafe.Pointer
}// CAFInfoStrings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFInfoStrings
type CAFInfoStrings struct {
	MNumEntries unsafe.Pointer
}// CAFInstrumentChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFInstrumentChunk
type CAFInstrumentChunk struct {
	MBaseNote unsafe.Pointer
	MInstrumentID unsafe.Pointer
	MMIDIHighNote unsafe.Pointer
	MMIDIHighVelocity unsafe.Pointer
	MMIDILowNote unsafe.Pointer
	MMIDILowVelocity unsafe.Pointer
	MReleaseRegionID unsafe.Pointer
	MStartRegionID unsafe.Pointer
	MSustainRegionID unsafe.Pointer
	MdBGain unsafe.Pointer
}// CAFMarker
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFMarker
type CAFMarker struct {
	MChannel unsafe.Pointer
	MFramePosition unsafe.Pointer
	MMarkerID unsafe.Pointer
	MSMPTETime unsafe.Pointer
	MType unsafe.Pointer
}// CAFMarkerChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFMarkerChunk
type CAFMarkerChunk struct {
	MMarkers unsafe.Pointer
	MNumberMarkers unsafe.Pointer
	MSMPTE_TimeType unsafe.Pointer
}// CAFOverviewChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFOverviewChunk
type CAFOverviewChunk struct {
	MData unsafe.Pointer
	MEditCount unsafe.Pointer
	MNumFramesPerOVWSample unsafe.Pointer
}// CAFOverviewSample
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFOverviewSample
type CAFOverviewSample struct {
	MMaxValue unsafe.Pointer
	MMinValue unsafe.Pointer
}// CAFPacketTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPacketTableHeader
type CAFPacketTableHeader struct {
	MNumberPackets unsafe.Pointer
	MNumberValidFrames unsafe.Pointer
	MPacketDescriptions unsafe.Pointer
	MPrimingFrames unsafe.Pointer
	MRemainderFrames unsafe.Pointer
}// CAFPeakChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPeakChunk
type CAFPeakChunk struct {
	MEditCount unsafe.Pointer
	MPeaks unsafe.Pointer
}// CAFPositionPeak
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPositionPeak
type CAFPositionPeak struct {
	MFrameNumber unsafe.Pointer
	MValue unsafe.Pointer
}// CAFRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegion
type CAFRegion struct {
	MFlags unsafe.Pointer
	MMarkers unsafe.Pointer
	MNumberMarkers unsafe.Pointer
	MRegionID unsafe.Pointer
}// CAFRegionChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegionChunk
type CAFRegionChunk struct {
	MNumberRegions unsafe.Pointer
	MRegions unsafe.Pointer
	MSMPTE_TimeType unsafe.Pointer
}// CAFStringID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFStringID
type CAFStringID struct {
	MStringID unsafe.Pointer
	MStringStartByteOffset unsafe.Pointer
}// CAFStrings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFStrings
type CAFStrings struct {
	MNumEntries unsafe.Pointer
	MStringsIDs unsafe.Pointer
}// CAFUMIDChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFUMIDChunk
type CAFUMIDChunk struct {
	MBytes unsafe.Pointer
}// CAF_SMPTE_Time
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAF_SMPTE_Time
type CAF_SMPTE_Time struct {
	MFrames unsafe.Pointer
	MHours unsafe.Pointer
	MMinutes unsafe.Pointer
	MSeconds unsafe.Pointer
	MSubFrameSampleOffset unsafe.Pointer
}// CAF_UUID_ChunkHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAF_UUID_ChunkHeader
type CAF_UUID_ChunkHeader struct {
	MHeader unsafe.Pointer
	MUUID unsafe.Pointer
}// CAMeterTrackEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAMeterTrackEntry
type CAMeterTrackEntry struct {
	Beats ClockBeats
	MeterDenom unsafe.Pointer
	MeterNumer unsafe.Pointer
}// CATempoMapEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CATempoMapEntry
type CATempoMapEntry struct {
}// ExtendedAudioFormatInfo - A specifier for the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedAudioFormatInfo
type ExtendedAudioFormatInfo struct {
	MASBD unsafe.Pointer // A format specification for an audio stream.
	MClassDescription unsafe.Pointer // A structure that describes an audio codec.
	MMagicCookie unsafe.Pointer // Decompression information for the audio data format specified in the   field.
	MMagicCookieSize unsafe.Pointer // The size, in bytes, of the   field.
}// ExtendedControlEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedControlEvent
type ExtendedControlEvent struct {
	ControlID AudioUnitParameterID
	GroupID MusicDeviceGroupID
	Value AudioUnitParameterValue
}// ExtendedNoteOnEvent - Describes a note-on event with extended parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedNoteOnEvent
type ExtendedNoteOnEvent struct {
	Duration unsafe.Pointer
	ExtendedParams unsafe.Pointer
	GroupID MusicDeviceGroupID
	InstrumentID MusicDeviceInstrumentID
}// ExtendedTempoEvent - Describes a music track tempo in beats-per-minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedTempoEvent
type ExtendedTempoEvent struct {
	Bpm unsafe.Pointer // The number of beats-per-minute.
}// HostCallbackInfo - The time- and transport-related callback functions for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallbackInfo
type HostCallbackInfo struct {
	TransportStateProc2 HostCallback_GetTransportState2
}// MIDIChannelMessage - Describes a MIDI channel message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIChannelMessage
type MIDIChannelMessage struct {
	Data1 unsafe.Pointer
	Data2 unsafe.Pointer
	Reserved unsafe.Pointer
	Status unsafe.Pointer // Data specific to the channel message.
}// MIDIMetaEvent - Describes a MIDI metaevent such as lyric text, time signature, and so on.
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
}// MIDINoteMessage - Describes a MIDI note.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDINoteMessage
type MIDINoteMessage struct {
	Channel unsafe.Pointer // The MIDI channel to play the note on.
	Duration unsafe.Pointer // The duration for the note.
	Note unsafe.Pointer // The note to play.
	ReleaseVelocity unsafe.Pointer // The key-release velocity for the note. Use 0 if you don’t want to specify a particular value.
	Velocity unsafe.Pointer // The key-press velocity for the note.
}// MIDIRawData - Describes a MIDI system-exclusive (SysEx) message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIRawData
type MIDIRawData struct {
	Data unsafe.Pointer
	Length unsafe.Pointer
}// MixerDistanceParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MixerDistanceParams
type MixerDistanceParams struct {
	MMaxAttenuation unsafe.Pointer
	MMaxDistance unsafe.Pointer
	MReferenceDistance unsafe.Pointer
}// MusicDeviceNoteParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceNoteParams
type MusicDeviceNoteParams struct {
	ArgCount unsafe.Pointer
	MControls unsafe.Pointer
	MPitch unsafe.Pointer
	MVelocity unsafe.Pointer
}// MusicDeviceStdNoteParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStdNoteParams
type MusicDeviceStdNoteParams struct {
	ArgCount unsafe.Pointer
	MPitch unsafe.Pointer
	MVelocity unsafe.Pointer
}// MusicEventUserData - Describes a user-defined event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventUserData
type MusicEventUserData struct {
	Data unsafe.Pointer // User-defined data.
	Length unsafe.Pointer // The size, in bytes, of the user data.
}// MusicTrackLoopInfo - Supports control of the looping behavior of a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackLoopInfo
type MusicTrackLoopInfo struct {
	LoopDuration MusicTimeStamp // The point in a music track, measured in beats from the end of the music track, at which to begin playback during looped playback.
	NumberOfLoops unsafe.Pointer // The number of times to play the designated portion of the music track.
}// NoteParamsControlValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NoteParamsControlValue
type NoteParamsControlValue struct {
	MID AudioUnitParameterID
	MValue AudioUnitParameterValue
}// ParameterEvent - Describes an audio unit parameter automation event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ParameterEvent
type ParameterEvent struct {
	Element AudioUnitElement
	ParameterID AudioUnitParameterID
	Scope AudioUnitScope
	Value AudioUnitParameterValue
}// ScheduledAudioFileRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioFileRegion
type ScheduledAudioFileRegion struct {
	MCompletionProc ScheduledAudioFileRegionCompletionProc // may be 
	MTimeStamp unsafe.Pointer
}// ScheduledAudioSlice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioSlice
type ScheduledAudioSlice struct {
	MBufferList unsafe.Pointer
	MCompletionProc ScheduledAudioSliceCompletionProc
	MCompletionProcUserData unsafe.Pointer
	MFlags unsafe.Pointer
	MNumberFrames unsafe.Pointer
	MReserved unsafe.Pointer
	MReserved2 unsafe.Pointer
	MTimeStamp unsafe.Pointer
}



