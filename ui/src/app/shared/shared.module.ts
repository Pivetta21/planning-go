import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatSidenavModule } from '@angular/material/sidenav';

import { AuthService } from '../core/services/auth.service';

import { ContainerComponent } from './components/container/container.component';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatButtonModule } from '@angular/material/button';

@NgModule({
  declarations: [
    ContainerComponent,
  ],
  imports: [
    CommonModule,
    MatSidenavModule,
    MatSnackBarModule,
    MatButtonModule,
  ],
  exports: [
    ContainerComponent,
  ],
  providers: [
    AuthService,
  ]
})
export class SharedModule { }
