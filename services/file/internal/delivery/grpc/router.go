package grpc

import (
	pb "architecture_go/services/file/internal/delivery/grpc/interface"
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type routeGuideServer struct {
	pb.ContactServiceServer
}

func (s *routeGuideServer) DownloadVideo(ctx context.Context, req *pb.DownloadVideoRequest) (*pb.DownloadVideoResponse, error) {
	videoName := req.GetName()
	file, err := os.Open(videoName)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return nil, err
	}
	return &pb.DownloadVideoResponse{
		Response: &pb.Video{
			Name:  videoName,
			Bytes: string(bs),
		},
	}, nil
}

func (s *routeGuideServer) UploadVideo(ctx context.Context, req *pb.UploadVideoRequest) (*pb.UploadVideoResponse, error) {
	videoName := req.Request.GetName()
	bytes := req.Request.GetBytes()

	dst, err := os.Create(videoName + time.Now().String())
	defer dst.Close()
	if err != nil {
		return nil, err
	}
	file := strings.NewReader(bytes)

	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}

	return &pb.UploadVideoResponse{
		Response: &pb.Video{
			Name:  videoName,
			Bytes: bytes,
		},
	}, nil
}

func (s *routeGuideServer) DeleteVideo(ctx context.Context, req *pb.DeleteVideoRequest) (*pb.DeleteVideoResponse, error) {
	videoName := req.GetName()

	err := os.Remove(videoName)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteVideoResponse{Name: videoName}, nil
}
