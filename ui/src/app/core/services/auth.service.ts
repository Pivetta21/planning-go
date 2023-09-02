import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { BehaviorSubject, Observable, Subject, catchError, lastValueFrom, tap, throwError } from 'rxjs';

import { environment } from 'src/environments/environment';
import { IAuthRequest, IAuthResponse } from '../models/auth.model';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private authUrl = `${environment.apiUrl}/auth`
  private isAuthenticated$ = new BehaviorSubject(false)

  constructor(
    private http: HttpClient,
  ) { }

  login(request: IAuthRequest): Observable<IAuthResponse> {
    return this.http.post<IAuthResponse>(`${this.authUrl}/sign-in`, request).pipe(
      tap(() => { this.isAuthenticated$.next(true) }),
      catchError((err) => {
        this.isAuthenticated$.next(false)
        return throwError(() => new Error(err))
      })
    )
  }

  register(request: IAuthRequest): Observable<IAuthResponse> {
    return this.http.post<IAuthResponse>(`${this.authUrl}/sign-up`, request)
  }

  logout(): Observable<IAuthResponse> {
    return this.http.delete<IAuthResponse>(`${this.authUrl}/logout`).pipe(
      tap(() => { this.isAuthenticated$.next(false) }),
    )
  }

  refresh(): Observable<IAuthResponse> {
    return this.http.post<IAuthResponse>(`${this.authUrl}/refresh`, null).pipe(
      tap(() => { this.isAuthenticated$.next(true) }),
      catchError((err) => {
        this.isAuthenticated$.next(false)
        return throwError(() => new Error(err))
      })
    )
  }

  getIsAuthenticated(): Observable<boolean> {
    return this.isAuthenticated$.asObservable()
  }
}
