import { Component, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-modal-input',
  templateUrl: './modal-input.component.html',
  styleUrls: ['./modal-input.component.scss']
})
export class ModalInputComponent {
    inputTitle: string | undefined = "Nickname" 

    @Output() inputText = new EventEmitter<string>()

    active: boolean = false 

    updateInputText (event: Event): void {
        this.inputText.emit((event.target as HTMLInputElement).value)
    }

    updateModalActive (state: boolean): void {
        this.active = state
    }
}
