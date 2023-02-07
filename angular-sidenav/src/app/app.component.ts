import { Component, ViewChild } from '@angular/core';
import { MatSidenav } from '@angular/material/sidenav';
import {BreakpointObserver} from '@angular/cdk/layout';

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
}
