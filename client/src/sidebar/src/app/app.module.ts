import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatDialogModule } from '@angular/material/dialog';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatListModule } from '@angular/material/list';
import { MatSidenavModule} from '@angular/material/sidenav';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatToolbarModule } from '@angular/material/toolbar';
import { RouterModule, Routes }   from '@angular/router';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatTooltipModule } from '@angular/material/tooltip';

import { SettingComponent } from './setting.component';
import { AccountComponent } from './account.component';
import { ScheduleComponent } from './schedule.component';
import { AboutComponent } from './about.component';

import {MatFormFieldModule} from '@angular/material/form-field';
import {CdkAccordionModule} from '@angular/cdk/accordion';

@NgModule({
  declarations: [
    AppComponent, SettingComponent, AccountComponent, ScheduleComponent, AboutComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatDialogModule,
    FormsModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatListModule,
    MatSidenavModule,
    MatIconModule,
    MatMenuModule,
    MatToolbarModule,
    RouterModule,
    RouterModule.forRoot([]),
    MatExpansionModule,
    MatTooltipModule,
    MatFormFieldModule,
    CdkAccordionModule
  
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
