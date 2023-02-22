import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http'

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ModalInputComponent } from './components/modal-input/modal-input.component';
import { LoadingComponent } from './components/loading/loading.component';
import { GameComponent } from './pages/game/game.component';

@NgModule({
  declarations: [
    AppComponent,
    ModalInputComponent,
    LoadingComponent,
    GameComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
