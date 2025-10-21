// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox
import (
	"unsafe"
)

// AUEventListenerRef is a CoreGraphics opaque type.
type AUEventListenerRef unsafe.Pointer

// AUParameterListenerRef is a CoreGraphics opaque type.
type AUParameterListenerRef unsafe.Pointer

// AudioConverterRef is a CoreGraphics opaque type.
type AudioConverterRef unsafe.Pointer

// AudioQueueBufferRef is a CoreGraphics opaque type.
type AudioQueueBufferRef unsafe.Pointer

// AudioQueueProcessingTapRef is a CoreGraphics opaque type.
type AudioQueueProcessingTapRef unsafe.Pointer

// AudioQueueRef is a CoreGraphics opaque type.
type AudioQueueRef unsafe.Pointer

// AudioQueueTimelineRef is a CoreGraphics opaque type.
type AudioQueueTimelineRef unsafe.Pointer

// CAClockRef is a CoreGraphics opaque type.
type CAClockRef unsafe.Pointer

// CFArrayRef is a CoreGraphics opaque type.
type CFArrayRef unsafe.Pointer

// CFDataRef is a CoreGraphics opaque type.
type CFDataRef unsafe.Pointer

// CFDictionaryRef is a CoreGraphics opaque type.
type CFDictionaryRef unsafe.Pointer

// CFRunLoopRef is a CoreGraphics opaque type.
type CFRunLoopRef unsafe.Pointer

// CFStringRef is a CoreGraphics opaque type.
type CFStringRef unsafe.Pointer

// CFURLRef is a CoreGraphics opaque type.
type CFURLRef unsafe.Pointer

// ExtAudioFileRef is a CoreGraphics opaque type.
type ExtAudioFileRef unsafe.Pointer

// MIDIEndpointRef is a CoreGraphics opaque type.
type MIDIEndpointRef unsafe.Pointer


// TimeInterval for non-CoreGraphics frameworks
type TimeInterval = float64  // NSTimeInterval

// Fallback type aliases for undefined types
// These types are referenced in method signatures but not fully documented.
// Using unsafe.Pointer as fallback to allow code generation.
type AUAudioChannelCount unsafe.Pointer

type AUAudioFrameCount unsafe.Pointer

type AUAudioObjectID unsafe.Pointer

type AUAudioUnitViewConfiguration unsafe.Pointer

type AUEventListenerBlock unsafe.Pointer

type AUEventListenerProc unsafe.Pointer

type AUGraph unsafe.Pointer

type AUHostMusicalContextBlock unsafe.Pointer

type AUHostTransportStateBlock unsafe.Pointer

type AUImplementorDisplayNameWithLengthCallback unsafe.Pointer

type AUImplementorValueFromStringCallback unsafe.Pointer

type AUInputHandler unsafe.Pointer

type AUInternalRenderBlock unsafe.Pointer

type AUMIDICIProfileChangedBlock unsafe.Pointer

type AUMIDIEventListBlock unsafe.Pointer

type AUMIDIOutputEventBlock unsafe.Pointer

type AUNode unsafe.Pointer

type AUNodeInteraction unsafe.Pointer

type AUParameterAddress unsafe.Pointer

type AUParameterAutomationObserver unsafe.Pointer

type AUParameterListenerBlock unsafe.Pointer

type AUParameterListenerProc unsafe.Pointer

type AUParameterObserver unsafe.Pointer

type AUParameterObserverToken unsafe.Pointer

type AUParameterRecordingObserver unsafe.Pointer

type AUPresetEvent unsafe.Pointer

type AURenderBlock unsafe.Pointer

type AURenderCallback unsafe.Pointer

type AURenderContextObserver unsafe.Pointer

type AURenderObserver unsafe.Pointer

type AURenderPullInputBlock unsafe.Pointer

type AUScheduleMIDIEventBlock unsafe.Pointer

type AUScheduleParameterBlock unsafe.Pointer

type AUValue unsafe.Pointer

type AUViewControllerBase unsafe.Pointer

type AVAudioFormat unsafe.Pointer

type AudioBufferList unsafe.Pointer

type AudioChannelLayout unsafe.Pointer

type AudioClassDescription unsafe.Pointer

type AudioCodec unsafe.Pointer

type AudioCodecPropertyID unsafe.Pointer

type AudioComponent unsafe.Pointer

type AudioComponentDescription unsafe.Pointer

type AudioComponentFactoryFunction unsafe.Pointer

type AudioComponentInstance unsafe.Pointer

type AudioConverterComplexInputDataProc unsafe.Pointer

type AudioConverterComplexInputDataProcRealtimeSafe unsafe.Pointer

type AudioConverterInputDataProc unsafe.Pointer

type AudioConverterPropertyID unsafe.Pointer

type AudioFileComponent unsafe.Pointer

type AudioFileComponentPropertyID unsafe.Pointer

type AudioFileID unsafe.Pointer

type AudioFilePropertyID unsafe.Pointer

type AudioFileStreamID unsafe.Pointer

type AudioFileStreamPropertyID unsafe.Pointer

type AudioFileTypeID unsafe.Pointer

type AudioFormatPropertyID unsafe.Pointer

type AudioObjectID unsafe.Pointer

type AudioObjectPropertyAddress unsafe.Pointer

type AudioObjectPropertyListenerProc unsafe.Pointer

type AudioQueueInputCallback unsafe.Pointer

type AudioQueueInputCallbackBlock unsafe.Pointer

type AudioQueueOutputCallback unsafe.Pointer

type AudioQueueOutputCallbackBlock unsafe.Pointer

type AudioQueueParameterEvent unsafe.Pointer

type AudioQueueParameterID unsafe.Pointer

type AudioQueueParameterValue unsafe.Pointer

type AudioQueueProcessingTapCallback unsafe.Pointer

type AudioQueuePropertyID unsafe.Pointer

type AudioQueuePropertyListenerProc unsafe.Pointer

type AudioServicesPropertyID unsafe.Pointer

type AudioServicesSystemSoundCompletionProc unsafe.Pointer

type AudioSessionInterruptionListener unsafe.Pointer

type AudioSessionPropertyID unsafe.Pointer

type AudioSessionPropertyListener unsafe.Pointer

type AudioStreamBasicDescription unsafe.Pointer

type AudioStreamPacketDependencyDescription unsafe.Pointer

type AudioStreamPacketDescription unsafe.Pointer

type AudioTimeStamp unsafe.Pointer

type AudioUnitElement unsafe.Pointer

type AudioUnitEvent unsafe.Pointer

type AudioUnitParameter unsafe.Pointer

type AudioUnitParameterEvent unsafe.Pointer

type AudioUnitParameterID unsafe.Pointer

type AudioUnitParameterValue unsafe.Pointer

type AudioUnitPropertyID unsafe.Pointer

type AudioUnitPropertyListenerProc unsafe.Pointer

type AudioUnitScope unsafe.Pointer

type BOOL unsafe.Pointer

type Boolean unsafe.Pointer

type CABarBeatTime unsafe.Pointer

type CAClockBeats unsafe.Pointer

type CAClockListenerProc unsafe.Pointer

type CAClockSeconds unsafe.Pointer

type CAClockTempo unsafe.Pointer

type CAClockTime unsafe.Pointer

type CFAbsoluteTime unsafe.Pointer

type Class unsafe.Pointer

type ComponentDescription unsafe.Pointer

type ExtAudioFilePropertyID unsafe.Pointer

type ExtendedControlEvent unsafe.Pointer

type ExtendedNoteOnEvent unsafe.Pointer

type FILE unsafe.Pointer

type FSRef unsafe.Pointer

type Float32 unsafe.Pointer

type Float64 unsafe.Pointer

type MIDICIProfile unsafe.Pointer

type MIDICIProfileState unsafe.Pointer

type MIDIChannelMessage unsafe.Pointer

type MIDIChannelNumber unsafe.Pointer

type MIDIEventList unsafe.Pointer

type MIDIMetaEvent unsafe.Pointer

type MIDINoteMessage unsafe.Pointer

type MIDIPacketList unsafe.Pointer

type MIDIProtocolID unsafe.Pointer

type MIDIRawData unsafe.Pointer

type MusicDeviceComponent unsafe.Pointer

type MusicDeviceGroupID unsafe.Pointer

type MusicDeviceInstrumentID unsafe.Pointer

type MusicDeviceNoteParams unsafe.Pointer

type MusicEventIterator unsafe.Pointer

type MusicEventType unsafe.Pointer

type MusicEventUserData unsafe.Pointer

type MusicPlayer unsafe.Pointer

type MusicSequence unsafe.Pointer

type MusicSequenceUserCallback unsafe.Pointer

type MusicTimeStamp unsafe.Pointer

type MusicTrack unsafe.Pointer

type NSArray unsafe.Pointer

type NSDictionary unsafe.Pointer

type NSError unsafe.Pointer

type NSIndexSet unsafe.Pointer

type NSInteger unsafe.Pointer

type NSKeyValueObservingOptions unsafe.Pointer

type NSNumber unsafe.Pointer

type NSObject unsafe.Pointer

type NSString unsafe.Pointer

type NSTimeInterval unsafe.Pointer

type NSUInteger unsafe.Pointer

type NoteInstanceID unsafe.Pointer

type OSStatus unsafe.Pointer

type ParameterEvent unsafe.Pointer

type SInt16 unsafe.Pointer

type SInt64 unsafe.Pointer

type SInt8 unsafe.Pointer

type SMPTETime unsafe.Pointer

type SystemSoundID unsafe.Pointer

type UIImage unsafe.Pointer

type UInt16 unsafe.Pointer

type UInt32 unsafe.Pointer

type UInt64 unsafe.Pointer

type UInt8 unsafe.Pointer



