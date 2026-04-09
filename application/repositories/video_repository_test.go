package repositories_test

import (
	"testing"
	"time"

	"github.com/jeffersonfreitas-dev/encoder-api/application/repositories"
	"github.com/jeffersonfreitas-dev/encoder-api/domain"
	"github.com/jeffersonfreitas-dev/encoder-api/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func Test_VideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
