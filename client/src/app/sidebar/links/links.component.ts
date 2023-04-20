import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'links',
  templateUrl: './links.component.html',
  //styleUrls: ['./links.component.css']
})
export class LinksComponent {

  constructor(public dialog: MatDialog) { }
    openDialog() {
      const dialogRef = this.dialog.open(LinksComponent);

      console.log(dialogRef);
    }
}
