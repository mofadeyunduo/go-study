package ormbenchmark

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/mofadeyunduo/golang-study/framework-practice/ormbenchmark/model"
	"testing"
)

var (
	SelectAllSQL      = "select * from student"
	SelectLeftJoinSQL = "select * from student left join grade on student.id = grade.student_id"
	InsertSQL         = "insert into student(name, sex, address) values(?, ?, ?)"
	UpdateSQL         = "update student set name = ?, sex = ?, address = ? where id = ?"
	DeleteSQL         = "delete from student where id = ?"
)

func benchmarkForSelectAll(crud model.CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.SelectAll(SelectAllSQL)
	}
}

func benchmarkForSelectLeftJoin(crud model.CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.SelectLeftJoin(SelectLeftJoinSQL)
	}
}

func benchmarkForInsert(crud model.CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.Insert(InsertSQL)
	}
}

func benchmarkForUpdate(crud model.CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.Update(UpdateSQL)
	}
}

func benchmarkForDelete(crud model.CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.Delete(DeleteSQL)
	}
}

// sqlx
func BenchmarkSqlxcrud_SelectAll(b *testing.B) {
	benchmarkForSelectAll(model.Sqlxcrud{}, b)
}

func BenchmarkSqlxcrud_SelectLeftJoin(b *testing.B) {
	benchmarkForSelectLeftJoin(model.Sqlxcrud{}, b)
}

func BenchmarkSqlxcrud_Insert(b *testing.B) {
	benchmarkForInsert(model.Sqlxcrud{}, b)
}

func BenchmarkSqlxcrud_Update(b *testing.B) {
	benchmarkForUpdate(model.Sqlxcrud{}, b)
}

func BenchmarkSqlxcrud_Delete(b *testing.B) {
	benchmarkForDelete(model.Sqlxcrud{}, b)
}

// gorm
func BenchmarkGormcrud_SelectAll(b *testing.B) {
	benchmarkForSelectAll(model.Gormcrud{}, b)
}

func BenchmarkGormcrud_SelectLeftJoin(b *testing.B) {
	benchmarkForSelectLeftJoin(model.Gormcrud{}, b)
}

func BenchmarkGormcrud_Insert(b *testing.B) {
	benchmarkForInsert(model.Gormcrud{}, b)
}

func BenchmarkGormcrud_Update(b *testing.B) {
	benchmarkForUpdate(model.Gormcrud{}, b)
}

func BenchmarkGormcrud_Delete(b *testing.B) {
	benchmarkForDelete(model.Gormcrud{}, b)
}

// raw
func BenchmarkRawcrud_SelectAll(b *testing.B) {
	benchmarkForSelectAll(model.Rawcrud{}, b)
}

func BenchmarkRawcrud_SelectLeftJoin(b *testing.B) {
	benchmarkForSelectLeftJoin(model.Rawcrud{}, b)
}

func BenchmarkRawcrud_Insert(b *testing.B) {
	benchmarkForInsert(model.Rawcrud{}, b)
}

func BenchmarkRawcrud_Update(b *testing.B) {
	benchmarkForUpdate(model.Rawcrud{}, b)
}

func BenchmarkRawcrud_Delete(b *testing.B) {
	benchmarkForDelete(model.Rawcrud{}, b)
}
