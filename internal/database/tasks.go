package database

import (
    "database/sql"
    "fmt"

    "github.com/artromone/4me/internal/models"
)

func (d *Database) CreateTask(task *models.Task) (int, error) {
    query := `
        INSERT INTO tasks (title, description, status, list_id, due_date) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id
    `
    var id int
    err := d.Conn.QueryRow(
        query, 
        task.Title, 
        task.Description, 
        task.Status, 
        task.ListID, 
        task.DueDate,
    ).Scan(&id)

    if err != nil {
        return 0, fmt.Errorf("error creating task: %v", err)
    }

    return id, nil
}

func (d *Database) GetTask(id int) (*models.Task, error) {
    query := `
        SELECT id, title, description, status, list_id, created_at, due_date 
        FROM tasks 
        WHERE id = $1
    `
    task := &models.Task{}
    err := d.Conn.QueryRow(query, id).Scan(
        &task.ID, 
        &task.Title, 
        &task.Description, 
        &task.Status, 
        &task.ListID, 
        &task.CreatedAt, 
        &task.DueDate,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("task not found")
        }
        return nil, fmt.Errorf("error retrieving task: %v", err)
    }

    return task, nil
}

func (d *Database) ListTasks(listID int) ([]models.Task, error) {
    query := `
        SELECT id, title, description, status, list_id, created_at, due_date 
        FROM tasks 
        WHERE list_id = $1 
        ORDER BY created_at DESC
    `
    rows, err := d.Conn.Query(query, listID)
    if err != nil {
        return nil, fmt.Errorf("error listing tasks: %v", err)
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        err := rows.Scan(
            &task.ID, 
            &task.Title, 
            &task.Description, 
            &task.Status, 
            &task.ListID, 
            &task.CreatedAt, 
            &task.DueDate,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning task: %v", err)
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func (d *Database) UpdateTask(task *models.Task) error {
    query := `
        UPDATE tasks 
        SET title = $1, description = $2, status = $3, list_id = $4, due_date = $5 
        WHERE id = $6
    `
    _, err := d.Conn.Exec(
        query, 
        task.Title, 
        task.Description, 
        task.Status, 
        task.ListID, 
        task.DueDate, 
        task.ID,
    )

    if err != nil {
        return fmt.Errorf("error updating task: %v", err)
    }

    return nil
}

func (d *Database) DeleteTask(id int) error {
    query := `DELETE FROM tasks WHERE id = $1`
    _, err := d.Conn.Exec(query, id)

    if err != nil {
        return fmt.Errorf("error deleting task: %v", err)
    }

    return nil
}
