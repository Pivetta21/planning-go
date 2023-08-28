import { Component } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/core/services/auth.service';

@Component({
  selector: 'app-container',
  templateUrl: './container.component.html',
  styleUrls: ['./container.component.css']
})
export class ContainerComponent {
  isLogoutLoading: boolean = false

  constructor(
    private router: Router,
    private snackBar: MatSnackBar,
    private authService: AuthService,
  ) { }

  logout() {
    this.isLogoutLoading = true

    this.authService.logout().subscribe({
      next: () => this.isLogoutLoading = false,
      error: () => {
        this.isLogoutLoading = false
        this.snackBar.open('Something went wrong. Please try again later...', 'Close', { duration: 3000 })
      },
      complete: () => this.router.navigate([''], { replaceUrl: true }),
    })
  }
}
