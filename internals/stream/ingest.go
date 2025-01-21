package stream

import (
	"log"
	"os"
	"os/exec"
)

var (
	srtUrl     = "srt://:9000" // a latency of 250 is below our 300 max latency and is also high enough to reduce packet loss and tolerate more network jitters as per our requirement
	FolderName = "./srt-data"
)

func createDirIfNotExists(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func runFFemPegCommand(inputUrl string, outputUrl string) *exec.Cmd {
	return exec.Command(
		"ffmpeg",
		"-i", inputUrl, // input srt stream
		"-hls_time", "2", // 2 seconds segment duration
		"-hls_list_size", "10", // no limit on number of segments
		"-c:v", "libx264", // video codec encode with h264
		"-preset", "ultrafast", // ultrafast encoding, low latency
		"-tune", "zerolatency", // zero latency
		"-loglevel", "debug", // log level debug to see more info
		outputUrl, // output m3u8 file, hls playlist output
	)
}

// Takes a filename and create the srt folder if it does not exist
// then runs ffmpeg command to ingest the srt stream
// this runs ffmpeg
func Ingest(filename string) {
	err := createDirIfNotExists(FolderName)
	if err != nil {
		log.Println("Error creating folder", err)
		return
	}

	log.Println("Ingesting from", srtUrl)
	outpurUrl := FolderName + "/" + filename
	cmd := runFFemPegCommand(srtUrl, outpurUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Starting ffmpeg")
	err = cmd.Run()
	if err != nil {
		log.Println("Error running ffmpeg", err)
		return
	}

	log.Println("Ingestion complete")
}

// ffmpeg -re -i ./srt-data/output.mp4 -c:v libx264 -preset ultrafast -f mpegts srt://:9000
// ffmpeg -re -i ./srt-data/output.mp4 -c:v libx264 -preset ultrafast -f mpegts srt://127.0.0.1:9000
