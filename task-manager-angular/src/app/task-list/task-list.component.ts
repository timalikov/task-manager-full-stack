import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { RouterLink } from '@angular/router';
import { Task } from '../task';
import { TaskService } from '../task.service';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'app-task-list',
  standalone: true,
  imports: [MatButtonModule, MatCardModule, RouterLink, CommonModule, HttpClientModule],
  providers: [TaskService],
  templateUrl: './task-list.component.html',
  styleUrl: './task-list.component.css'
})
export class TaskListComponent {
  tasks: Task[] = [];
  totalTasks: number = 0;
  totalPages: number = 0;
  currentPage: number = 1;
  pageSize: number = 10;

  constructor(private taskService: TaskService) {}

  ngOnInit(): void {
    this.loadTasks();
  }

  loadTasks(): void {
    this.taskService.getTasks(this.currentPage, this.pageSize).subscribe({
      next: (response) => {
        console.log('Response:', response); // Log the response to check if data is being returned
        this.tasks = response.data;
        this.totalTasks = response.totalTasks;
        this.totalPages = response.totalPages;
      },
      error: (error) => {
        console.error('Error fetching tasks:', error);
      }
    });
  }  
  
  removeTask(id: number): void {
    this.taskService.deleteTask(id).subscribe({
      next: () => {
        console.log(`Task with id ${id} deleted successfully.`);
        // Remove the task from the tasks array in the frontend
        this.tasks = this.tasks.filter(task => task.id !== id);
      },
      error: (error) => {
        console.error('Error deleting task:', error);
      }
    });
  }

  trackById(index: number, task: Task): number {
    return task.id;
  }
}
