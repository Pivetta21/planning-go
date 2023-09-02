import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProfileRoutingModule } from './profile-routing.module';
import { FindComponent } from './pages/find/find.component';
import { SharedModule } from 'src/app/shared/shared.module';
import { HttpClientModule } from '@angular/common/http';
import { ProfileService } from './services/profile.service';

@NgModule({
  declarations: [
    FindComponent,
  ],
  imports: [
    CommonModule,
    SharedModule,
    HttpClientModule,
    ProfileRoutingModule,
  ],
  providers: [
    ProfileService,
  ]
})
export class ProfileModule { }
