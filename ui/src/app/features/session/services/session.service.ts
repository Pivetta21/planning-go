import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from 'src/environments/environment';
import { ISessionModel } from '../models/session.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SessionService {
  private sessionUrl = `${environment.apiUrl}/session`;

  constructor(
    private http: HttpClient,
  ) { }

  list(): Observable<ISessionModel[]> {
    return this.http.get<ISessionModel[]>(this.sessionUrl);
  }

  delete(identifier: string): Observable<any> {
    return this.http.delete(`${this.sessionUrl}/${identifier}`);
  }
}
