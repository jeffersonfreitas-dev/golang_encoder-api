package domain_test

import (
	"testing"
	"time"

	"github.com/jeffersonfreitas-dev/encoder-api/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func Test_Validate_VideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func Test_Validade_VideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "abc"
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func Test_Validade_VideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
