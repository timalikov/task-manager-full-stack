import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { TabGroupComponent } from './tab-group/tab-group.component';
import { AuthService } from './auth.service';
import { HttpClientModule } from '@angular/common/http'; 


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, TabGroupComponent, HttpClientModule], 
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'task-manager';

  constructor(public authService: AuthService){}
}
