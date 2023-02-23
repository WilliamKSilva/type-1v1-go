import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { firstValueFrom } from 'rxjs';
import { Game } from 'src/app/game';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.scss']
})
export class GameComponent implements OnInit {
    constructor(private route: ActivatedRoute, private http: HttpClient) {}

    game: Game | undefined

    loading: boolean = false

    ngOnInit() {
        this.loading = true

        let gameId: number | undefined

        this.route.paramMap.subscribe(params => gameId = Number(params.get('id')))

        this.getGame(gameId) 

        this.loading = false
    }

    async getGame(id: number | undefined): Promise<void> {
        this.game = await firstValueFrom(this.http.get<Game>(`http://localhost:3000/games?id=${id}`))
    }
}
