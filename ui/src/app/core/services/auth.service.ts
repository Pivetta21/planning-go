import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

import { environment } from 'src/environments/environment';
import { IAuthRequest, IAuthResponse } from '../models/auth.model';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private authUrl = `${environment.apiUrl}/auth`;

  private options = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    }),
    withCredentials: true,
  };

  constructor(
    private http: HttpClient,
  ) { }

  login(request: IAuthRequest): Observable<IAuthResponse> {
    return this.http.post<IAuthResponse>(`${this.authUrl}/sign-in`, request, this.options)
  }

  register(request: IAuthRequest): Observable<IAuthResponse> {
    return this.http.post<IAuthResponse>(`${this.authUrl}/sign-up`, request, this.options)
  }

  logout(): Observable<IAuthResponse> {
    return this.http.delete<IAuthResponse>(`${this.authUrl}/logout`, this.options)
  }

  refresh(): Observable<IAuthResponse> {
    return this.http.post<IAuthResponse>(`${this.authUrl}/refresh`, this.options)
  }
}
