import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { environment } from 'src/environments/environment';
import { IProfileModel } from '../models/profile.model';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {
  private profileUrl = `${environment.apiUrl}/profile`;

  constructor(
    private http: HttpClient,
  ) { }

  find(): Observable<IProfileModel> {
    return this.http.get<IProfileModel>(this.profileUrl)
  }
}
