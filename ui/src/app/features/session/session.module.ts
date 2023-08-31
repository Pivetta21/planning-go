import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';

import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatTooltipModule } from '@angular/material/tooltip';

import { SharedModule } from 'src/app/shared/shared.module';

import { ListComponent } from './pages/list/list.component';
import { SessionRoutingModule } from './session-routing.module';
import { SessionService } from './services/session.service';

@NgModule({
  declarations: [
    ListComponent,
  ],
  imports: [
    CommonModule,
    HttpClientModule,
    SharedModule,
    MatCardModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatSnackBarModule,
    SessionRoutingModule,
  ],
  providers: [
    SessionService,
  ]
})
export class SessionModule { }
