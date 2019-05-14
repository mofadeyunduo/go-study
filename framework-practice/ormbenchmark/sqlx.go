package ormbenchmark

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mofadeyunduo/golang-study/framework-practice/ormbenchmark/model"
	"log"
)

type Sqlxcrud struct {
}

var sqlxdb *sqlx.DB

func init() {
	sqlxDB, err := sqlx.Open("mysql", "root:Tomato-123@/student")
	sqlxdb = sqlxDB
	if err != nil {
		log.Panic(err)
	}
}

func (x Sqlxcrud) SelectAll(sql string) {
	idStu := []model.Student{}
	err := sqlxdb.Select(&idStu, sql)

	if err != nil {
		log.Panic(err)
	}
}

func (x Sqlxcrud) SelectLeftJoin(sql string) {
	var idStu []model.StudentGrade
	err := sqlxdb.Select(&idStu, sql)

	if err != nil {
		log.Panic(err)
	}
}

func (x Sqlxcrud) Insert(sql string) {
	s := model.Student{
		0, "Piers", 1, "China",
	}
	tx := sqlxdb.MustBegin()
	tx.MustExec(sql, s.Name, s.Sex, s.Address)
	err := tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}

func (x Sqlxcrud) Update(sql string) {
	s := model.Student{
		2, "Fiers", 1, "America",
	}
	tx := sqlxdb.MustBegin()
	tx.MustExec(sql, s.Name, s.Sex, s.Address, s.Id)
	if err := tx.Commit(); err != nil {
		log.Panic(err)
	}
}

func (x Sqlxcrud) Delete(sql string) {
	tx := sqlxdb.MustBegin()
	tx.MustExec(sql, 1)
	if err := tx.Commit(); err != nil {
		log.Panic(err)
	}
}
