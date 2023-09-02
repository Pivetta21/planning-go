import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';

import { IProfileModel, ProfileUpdateForm, ProfileUpdateRequest } from '../../models/profile.model';
import { ProfileService } from '../../services/profile.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-find',
  templateUrl: './find.component.html',
  styleUrls: ['./find.component.css']
})
export class FindComponent implements OnInit {
  profile$!: Observable<IProfileModel>
  hasError: boolean = false

  isEditView: boolean = false

  profileUpdateForm?: FormGroup<ProfileUpdateForm>
  isUpdateLoading: boolean = false

  constructor(
    private formBuilder: FormBuilder,
    private profileService: ProfileService,
    private snackbar: MatSnackBar,
  ) { }

  ngOnInit() {
    this.fetchProfile()
  }

  fetchProfile() {
    this.profile$ = this.profileService.find()
  }

  toggleEditView(username: string) {
    this.isEditView = !this.isEditView

    if (this.isEditView) {
      this.profileUpdateForm = this.formBuilder.nonNullable.group({
        username: [username, [Validators.required, Validators.minLength(3), Validators.maxLength(25)]],
      })
    }
  }

  onSubmit() {
    if (!this.profileUpdateForm)
      return

    const req = new ProfileUpdateRequest(this.profileUpdateForm.value)
    this.profileService.patch(req).subscribe({
      next: () => { this.fetchProfile() },
      error: (err) => this.snackbar.open(err, 'Close', { duration: 3000 }),
      complete: () => { this.isEditView = !this.isEditView }
    })
  }

  get username() {
    return this.profileUpdateForm?.get('username')!
  }
}
