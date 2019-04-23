package tserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Sex     int    `json:"sex"`
	Address string `json:"address"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	sid, ok := r.URL.Query()["id"]
	if !ok {
	}

	// fixme not good
	id, err := strconv.Atoi(sid[0])
	if err != nil {
		log.Fatal(err)
	}

	student, err := get(id)
	if err != nil {
		log.Fatal()
	}

	bs, err := json.Marshal(student)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	sids, ok := r.URL.Query()["ids"]
	if !ok {
	}

	students, err := list(strings.Split(sids[0], ","))
	if err != nil {
		log.Fatal()
	}

	bs, err := json.Marshal(students)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
}

func saveOrUpdateHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	var stu Student
	err = json.Unmarshal(bytes, &stu)
	if err != nil {
		log.Panic(err)
	}
	saveOrUpdate(stu)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("{}"))
	if err != nil {
		log.Panic(err)
	}
}

func get(id int) (Student, error) {
	var s Student
	stat, err := MySQL.Prepare("select * from student where id = ?")
	if err != nil {
		return s, err
	}

	defer stat.Close()

	err = stat.QueryRow(id).Scan(&s.Id, &s.Name, &s.Sex, &s.Address)
	if err != nil {
		log.Panic(err)
	}
	return s, nil
}

func list(sids []string) ([]Student, error) {
	if len(sids) == 0 {
		return nil, errors.New("INVALID ARGUMENT")
	}

	ss := []Student{}
	// fixme wo we don't prevent sql inject
	query := "(" + strings.Join(sids, ",") + ")"
	rows, err := MySQL.Query("select * from student where id in " + query)

	if err != nil {
		return ss, err
	}
	defer rows.Close()

	for rows.Next() {
		s := Student{}
		if err := rows.Scan(&s.Id, &s.Name, &s.Sex, &s.Address); err != nil {
			return nil, err
		}
		ss = append(ss, s)
		fmt.Print(s, " ", query)
	}
	return ss, nil
}

func saveOrUpdate(s Student) {
	id := s.Id

	if id == 0 {
		_, err := MySQL.Exec("insert into student(name, sex, address) values(?, ?, ?)", s.Name, s.Sex, s.Address)
		if err != nil {
			log.Panic(err)
		}
	} else {
		_, err := MySQL.Exec("update student set name = ?, sex = ?, address = ? where id = ?", s.Name, s.Sex, s.Address, s.Id)
		if err != nil {
			log.Panic(err)
		}
	}
}

func delete() {

}
