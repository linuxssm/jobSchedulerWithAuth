package api

import (
	"github.com/syndtr/goleveldb/leveldb"
 	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/linuxssm/frontendServer/conf"
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
		db, err := sql.Open("mysql", conf.DB_AUTH_OPEN_INFO)
		checkErr(err)
		//err = mysqlCreateTable(db,"jobScheduler")

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
func (s *Storage) SaveEntry(e *User) error {
	if dbConfig.isSql{
		stmt, err := s.SqlDb.Prepare("INSERT job SET ID=?,Expression=?,Endpoint=?,Payload=?")
		checkErr(err)
		_, err = stmt.Exec(e.Id, e.Email)
		checkErr(err)
		return nil
	}else{
		err := s.DB.Put([]byte(e.Id), e.Bytes(), nil)
		if err != nil {
			return err
		}
		return nil

	}
}

func (s *Storage) AllEntries() []*User {
	if dbConfig.isSql{
		result := make([]*User, 0)
		// query
		rows, err := s.SqlDb.Query("SELECT * FROM goauth")
		checkErr(err)

		var hash string
		var role string
		for rows.Next() {
			entry := &User{}
			err = rows.Scan(&entry.Id, &entry.Email, &hash, &role)
			checkErr(err)
			fmt.Println(entry)
			result = append(result, entry)
		}

		return result

	}else{
		result := make([]*User, 0)
		iter := s.DB.NewIterator(nil, nil)
		for iter.Next() {
			entry, _ := NewUserFromBytes(iter.Value())
			result = append(result, entry)
		}
		return result

	}
}

func (s *Storage) GetEntry(id string) *User {
	if dbConfig.isSql{
		rows, err := s.SqlDb.Query("SELECT * FROM job where Id = ?", id)
		checkErr(err)
		entry := &User{}

		for rows.Next() {
			err = rows.Scan(&entry.Id, &entry.Email)
			checkErr(err)
			fmt.Println(entry)
		}
		return entry

	}else{
		data, err := s.DB.Get([]byte(id), nil)
		if err != nil {
			return nil
		}
		entry, err := NewUserFromBytes(data)
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
