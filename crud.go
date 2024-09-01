package main

func createTask(task string) error {
    _, err := db.Exec("INSERT INTO todos (task) VALUES ($1)", task)
    return err
}

func getTasks() ([]Todo, error) {
    rows, err := db.Query("SELECT id, task, completed FROM todos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []Todo
    for rows.Next() {
        var t Todo
        if err := rows.Scan(&t.ID, &t.Task, &t.Completed); err != nil {
            return nil, err
        }
        todos = append(todos, t)
    }
    return todos, nil
}

func updateTask(id int, completed bool) error {
    _, err := db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", completed, id)
    return err
}

func deleteTask(id int) error {
    _, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
    return err
}

type Todo struct {
    ID        int
    Task      string
    Completed bool
}
