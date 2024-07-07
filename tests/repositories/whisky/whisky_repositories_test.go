package whisky

import (
	"context"
	"github.com/GagulProject/go-whisky/generated/models"
	"github.com/GagulProject/go-whisky/internal/model/whisky"
	whiskyR "github.com/GagulProject/go-whisky/internal/repository/whisky"
	"github.com/GagulProject/go-whisky/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WhiskyRepositoriesTestSuite struct {
	suite.Suite
	app  *server.TestServer
	repo whiskyR.WhiskyRepositories
}

func (s *WhiskyRepositoriesTestSuite) SetupTest() {
	s.app = server.
		NewTestServer(server.Populate(&s.repo)).
		WithPreparedTables(models.TableNames.Whisky)
	defer func() {
		s.app.Stop()
	}()
}

func (s *WhiskyRepositoriesTestSuite) TestServer() {
	whiskyM, err := s.repo.Create(context.Background(), &whisky.Whisky{
		Strength: 30,
		Size:     500,
	})
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), whiskyM.Strength, 30)
}

func TestIWhiskyRepositoriesTestSuite(t *testing.T) {
	suite.Run(t, new(WhiskyRepositoriesTestSuite))
}
