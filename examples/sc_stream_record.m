#import <Foundation/Foundation.h>
#import <ScreenCaptureKit/ScreenCaptureKit.h>
#import <CoreGraphics/CoreGraphics.h>
#import <CoreMedia/CoreMedia.h>
#import <CoreVideo/CoreVideo.h>
#import <CoreImage/CoreImage.h>
#import <ImageIO/ImageIO.h>
#import <UniformTypeIdentifiers/UniformTypeIdentifiers.h>

// Simple delegate to receive frames
@interface StreamOutputDelegate : NSObject <SCStreamOutput>
@property (atomic) int frameCount;
@property (copy) NSString *outputDir;
@property (atomic) BOOL saveFrames;
@end

@implementation StreamOutputDelegate

- (instancetype)initWithOutputDir:(NSString *)dir saveFrames:(BOOL)save {
    self = [super init];
    if (self) {
        _frameCount = 0;
        _outputDir = [dir copy];
        _saveFrames = save;
    }
    return self;
}

- (void)stream:(SCStream *)stream didOutputSampleBuffer:(CMSampleBufferRef)sampleBuffer ofType:(SCStreamOutputType)type {
    if (type != SCStreamOutputTypeScreen) {
        return;
    }

    self.frameCount++;

    if (self.frameCount % 30 == 0) {  // Log every 30 frames
        NSLog(@"Received frame %d", self.frameCount);
    }

    if (!self.saveFrames) {
        return;
    }

    // Get the pixel buffer from the sample buffer
    CVPixelBufferRef pixelBuffer = CMSampleBufferGetImageBuffer(sampleBuffer);
    if (!pixelBuffer) {
        NSLog(@"No pixel buffer in sample");
        return;
    }

    // Create CGImage from pixel buffer
    CIImage *ciImage = [CIImage imageWithCVPixelBuffer:pixelBuffer];
    CIContext *context = [CIContext context];
    CGImageRef cgImage = [context createCGImage:ciImage fromRect:ciImage.extent];

    if (!cgImage) {
        NSLog(@"Failed to create CGImage");
        return;
    }

    // Save every 30th frame as PNG
    if (self.frameCount % 30 == 0) {
        NSString *filename = [NSString stringWithFormat:@"%@/frame_%04d.png", self.outputDir, self.frameCount];
        NSURL *fileURL = [NSURL fileURLWithPath:filename];

        CGImageDestinationRef destination = CGImageDestinationCreateWithURL((__bridge CFURLRef)fileURL, (__bridge CFStringRef)UTTypePNG.identifier, 1, NULL);
        if (destination) {
            CGImageDestinationAddImage(destination, cgImage, NULL);
            CGImageDestinationFinalize(destination);
            CFRelease(destination);
            NSLog(@"Saved frame to: %@", filename);
        }
    }

    CGImageRelease(cgImage);
}

@end

int main(int argc, const char * argv[]) {
    @autoreleasepool {
        NSLog(@"=== SCStream Recording Test ===\n");

        // Parse arguments
        BOOL saveFrames = NO;
        int duration = 5;

        for (int i = 1; i < argc; i++) {
            if (strcmp(argv[i], "--save") == 0) {
                saveFrames = YES;
            } else if (strcmp(argv[i], "--duration") == 0 && i + 1 < argc) {
                duration = atoi(argv[i + 1]);
                i++;
            }
        }

        // Create output directory if saving frames
        NSString *outputDir = nil;
        if (saveFrames) {
            outputDir = [NSString stringWithFormat:@"%@/sc_frames", NSTemporaryDirectory()];
            [[NSFileManager defaultManager] createDirectoryAtPath:outputDir
                                       withIntermediateDirectories:YES
                                                        attributes:nil
                                                             error:nil];
            NSLog(@"Output directory: %@\n", outputDir);
        }

        // First, get shareable content to get a display
        dispatch_semaphore_t contentSem = dispatch_semaphore_create(0);
        __block SCShareableContent *shareableContent = nil;
        __block NSError *contentError = nil;

        NSLog(@"Getting shareable content...");
        [SCShareableContent getShareableContentWithCompletionHandler:^(SCShareableContent *content, NSError *error) {
            if (error) {
                contentError = error;
            } else {
                shareableContent = [content retain];
            }
            dispatch_semaphore_signal(contentSem);
        }];

        dispatch_semaphore_wait(contentSem, DISPATCH_TIME_FOREVER);

        if (contentError) {
            NSLog(@"Error getting content: %@", contentError.localizedDescription);
            return 1;
        }

        if (!shareableContent || shareableContent.displays.count == 0) {
            NSLog(@"No displays found");
            return 1;
        }

        SCDisplay *display = shareableContent.displays.firstObject;
        NSLog(@"Recording display: %u (%lu x %lu)", display.displayID,
              (unsigned long)display.width, (unsigned long)display.height);

        // Create content filter
        NSArray *emptyArray = @[];
        SCContentFilter *filter = [[SCContentFilter alloc] initWithDisplay:display
                                                       excludingApplications:emptyArray
                                                          exceptingWindows:emptyArray];

        // Create stream configuration
        SCStreamConfiguration *config = [[SCStreamConfiguration alloc] init];
        config.width = display.width;
        config.height = display.height;
        config.pixelFormat = kCVPixelFormatType_32BGRA;
        config.queueDepth = 5;
        config.showsCursor = YES;

        // Create delegate
        StreamOutputDelegate *delegate = [[StreamOutputDelegate alloc] initWithOutputDir:outputDir saveFrames:saveFrames];

        // Create stream
        NSError *streamError = nil;
        SCStream *stream = [[SCStream alloc] initWithFilter:filter
                                              configuration:config
                                                   delegate:nil];

        if (streamError) {
            NSLog(@"Error creating stream: %@", streamError.localizedDescription);
            return 1;
        }

        // Create a dispatch queue for the stream output
        dispatch_queue_t outputQueue = dispatch_queue_create("com.example.scstream.output",
                                                             DISPATCH_QUEUE_SERIAL);

        // Add stream output
        [stream addStreamOutput:delegate
                           type:SCStreamOutputTypeScreen
            sampleHandlerQueue:outputQueue
                         error:&streamError];

        if (streamError) {
            NSLog(@"Error adding stream output: %@", streamError.localizedDescription);
            return 1;
        }

        // Start streaming
        dispatch_semaphore_t startSem = dispatch_semaphore_create(0);
        [stream startCaptureWithCompletionHandler:^(NSError *error) {
            if (error) {
                NSLog(@"Error starting capture: %@", error.localizedDescription);
            } else {
                NSLog(@"✓ Capture started!");
            }
            dispatch_semaphore_signal(startSem);
        }];

        dispatch_semaphore_wait(startSem, DISPATCH_TIME_FOREVER);

        // Record for specified duration
        NSLog(@"Recording for %d seconds...\n", duration);
        [NSThread sleepForTimeInterval:duration];

        // Stop streaming
        dispatch_semaphore_t stopSem = dispatch_semaphore_create(0);
        [stream stopCaptureWithCompletionHandler:^(NSError *error) {
            if (error) {
                NSLog(@"Error stopping capture: %@", error.localizedDescription);
            } else {
                NSLog(@"\n✓ Capture stopped!");
            }
            dispatch_semaphore_signal(stopSem);
        }];

        dispatch_semaphore_wait(stopSem, DISPATCH_TIME_FOREVER);

        NSLog(@"\n=== Recording Complete ===");
        NSLog(@"Total frames received: %d", delegate.frameCount);
        NSLog(@"Frame rate: %.1f fps", (float)delegate.frameCount / duration);

        if (saveFrames) {
            NSLog(@"Frames saved to: %@", outputDir);
            NSLog(@"Saved every 30th frame as PNG");
        }

        [stream release];
        [delegate release];
        [filter release];
        [config release];
        [shareableContent release];

        NSLog(@"\n=== Test PASSED ===");
    }
    return 0;
}
