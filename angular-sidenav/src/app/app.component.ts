import { Component, ViewChild } from '@angular/core';
import { MatSidenav } from '@angular/material/sidenav';
import { MatDialog } from '@angular/material/dialog';
import { PopupComponent } from './popup/popup.component';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  
  title = 'angular-sidenav'; 

  isSidebarOpen = false;

  openSidebar(){
    this.isSidebarOpen = true;
  }

  closeSidebar(){
    this.isSidebarOpen = false;
  }


  constructor(private dialogRef : MatDialog){}
  openDialog(){
    this.dialogRef.open(PopupComponent)
  }


}
