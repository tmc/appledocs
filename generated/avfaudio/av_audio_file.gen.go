// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioFile] class.
var (
	AudioFileClass     _AudioFileClass
	AudioFileClassOnce sync.Once
)

func getAudioFileClass() _AudioFileClass {
	AudioFileClassOnce.Do(func() {
		AudioFileClass = _AudioFileClass{objc.GetClass("AVAudioFile")}
	})
	return AudioFileClass
}

type _AudioFileClass struct {
	class objc.Class
}





// An interface definition for the [AudioFile] class.
type IAudioFile interface {
	objectivec.IObject
	

	// properties:
	FileFormat() IAVAudioFormat
	FramePosition() AudioFramePosition /* typedef */
	SetFramePosition(value AudioFramePosition /* typedef */)
	IsOpen() bool
	Length() AudioFramePosition /* typedef */
	ProcessingFormat() IAVAudioFormat
	Url() objc.IObject /* cross-framework: NSURL */
	AVAudioFileTypeKey() objc.IObject /* cross-framework: NSString */


	

	// methods:
	Close()
	ReadIntoBufferError(buffer IAVAudioPCMBuffer, outError objectivec.IObject) bool
	ReadIntoBufferFrameCountError(buffer IAVAudioPCMBuffer, frames AudioFrameCount /* typedef */, outError objectivec.IObject) bool
	WriteFromBufferError(buffer IAVAudioPCMBuffer, outError objectivec.IObject) bool


}





// Alloc allocates a new instance without initialization.
func (ac _AudioFileClass) Alloc() AudioFile {
	rv := objc.Send[AudioFile](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioFileClass) New() AudioFile {
	rv := objc.Send[AudioFile](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioFile) Init() AudioFile {
	rv := objc.Send[AudioFile](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioFile) Autorelease() AudioFile {
	rv := objc.Send[AudioFile](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioFile creates a new AudioFile instance.
func NewAudioFile() AudioFile {
	return getAudioFileClass().New()
}





// An object that represents an audio file that the system can open for reading or writing.
//
// Regardless of the file format, you read and write using objects. These objects contain samples as that the framework refers to as the file’s processing format. You convert to and from using the file’s actual format. Reads and writes are always sequential. Random access is possible by setting the property.


// An object that represents an audio file that the system can open for reading or writing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile
type AudioFile struct {
	objectivec.Object
}

// AudioFileFrom constructs a [AudioFile] from an unsafe.Pointer.
//
// An object that represents an audio file that the system can open for reading or writing.
func AudioFileFrom(ptr unsafe.Pointer) AudioFile {
	return AudioFile{objectivec.Object{objc.ID(ptr)}}
}






// Opens a file for reading using the specified processing format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forReading:commonFormat:interleaved:)
func NewAudioFileForReadingCommonFormatInterleavedError(fileURL objc.IObject /* cross-framework: NSURL */, format AudioCommonFormat, interleaved bool, outError objectivec.IObject) AudioFile {
	instance := getAudioFileClass().Alloc()
	rv := objc.Send[AudioFile](instance.ID, objc.Sel("initForReading:commonFormat:interleaved:error:"), fileURL, format, interleaved, outError)
	rv.Autorelease()
	return rv
}


// Opens a file for reading using the standard, deinterleaved floating point format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forReading:)
func NewAudioFileForReadingError(fileURL objc.IObject /* cross-framework: NSURL */, outError objectivec.IObject) AudioFile {
	instance := getAudioFileClass().Alloc()
	rv := objc.Send[AudioFile](instance.ID, objc.Sel("initForReading:error:"), fileURL, outError)
	rv.Autorelease()
	return rv
}


// Opens a file for writing using a specified processing format and settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forWriting:settings:commonFormat:interleaved:)
func NewAudioFileForWritingSettingsCommonFormatInterleavedError(fileURL objc.IObject /* cross-framework: NSURL */, settings foundation.IDictionary, format AudioCommonFormat, interleaved bool, outError objectivec.IObject) AudioFile {
	instance := getAudioFileClass().Alloc()
	rv := objc.Send[AudioFile](instance.ID, objc.Sel("initForWriting:settings:commonFormat:interleaved:error:"), fileURL, settings, format, interleaved, outError)
	rv.Autorelease()
	return rv
}


// Opens a file for writing using the specified settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forWriting:settings:)
func NewAudioFileForWritingSettingsError(fileURL objc.IObject /* cross-framework: NSURL */, settings foundation.IDictionary, outError objectivec.IObject) AudioFile {
	instance := getAudioFileClass().Alloc()
	rv := objc.Send[AudioFile](instance.ID, objc.Sel("initForWriting:settings:error:"), fileURL, settings, outError)
	rv.Autorelease()
	return rv
}

















// Closes the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/close()
func (a_ AudioFile) Close() {
	objc.Send[objc.ID](a_.ID, objc.Sel("close"))
}


// Reads an entire audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/read(into:)
func (a_ AudioFile) ReadIntoBufferError(buffer IAVAudioPCMBuffer, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("readIntoBuffer:error:"), buffer, outError)
	return rv
}


// Reads a portion of an audio buffer using the number of frames you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/read(into:frameCount:)
func (a_ AudioFile) ReadIntoBufferFrameCountError(buffer IAVAudioPCMBuffer, frames AudioFrameCount /* typedef */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("readIntoBuffer:frameCount:error:"), buffer, frames, outError)
	return rv
}


// Writes an audio buffer sequentially.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/write(from:)
func (a_ AudioFile) WriteFromBufferError(buffer IAVAudioPCMBuffer, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("writeFromBuffer:error:"), buffer, outError)
	return rv
}







// The on-disk format of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/fileFormat
func (a_ AudioFile) FileFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("fileFormat"))
	return rv
}


// The position in the file where the next read or write operation occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/framePosition
func (a_ AudioFile) FramePosition() AudioFramePosition /* typedef */ {
	rv := objc.Send[int64](a_.ID, objc.Sel("framePosition"))
	return rv
}


// The position in the file where the next read or write operation occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/framePosition
func (a_ AudioFile) SetFramePosition(value AudioFramePosition /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFramePosition:"), value)
}


// A Boolean value that indicates whether the file is open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/isOpen
func (a_ AudioFile) IsOpen() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isOpen"))
	return rv
}


// The number of sample frames in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/length
func (a_ AudioFile) Length() AudioFramePosition /* typedef */ {
	rv := objc.Send[int64](a_.ID, objc.Sel("length"))
	return rv
}


// The processing format of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/processingFormat
func (a_ AudioFile) ProcessingFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("processingFormat"))
	return rv
}


// The location of the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/url
func (a_ AudioFile) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("url"))
	return rv
}


// A string that indicates the audio file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofiletypekey
func (a_ AudioFile) AVAudioFileTypeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioFileTypeKey"))
	return rv
}







