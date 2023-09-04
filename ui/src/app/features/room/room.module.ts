import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatSnackBarModule } from '@angular/material/snack-bar';

import { SharedModule } from 'src/app/shared/shared.module';

import { RoomRoutingModule } from './room-routing.module';
import { WebsocketComponent } from './pages/websocket/websocket.component';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    WebsocketComponent,
  ],
  imports: [
    CommonModule,
    SharedModule,
    RoomRoutingModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatSnackBarModule,
    FormsModule,
    HttpClientModule,
  ],
})
export class RoomModule { }
