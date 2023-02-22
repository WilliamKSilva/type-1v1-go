export interface Game {
    id: number;
    playerOne: string;
    playerTwo: string;
    status: string;
    winner: string;
    text: string;
}

export type NewGameData = {
    player_one?: string
}
