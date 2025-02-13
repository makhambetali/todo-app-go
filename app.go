package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Done     bool   `json:"done"`
	Priority string `json:"priority"`
	DueDate  string `json:"due_date"`
}

type App struct {
	ctx context.Context
	db  *sql.DB
}

func NewApp() *App {
	app := &App{}
	app.InitDB()
	return app
}

func (a *App) InitDB() {
	var err error
	a.db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		done BOOLEAN DEFAULT FALSE,
		priority TEXT DEFAULT 'Normal',
		due_date TEXT
	);`
	_, err = a.db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AddTask(text, priority, dueDate string) Task {
	res, err := a.db.Exec("INSERT INTO tasks (text, done, priority, due_date) VALUES (?, ?, ?, ?)", text, false, priority, dueDate)
	if err != nil {
		log.Println("Ошибка добавления задачи:", err)
		return Task{}
	}

	id, _ := res.LastInsertId()
	return Task{ID: int(id), Text: text, Done: false, Priority: priority, DueDate: dueDate}
}

func (a *App) GetTasks() []Task {
	rows, err := a.db.Query("SELECT id, text, done, priority, due_date FROM tasks")
	if err != nil {
		log.Println("Ошибка получения задач:", err)
		return nil
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Text, &task.Done, &task.Priority, &task.DueDate)
		if err != nil {
			log.Println("Ошибка при сканировании данных:", err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (a *App) ToggleTask(id int) {
	_, err := a.db.Exec("UPDATE tasks SET done = NOT done WHERE id = ?", id)
	if err != nil {
		log.Println("Ошибка обновления задачи:", err)
	}
}

func (a *App) DeleteTask(id int) {
	_, err := a.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Println("Ошибка удаления задачи:", err)
	}
}
