FROM golang:1.22.2-bullseye AS builder

WORKDIR /app

# Install FFmpeg
RUN apt-get update && \
    apt-get install -y ffmpeg && \
    apt-get clean

COPY . .
RUN go mod download
RUN go build -o /app/stream -x cmd/main.go

FROM debian:bullseye-slim
# copy the binary from builder then run it as an executbale

RUN apt-get update && \
    apt-get install -y ffmpeg && \
    apt-get clean

COPY --from=builder /app/stream /app/stream
COPY --from=builder /app/srt-data /app/srt-data

EXPOSE 8080
EXPOSE 9000
CMD ["/app/stream"]
