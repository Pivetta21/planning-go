import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { WebsocketComponent } from './pages/websocket/websocket.component';

const routes: Routes = [
  {
    path: '',
    component: WebsocketComponent,
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class RoomRoutingModule { }
