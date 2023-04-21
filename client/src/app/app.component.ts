import { Component, ViewEncapsulation } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

import { AccountComponent } from './sidebar/account/account.component';
import { ScheduleComponent } from './sidebar/schedule/schedule.component';
import { AboutComponent } from './sidebar/about/about.component';
import { LinksComponent } from './sidebar/links/links.component';
import { SignupComponent } from './sidebar/signup/signup.component';
import { LandingComponent } from './sidebar/landing/landing.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./app.component.css']
})

export class AppComponent {

  title = 'sidebar';

  constructor(public dialog: MatDialog) { }

  openLinks() {
    const dialogRef = this.dialog.open(LinksComponent, {
      id: 'links',
      panelClass: 'link-style'
    });

    console.log(dialogRef);
  }

  openAccount() {
    const dialogRef = this.dialog.open(AccountComponent, {
      id: 'account',
      panelClass: 'acct-style'
    });

    console.log(dialogRef);
  }

  openSchedule() {
    const dialogRef = this.dialog.open(ScheduleComponent, {
      id: 'schedule',
      panelClass: 'sched-style'
    });

    console.log(dialogRef);
  }

  openAbout() {
    const dialogRef = this.dialog.open(AboutComponent, {
      id: 'about',
      panelClass: 'about-style'
    });

    console.log(dialogRef);
  }



} 
