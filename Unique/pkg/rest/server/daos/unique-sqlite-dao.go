package daos

import (
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/Unique/unique/pkg/rest/server/daos/clients/sqls"
	"github.com/unnagirirao/Unique/unique/pkg/rest/server/models"
)

type UniqueDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateUniques(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS uniques(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Unique TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewUniqueDao() (*UniqueDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateUniques(sqlClient)
	if err != nil {
		return nil, err
	}
	return &UniqueDao{
		sqlClient,
	}, nil
}

func (uniqueDao *UniqueDao) CreateUnique(m *models.Unique) (*models.Unique, error) {
	insertQuery := "INSERT INTO uniques(Unique)values(?)"
	res, err := uniqueDao.sqlClient.DB.Exec(insertQuery, m.Unique)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("unique created")
	return m, nil
}

func (uniqueDao *UniqueDao) UpdateUnique(id int64, m *models.Unique) (*models.Unique, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	unique, err := uniqueDao.GetUnique(id)
	if err != nil {
		return nil, err
	}
	if unique == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE uniques SET Unique = ? WHERE Id = ?"
	res, err := uniqueDao.sqlClient.DB.Exec(updateQuery, m.Unique, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("unique updated")
	return m, nil
}

func (uniqueDao *UniqueDao) DeleteUnique(id int64) error {
	deleteQuery := "DELETE FROM uniques WHERE Id = ?"
	res, err := uniqueDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("unique deleted")
	return nil
}

func (uniqueDao *UniqueDao) ListUniques() ([]*models.Unique, error) {
	selectQuery := "SELECT * FROM uniques"
	rows, err := uniqueDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var uniques []*models.Unique
	for rows.Next() {
		m := models.Unique{}
		if err = rows.Scan(&m.Id, &m.Unique); err != nil {
			return nil, err
		}
		uniques = append(uniques, &m)
	}
	if uniques == nil {
		uniques = []*models.Unique{}
	}

	log.Debugf("unique listed")
	return uniques, nil
}

func (uniqueDao *UniqueDao) GetUnique(id int64) (*models.Unique, error) {
	selectQuery := "SELECT * FROM uniques WHERE Id = ?"
	row := uniqueDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Unique{}
	if err := row.Scan(&m.Id, &m.Unique); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("unique retrieved")
	return &m, nil
}
