import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
//import { SettingComponent } from './setting.component';

@Component({
  selector: 'setting',
  templateUrl: './setting.component.html',
  //styleUrls: ['/src/styles.css']

})
export class SettingComponent  {
  constructor(public dialog: MatDialog) { }

    openDialog() {
      const dialogRef = this.dialog.open(SettingComponent);

      console.log(dialogRef);
    }
}
