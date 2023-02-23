import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
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

    getGame(id: number | undefined) {
        this.http.get<Game>(`http://localhost:3000/games/${id}`).subscribe(data => this.game = data)
    }

}
