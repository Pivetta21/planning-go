import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { MatSnackBar } from '@angular/material/snack-bar';

import { LoginForm, LoginRequest } from '../../models/login.model';
import { AuthService } from '../../../../core/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  loginForm!: FormGroup<LoginForm>
  isLoading: boolean = false

  constructor(
    private router: Router,
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private snackBar: MatSnackBar,
  ) { }

  ngOnInit() {
    this.initializeForm()
  }

  onSubmit() {
    if (this.loginForm.valid) {
      this.isLoading = true;

      const request = new LoginRequest(this.loginForm.value)
      this.authService.login(request).subscribe({
        next: () => this.isLoading = false,
        error: () => {
          this.isLoading = false
          this.snackBar.open('Please check your credentials', 'Close', { duration: 3000 })
        },
        complete: () => this.router.navigate(['auth', 'register']),
      })
    }
  }

  private initializeForm() {
    this.loginForm = this.formBuilder.nonNullable.group({
      username: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(25)]],
      password: ['', [Validators.required, Validators.minLength(4), Validators.maxLength(60)]],
    })
  }

  get username() {
    return this.loginForm.get('username')!
  }

  get password() {
    return this.loginForm.get('password')!
  }
}
