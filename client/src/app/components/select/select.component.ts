import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { DecalService } from 'src/app/services/decal.service';

@Component({
  selector: 'app-select',
  templateUrl: './select.component.html',
  styleUrls: ['./select.component.css']
})
export class SelectComponent implements OnInit {

decalList: any;
constructor(private myDecalService: DecalService) {}

 ngOnInit(): void { 
  this.fetchDecals();
  }
 
  private fetchDecals() {
    this.myDecalService.getDecal().subscribe(response => {
      this.decalList = response;
      console.log(this.decalList)
    });
  }
  // decals = new FormControl('');

  // decalList: string[] = ['Brown', 'Brown3', 'Commuter', 'Gated', 'Gold', 'Silver']
   
}
