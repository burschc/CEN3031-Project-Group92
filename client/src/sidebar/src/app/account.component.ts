import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import {FormBuilder, FormControl} from '@angular/forms';
import {FloatLabelType} from '@angular/material/form-field';
import {MatFormFieldModule} from '@angular/material/form-field';

@Component({
  selector: 'account',
  templateUrl: './account.component.html',
  //styleUrls: ['/src/styles.css']

})
export class AccountComponent  {
  /*constructor(public dialog: MatDialog) { }

    openDialog() {
      const dialogRef = this.dialog.open(AccountComponent);

      console.log(dialogRef);
    } */

    /*panelOpenState: boolean = false;

    togglePanel() {
        this.panelOpenState = !this.panelOpenState
    } */
    panelOpenState = false;


    hideRequiredControl = new FormControl(false);
    floatLabelControl = new FormControl('auto' as FloatLabelType);
    options = this._formBuilder.group({
      hideRequired: this.hideRequiredControl,
      floatLabel: this.floatLabelControl,
    });
  
    constructor(private _formBuilder: FormBuilder) {}
  
    getFloatLabelValue(): FloatLabelType {
      return this.floatLabelControl.value || 'auto';
    }

    
}
