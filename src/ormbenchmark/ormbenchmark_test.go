package ormbenchmark

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var (
	SelectAllSQL      = "select * from student"
	SelectLeftJoinSQL = "select * from student left join grade on student.id = grade.student_id"
	InsertSQL         = "insert into student(name, sex, address) values(?, ?, ?)"
	UpdateSQL         = "update student set name = ?, sex = ?, address = ? where id = ?"
	DeleteSQL         = "delete from student where id = ?"
)

func benchmarkForSelectAll(crud CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.SelectAll(SelectAllSQL)
	}
}

func benchmarkForSelectLeftJoin(crud CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.SelectLeftJoin(SelectLeftJoinSQL)
	}
}

func benchmarkForInsert(crud CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.Insert(InsertSQL)
	}
}

func benchmarkForUpdate(crud CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.Update(UpdateSQL)
	}
}

func benchmarkForDelete(crud CRUD, b *testing.B) {
	for i := 0; i < b.N; i++ {
		crud.Delete(DeleteSQL)
	}
}

// sqlx
func BenchmarkSqlxcrud_SelectAll(b *testing.B) {
	benchmarkForSelectAll(Sqlxcrud{}, b)
}

func BenchmarkSqlxcrud_SelectLeftJoin(b *testing.B) {
	benchmarkForSelectLeftJoin(Sqlxcrud{}, b)
}

func BenchmarkSqlxcrud_Insert(b *testing.B) {
	benchmarkForInsert(Sqlxcrud{}, b)
}

func BenchmarkSqlxcrud_Update(b *testing.B) {
	benchmarkForUpdate(Sqlxcrud{}, b)
}

func BenchmarkSqlxcrud_Delete(b *testing.B) {
	benchmarkForDelete(Sqlxcrud{}, b)
}

// gorm
func BenchmarkGormcrud_SelectAll(b *testing.B) {
	benchmarkForSelectAll(Gormcrud{}, b)
}

func BenchmarkGormcrud_SelectLeftJoin(b *testing.B) {
	benchmarkForSelectLeftJoin(Gormcrud{}, b)
}

func BenchmarkGormcrud_Insert(b *testing.B) {
	benchmarkForInsert(Gormcrud{}, b)
}

func BenchmarkGormcrud_Update(b *testing.B) {
	benchmarkForUpdate(Gormcrud{}, b)
}

func BenchmarkGormcrud_Delete(b *testing.B) {
	benchmarkForDelete(Gormcrud{}, b)
}

// raw
func BenchmarkRawcrud_SelectAll(b *testing.B) {
	benchmarkForSelectAll(Rawcrud{}, b)
}

func BenchmarkRawcrud_SelectLeftJoin(b *testing.B) {
	benchmarkForSelectLeftJoin(Rawcrud{}, b)
}

func BenchmarkRawcrud_Insert(b *testing.B) {
	benchmarkForInsert(Rawcrud{}, b)
}

func BenchmarkRawcrud_Update(b *testing.B) {
	benchmarkForUpdate(Rawcrud{}, b)
}

func BenchmarkRawcrud_Delete(b *testing.B) {
	benchmarkForDelete(Rawcrud{}, b)
}
