import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';

import { IProfileModel } from '../../models/profile.model';
import { ProfileService } from '../../services/profile.service';

@Component({
  selector: 'app-find',
  templateUrl: './find.component.html',
  styleUrls: ['./find.component.css']
})
export class FindComponent implements OnInit {
  profile$!: Observable<IProfileModel>
  hasError: boolean = false

  constructor(
    private profileService: ProfileService,
  ) { }

  ngOnInit() {
    this.fetchProfile()
  }

  fetchProfile() {
    this.profile$ = this.profileService.find()
  }
}
