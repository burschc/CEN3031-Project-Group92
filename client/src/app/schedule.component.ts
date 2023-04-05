import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'schedule',
  templateUrl: './schedule.component.html',
  //styleUrls: ['/src/styles.css']

})
export class ScheduleComponent  {
  constructor(public dialog: MatDialog) { }

    openDialog() {
      const dialogRef = this.dialog.open(ScheduleComponent);

      console.log(dialogRef);
    }
}
