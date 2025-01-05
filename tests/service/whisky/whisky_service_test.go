package whisky

import (
	"context"
	"github.com/GagulProject/go-whisky/generated/models"
	"github.com/GagulProject/go-whisky/internal/model/whisky"
	whiskySvc "github.com/GagulProject/go-whisky/internal/service/whisky"
	"github.com/GagulProject/go-whisky/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WhiskyServiceTestSuite struct {
	suite.Suite
	app *server.TestServer
	svc whiskySvc.WhiskyService
}

func (s *WhiskyServiceTestSuite) SetupTest() {
	s.app = server.
		NewTestServer(server.Populate(&s.svc)).
		WithPreparedTables(models.TableNames.Whisky)
	defer func() {
		s.app.Stop()
	}()
}

func (s *WhiskyServiceTestSuite) TestCreate() {
	whiskyM, err := s.svc.Create(context.Background(), &whisky.Whisky{
		Strength: 30,
		Size:     500,
	})
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), whiskyM.Strength, 30)
}

func (s *WhiskyServiceTestSuite) TestScroll() {
	s.createMockWhiskies()

	whiskies, err := s.svc.ScrollAll(context.Background())
}

func (s *WhiskyServiceTestSuite) createMockWhiskies() {
	
}

func TestIWhiskyServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WhiskyServiceTestSuite))
}
