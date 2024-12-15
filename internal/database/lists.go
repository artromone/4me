package database

import (
    "fmt"

    "github.com/artromone/4me/internal/models"
)

func (d *Database) CreateList(list *models.List) (int, error) {
    query := `
        INSERT INTO lists (name, description, group_id) 
        VALUES ($1, $2, $3) 
        RETURNING id
    `
    var id int
    err := d.Conn.QueryRow(
        query, 
        list.Name, 
        list.Description, 
        list.GroupID,
    ).Scan(&id)

    if err != nil {
        return 0, fmt.Errorf("error creating list: %v", err)
    }

    return id, nil
}

func (d *Database) GetList(id int) (*models.List, error) {
    query := `
        SELECT id, name, description, group_id 
        FROM lists 
        WHERE id = $1
    `
    list := &models.List{}
    err := d.Conn.QueryRow(query, id).Scan(
        &list.ID, 
        &list.Name, 
        &list.Description, 
        &list.GroupID,
    )

    if err != nil {
        return nil, fmt.Errorf("error retrieving list: %v", err)
    }

    // Fetch tasks for this list
    tasks, err := d.ListTasks(id)
    if err == nil {
        list.Tasks = tasks
    }

    return list, nil
}

func (d *Database) ListLists(groupID int) ([]models.List, error) {
    query := `
        SELECT id, name, description, group_id 
        FROM lists 
        WHERE group_id = $1 
        ORDER BY name
    `
    rows, err := d.Conn.Query(query, groupID)
    if err != nil {
        return nil, fmt.Errorf("error listing lists: %v", err)
    }
    defer rows.Close()

    var lists []models.List
    for rows.Next() {
        var list models.List
        err := rows.Scan(
            &list.ID, 
            &list.Name, 
            &list.Description, 
            &list.GroupID,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning list: %v", err)
        }
        lists = append(lists, list)
    }

    return lists, nil
}
