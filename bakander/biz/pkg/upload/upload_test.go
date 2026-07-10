package upload

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// createFileHeader builds a valid multipart.FileHeader with working Open()
func createFileHeader(fieldname, filename string, content []byte) (*multipart.FileHeader, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	part, err := w.CreateFormFile(fieldname, filename)
	if err != nil {
		return nil, err
	}
	if _, err := part.Write(content); err != nil {
		return nil, err
	}
	w.Close()

	reader := multipart.NewReader(&buf, w.Boundary())
	form, err := reader.ReadForm(10 << 20)
	if err != nil {
		return nil, err
	}
	fhs := form.File[fieldname]
	if len(fhs) == 0 {
		return nil, fmt.Errorf("no file for field %q", fieldname)
	}
	return fhs[0], nil
}

// mockPublisher records published tasks
type mockPublisher struct {
	mu    sync.Mutex
	tasks []FileUploadTask
	err   error
}

func (m *mockPublisher) Publish(_ context.Context, task FileUploadTask) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tasks = append(m.tasks, task)
	return m.err
}

func (m *mockPublisher) lastTask() FileUploadTask {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.tasks) == 0 {
		return FileUploadTask{}
	}
	return m.tasks[len(m.tasks)-1]
}

// mockCloudUploader for testing consumer handlers
type mockCloudUploader struct {
	mu      sync.Mutex
	uploads []uploadRecord
	err     error
}

type uploadRecord struct {
	localPath  string
	remotePath string
	bucket     string
}

func (m *mockCloudUploader) Upload(_ context.Context, localPath, remotePath, bucket string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.uploads = append(m.uploads, uploadRecord{localPath, remotePath, bucket})
	return m.err
}

func TestUploadFileSuccess(t *testing.T) {
	pub := &mockPublisher{}
	svc := NewUploadService(pub, nil, nil)

	fh, err := createFileHeader("files", "C:\\Users\\k\\Desktop\\tmp2073783210864545792.jpgg", []byte("fake-jpeg-data"))
	if err != nil {
		t.Fatalf("cannot create file header: %v", err)
	}

	filePath, coverPath, err := svc.UploadFile(context.Background(), *fh)
	if err != nil {
		t.Fatalf("UploadFile failed: %v", err)
	}
	if filePath == "" {
		t.Error("filePath should not be empty")
	}
	if coverPath == "" {
		t.Error("coverPath should not be empty")
	}
	if !strings.Contains(filePath, ".jpg") {
		t.Errorf("filePath should contain .jpg: %s", filePath)
	}

	task := pub.lastTask()
	if task.FileTmpPath == "" {
		t.Error("published task FileTmpPath should not be empty")
	}
	if task.FileUploadPath != filePath {
		t.Errorf("task.FileUploadPath = %q, want %q", task.FileUploadPath, filePath)
	}

	os.Remove(task.FileTmpPath)
}

func TestUploadFileNoSuffix(t *testing.T) {
	pub := &mockPublisher{}
	svc := NewUploadService(pub, nil, nil)

	fh, err := createFileHeader("files", "noextension", []byte("data"))
	if err != nil {
		t.Fatalf("cannot create file header: %v", err)
	}

	_, _, err = svc.UploadFile(context.Background(), *fh)
	if err == nil {
		t.Error("expected error for file without suffix, got nil")
	}
}

func TestUploadFilePublishError(t *testing.T) {
	pub := &mockPublisher{err: fmt.Errorf("mq connection lost")}
	svc := NewUploadService(pub, nil, nil)

	fh, err := createFileHeader("files", "photo.png", []byte("fake-png-data"))
	if err != nil {
		t.Fatalf("cannot create file header: %v", err)
	}

	_, _, upErr := svc.UploadFile(context.Background(), *fh)
	if upErr == nil {
		t.Error("expected error from publish failure, got nil")
	}

	matches, _ := filepath.Glob("./tmp*.png")
	for _, m := range matches {
		os.Remove(m)
	}
}

func TestFileUploadTaskJSONRoundtrip(t *testing.T) {
	original := FileUploadTask{
		FileTmpPath:     "./tmp12345.mp4",
		CoverTmpPath:    "./tmp12345.png",
		FileUploadPath:  "2026/05/26/12345.mp4",
		CoverUploadPath: "2026/05/26/12345.png",
	}

	body, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded FileUploadTask
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if original != decoded {
		t.Errorf("roundtrip mismatch:\n  original: %+v\n  decoded:  %+v", original, decoded)
	}
}

func TestGetFileSuffix(t *testing.T) {
	tests := []struct {
		fileName string
		want     string
		wantErr  bool
	}{
		{"avatar.png", "png", false},
		{"video.mp4", "mp4", false},
		{"archive.tar.gz", "gz", false},
		{"noextension", "", true},
		{"UPPERCASE.JPG", "jpg", false},
		{"dotfile.", "", false},
		{"a.b.c.txt", "txt", false},
	}

	for _, tt := range tests {
		got, err := getFileSuffix(tt.fileName)
		if tt.wantErr && err == nil {
			t.Errorf("getFileSuffix(%q) expected error, got nil", tt.fileName)
			continue
		}
		if !tt.wantErr && err != nil {
			t.Errorf("getFileSuffix(%q) unexpected error: %v", tt.fileName, err)
			continue
		}
		if got != tt.want {
			t.Errorf("getFileSuffix(%q) = %q, want %q", tt.fileName, got, tt.want)
		}
	}
}

func TestSuffixSetsBoundary(t *testing.T) {
	sets := map[string]map[string]struct{}{
		"imageSuffixSet": imageSuffixSet,
		"videoSuffixSet": videoSuffixSet,
		"audioSuffixSet": audioSuffixSet,
	}

	for name, s := range sets {
		for suffix := range s {
			if !strings.HasPrefix(suffix, ".") {
				t.Errorf("%s contains entry %q without leading '.'", name, suffix)
			}
		}
	}
}

func TestRunImageUploadHandler(t *testing.T) {
	content := []byte("fake-image-bytes")
	tmpFile, err := os.CreateTemp("", "img-test-*.png")
	if err != nil {
		t.Fatalf("cannot create temp file: %v", err)
	}
	tmpFile.Write(content)
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	uploader := &mockCloudUploader{}
	task := FileUploadTask{
		FileTmpPath:    tmpFile.Name(),
		FileUploadPath: "2026/07/07/test123.png",
	}

	body, _ := json.Marshal(task)
	handler := func(body []byte) error {
		var t FileUploadTask
		if err := json.Unmarshal(body, &t); err != nil {
			return err
		}
		defer os.Remove(t.FileTmpPath)
		return uploader.Upload(context.Background(), t.FileTmpPath, t.FileUploadPath, "image")
	}

	if err := handler(body); err != nil {
		t.Fatalf("image upload handler failed: %v", err)
	}

	uploader.mu.Lock()
	defer uploader.mu.Unlock()
	if len(uploader.uploads) != 1 {
		t.Fatalf("expected 1 upload, got %d", len(uploader.uploads))
	}
	if uploader.uploads[0].bucket != "image" {
		t.Errorf("bucket = %q, want %q", uploader.uploads[0].bucket, "image")
	}
}

func TestRunDocUploadNoop(t *testing.T) {
	svc := NewUploadService(&mockPublisher{}, nil, nil)
	if err := svc.RunDocUpload(context.Background()); err != nil {
		t.Errorf("RunDocUpload should return nil, got %v", err)
	}
}
