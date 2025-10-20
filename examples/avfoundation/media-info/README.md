# AVFoundation Framework Example

This example demonstrates the AVFoundation framework for audio/video playback, recording, and media processing in Go.

## What it demonstrates

- AVPlayer creation
- Media playback workflow
- Camera capture workflow
- Supported formats
- Capture outputs
- Common use cases
- Framework integration
- Audio session management

## Running the example

```bash
go run main.go
# or with e2e flag
go run main.go -e2e
```

## Key Concepts

### AVFoundation Overview

AVFoundation is Apple's comprehensive media framework providing:
- Audio and video playback
- Audio and video recording
- Camera and microphone capture
- Media editing and composition
- Asset management and export
- Real-time media processing
- Live streaming support

### Architecture

AVFoundation uses a pipeline-based architecture:
- **Assets** represent media resources
- **Players** control playback
- **Capture Sessions** coordinate recording
- **Compositions** enable editing

## Media Playback

### Basic Playback

```go
import (
    "github.com/tmc/appledocs/generated/avfoundation"
    "github.com/tmc/appledocs/generated/foundation"
)

// Create URL to media
url := foundation.NewURLWithString("file:///path/to/video.mp4")

// Create asset
asset := avfoundation.NewAssetWithURL(url)

// Create player item
playerItem := avfoundation.NewPlayerItemWithAsset(asset)

// Create player
player := avfoundation.NewPlayerWithPlayerItem(playerItem)

// Start playback
player.Play()
```

### Video Display

```go
// For video playback, create a player layer
playerLayer := avfoundation.NewPlayerLayer()
playerLayer.SetPlayer(player)

// Add to view's layer hierarchy
view.Layer().AddSublayer(playerLayer)
playerLayer.SetFrame(view.Bounds())
```

### Playback Control

```go
// Play/Pause
player.Play()
player.Pause()

// Seek to time
time := avfoundation.CMTimeMake(30, 1) // 30 seconds
player.SeekToTime(time)

// Set playback rate
player.SetRate(1.0)  // Normal speed
player.SetRate(2.0)  // 2x speed
player.SetRate(0.5)  // Half speed

// Get current time
currentTime := player.CurrentTime()

// Get duration
duration := player.CurrentItem().Duration()
```

### Observing Playback

```go
// Observe player status (requires KVO)
// player.addObserver(..., forKeyPath: "status", ...)

// Observe playback position
player.AddPeriodicTimeObserver(interval, queue, func(time CMTime) {
    // Update UI with current time
})

// Observe when playback ends
// NotificationCenter.default.addObserver(...,
//     name: .AVPlayerItemDidPlayToEndTime, ...)
```

## Camera Capture

### Setting Up Capture Session

```go
// Create capture session
session := avfoundation.NewCaptureSession()

// Begin configuration
session.BeginConfiguration()

// Set session preset
session.SetSessionPreset("AVCaptureSessionPresetHigh")

// Get camera device
devices := avfoundation.CaptureDeviceDiscoverySession(
    deviceTypes: [.builtInWideAngleCamera],
    mediaType: .video,
    position: .back,
)
device := devices.Devices().First()

// Create input
input := avfoundation.NewCaptureDeviceInput(device)
if session.CanAddInput(input) {
    session.AddInput(input)
}

// Create output
output := avfoundation.NewCaptureVideoDataOutput()
output.SetSampleBufferDelegate(delegate, queue)

if session.CanAddOutput(output) {
    session.AddOutput(output)
}

// Commit configuration
session.CommitConfiguration()

// Start capture
session.StartRunning()
```

### Handling Captured Frames

```go
// Implement AVCaptureVideoDataOutputSampleBufferDelegate

func CaptureOutput(
    output: AVCaptureOutput,
    sampleBuffer: CMSampleBuffer,
    connection: AVCaptureConnection,
) {
    // Get pixel buffer
    pixelBuffer := CMSampleBufferGetImageBuffer(sampleBuffer)

    // Process frame with Vision, Core Image, Metal, etc.

    // Example: Apply Core Image filter
    ciImage := coreimage.NewImageWithCVPixelBuffer(pixelBuffer)
    filtered := filter.OutputImage()

    // Render or display result
}
```

### Photo Capture

```go
// Create photo output
photoOutput := avfoundation.NewCapturePhotoOutput()

// Add to session
if session.CanAddOutput(photoOutput) {
    session.AddOutput(photoOutput)
}

// Configure photo settings
settings := avfoundation.NewCapturePhotoSettings()
settings.SetFlashMode(.auto)
settings.SetHighResolutionPhotoEnabled(true)

// Capture photo
photoOutput.CapturePhoto(settings, delegate)
```

## Media Recording

### Recording to File

```go
// Create movie file output
fileOutput := avfoundation.NewCaptureMovieFileOutput()

// Add to capture session
if session.CanAddOutput(fileOutput) {
    session.AddOutput(fileOutput)
}

// Start recording
outputURL := foundation.NewURLWithString("file:///path/to/output.mov")
fileOutput.StartRecordingToOutputFileURL(outputURL, delegate)

// Stop recording
fileOutput.StopRecording()
```

### Audio Recording

```go
// Setup audio session
audioSession := avfoundation.SharedAudioSession()
audioSession.SetCategory(.record)
audioSession.SetActive(true)

// Create audio recorder
settings := map[string]interface{}{
    AVFormatIDKey: kAudioFormatMPEG4AAC,
    AVSampleRateKey: 44100.0,
    AVNumberOfChannelsKey: 2,
    AVEncoderBitRateKey: 128000,
}

recorder := avfoundation.NewAudioRecorder(url, settings)
recorder.Record()

// Stop recording
recorder.Stop()
```

## Media Editing

### Creating Compositions

```go
// Create composition
composition := avfoundation.NewMutableComposition()

// Create video track
videoTrack := composition.AddMutableTrack(
    withMediaType: .video,
    preferredTrackID: kCMPersistentTrackID_Invalid,
)

// Load source asset
sourceAsset := avfoundation.NewAssetWithURL(sourceURL)
sourceVideoTrack := sourceAsset.TracksWithMediaType(.video).First()

// Insert time range
timeRange := CMTimeRange{
    start: CMTime.zero,
    duration: sourceAsset.Duration(),
}

videoTrack.InsertTimeRange(
    timeRange,
    of: sourceVideoTrack,
    at: CMTime.zero,
)

// Create audio track
audioTrack := composition.AddMutableTrack(
    withMediaType: .audio,
    preferredTrackID: kCMPersistentTrackID_Invalid,
)

sourceAudioTrack := sourceAsset.TracksWithMediaType(.audio).First()
audioTrack.InsertTimeRange(timeRange, of: sourceAudioTrack, at: CMTime.zero)
```

### Applying Video Effects

```go
// Create video composition
videoComposition := avfoundation.NewMutableVideoComposition()
videoComposition.SetRenderSize(CGSize{1920, 1080})
videoComposition.SetFrameDuration(CMTime{value: 1, timescale: 30}) // 30fps

// Create layer instruction
layerInstruction := avfoundation.NewMutableVideoCompositionLayerInstruction(
    trackID: videoTrack.TrackID(),
)

// Apply transform
layerInstruction.SetTransform(transform, at: CMTime.zero)

// Create instruction
instruction := avfoundation.NewMutableVideoCompositionInstruction()
instruction.SetTimeRange(timeRange)
instruction.SetLayerInstructions([layerInstruction])

videoComposition.SetInstructions([instruction])
```

## Media Export

### Exporting Media

```go
// Create export session
exportSession := avfoundation.NewAssetExportSession(
    asset: composition,
    presetName: AVAssetExportPreset1920x1080,
)

// Configure export
exportSession.SetOutputURL(outputURL)
exportSession.SetOutputFileType(.mp4)
exportSession.SetShouldOptimizeForNetworkUse(true)

// Optional: Apply video composition
exportSession.SetVideoComposition(videoComposition)

// Optional: Apply audio mix
exportSession.SetAudioMix(audioMix)

// Export asynchronously
exportSession.ExportAsynchronously {
    switch exportSession.Status() {
    case .completed:
        // Export successful
    case .failed:
        // Handle error: exportSession.Error()
    case .cancelled:
        // Export was cancelled
    }
}
```

### Export Progress

```go
// Monitor export progress
timer := Timer.scheduledTimer(withTimeInterval: 0.1) { _ in
    progress := exportSession.Progress()
    // Update UI with progress (0.0 to 1.0)
}

// Cancel export
exportSession.CancelExport()
```

## Audio Processing

### AVAudioEngine

```go
// Create audio engine
audioEngine := avfoundation.NewAudioEngine()

// Get player node
playerNode := audioEngine.MainMixerNode()

// Create effect nodes
reverb := avfoundation.NewAudioUnitReverb()
reverb.LoadFactoryPreset(.cathedral)
reverb.SetWetDryMix(50.0)

// Attach nodes
audioEngine.AttachNode(reverb)

// Connect nodes
audioEngine.Connect(playerNode, to: reverb, format: format)
audioEngine.Connect(reverb, to: audioEngine.OutputNode(), format: format)

// Start engine
audioEngine.Start()
```

### Audio File Playback

```go
// Create audio file
audioFile := avfoundation.NewAudioFile(url)

// Schedule file
playerNode.ScheduleFile(audioFile, at: nil) {
    // Playback completed
}

// Start playback
playerNode.Play()
```

## Advanced Features

### Time-Based Metadata

```go
// Add timed metadata
metadataItem := avfoundation.NewMutableMetadataItem()
metadataItem.SetIdentifier(.commonIdentifierTitle)
metadataItem.SetValue("Chapter 1")
metadataItem.SetTime(CMTime{value: 0, timescale: 1})
metadataItem.SetDuration(CMTime{value: 60, timescale: 1})

// Add to asset
composition.SetMetadata([metadataItem])
```

### Custom Video Compositing

```go
// Implement AVVideoCompositing protocol
// Allows custom GPU-based video effects using Metal

class CustomCompositor: NSObject, AVVideoCompositing {
    func startRequest(_ request: AVAsynchronousVideoCompositionRequest) {
        // Get source frames
        let sourceFrame = request.sourceFrame(byTrackID: trackID)

        // Apply custom Metal processing
        let processedFrame = applyMetalEffect(sourceFrame)

        // Finish request
        request.finish(withComposedVideoFrame: processedFrame)
    }
}
```

### Face Detection in Video

```go
// Create metadata output for face detection
metadataOutput := avfoundation.NewCaptureMetadataOutput()
metadataOutput.SetMetadataObjectTypes([.face])
metadataOutput.SetMetadataObjectsDelegate(delegate, queue)

// Add to session
session.AddOutput(metadataOutput)

// Handle detected faces
func metadataOutput(
    output: AVCaptureMetadataOutput,
    didOutput metadataObjects: [AVMetadataObject],
    connection: AVCaptureConnection,
) {
    for metadata in metadataObjects {
        if let face = metadata as? AVMetadataFaceObject {
            // Process detected face
            let bounds = face.Bounds()
        }
    }
}
```

## Best Practices

### Performance

1. **Run on Background Threads**: Capture and processing should not block main thread
2. **Reuse Buffer Pools**: Minimize allocations for video frames
3. **Optimize Video Settings**: Match resolution to display needs
4. **Use Hardware Acceleration**: Leverage VideoToolbox for encoding/decoding
5. **Minimize Copies**: Use IOSurface for zero-copy frame sharing

### Memory Management

```go
// Properly release resources
player.Pause()
player.ReplaceCurrentItemWithPlayerItem(nil)
session.StopRunning()

// Remove observers
player.RemoveTimeObserver(observer)
NotificationCenter.default.removeObserver(observer)
```

### Audio Session

```go
// Configure audio session appropriately
audioSession := AVAudioSession.sharedInstance()

// For playback
audioSession.setCategory(.playback, mode: .moviePlayback)

// For recording
audioSession.setCategory(.record, mode: .videoRecording)

// For both (e.g., video calls)
audioSession.setCategory(.playAndRecord, mode: .videoChat)

// Activate session
audioSession.setActive(true)
```

### Error Handling

```go
// Always check for errors
var error: NSError?
let input = try? AVCaptureDeviceInput(device: device)
if input == nil {
    // Handle error
    return
}

// Observe player item status
playerItem.addObserver(..., forKeyPath: "status") {
    switch playerItem.status {
    case .readyToPlay:
        // Ready
    case .failed:
        // Handle error: playerItem.error
    }
}
```

## Common Patterns

### Video Player with Controls

```go
// Setup player
player := AVPlayer(url: url)
playerLayer := AVPlayerLayer(player: player)
view.layer.addSublayer(playerLayer)

// Add playback controls
player.addPeriodicTimeObserver(interval) { time in
    // Update progress slider
}

// Handle play/pause
@objc func togglePlayPause() {
    if player.rate == 0 {
        player.play()
    } else {
        player.pause()
    }
}
```

### Live Camera Preview

```go
// Setup session
session := AVCaptureSession()
session.sessionPreset = .high

// Get camera
let device = AVCaptureDevice.default(.builtInWideAngleCamera, for: .video, position: .back)
let input = try! AVCaptureDeviceInput(device: device!)

// Add input
session.addInput(input)

// Create preview layer
previewLayer := AVCaptureVideoPreviewLayer(session: session)
previewLayer.videoGravity = .resizeAspectFill
view.layer.addSublayer(previewLayer)

// Start
session.startRunning()
```

### Screen Recording

```go
// Use ScreenCaptureKit (modern) or AVCaptureScreenInput (legacy)
// See ScreenCaptureKit examples for recommended approach
```

## Integration Examples

### With Vision Framework

```go
// Process camera frames with Vision
func captureOutput(...) {
    let pixelBuffer = CMSampleBufferGetImageBuffer(sampleBuffer)

    let request = VNDetectFaceRectanglesRequest { request, error in
        guard let results = request.results as? [VNFaceObservation] else { return }
        // Handle detected faces
    }

    let handler = VNImageRequestHandler(cvPixelBuffer: pixelBuffer)
    try? handler.perform([request])
}
```

### With Core Image

```go
// Apply filters to video
let ciImage = CIImage(cvPixelBuffer: pixelBuffer)
let filtered = filter.outputImage

// Render back to pixel buffer
ciContext.render(filtered, to: outputPixelBuffer)
```

### With Metal

```go
// Use Metal for custom video processing
let textureCache = CVMetalTextureCacheCreate(...)
let texture = CVMetalTextureGetTexture(...)

// Apply Metal shader
// Render to output texture
```

## Use Cases

### Video Player App
- Load and play video files
- Streaming media support
- Playback controls (play, pause, seek)
- Picture-in-picture mode

### Camera App
- Real-time camera preview
- Photo and video capture
- Apply live filters
- Face/QR code detection

### Video Conferencing
- Camera/microphone capture
- Real-time video/audio transmission
- Hardware encoding
- Echo cancellation

### Screen Recording
- Capture screen/windows
- Add camera overlay
- Record system audio
- Export to video file

### Video Editor
- Multi-track timeline
- Trim, crop, rotate
- Add transitions
- Apply effects and filters
- Export in multiple formats

## References

- [AVFoundation Programming Guide](https://developer.apple.com/documentation/avfoundation)
- [AVPlayer Documentation](https://developer.apple.com/documentation/avfoundation/avplayer)
- [AVCaptureSession Documentation](https://developer.apple.com/documentation/avfoundation/avcapturesession)
- [Media Playback Programming Guide](https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/MediaPlaybackGuide/)
- [Camera Capture Best Practices](https://developer.apple.com/documentation/avfoundation/capture_setup)
