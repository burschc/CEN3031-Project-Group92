import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { initializeMap } from 'src/assets/schedule.js';

@Component({
  selector: 'schedule',
  templateUrl: './schedule.component.html',
  //styleUrls: ['./schedule.js']
})
export class ScheduleComponent {
  constructor(public dialog: MatDialog) {  }
    openDialog() {
      const dialogRef = this.dialog.open(ScheduleComponent);
      console.log(dialogRef);
    }

    ngOnInit(): void {
      initializeMap();
    }

}
