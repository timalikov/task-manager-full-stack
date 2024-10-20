import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { TaskService } from '../task.service'; 
import { Task } from '../task';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatDatepickerModule } from '@angular/material/datepicker'; 
import { MatNativeDateModule } from '@angular/material/core'; 


@Component({
  selector: 'app-add-task',
  standalone: true,
  imports: [MatButtonModule, MatCardModule, MatInputModule, MatFormFieldModule, FormsModule, CommonModule, MatDatepickerModule, MatNativeDateModule],
  providers: [TaskService],
  templateUrl: './add-task.component.html',
  styleUrls: ['./add-task.component.css']
})
export class AddTaskComponent {
  taskTitle: string = '';
  taskDescription: string = '';
  taskStatus: string = 'pending'; 
  taskPriority: number = 3; 
  taskDeadline: Date = new Date();

  constructor(private taskService: TaskService) {} 

  addTask(): void {
    let task: Task = {
      id: 0, 
      title: this.taskTitle,
      description: this.taskDescription,
      status: this.taskStatus,
      priority: this.taskPriority,
      deadline: this.taskDeadline
    };

    this.taskService.addTask(task).subscribe({
      next: (response) => {
        console.log('Task added successfully:', response);
        this.taskTitle = '';
        this.taskDescription = '';
      },
      error: (error) => {
        console.error('Error adding task:', error);
      }
    });
  }
}
