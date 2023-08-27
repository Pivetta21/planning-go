import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { environment } from 'src/environments/environment';
import { LoginRequest, LoginResponse } from '../models/login.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private authUrl = `${environment.apiUrl}/auth`;

  constructor(
    private http: HttpClient,
  ) { }

  login(request: LoginRequest): Observable<LoginResponse> {
    const httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    }

    return this.http.post<LoginResponse>(`${this.authUrl}/sign-in`, request, httpOptions)
  }
}
