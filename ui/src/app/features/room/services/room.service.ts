import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class RoomService {
  private roomUrl = `${environment.apiUrl}/room`;
  private roomWsUrl = `${environment.wsUrl}/room`;

  constructor(
    private http: HttpClient,
  ) { }
}
