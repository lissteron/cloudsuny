package images

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/lissteron/simplerr"
	"github.com/loghole/tracing"
	"github.com/loghole/tracing/tracelog"

	"github.com/lissteron/cloudsuny/internal/app/codes"
	"github.com/lissteron/cloudsuny/internal/app/domain"
)

const (
	imageFolder = "images"
	maxRandInt  = 31
)

type Service struct {
	dir    string
	logger tracelog.Logger
	rnd    *rand.Rand
}

func NewService(dir string, logger tracelog.Logger) *Service {
	return &Service{
		dir:    dir,
		logger: logger,
		rnd:    rand.New(rand.NewSource(time.Now().UnixNano())), // nolint:gosec //name rand
	}
}

func (s *Service) UploadImage(ctx context.Context, img *domain.Image) (*domain.Image, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	name := fmt.Sprintf("%s.%s", s.generateRandomID(), img.Format)
	img.Name = path.Join(imageFolder, name)

	file, err := os.Create(path.Join(s.dir, name))
	if err != nil {
		s.logger.Errorf(ctx, "open file failed: %v", err)

		return nil, simplerr.WrapWithCode(err, codes.SystemError, "open file failed")
	}

	defer file.Close()

	if _, err := file.Write(img.Bytes); err != nil {
		s.logger.Errorf(ctx, "encode image failed: %v", err)

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
	buf = append(buf, strconv.FormatInt(s.rnd.Int63n(maxRandInt), 32)...)

	return string(buf)
}
