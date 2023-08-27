import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { MatSnackBar } from '@angular/material/snack-bar';

import { RegisterForm, RegisterRequest } from '../../models/register.model';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  registerForm!: FormGroup<RegisterForm>
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
    if (this.registerForm.valid) {
      this.isLoading = true;

      const request = new RegisterRequest(this.registerForm.value)
      this.authService.register(request).subscribe({
        next: () => this.isLoading = false,
        error: () => {
          this.isLoading = false
          this.snackBar.open('Please check your credentials', 'Close', { duration: 3000 })
        },
        complete: () => this.router.navigate(['auth', 'login']),
      })
    }
  }

  private initializeForm() {
    this.registerForm = this.formBuilder.nonNullable.group({
      username: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(25)]],
      password: ['', [Validators.required, Validators.minLength(4), Validators.maxLength(60)]],
    })
  }

  get username() {
    return this.registerForm.get('username')!
  }

  get password() {
    return this.registerForm.get('password')!
  }
}
