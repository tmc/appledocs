package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/avkit"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("AVKit Framework Examples")
	fmt.Println("========================")

	// Example 1: Create AVPlayerViewController
	fmt.Println("\n1. Creating AVPlayerViewController:")

	playerViewController := avkit.NewPlayerViewController()
	fmt.Printf("   Player view controller created: %v\n", playerViewController)

	// Example 2: AVKit components
	fmt.Println("\n2. AVKit Components:")

	components := map[string]string{
		"AVPlayerViewController":      "Full-featured video player UI",
		"AVPlayerView":                "macOS video player view",
		"AVRoutePickerView":           "AirPlay device picker",
		"AVPictureInPictureController": "Picture-in-picture mode",
		"AVPlayerLayer":               "Core rendering layer (from AVFoundation)",
	}

	for component, desc := range components {
		fmt.Printf("   %-30s: %s\n", component, desc)
	}

	// Example 3: Player features
	fmt.Println("\n3. AVPlayerViewController Features:")

	features := []string{
		"Built-in playback controls",
		"Scrubbing and seeking",
		"Volume control",
		"Full-screen mode",
		"Picture-in-picture support",
		"AirPlay streaming",
		"Subtitles/closed captions",
		"Playback speed control",
		"Chapter navigation",
		"Live streaming support",
		"HDR and Dolby Vision",
		"Spatial audio support",
	}

	for i, feature := range features {
		fmt.Printf("   %2d. %s\n", i+1, feature)
	}

	// Example 4: Typical workflow
	fmt.Println("\n4. Typical Video Player Workflow:")

	workflow := []string{
		"1. Create AVPlayer with media URL (AVFoundation)",
		"2. Create AVPlayerViewController",
		"3. Set player on view controller",
		"4. Present view controller modally or embed in view",
		"5. Player UI handles all user interaction",
		"6. Optional: Observe playback state",
		"7. Optional: Customize controls appearance",
		"8. Dismiss when done",
	}

	for _, step := range workflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 5: Picture-in-Picture workflow
	fmt.Println("\n5. Picture-in-Picture Workflow:")

	pipWorkflow := []string{
		"1. Check if PiP is supported on device",
		"2. Create AVPictureInPictureController with player layer",
		"3. Set delegate for PiP events",
		"4. Call startPictureInPicture() when user taps PiP button",
		"5. Video continues playing in floating window",
		"6. User can resize, move, or restore window",
		"7. Handle delegate callbacks for state changes",
		"8. Stop PiP when video ends or user closes",
	}

	for _, step := range pipWorkflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 6: Customization options
	fmt.Println("\n6. Customization Options:")

	customization := map[string]string{
		"showsPlaybackControls":      "Show/hide player controls",
		"allowsPictureInPicturePlayback": "Enable PiP mode",
		"updatesNowPlayingInfoCenter": "Update system Now Playing info",
		"entersFullScreenWhenPlaybackBegins": "Auto full-screen",
		"exitsFullScreenWhenPlaybackEnds": "Auto exit full-screen",
		"videoGravity":               "Content scaling (aspect fill/fit)",
		"requiresLinearPlayback":     "Disable seeking/scrubbing",
	}

	for property, desc := range customization {
		fmt.Printf("   %-40s: %s\n", property, desc)
	}

	// Example 7: AirPlay integration
	fmt.Println("\n7. AirPlay Integration:")

	airplaySteps := []string{
		"1. AVKit automatically includes AirPlay button in controls",
		"2. User taps AirPlay button",
		"3. AVRoutePickerView shows available devices",
		"4. User selects Apple TV or AirPlay 2 device",
		"5. Video streams to selected device",
		"6. Playback controls remain on source device",
		"7. User can disconnect from AirPlay button",
	}

	for _, step := range airplaySteps {
		fmt.Printf("   %s\n", step)
	}

	// Example 8: Playback states
	fmt.Println("\n8. Player States and Delegate Callbacks:")

	states := []string{
		"willBeginFullScreenPresentation - About to enter full-screen",
		"didBeginFullScreenPresentation - Entered full-screen",
		"willEndFullScreenPresentation - About to exit full-screen",
		"didEndFullScreenPresentation - Exited full-screen",
		"pictureInPictureWillStart - PiP starting",
		"pictureInPictureDidStart - PiP active",
		"pictureInPictureWillStop - PiP stopping",
		"pictureInPictureDidStop - PiP ended",
	}

	for i, state := range states {
		fmt.Printf("   %2d. %s\n", i+1, state)
	}

	// Example 9: Common use cases
	fmt.Println("\n9. Common Use Cases:")

	useCases := map[string]string{
		"Video Player App":     "Full-featured video playback",
		"Educational Content":  "Course videos with chapters",
		"Live Streaming":       "Sports, news, events",
		"Video Preview":        "Quick video inspection",
		"Media Browser":        "Gallery with video playback",
		"Background Video":     "PiP while using other apps",
		"AirPlay Streaming":    "Cast to Apple TV",
		"Accessibility":        "Subtitles and audio descriptions",
	}

	for useCase, desc := range useCases {
		fmt.Printf("   %-20s: %s\n", useCase, desc)
	}

	// Example 10: Platform differences
	fmt.Println("\n10. Platform-Specific Features:")

	platforms := map[string][]string{
		"iOS/iPadOS": {
			"AVPlayerViewController",
			"Native full-screen mode",
			"Picture-in-picture",
			"SharePlay integration",
			"Control Center integration",
		},
		"macOS": {
			"AVPlayerView (native AppKit)",
			"Window-based full-screen",
			"Picture-in-picture",
			"Touch Bar support",
			"Menu bar integration",
		},
		"tvOS": {
			"AVPlayerViewController optimized for TV",
			"Siri Remote navigation",
			"Info panel with metadata",
			"Top Shelf integration",
		},
	}

	for platform, items := range platforms {
		fmt.Printf("\n   %s:\n", platform)
		for _, item := range items {
			fmt.Printf("     • %s\n", item)
		}
	}

	// Example 11: Video formats supported
	fmt.Println("\n11. Supported Video Formats:")

	formats := []string{
		"H.264/AVC - Standard HD video",
		"HEVC/H.265 - 4K and HDR video",
		"ProRes - Professional video editing",
		"Dolby Vision - HDR with dynamic metadata",
		"HDR10 - Standard HDR",
		"HLS - Adaptive streaming protocol",
		"MPEG-4 - Legacy format",
	}

	for i, format := range formats {
		fmt.Printf("   %2d. %s\n", i+1, format)
	}

	// Example 12: Accessibility features
	fmt.Println("\n12. Accessibility Features:")

	accessibility := []string{
		"VoiceOver support for all controls",
		"Closed captions (CC) support",
		"Subtitles for the Deaf and Hard of Hearing (SDH)",
		"Audio descriptions",
		"Keyboard navigation (macOS)",
		"Reduced motion support",
		"High contrast mode",
		"Accessibility labels on all controls",
	}

	for i, feature := range accessibility {
		fmt.Printf("   %2d. %s\n", i+1, feature)
	}

	fmt.Println("\n✓ AVKit framework examples completed!")
	fmt.Println("\nNote: AVKit provides ready-to-use video player UI:")
	fmt.Println("  - Full-featured playback controls")
	fmt.Println("  - Picture-in-picture support")
	fmt.Println("  - AirPlay integration")
	fmt.Println("  - Accessibility built-in")
	fmt.Println("  - Native platform integration")
	fmt.Println("\nReal applications would:")
	fmt.Println("  - Create AVPlayer with video URL")
	fmt.Println("  - Set player on AVPlayerViewController")
	fmt.Println("  - Present modally or embed in view hierarchy")
	fmt.Println("  - Customize controls and behavior")
	fmt.Println("  - Implement PiP for background playback")
	fmt.Println("  - Handle full-screen transitions")
}
