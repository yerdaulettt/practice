package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var taskDatabase = []task{}
var uniqueId int = 1

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Query().Has("id") {
		taskId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"invalid id"}`))
			return
		}

		realTaskId := -1
		for i := range taskDatabase {
			if taskDatabase[i].Id == taskId {
				realTaskId = i
				break
			}
		}

		if taskId < 1 || realTaskId == -1 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"task not found"}`))
			return
		}

		err = json.NewEncoder(w).Encode(taskDatabase[realTaskId])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"json error"}`))
		}
	} else if r.URL.Query().Has("done") {
		status := r.URL.Query().Get("done")
		var tasksResult []task

		switch status {
		case "true":
			for i := range taskDatabase {
				if taskDatabase[i].Done == true {
					tasksResult = append(tasksResult, taskDatabase[i])
				}
			}
		case "false":
			for i := range taskDatabase {
				if taskDatabase[i].Done == false {
					tasksResult = append(tasksResult, taskDatabase[i])
				}
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"no options"}`))
			return
		}

		if tasksResult != nil {
			json.NewEncoder(w).Encode(&tasksResult)
		} else {
			w.Write([]byte(`{"error":"empty"}`))
		}
	} else if r.URL.Query().Has("sort") {
		option := r.URL.Query().Get("sort")

		switch option {
		case "asc":
			sort.Slice(taskDatabase, func(i, j int) bool {
				return taskDatabase[i].Id < taskDatabase[j].Id
			})
		case "desc":
			sort.Slice(taskDatabase, func(i, j int) bool {
				return taskDatabase[i].Id > taskDatabase[j].Id
			})
		default:
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(`{"error":"no optoins"}`))
			return
		}

		err := json.NewEncoder(w).Encode(&taskDatabase)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		err := json.NewEncoder(w).Encode(&taskDatabase)
		if err != nil {
			// http.Error(w, `{"error":"some error"}`, http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"json error"}`))
		}
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Query().Has("call") {
		var testDB []task

		response, err := http.Get("https://jsonplaceholder.typicode.com/todos")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		err = json.NewDecoder(response.Body).Decode(&testDB)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		var filteredDB []task
		for i := range testDB {
			if len(testDB[i].Title) > 20 {
				continue
			}

			var t task = testDB[i]
			t.Id = uniqueId
			uniqueId++

			filteredDB = append(filteredDB, t)
		}

		taskDatabase = append(taskDatabase, filteredDB...)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "added tasks from external source"}`))
	} else {
		var t task
		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"not valid"}`))
			return
		}

		if t.Title == "" {
			// http.Error(w, `{"error": "invalid title"}`, http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"invalid title"}`))
			return
		}

		if len(t.Title) > 20 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"title length more than 20 characters"}`))
			return
		}

		t.Id = uniqueId
		uniqueId++
		t.Done = false
		taskDatabase = append(taskDatabase, t)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&t)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(r.URL.Query().Get("id"))
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid id"}`))
		return
	}

	realTaskId := -1
	for i := range taskDatabase {
		if taskDatabase[i].Id == taskId {
			realTaskId = i
			break
		}
	}

	if taskId < 1 || realTaskId == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"task not found"}`))
		return
	}

	var task_test task
	errr := json.NewDecoder(r.Body).Decode(&task_test)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"error occured"}`))
		return
	}
	taskDatabase[realTaskId].Done = task_test.Done

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"updated":` + strconv.FormatBool(task_test.Done) + `}`))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Query().Has("id") {
		taskId, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"invalid id"}`))
			return
		}

		realTaskId := -1
		for i := range taskDatabase {
			if taskDatabase[i].Id == taskId {
				realTaskId = i
				break
			}
		}

		if taskId < 0 || realTaskId == -1 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"task not found"}`))
			return
		}

		taskDatabase = append(taskDatabase[:realTaskId], taskDatabase[(realTaskId+1):]...)
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message":"task deleted"}`))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"id not provided"}`))
	}
}
