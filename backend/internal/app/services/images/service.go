package images

import (
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/lissteron/cloudsuny/internal/app/domain"
)

const (
	maxRandInt = 31
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
}

type Service struct {
	imagesPath string
	rnd        rand.Source
	logger     Logger
}

func NewService(path string) *Service {
	return &Service{
		imagesPath: path,
		rnd:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *Service) UploadImage(data []byte) (*domain.Image, error) {
	imageID := s.generateRandomID()

	file, err := os.Create(path.Join(s.imagesPath, imageID))
	if err != nil {
		return nil, err
	}

	if _, err := file.Write(data); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Service) GetStoragePath() string {
	return s.imagesPath
}

func (s *Service) generateRandomID() string {
	buf := make([]byte, 0)
	buf = append(buf, strconv.FormatInt(time.Now().UnixNano(), 32)...)
	buf = append(buf, strconv.FormatInt(rand.New(s.rnd).Int63n(maxRandInt), 32)...)

	return string(buf)
}
