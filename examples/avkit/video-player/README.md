# AVKit Framework Example

This example demonstrates the AVKit framework for creating video player user interfaces in Go.

## What it demonstrates

- AVPlayerViewController creation
- Built-in player controls
- Picture-in-picture mode
- AirPlay integration
- Customization options
- Platform-specific features
- Accessibility support
- Common workflows

## Running the example

```bash
go run main.go
# or with e2e flag
go run main.go -e2e
```

## Key Concepts

### AVKit Overview

AVKit provides ready-to-use video player interfaces:
- **AVPlayerViewController** (iOS/tvOS) - Full-featured player UI
- **AVPlayerView** (macOS) - Native AppKit player view
- Built-in playback controls
- Picture-in-picture support
- AirPlay streaming
- Accessibility features
- System integration (Now Playing, Control Center)

### Why Use AVKit?

**Instead of building custom video UI:**
- Get playback controls for free
- Automatic platform integration
- Accessibility built-in
- Picture-in-picture handled
- AirPlay support included
- Consistent user experience
- Less code to maintain

**AVKit handles:**
- Play/pause buttons
- Scrubbing/seeking
- Volume control
- Full-screen mode
- Loading indicators
- Error states
- Subtitles/captions
- Chapter markers

## Basic Video Player

### iOS/iPadOS Implementation

```go
import (
    "github.com/tmc/appledocs/generated/avkit"
    "github.com/tmc/appledocs/generated/avfoundation"
    "github.com/tmc/appledocs/generated/foundation"
)

// Create URL to video
videoURL := foundation.NewURLWithString(
    "https://example.com/video.mp4",
)

// Create AVPlayer
player := avfoundation.NewPlayerWithURL(videoURL)

// Create player view controller
playerVC := avkit.NewPlayerViewController()
playerVC.SetPlayer(player)

// Present modally
parentVC.PresentViewController(playerVC, animated: true) {
    // Play when presented
    player.Play()
}

// Or embed in view hierarchy
addChildViewController(playerVC)
view.AddSubview(playerVC.View())
playerVC.View().SetFrame(containerView.Bounds())
playerVC.DidMoveToParentViewController(self)
```

### macOS Implementation

```go
// Create player view (macOS)
playerView := avkit.NewPlayerView()
playerView.SetPlayer(player)

// Add to window
window.ContentView().AddSubview(playerView)
playerView.SetFrame(window.ContentView().Bounds())

// Configure controls
playerView.SetShowsFullScreenToggleButton(true)
playerView.SetControlsStyle(.floating)

// Start playback
player.Play()
```

## Customization

### Control Visibility

```go
// Hide controls
playerVC.SetShowsPlaybackControls(false)

// Show controls
playerVC.SetShowsPlaybackControls(true)

// Auto-hide after inactivity (default behavior)
```

### Video Gravity

```go
// Aspect fill (crop to fill)
playerVC.SetVideoGravity("AVLayerVideoGravityResizeAspectFill")

// Aspect fit (letterbox)
playerVC.SetVideoGravity("AVLayerVideoGravityResizeAspect")

// Stretch to fill
playerVC.SetVideoGravity("AVLayerVideoGravityResize")
```

### Full-Screen Behavior

```go
// Auto enter full-screen when playback begins
playerVC.SetEntersFullScreenWhenPlaybackBegins(true)

// Auto exit full-screen when playback ends
playerVC.SetExitsFullScreenWhenPlaybackEnds(true)

// Disable full-screen
playerVC.SetAllowsFullScreen(false) // macOS only
```

### Playback Restrictions

```go
// Disable seeking/scrubbing
playerVC.SetRequiresLinearPlayback(true)

// Enable all controls (default)
playerVC.SetRequiresLinearPlayback(false)
```

## Picture-in-Picture

### Setting Up PiP

```go
import "github.com/tmc/appledocs/generated/avkit"

// Check if PiP is supported
if !avkit.IsPictureInPictureSupported() {
    // Device doesn't support PiP
    return
}

// Get player layer from AVPlayerViewController
playerLayer := playerVC.ContentOverlayView().Layer().Sublayers()[0]

// Create PiP controller
pipController := avkit.NewPictureInPictureController(playerLayer)
pipController.SetDelegate(self)

// Enable PiP in player view controller
playerVC.SetAllowsPictureInPicturePlayback(true)
```

### Starting/Stopping PiP

```go
// Start PiP
if pipController.IsPictureInPicturePossible() {
    pipController.StartPictureInPicture()
}

// Stop PiP
pipController.StopPictureInPicture()

// Check if PiP is active
isActive := pipController.IsPictureInPictureActive()
```

### PiP Delegate Methods

```go
// Implement AVPictureInPictureControllerDelegate

func PictureInPictureControllerWillStartPictureInPicture(
    controller: AVPictureInPictureController,
) {
    // PiP will start
}

func PictureInPictureControllerDidStartPictureInPicture(
    controller: AVPictureInPictureController,
) {
    // PiP is now active
}

func PictureInPictureControllerWillStopPictureInPicture(
    controller: AVPictureInPictureController,
) {
    // PiP will stop
}

func PictureInPictureControllerDidStopPictureInPicture(
    controller: AVPictureInPictureController,
) {
    // PiP has ended
}

func PictureInPictureControllerFailedToStartPictureInPicture(
    controller: AVPictureInPictureController,
    error: NSError,
) {
    // Handle PiP start failure
}

func PictureInPictureRestoreUserInterfaceForStop(
    controller: AVPictureInPictureController,
    completion: (Bool) -> Void,
) {
    // Restore full app UI
    // Call completion(true) when ready
}
```

### PiP Best Practices

1. **Check Support**: Always check if PiP is available
2. **Handle Restoration**: Implement restore delegate method
3. **Continue Playback**: Keep playing when app backgrounds
4. **Audio Session**: Configure for background audio
5. **Resource Management**: Clean up when PiP ends

## AirPlay Integration

### Built-in AirPlay Button

```go
// AirPlay button is automatically included in AVPlayerViewController
// User can tap it to select AirPlay devices

// Disable AirPlay
playerVC.SetAllowsAirPlayPlayback(false) // Not available on AVPlayerViewController

// For AVPlayer (AVFoundation level)
player.SetAllowsExternalPlayback(false) // Disables AirPlay
```

### Custom AirPlay Button

```go
// Create route picker view
routePickerView := avkit.NewRoutePickerView()
routePickerView.SetTintColor(customColor)

// Add to view hierarchy
view.AddSubview(routePickerView)
routePickerView.SetFrame(rect)

// Or create route picker button
routePickerButton := routePickerView.RoutePickerButton()
```

### Detecting AirPlay

```go
// Observe external playback active
player.AddObserver(self, forKeyPath: "externalPlaybackActive") {
    if player.IsExternalPlaybackActive() {
        // Video is playing on AirPlay device
    } else {
        // Video is playing locally
    }
}
```

## Delegate Methods

### AVPlayerViewController Delegate

```go
// Implement AVPlayerViewControllerDelegate

func PlayerViewControllerWillStartPictureInPicture(
    controller: AVPlayerViewController,
) {
    // PiP starting
}

func PlayerViewControllerDidStartPictureInPicture(
    controller: AVPlayerViewController,
) {
    // PiP started
}

func PlayerViewControllerWillStopPictureInPicture(
    controller: AVPlayerViewController,
) {
    // PiP stopping
}

func PlayerViewControllerDidStopPictureInPicture(
    controller: AVPlayerViewController,
) {
    // PiP stopped
}

func PlayerViewControllerShouldAutomaticallyDismissAtPictureInPictureStart(
    controller: AVPlayerViewController,
) -> Bool {
    // Return true to auto-dismiss when PiP starts
    return true
}

// Full-screen callbacks
func PlayerViewControllerWillBeginFullScreenPresentation(
    controller: AVPlayerViewController,
) {
    // About to enter full-screen
}

func PlayerViewControllerDidBeginFullScreenPresentation(
    controller: AVPlayerViewController,
) {
    // Entered full-screen
}

func PlayerViewControllerWillEndFullScreenPresentation(
    controller: AVPlayerViewController,
) {
    // About to exit full-screen
}

func PlayerViewControllerDidEndFullScreenPresentation(
    controller: AVPlayerViewController,
) {
    // Exited full-screen
}
```

## Now Playing Integration

### Update Now Playing Info

```go
// AVPlayerViewController automatically updates Now Playing
// To customize, set metadata on AVPlayerItem

playerItem := player.CurrentItem()

// Create metadata items
var metadataItems []AVMetadataItem

// Title
titleItem := avfoundation.NewMutableMetadataItem()
titleItem.SetIdentifier(.commonIdentifierTitle)
titleItem.SetValue("Video Title")

// Artist
artistItem := avfoundation.NewMutableMetadataItem()
artistItem.SetIdentifier(.commonIdentifierArtist)
artistItem.SetValue("Artist Name")

// Artwork
artworkItem := avfoundation.NewMutableMetadataItem()
artworkItem.SetIdentifier(.commonIdentifierArtwork)
artworkItem.SetValue(imageData)

metadataItems = [titleItem, artistItem, artworkItem]
playerItem.SetExternalMetadata(metadataItems)

// Enable Now Playing updates
playerVC.SetUpdatesNowPlayingInfoCenter(true)
```

## Advanced Features

### Custom Content Overlays

```go
// Add custom views over video content
overlayView := playerVC.ContentOverlayView()

// Add custom controls
customButton := UIButton()
overlayView.AddSubview(customButton)

// Note: Overlays appear above video, below standard controls
```

### Programmatic Full-Screen

```go
// Enter full-screen (iOS/iPadOS)
// Not directly available, use modalPresentationStyle

// Exit full-screen
playerVC.DismissViewController(animated: true, nil)
```

### Playback Speed

```go
// Get playback menu controller (iOS 15+)
// Allows user to change playback speed

// Standard speeds: 0.5x, 1.0x, 1.25x, 1.5x, 2.0x
// User selects from built-in menu
```

## Platform-Specific Features

### iOS/iPadOS

```go
// SharePlay (iOS 15+)
// Automatically supported when using AVPlayerViewController
// Users can start SharePlay from system controls

// Control Center integration
// Automatic when updatesNowPlayingInfoCenter = true

// Lock screen controls
// Automatic with Now Playing info
```

### macOS

```go
// AVPlayerView specific features
playerView := avkit.NewPlayerView()

// Control style
playerView.SetControlsStyle(.inline)   // Standard controls
playerView.SetControlsStyle(.floating)  // Floating controls
playerView.SetControlsStyle(.minimal)   // Minimal controls
playerView.SetControlsStyle(.none)      // No controls

// Full-screen toggle
playerView.SetShowsFullScreenToggleButton(true)

// Sharing button
playerView.SetShowsSharingServiceButton(true)

// Frame stepping
playerView.SetShowsFrameSteppingButtons(true)

// Timeline
playerView.SetShowsTimecodes(true)
```

### tvOS

```go
// AVPlayerViewController optimized for tvOS
// Siri Remote navigation built-in
// Info panel shows metadata
// Top Shelf integration for featured content
```

## Accessibility

### Built-in Accessibility

AVKit automatically provides:
- VoiceOver support for all controls
- Accessibility labels on buttons
- Keyboard navigation (macOS)
- Closed caption support
- Subtitle support

### Custom Accessibility Labels

```go
// Customize accessibility
playerVC.View().SetAccessibilityLabel("Main video player")
playerVC.View().SetAccessibilityHint("Double tap to play or pause")
```

### Subtitles and Captions

```go
// Subtitles are automatically available if present in media
// User can toggle from standard controls

// Check if closed captions are available
mediaOptions := playerItem.Asset().AvailableMediaCharacteristics(
    withMediaCharacteristic: .legible,
)

// Programmatically select subtitle track
legibleGroup := playerItem.Asset().MediaSelectionGroup(
    forMediaCharacteristic: .legible,
)
if legibleGroup != nil {
    options := legibleGroup.Options()
    // Select option
    playerItem.Select(options[0], in: legibleGroup)
}
```

## Best Practices

### Memory Management

```go
// Pause and clean up when done
player.Pause()
playerVC.SetPlayer(nil)

// Remove observers
player.RemoveTimeObserver(timeObserver)

// Release PiP controller
pipController.SetDelegate(nil)
```

### Performance

```go
// Preload video
playerItem.SetPreferredForwardBufferDuration(30.0) // 30 seconds

// Use appropriate quality
playerItem.SetPreferredPeakBitRate(5_000_000) // 5 Mbps

// Enable automatic waiting
player.SetAutomaticallyWaitsToMinimizeStalling(true)
```

### Error Handling

```go
// Observe player item status
playerItem.AddObserver(self, forKeyPath: "status") {
    switch playerItem.Status() {
    case .readyToPlay:
        // Ready to play
    case .failed:
        // Handle error
        error := playerItem.Error()
        // Show error to user
    }
}
```

### Audio Session (iOS)

```go
// Configure for video playback
audioSession := avfoundation.SharedAudioSession()
audioSession.SetCategory(.playback, mode: .moviePlayback)
audioSession.SetActive(true)

// For PiP, use .playAndRecord or enable background modes
```

## Common Patterns

### Modal Video Player

```go
// Present player modally
let playerVC = AVPlayerViewController()
playerVC.player = player
present(playerVC, animated: true) {
    player.play()
}

// Dismiss when done
NotificationCenter.default.addObserver(
    forName: .AVPlayerItemDidPlayToEndTime,
) { _ in
    playerVC.dismiss(animated: true)
}
```

### Embedded Player

```go
// Add as child view controller
addChild(playerVC)
containerView.addSubview(playerVC.view)
playerVC.view.frame = containerView.bounds
playerVC.didMove(toParent: self)

// Remove when done
playerVC.willMove(toParent: nil)
playerVC.view.removeFromSuperview()
playerVC.removeFromParent()
```

### Picture-in-Picture Background Playback

```go
// Enable background modes in app capabilities
// Enable PiP
playerVC.allowsPictureInPicturePlayback = true

// Configure audio session
let session = AVAudioSession.sharedInstance()
try session.setCategory(.playback, mode: .moviePlayback)
try session.setActive(true)

// Start PiP before app backgrounds
pipController.startPictureInPicture()
```

## Use Cases

### Video Player App
- Full-featured video playback
- Playlist support
- Download and offline viewing
- Streaming video support

### Educational Content
- Course videos
- Chapter navigation
- Speed control for lectures
- Subtitles for accessibility

### Live Streaming
- Sports events
- News broadcasts
- Webinars
- Conference talks

### Video Preview
- Gallery apps
- File browsers
- Social media feeds
- Video thumbnails with playback

### Background Video (PiP)
- Multitasking while watching
- Video calls with other apps
- Tutorial videos while coding
- Sports highlights while browsing

## Integration with Other Frameworks

### AVFoundation

AVKit builds on AVFoundation:
- Use AVPlayer for playback control
- Use AVPlayerItem for media configuration
- Use AVAsset for media inspection
- Use AVComposition for editing

### Core Media

Access media samples:
- CMTime for precise timing
- CMSampleBuffer for raw data
- CMTimeRange for time ranges

### Photos Framework

Play videos from Photos library:
```go
asset := PHAsset.fetchAssets(...).firstObject
PHImageManager.default().requestPlayerItem(forVideo: asset) { item, info in
    player.replaceCurrentItem(with: item)
}
```

## References

- [AVKit Documentation](https://developer.apple.com/documentation/avkit)
- [AVPlayerViewController](https://developer.apple.com/documentation/avkit/avplayerviewcontroller)
- [AVPlayerView (macOS)](https://developer.apple.com/documentation/avkit/avplayerview)
- [Picture in Picture](https://developer.apple.com/documentation/avkit/adopting_picture_in_picture_in_a_custom_player)
- [AVFoundation Integration](https://developer.apple.com/documentation/avfoundation)
