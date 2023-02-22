import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { Game, NewGameData } from 'src/app/game';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent {
    constructor (private http: HttpClient, private router: Router) {}

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