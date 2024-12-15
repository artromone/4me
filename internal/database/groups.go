package database

import (
    "fmt"

    "github.com/artromone/4me/internal/models"
)

func (d *Database) CreateGroup(group *models.Group) (int, error) {
    query := `
        INSERT INTO groups (name, description) 
        VALUES ($1, $2) 
        RETURNING id
    `
    var id int
    err := d.Conn.QueryRow(
        query, 
        group.Name, 
        group.Description,
    ).Scan(&id)

    if err != nil {
        return 0, fmt.Errorf("error creating group: %v", err)
    }

    return id, nil
}

func (d *Database) GetGroup(id int) (*models.Group, error) {
    query := `
        SELECT id, name, description 
        FROM groups 
        WHERE id = $1
    `
    group := &models.Group{}
    err := d.Conn.QueryRow(query, id).Scan(
        &group.ID, 
        &group.Name, 
        &group.Description,
    )

    if err != nil {
        return nil, fmt.Errorf("error retrieving group: %v", err)
    }

    // Fetch lists for this group
    lists, err := d.ListLists(id)
    if err == nil {
        group.Lists = lists
    }

    return group, nil
}

func (d *Database) ListGroups() ([]models.Group, error) {
    query := `
        SELECT id, name, description 
        FROM groups 
        ORDER BY name
    `
    rows, err := d.Conn.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error listing groups: %v", err)
    }
    defer rows.Close()

    var groups []models.Group
    for rows.Next() {
        var group models.Group
        err := rows.Scan(
            &group.ID, 
            &group.Name, 
            &group.Description,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning group: %v", err)
        }
        groups = append(groups, group)
    }

    return groups, nil
}
