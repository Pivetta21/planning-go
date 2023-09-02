import { Location } from "@angular/common";
import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { MatSnackBar } from "@angular/material/snack-bar";
import { Router } from "@angular/router";
import { Observable, catchError, firstValueFrom, lastValueFrom, switchMap, tap, throwError } from "rxjs";
import { AuthService } from "../services/auth.service";

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
    private isAuthenticated: boolean = false

    constructor(
        private router: Router,
        private location: Location,
        private snackBar: MatSnackBar,
        private authService: AuthService,
    ) {
        authService.getIsAuthenticated().subscribe(value => {
            this.isAuthenticated = value
        })
    }

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        if (req.url.search(/.*\/auth\/.*/) !== -1) {
            return next.handle(req)
        }

        return next.handle(req).pipe(
            catchError((err: HttpErrorResponse) => {
                if (!this.isAuthenticated) {
                    this.navigateToLogin()
                    return next.handle(req)
                }

                if (err.status === 401) {
                    return this.authService.refresh().pipe(
                        switchMap(() => { 
                            return next.handle(req) 
                        }),
                        catchError((refreshErr) => {
                            this.navigateToLogin()
                            return throwError(() => new Error(refreshErr.message))
                        })
                    )
                }

                if (err.status === 403) {
                    this.navigateBack()
                    return throwError(() => new Error("Forbidden"))
                }

                this.snackBar.open(err.message, 'Close', { duration: 5000 })
                return throwError(() => new Error(err.message))
            })
        )
    }

    private navigateToLogin() {
        this.router.navigate(['auth', 'login'], { replaceUrl: true })
            .then(success => {
                if (!success) {
                    this.location.back()
                }
            })
    }

    private navigateBack() {
        this.router.navigate(['..'], { replaceUrl: true })
            .then(success => {
                if (!success) {
                    this.location.back()
                }
            })
    }
}