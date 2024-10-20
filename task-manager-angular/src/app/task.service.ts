import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Task } from './task';  
import { map } from 'rxjs/operators'; 

@Injectable({
  providedIn: 'root',
})
export class TaskService {
  private apiUrl = 'http://localhost:8080/tasks'; 

  constructor(private http: HttpClient) {}

  getTasks(page: number = 1, limit: number = 10): Observable<{ data: Task[], totalTasks: number, totalPages: number }> {
    const url = `${this.apiUrl}?page=${page}&page_size=${limit}`;
    return this.http.get<any>(url).pipe(
      map(response => ({
        data: response.data.map((task: any) => ({
          id: task.ID,
          title: task.Title,
          description: task.Description,
          status: task.Status,
          priority: task.Priority,
          deadline: new Date(task.Deadline),
        })),
        totalTasks: response.totalTasks,
        totalPages: response.totalPages
      }))
    );
  }

  addTask(task: Task): Observable<Task> {
    return this.http.post<Task>(this.apiUrl, task);
  }

  getTaskById(id: number): Observable<Task> {
    const url = `${this.apiUrl}/${id}`;
    return this.http.get<Task>(url);
  }

  updateTask(task: Task): Observable<Task> {
    const url = `${this.apiUrl}/${task.id}`;
    return this.http.put<Task>(url, task);
  }
  
  deleteTask(id: number): Observable<void> {
    const url = `${this.apiUrl}/${id}`;
    return this.http.delete<void>(url); 
  }

}
