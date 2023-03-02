import { Component } from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { MatDialog } from '@angular/material/dialog';
import { SettingsComponent } from '../settings/settings.component';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent {
[x: string]: any;

  isHandset$: Observable<boolean> = this.breakpointObserver.observe(Breakpoints.Handset)
    .pipe(
      map(result => result.matches),
      shareReplay()
    );

  constructor(private breakpointObserver: BreakpointObserver) {}

  openSettings(){
    this['dialogRef'].open(SettingsComponent)
  }}


//TODO: use service component (ng generate service) to see if modal window can be generated 
/*
constructor(private dialogRef : MatDialog){}
  openDialog(){
    
    if (this.dialogRef instanceof PopupComponent) {
          this.dialogRef.open(PopupComponent)

    }

    else if (this.dialogRef instanceof AppComponent) {
      this.dialogRef.open(AppComponent)
    }
  } 
*/