import { Component } from '@angular/core';

@Component({
  selector: 'app-modal-input',
  templateUrl: './modal-input.component.html',
  styleUrls: ['./modal-input.component.scss']
})
export class ModalInputComponent {
    inputTitle: string | undefined = "Nickname" 

    active: boolean = true 

    updateModalActive (state: boolean): void {
        this.active = state
    }
}
