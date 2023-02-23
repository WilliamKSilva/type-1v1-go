import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { firstValueFrom } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket'
import { Game } from 'src/app/game';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.scss']
})
export class GameComponent implements OnInit {
    constructor(private route: ActivatedRoute, private http: HttpClient) {}

    game: Game | undefined

    socketConn: WebSocketSubject<unknown> | undefined

    loading: boolean = false

    playerText: string = ""

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

    updateInputText (event: Event): void {
        this.playerText = this.playerText + (event.target as HTMLInputElement).value

        if (!this.socketConn) {
             this.socketConn = webSocket('ws://localhost:3000/games/run')
        }

        this.socketConn.subscribe({
            next: msg => console.log(msg),
            error: err => console.log(err),
            complete: () => console.log('complete')
        }) 
    }
}
