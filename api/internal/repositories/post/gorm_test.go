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
	"github.com/edwardkerckhof/goblog/pkg/utils"
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
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
	s.DB, err = gorm.Open(dialector, &gorm.Config{Logger: newLogger})
	require.NoError(s.T(), err)

	s.repository = NewGormRepository(s.DB)

	s.post = &domain.Post{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Title: "Test title",
		Body:  "Test Body",
	}
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Get() {
	query := `SELECT * FROM "posts" WHERE "posts"."id" = $1`

	rows := s.mock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "title", "body"}).
		AddRow(s.post.ID, s.post.CreatedAt, s.post.UpdatedAt, s.post.DeletedAt, s.post.Title, s.post.Body)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.post.ID).
		WillReturnRows(rows)

	res, err := s.repository.Get(s.post.ID)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	requireEqual(s.T(), s.post, res)
}

func (s *Suite) Test_repository_GetAll() {
	query := `SELECT * FROM "posts"`

	rows := s.mock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "title", "body"}).
		AddRow(s.post.ID, s.post.CreatedAt, s.post.UpdatedAt, s.post.DeletedAt, s.post.Title, s.post.Body)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := s.repository.GetAll(0)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	requireEquals(s.T(), s.post, res)
}

func (s *Suite) Test_repository_Create() {
	query := `
		INSERT INTO "posts" ("created_at","updated_at","deleted_at","title","body","id")
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING "id"
	`

	rows := s.mock.NewRows([]string{"id"}).
		AddRow(s.post.ID)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.post.CreatedAt, s.post.UpdatedAt, s.post.DeletedAt, s.post.Title, s.post.Body, s.post.ID).
		WillReturnRows(rows)
	s.mock.ExpectCommit()

	res, err := s.repository.Create(s.post)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	requireEqual(s.T(), s.post, res)
}

func (s *Suite) Test_repository_Update() {
	postToUpdate := *s.post
	postToUpdate.Title = "Updated title"
	postToUpdate.Body = "Updated body"

	s.mock.ExpectBegin()
	s.mock.ExpectExec("UPDATE").
		WithArgs(postToUpdate.CreatedAt, utils.AnyTime{}, gorm.DeletedAt{}, postToUpdate.Title, postToUpdate.Body, postToUpdate.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	res, err := s.repository.Update(&postToUpdate)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	requireEqual(s.T(), &postToUpdate, res)
}

func (s *Suite) Test_repository_Delete() {
	postToDelete := *s.post
	postToDelete.ID = 2
	softDeleteQuery := `UPDATE "posts" SET "deleted_at"=$1 WHERE "posts"."id" = $2`

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(softDeleteQuery)).
		WithArgs(utils.AnyTime{}, postToDelete.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	s.repository.Delete(&postToDelete)
}

func requireEqual(t *testing.T, post *domain.Post, res *domain.Post) {
	require.Equal(t, post.ID, res.ID)
	require.Equal(t, post.Body, res.Body)
	require.Equal(t, post.Title, res.Title)
	require.Equal(t, post.CreatedAt, res.CreatedAt)
	require.Equal(t, post.UpdatedAt, res.UpdatedAt)
	require.Equal(t, post.DeletedAt, res.DeletedAt)
}

func requireEquals(t *testing.T, post *domain.Post, res []*domain.Post) {
	for _, result := range res {
		requireEqual(t, post, result)
	}
}
