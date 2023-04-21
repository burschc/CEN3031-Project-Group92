import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { LeafletModule } from '@asymmetrik/ngx-leaflet';
import { AppRoutingModule } from './app-routing.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSelectModule } from '@angular/material/select';
import { MatDialogModule } from '@angular/material/dialog';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';
import { HttpClientModule} from '@angular/common/http';
import { MatAutocompleteModule} from '@angular/material/autocomplete';
import { MatInputModule} from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatListModule } from '@angular/material/list';
import { MatSidenavModule} from '@angular/material/sidenav';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatToolbarModule } from '@angular/material/toolbar';
import { RouterModule, Routes }   from '@angular/router';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatTooltipModule } from '@angular/material/tooltip';

import {MatFormFieldModule} from '@angular/material/form-field';
import {CdkAccordionModule} from '@angular/cdk/accordion';
import { AccountComponent } from './sidebar/account/account.component';
import { ScheduleComponent } from './sidebar/schedule/schedule.component';
import { AboutComponent } from './sidebar/about/about.component';
import { LinksComponent } from './sidebar/links/links.component';
import { SignupComponent } from './sidebar/signup/signup.component';
import { LandingComponent } from './sidebar/landing/landing.component';

import { AppComponent } from './app.component';
import { MapComponent } from './components/map/map.component';
import { HeaderComponent } from './components/header/header.component';
import { SelectComponent } from './components/select/select.component';
import { ButtonComponent } from './components/button/button.component';
import { DecalService } from './services/decal.service';
import { SearchComponent } from './components/search/search.component';

 
@NgModule({
  declarations: [
    AppComponent,
    MapComponent,
    HeaderComponent,
    SelectComponent,
    ButtonComponent,
    SearchComponent,
    AccountComponent, 
    ScheduleComponent, 
    AboutComponent, 
    LinksComponent, 
    SignupComponent, 
    LandingComponent
  ],

  imports: [
    BrowserModule,
    LeafletModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSelectModule,
    MatAutocompleteModule,
    MatInputModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    MatDialogModule,
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
  providers: [DecalService],
  bootstrap: [AppComponent],
  
})
export class AppModule { }
