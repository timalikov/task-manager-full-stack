package main

import (
	"log"
	"task_management_backend/routes"
)

func main() {

	// Migrations: it is enough to run it one time
	//migrations.Migrate()

	router := routes.SetupRoutes()

	// Start the server
	log.Println("Starting server on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	//database, err := db.ConnectToDb()
	//if err != nil {
	//	log.Fatalf("failed to connect to the database: %v", err)
	//}
	//
	//// Create a new instance of TaskDAO
	//taskDAO := dao.NewTaskDAO(database)
	//
	//// Create a new task
	//newTask := &models.Task{
	//	ID:          1,
	//	Title:       "New Task",
	//	Description: "This is a new task",
	//	Priority:    1,
	//	Deadline:    time.Now().AddDate(0, 0, 7), // Deadline 7 days from now
	//}
	////err = taskDAO.CreateTask(newTask)
	////if err != nil {
	////	log.Fatalf("failed to create task: %v", err)
	////}
	////fmt.Println("Task created successfully:", newTask)
	//
	//// Get a task by ID
	//task, err := taskDAO.GetTaskByID(newTask.ID)
	//if err != nil {
	//	log.Fatalf("failed to retrieve task: %v", err)
	//}
	//fmt.Println("Retrieved task:", task)
	//
	//// Update the task
	//task.Status = "completed"
	//err = taskDAO.UpdateTask(task)
	//if err != nil {
	//	log.Fatalf("failed to update task: %v", err)
	//}
	//fmt.Println("Task updated successfully")
	//
	//// Get all tasks
	//tasks, err := taskDAO.GetAllTasks()
	//if err != nil {
	//	log.Fatalf("failed to retrieve tasks: %v", err)
	//}
	//fmt.Println("All tasks:", tasks)
	//
	//// Delete the task
	//err = taskDAO.DeleteTask(task.ID)
	//if err != nil {
	//	log.Fatalf("failed to delete task: %v", err)
	//}
	//fmt.Println("Task deleted successfully")

}
