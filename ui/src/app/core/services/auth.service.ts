import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

import { environment } from 'src/environments/environment';

import { LoginRequest, LoginResponse } from '../../features/auth/models/login.model';
import { RegisterRequest, RegisterResponse } from '../../features/auth/models/register.model';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private authUrl = `${environment.apiUrl}/auth`;

  private httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    }),
    withCredentials: true,
  }

  constructor(
    private http: HttpClient,
  ) { }

  login(request: LoginRequest): Observable<LoginResponse> {
    return this.http.post<LoginResponse>(`${this.authUrl}/sign-in`, request, this.httpOptions)
  }

  register(request: RegisterRequest): Observable<RegisterResponse> {
    return this.http.post<RegisterResponse>(`${this.authUrl}/sign-up`, request, this.httpOptions)
  }

  logout(): Observable<{ message: string }> {
    return this.http.delete<{ message: string }>(`${this.authUrl}/logout`, this.httpOptions)
  }
}
