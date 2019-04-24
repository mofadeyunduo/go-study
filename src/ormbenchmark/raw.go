package ormbenchmark

import (
	"database/sql"
	"log"
)

type Rawcrud struct {
}

var rawDB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:Tomato-123@/student")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rawDB = db
}

func (Rawcrud) SelectAll(sql string) {
	stat, err := rawDB.Query(sql)
	if err != nil {
		log.Panic(err)
	}
	defer stat.Close()

	var students []Student
	for stat.Next() {
		s := Student{}
		err = stat.Scan(&s.Id, &s.Name, &s.Sex, &s.Address)
		if err != nil {
			log.Panic(err)
		}
		students = append(students, s)
	}
}

func (Rawcrud) SelectLeftJoin(lsql string) {
	stat, err := rawDB.Query(lsql)
	if err != nil {
		log.Panic(err)
	}
	defer stat.Close()

	var sgs []StudentGrade
	for stat.Next() {
		sg := StudentGrade{}
		err = stat.Scan(&sg.Id, &sg.Name, &sg.Sex, &sg.Address, &sg.StudentId, &sg.GradeName)
		if err != nil {
			log.Panic(err)
		}
		sgs = append(sgs, sg)
	}
}

func (Rawcrud) Insert(isql string) {
	tx, err := rawDB.Begin()
	if err != nil {
		log.Panic(err)
	}

	s := Student{
		0, "INC.COM", 1, "Paris",
	}
	_, err = tx.Exec(isql, s.Name, s.Sex, s.Address)
	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}

func (Rawcrud) Update(usql string) {
	tx, err := rawDB.Begin()
	if err != nil {
		log.Panic(err)
	}

	s := Student{
		4, "GODFATHER.COM", 1, "ICELAND",
	}
	_, err = tx.Exec(usql, s.Name, s.Sex, s.Address, s.Id)
	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}

func (Rawcrud) Delete(dsql string) {
	tx, err := rawDB.Begin()
	if err != nil {
		log.Panic(err)
	}

	s := Student{
		Id: 3,
	}
	_, err = tx.Exec(dsql, s.Id)
	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}
