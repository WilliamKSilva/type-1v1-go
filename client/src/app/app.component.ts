import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
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

  playerOne: string | undefined

  isLoading: boolean = false

  modalActive: boolean = false

  createdGame: Game | undefined

  updatePlayerOne(name: string) {
    this.playerOne = name
  } 

  openModal(): void {
      this.modalActive = true
  }
 
  newGame (data: NewGameData): void  {
    this.modalActive = false 

    this.isLoading = true

    this.http.post<Game>(this.gameURL, data).subscribe(data => this.createdGame = data)

    this.isLoading = false
  }

}
