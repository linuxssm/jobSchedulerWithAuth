package jobScheduler

import (
	"github.com/syndtr/goleveldb/leveldb"
 	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const defaujltDbType string = "leveldb"
type Storage struct {
	dbtype string
	DB *leveldb.DB
	SqlDb *sql.DB
}

type DatabaseConfig struct {
	dbtype string "defaujltDbType"
	isSql bool
}
var dbConfig DatabaseConfig

func ConfigDatabase(_dbtype string) *DatabaseConfig{
	//dbConfig = &DatabaseConfig{}

	if _dbtype == "mysql" {
		dbConfig.isSql = true
	}
	dbConfig.dbtype = _dbtype

	return &dbConfig
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func mysqlCreateTable(db *sql.DB, dbname string) error{
	_,err := db.Exec("CREATE DATABASE "+dbname)
	if err != nil {
//		panic(err)
	}

	_,err = db.Exec("USE "+dbname)
	if err != nil {
	//	panic(err)
	}

	_,err = db.Exec("CREATE TABLE if not exists job (uid INT(10) NOT NULL AUTO_INCREMENT,ID VARCHAR(64) NULL DEFAULT NULL,"+
	"Expression VARCHAR(64),"+
	"Endpoint VARCHAR(64),"+
	"Payload VARCHAR(64)," +
	"PRIMARY KEY (uid))")
	if err != nil {
		panic(err)
	}
	return err
}

func NewStorage(dbname string) (*Storage, error) {
	if dbConfig.isSql{
		db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/")
		checkErr(err)
		err = mysqlCreateTable(db,"jobScheduler")

		return &Storage{SqlDb: db}, nil
	}else{
		db, err := leveldb.OpenFile(dbname, nil)
		if err != nil {
			return nil, err
		}
		return &Storage{DB: db}, nil
	}
}

func (s *Storage) Close() {
	if dbConfig.isSql{
		s.SqlDb.Close()
	}else{
		s.DB.Close()
	}

}

// Insert or Put
//ID         string
//Expression string
//Endpoint   string
//Payload    string
func (s *Storage) SaveEntry(e *Entry) error {
	if dbConfig.isSql{
		stmt, err := s.SqlDb.Prepare("INSERT job SET ID=?,Expression=?,Endpoint=?,Payload=?")
		checkErr(err)
		_, err = stmt.Exec(e.ID, e.Expression, e.Endpoint, e.Payload)
		checkErr(err)
		return nil
	}else{
		err := s.DB.Put([]byte(e.ID), e.Bytes(), nil)
		if err != nil {
			return err
		}
		return nil

	}
}

func (s *Storage) AllEntries() []*Entry {
	if dbConfig.isSql{
		result := make([]*Entry, 0)
		// query
		rows, err := s.SqlDb.Query("SELECT * FROM job")
		checkErr(err)

		var uid int
		for rows.Next() {
			entry := &Entry{}
			err = rows.Scan(&uid, &entry.ID, &entry.Expression, &entry.Endpoint, &entry.Payload)
			checkErr(err)
			fmt.Println(entry)
			result = append(result, entry)
		}

		return result

	}else{
		result := make([]*Entry, 0)
		iter := s.DB.NewIterator(nil, nil)
		for iter.Next() {
			entry, _ := NewEntryFromBytes(iter.Value())
			result = append(result, entry)
		}
		return result

	}
}

func (s *Storage) GetEntry(id string) *Entry {
	if dbConfig.isSql{
		rows, err := s.SqlDb.Query("SELECT * FROM job where Id = ?", id)
		checkErr(err)
		entry := &Entry{}

		for rows.Next() {
			err = rows.Scan(&entry.ID, &entry.Expression, &entry.Endpoint, &entry.Payload)
			checkErr(err)
			fmt.Println(entry)
		}
		return entry

	}else{
		data, err := s.DB.Get([]byte(id), nil)
		if err != nil {
			return nil
		}
		entry, err := NewEntryFromBytes(data)
		if err != nil {
			return nil
		}
		return entry
	}

}

func (s *Storage) DeleteEntry(id string) error {
	if dbConfig.isSql{
		stmt, err := s.SqlDb.Prepare("delete from job where Id=?")
		checkErr(err)

		res, err := stmt.Exec(id)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)
		fmt.Println(affect)
		return nil

	}else{
		err := s.DB.Delete([]byte(id), nil)
		if err != nil {
			return err
		}
		return nil
	}

}
