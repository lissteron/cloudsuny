package images

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/lissteron/simplerr"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

const (
	imageFolder = "images"
	maxRandInt  = 31
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
	dir    string
	logger Logger
	rnd    rand.Source
}

func NewService(dir string, logger Logger) *Service {
	return &Service{
		dir:    dir,
		logger: logger,
		rnd:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *Service) UploadImage(img *domain.Image) (*domain.Image, error) {
	name := fmt.Sprintf("%s.%s", s.generateRandomID(), img.Format)
	img.Name = path.Join(imageFolder, name)

	file, err := os.Create(path.Join(s.dir, name))
	if err != nil {
		s.logger.Errorf("open file failed: %v", err)
		return nil, simplerr.WrapWithCode(err, codes.SystemError, "open file failed")
	}

	defer file.Close()

	if _, err := file.Write(img.Bytes); err != nil {
		s.logger.Errorf("encode image failed: %v", err)
		return nil, simplerr.WrapWithCode(err, codes.SystemError, "write image failed")
	}

	return img, nil
}

func (s *Service) GetStoragePath() string {
	return s.dir
}

func (s *Service) generateRandomID() string {
	buf := make([]byte, 0)
	buf = append(buf, strconv.FormatInt(time.Now().UnixNano(), 32)...)
	buf = append(buf, strconv.FormatInt(rand.New(s.rnd).Int63n(maxRandInt), 32)...)

	return string(buf)
}
