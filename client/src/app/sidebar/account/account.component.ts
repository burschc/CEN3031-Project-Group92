import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { FormGroup, FormControl, Validators, FormArray } from '@angular/forms';

@Component({
  selector: 'account',
  templateUrl: './account.component.html',
  //styleUrls: ['/src/styles.css']

})
export class AccountComponent  {

    
  constructor(public dialog: MatDialog) { }
    openDialog() {
      const dialogRef = this.dialog.open(AccountComponent);
      console.log(dialogRef);
    }    

  //form stuff
  LogInForm!: FormGroup;
  ngOnInit() {
    this.LogInForm = new FormGroup({
      username: new FormControl(null, [Validators.required]),
      password: new FormControl(null, [Validators.required])
    });
  }

  onSubmit() {
    if (this.LogInForm.invalid) {
      alert(`Username or password is invalid!`);
      return;
    }
  }

    /*username : string ="";
    password : string ="";
    show: boolean= false;
    submit(){
      console.log("user name is " + this.username)
      this.clear();
    }

    clear(){
      this.username ="";
      this.password = "";
      this.show = true;
    } */
    

    
}
