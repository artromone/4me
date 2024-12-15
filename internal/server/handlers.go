package server

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/artromone/4me/internal/models"
)

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := s.db.CreateTask(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    task.ID = id
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

func (s *Server) listTasks(w http.ResponseWriter, r *http.Request) {
    listIDStr := r.URL.Query().Get("list_id")
    listID, err := strconv.Atoi(listIDStr)
    if err != nil {
        http.Error(w, "Invalid list ID", http.StatusBadRequest)
        return
    }

    tasks, err := s.db.ListTasks(listID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func (s *Server) getTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    task, err := s.db.GetTask(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func (s *Server) updateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    task.ID = id

    if err := s.db.UpdateTask(&task); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    if err := s.db.DeleteTask(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (s *Server) createList(w http.ResponseWriter, r *http.Request) {
    var list models.List
    if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := s.db.CreateList(&list)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    list.ID = id
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(list)
}

func (s *Server) listLists(w http.ResponseWriter, r *http.Request) {
    groupIDStr := r.URL.Query().Get("group_id")
    groupID, err := strconv.Atoi(groupIDStr)
    if err != nil {
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }

    lists, err := s.db.ListLists(groupID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(lists)
}

func (s *Server) createGroup(w http.ResponseWriter, r *http.Request) {
    var group models.Group
    if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := s.db.CreateGroup(&group)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    group.ID = id
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(group)
}

func (s *Server) listGroups(w http.ResponseWriter, r *http.Request) {
    groups, err := s.db.ListGroups()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(groups)
}
