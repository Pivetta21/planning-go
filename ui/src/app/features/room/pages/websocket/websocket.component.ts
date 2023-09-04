import { Component, ElementRef, OnDestroy, OnInit, Renderer2, ViewChild } from '@angular/core';
import { WebSocketSubject, webSocket } from 'rxjs/webSocket'
import { RoomService } from '../../services/room.service';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-websocket',
  templateUrl: './websocket.component.html',
  styleUrls: ['./websocket.component.css']
})
export class WebsocketComponent implements OnInit, OnDestroy {
  message?: string;
  cachedMessage?: string;
  subject?: WebSocketSubject<string>

  @ViewChild('div') div!: ElementRef

  constructor(
    private render: Renderer2,
    private roomService: RoomService,
  ) { }

  ngOnInit(): void {
    this.subject = webSocket(`${environment.wsUrl}/room/conn?id=1`);

    this.subject.subscribe({
      next: (msg) => {
        const p = this.render.createElement('p') as HTMLParagraphElement
        p.innerHTML = msg
        p.classList.add('mat-body-1', 'message')
        if (msg == this.cachedMessage) {
          p.classList.add('sent')
        } else {
          p.classList.add('received')
        }
        this.render.appendChild(this.div.nativeElement, p)
        this.cachedMessage = ''
      },
    })
  }

  ngOnDestroy(): void {
    this.subject?.complete()
    this.subject?.unsubscribe()
  }

  sendToServer($event: Event) {
    $event.preventDefault();

    if (!this.subject || !this.message)
      return

    this.subject.next(this.message)
    this.cachedMessage = this.message
    this.clearMessage()
  }

  clearMessage() {
    this.message = ''
  }
}
