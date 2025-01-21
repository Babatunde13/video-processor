# Video Processor
This API uses ffmpeg to stream video with low latency.

# Codebase Setup
1. Clone this repository using the comand below
```bash
git clone https://github.com/Babatunde13/video-processor.git
```

**If you have ffmpeg locally you can run this app directly, otherwise you can use the docker setup to start the app.**

2. Install the dependencies
```bash
go mod download
```
3. Run application
```bash
go run cmd/main.go
```

**If you do not have ffmpeg locally, you can download it manually and use the step above to start the app or use docker to set it up**
```bash
docker-compose up
```


## API Endpoints

1. `POST /upload`: expects a multipart form upload which saves the file to `./srt-data/output.mp4`
2. `POST /start-processing`: this starts the processing and generates the `.m3u8` and `.ts` files
3. `GET /hls/stream`: this streams the `.m3u8` and `.ts` files
