import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { Game, NewGameData } from './game';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent {
  title = 'client';
  
  constructor (private http: HttpClient) {}

  gameURL: string = 'http://localhost:3000/games'

  newGame (data: NewGameData): Observable<Game> {
   return this.http.post<Game>(this.gameURL, data)
  }

}
