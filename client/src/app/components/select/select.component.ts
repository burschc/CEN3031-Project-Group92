import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { DecalService } from 'src/app/services/decal.service';
import { Decal } from 'src/app/Decal';

@Component({
  selector: 'app-select',
  templateUrl: './select.component.html',
  styleUrls: ['./select.component.css']
})
export class SelectComponent implements OnInit {
decals = new FormControl('');
decalList: Decal[] = [];

constructor(private decalService: DecalService) {}

ngOnInit(): void { 
  this.decalService.getDecals().subscribe((decals) => this.decalList = decals);
}
}
