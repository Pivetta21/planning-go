import { Component, OnInit } from '@angular/core';
import { Observable, tap } from 'rxjs';
import { ISessionModel } from '../../models/session.model';
import { SessionService } from '../../services/session.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {
  sessions$!: Observable<ISessionModel[]>
  isLoading: boolean = true

  constructor(
    private sessionService: SessionService,
    private snackbar: MatSnackBar,
  ) { }

  ngOnInit() {
    this.fetchSessions()
  }

  fetchSessions() {
    this.isLoading = true
    this.sessions$ = this.sessionService.list().pipe(
      tap(() => this.isLoading = false)
    )
  }

  deleteSession(identifier: string) {
    this.sessionService.delete(identifier).subscribe({
      next: () => this.fetchSessions(),
      error: (err) => this.snackbar.open(err, 'Close', { duration: 3000 })
    })
  }

  showDeleteTooltip(current: boolean): string {
    return current
      ? "Current session, this action is only possible through logout."
      : "This action will delete the selected session."
  }
}
