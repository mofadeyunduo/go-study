package ormbenchmark

type CRUD interface {
	SelectAll(sql string)

	SelectLeftJoin(sql string)

	Insert(sql string)

	Update(sql string)

	Delete(sql string)
}
