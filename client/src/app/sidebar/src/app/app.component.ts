import { Component, ViewEncapsulation } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
//import { FormsModule } from '@angular/forms';
//import { MatButtonModule } from '@angular/material/button';

import { AccountComponent } from './account/account.component';
import { ScheduleComponent } from './schedule/schedule.component';
import { AboutComponent } from './about/about.component';
import { LinksComponent } from './links/links.component';
//import {MatFormFieldModule} from '@angular/material/form-field';
//import {CdkAccordionModule} from '@angular/cdk/accordion';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  title = 'final';

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

