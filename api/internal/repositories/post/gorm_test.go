package post_repository

import (
	"database/sql"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/edwardkerckhof/goblog/internal/core/domain"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository ports.PostRepository
	post       *domain.Post
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
	s.DB, err = gorm.Open(dialector, &gorm.Config{Logger: newLogger})
	require.NoError(s.T(), err)

	s.repository = NewGormRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

var post = domain.Post{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	},
	Title: "Test title",
	Body:  "Test Body",
}

func (s *Suite) Test_repository_Get() {
	query := `SELECT * FROM "posts" WHERE "posts"."id" = $1`

	rows := s.mock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "title", "body"}).
		AddRow(post.ID, post.CreatedAt, post.UpdatedAt, post.DeletedAt, post.Title, post.Body)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(post.ID).
		WillReturnRows(rows)

	res, err := s.repository.Get(post.ID)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}
