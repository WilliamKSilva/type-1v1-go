import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { firstValueFrom } from 'rxjs';
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

    isNewGameLoading: boolean = false 
    isEnterGameLoading: boolean = false

    modalActive: boolean = false

    updatePlayerOne(name: string) {
        this.playerOne = name
    } 

    openModal(): void {
        this.modalActive = true
    }
 
    async newGame (data: NewGameData): Promise<void> {
        this.modalActive = false 

        this.isNewGameLoading = true 

        const response = await firstValueFrom(this.http.post<Game>(this.gameURL, data))

        this.isNewGameLoading = false

        this.router.navigate(['/games', { id: response.id }])
    }
}
