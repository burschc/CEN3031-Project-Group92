import { Component, ViewEncapsulation} from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'about',
  templateUrl: './about.component.html',
  //styleUrls: ['/src/styles.css'],
})
export class AboutComponent  {
  constructor(public dialog: MatDialog) { }

    openDialog() {
      const dialogRef = this.dialog.open(AboutComponent);

      console.log(dialogRef);
    }

}
