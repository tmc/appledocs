// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	SetFileFormat(value IAVAudioFormat)
	FramePosition() AudioFramePosition /* not a class type */
	SetFramePosition(value AudioFramePosition /* not a class type */)
	IsOpen() bool
	SetIsOpen(value bool)
	Length() AudioFramePosition /* not a class type */
	SetLength(value AudioFramePosition /* not a class type */)
	ProcessingFormat() IAVAudioFormat
	SetProcessingFormat(value IAVAudioFormat)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	AVAudioFileTypeKey() objc.IObject /* cross-framework: NSString */
	// methods:
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

// Alloc allocates a new instance without initialization.
func (ac _AudioFileClass) Alloc() AudioFile {
	rv := objc.Send[AudioFile](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The on-disk format of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/fileformat
func (a_ AudioFile) FileFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("fileFormat"))
	return rv
}


// The on-disk format of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/fileformat
func (a_ AudioFile) SetFileFormat(value IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFileFormat:"), value)
}


// The position in the file where the next read or write operation occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/frameposition
func (a_ AudioFile) FramePosition() AudioFramePosition /* not a class type */ {
	rv := objc.Send[AudioFramePosition](a_.ID, objc.Sel("framePosition"))
	return rv
}


// The position in the file where the next read or write operation occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/frameposition
func (a_ AudioFile) SetFramePosition(value AudioFramePosition /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFramePosition:"), value)
}


// A Boolean value that indicates whether the file is open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/isopen
func (a_ AudioFile) IsOpen() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isOpen"))
	return rv
}


// A Boolean value that indicates whether the file is open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/isopen
func (a_ AudioFile) SetIsOpen(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsOpen:"), value)
}


// The number of sample frames in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/length
func (a_ AudioFile) Length() AudioFramePosition /* not a class type */ {
	rv := objc.Send[AudioFramePosition](a_.ID, objc.Sel("length"))
	return rv
}


// The number of sample frames in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/length
func (a_ AudioFile) SetLength(value AudioFramePosition /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLength:"), value)
}


// The processing format of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/processingformat
func (a_ AudioFile) ProcessingFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("processingFormat"))
	return rv
}


// The processing format of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/processingformat
func (a_ AudioFile) SetProcessingFormat(value IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProcessingFormat:"), value)
}


// The location of the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/url
func (a_ AudioFile) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("url"))
	return rv
}


// The location of the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofile/url
func (a_ AudioFile) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUrl:"), value)
}


// A string that indicates the audio file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiofiletypekey
func (a_ AudioFile) AVAudioFileTypeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAudioFileTypeKey"))
	return rv
}



