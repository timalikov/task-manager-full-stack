import { Component, OnInit } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ActivatedRoute, Router } from '@angular/router';
import { Task } from '../task';
import { FormsModule } from '@angular/forms';
import { TaskService } from '../task.service'; 
import { CommonModule } from '@angular/common';
import { MatDatepickerModule } from '@angular/material/datepicker'; 
import { MatNativeDateModule } from '@angular/material/core'; 


@Component({
  selector: 'app-edit-task',
  standalone: true,
  imports: [MatButtonModule, MatCardModule, MatInputModule, MatFormFieldModule, FormsModule, CommonModule, MatDatepickerModule, MatNativeDateModule],
  providers: [TaskService],
  templateUrl: './edit-task.component.html',
  styleUrls: ['./edit-task.component.css']
})
export class EditTaskComponent implements OnInit {
  taskId: number = -1;
  taskTitle: string = '';
  taskDescription: string = '';
  taskStatus: string = 'pending'; 
  taskPriority: number = 3; 
  taskDeadline: Date = new Date();

  constructor(
    private route: ActivatedRoute, 
    private router: Router, 
    private taskService: TaskService 
  ) { }

  ngOnInit(): void {
    this.taskId = Number(this.route.snapshot.paramMap.get('id'));

    this.taskService.getTaskById(this.taskId).subscribe({
      next: (task) => {
        this.taskTitle = task.title;
        this.taskDescription = task.description;
        this.taskStatus = task.status;
        this.taskPriority = task.priority;
        this.taskDeadline = new Date(task.deadline);
      },
      error: (error) => {
        console.error('Error fetching task:', error);
      }
    });
  }

  editTask(): void {
    let updatedTask: Task = {
      id: this.taskId,
      title: this.taskTitle,
      description: this.taskDescription,
      status: this.taskStatus,
      priority: this.taskPriority,
      deadline: this.taskDeadline
    };

    this.taskService.updateTask(updatedTask).subscribe({
      next: (response) => {
        console.log('Task updated successfully:', response);
        this.router.navigate(['/']); 
      },
      error: (error) => {
        console.error('Error updating task:', error);
      }
    });
  }
}
