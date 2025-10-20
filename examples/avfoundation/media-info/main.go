package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/avfoundation"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("AVFoundation Framework Examples")
	fmt.Println("===============================")

	// Example 1: Create AVPlayer
	fmt.Println("\n1. Creating AVPlayer:")

	player := avfoundation.NewPlayer()
	fmt.Printf("   Player created: %v\n", player)

	// Example 2: AVFoundation capabilities
	fmt.Println("\n2. AVFoundation Capabilities:")

	capabilities := []string{
		"Audio/Video Playback",
		"Audio/Video Recording",
		"Camera Capture",
		"Media Editing",
		"Media Composition",
		"Asset Management",
		"Media Export",
		"Speech Synthesis",
		"Speech Recognition",
		"Spatial Audio",
		"Video Effects",
		"Live Streaming",
	}

	for i, cap := range capabilities {
		fmt.Printf("   %2d. %s\n", i+1, cap)
	}

	// Example 3: Core classes
	fmt.Println("\n3. Core AVFoundation Classes:")

	coreClasses := map[string]string{
		"AVPlayer":            "Playback controller for media",
		"AVPlayerItem":        "Media asset being played",
		"AVAsset":             "Media resource (file, stream, etc.)",
		"AVCaptureSession":    "Capture pipeline coordinator",
		"AVCaptureDevice":     "Hardware capture device (camera, mic)",
		"AVCaptureInput":      "Capture device input",
		"AVCaptureOutput":     "Capture output (video, photo, etc.)",
		"AVAssetExportSession": "Media export/transcoding",
		"AVComposition":       "Timeline-based media editing",
		"AVAudioEngine":       "Audio processing graph",
	}

	for class, desc := range coreClasses {
		fmt.Printf("   %-25s: %s\n", class, desc)
	}

	// Example 4: Media playback workflow
	fmt.Println("\n4. Media Playback Workflow:")

	playbackSteps := []string{
		"1. Create URL to media file or stream",
		"2. Create AVAsset from URL",
		"3. Create AVPlayerItem from asset",
		"4. Create AVPlayer with player item",
		"5. Optional: Create AVPlayerLayer for video",
		"6. Call player.play() to start playback",
		"7. Observe playback status and progress",
		"8. Handle playback completion",
	}

	for _, step := range playbackSteps {
		fmt.Printf("   %s\n", step)
	}

	// Example 5: Camera capture workflow
	fmt.Println("\n5. Camera Capture Workflow:")

	captureSteps := []string{
		"1. Create AVCaptureSession",
		"2. Get camera device (AVCaptureDevice)",
		"3. Create device input from camera",
		"4. Add input to session",
		"5. Create output (video, photo, etc.)",
		"6. Add output to session",
		"7. Set output delegate for callbacks",
		"8. Start session running",
		"9. Receive frames in delegate callbacks",
		"10. Stop session when done",
	}

	for _, step := range captureSteps {
		fmt.Printf("   %s\n", step)
	}

	// Example 6: Media formats
	fmt.Println("\n6. Supported Media Formats:")

	formats := map[string][]string{
		"Video": {
			"H.264/AVC",
			"HEVC/H.265",
			"ProRes",
			"MPEG-4",
			"QuickTime",
		},
		"Audio": {
			"AAC",
			"MP3",
			"Apple Lossless (ALAC)",
			"PCM/LPCM",
			"Opus",
		},
		"Containers": {
			"MP4",
			"MOV",
			"M4A",
			"WAV",
			"CAF",
		},
	}

	for category, items := range formats {
		fmt.Printf("\n   %s:\n", category)
		for _, item := range items {
			fmt.Printf("     • %s\n", item)
		}
	}

	// Example 7: Capture outputs
	fmt.Println("\n7. AVCapture Output Types:")

	outputs := map[string]string{
		"AVCaptureVideoDataOutput":    "Video frames (CVPixelBuffer)",
		"AVCapturePhotoOutput":        "Still photo capture",
		"AVCaptureMovieFileOutput":    "Record to movie file",
		"AVCaptureAudioDataOutput":    "Audio sample buffers",
		"AVCaptureMetadataOutput":     "Face detection, barcodes, etc.",
		"AVCaptureDepthDataOutput":    "Depth map data (TrueDepth)",
		"AVCaptureFileOutput":         "Generic file output",
	}

	for output, desc := range outputs {
		fmt.Printf("   %-30s: %s\n", output, desc)
	}

	// Example 8: Player states
	fmt.Println("\n8. AVPlayer States:")

	states := []string{
		"Unknown - Initial state",
		"ReadyToPlay - Media loaded and ready",
		"Failed - Error occurred",
		"Paused - Playback paused",
		"Playing - Currently playing",
		"Seeking - Seeking to new position",
		"Buffering - Loading data",
	}

	for i, state := range states {
		fmt.Printf("   %d. %s\n", i+1, state)
	}

	// Example 9: Common use cases
	fmt.Println("\n9. Common Use Cases:")

	useCases := map[string]string{
		"Video Player":       "Media playback app",
		"Camera App":         "Photo/video capture",
		"Video Conferencing": "Real-time video/audio",
		"Screen Recording":   "Capture screen content",
		"Live Streaming":     "RTMP/HLS streaming",
		"Media Editor":       "Video editing/composition",
		"Audio Recording":    "Voice memos, podcasts",
		"QR Scanner":         "Metadata detection",
		"AR Effects":         "Video frame processing",
		"Video Transcoding":  "Format conversion",
	}

	for useCase, desc := range useCases {
		fmt.Printf("   %-20s: %s\n", useCase, desc)
	}

	// Example 10: Integration with other frameworks
	fmt.Println("\n10. Framework Integration:")

	integrations := map[string]string{
		"Core Media":        "Media sample buffers and timing",
		"Core Video":        "Pixel buffer management",
		"Core Audio":        "Low-level audio processing",
		"Metal":             "GPU-accelerated video processing",
		"Core Image":        "Video filters and effects",
		"Vision":            "Computer vision on video frames",
		"Photos":            "Photo library integration",
		"ScreenCaptureKit":  "Screen/window recording",
		"IOSurface":         "Zero-copy frame sharing",
		"QuartzCore":        "AVPlayerLayer for display",
	}

	for framework, desc := range integrations {
		fmt.Printf("   %-20s: %s\n", framework, desc)
	}

	// Example 11: Camera device types
	fmt.Println("\n11. Camera Device Types:")

	cameras := []string{
		"Built-in Wide Angle - Standard rear camera",
		"Built-in Ultra Wide - Wide field of view",
		"Built-in Telephoto - Optical zoom camera",
		"Built-in Dual Camera - Wide + telephoto",
		"Built-in Triple Camera - Wide + ultra wide + telephoto",
		"Built-in TrueDepth - Front Face ID camera",
		"Built-in LiDAR - Depth sensing camera",
		"Continuity Camera - iPhone as Mac webcam",
		"External - USB/Thunderbolt cameras",
	}

	for i, camera := range cameras {
		fmt.Printf("   %2d. %s\n", i+1, camera)
	}

	// Example 12: Audio session categories
	fmt.Println("\n12. Audio Session Categories:")

	audioCategories := map[string]string{
		"Playback":           "Background music/video playback",
		"Record":             "Audio recording only",
		"PlayAndRecord":      "Simultaneous playback and recording",
		"MultiRoute":         "Multiple audio routes",
		"Ambient":            "Mix with other audio",
		"SoloAmbient":        "Default category",
	}

	for category, desc := range audioCategories {
		fmt.Printf("   %-20s: %s\n", category, desc)
	}

	// Example 13: Export presets
	fmt.Println("\n13. AVAssetExportSession Presets:")

	presets := []string{
		"LowQuality - Smallest file size",
		"MediumQuality - Balanced size/quality",
		"HighestQuality - Best quality",
		"640x480 - Legacy resolution",
		"960x540 - qHD resolution",
		"1280x720 - 720p HD",
		"1920x1080 - 1080p Full HD",
		"3840x2160 - 4K UHD",
		"HEVC 1920x1080 - H.265 Full HD",
		"HEVC 3840x2160 - H.265 4K",
		"AppleProRes422LPCM - Professional codec",
	}

	for i, preset := range presets {
		fmt.Printf("   %2d. %s\n", i+1, preset)
	}

	// Example 14: Key-Value Observing properties
	fmt.Println("\n14. Common KVO Properties:")

	kvoProps := []string{
		"player.status - Ready/failed state",
		"player.rate - Playback speed (0.0 = paused)",
		"player.currentItem - Active player item",
		"playerItem.status - Item ready state",
		"playerItem.duration - Media duration",
		"playerItem.loadedTimeRanges - Buffered ranges",
		"playerItem.isPlaybackLikelyToKeepUp - Buffering state",
		"asset.duration - Asset total duration",
	}

	for i, prop := range kvoProps {
		fmt.Printf("   %2d. %s\n", i+1, prop)
	}

	fmt.Println("\n✓ AVFoundation framework examples completed!")
	fmt.Println("\nNote: AVFoundation is Apple's comprehensive media framework for:")
	fmt.Println("  - Audio/video playback and recording")
	fmt.Println("  - Camera and microphone capture")
	fmt.Println("  - Media editing and composition")
	fmt.Println("  - Live streaming and real-time processing")
	fmt.Println("\nReal applications would:")
	fmt.Println("  - Create AVPlayer for media playback")
	fmt.Println("  - Setup AVCaptureSession for camera access")
	fmt.Println("  - Use AVAssetExportSession for transcoding")
	fmt.Println("  - Observe player/item status with KVO")
	fmt.Println("  - Handle audio session configuration")
	fmt.Println("  - Integrate with Metal for GPU processing")
}
