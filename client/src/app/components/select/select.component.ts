import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-select',
  templateUrl: './select.component.html',
  styleUrls: ['./select.component.css']
})
export class SelectComponent {
  constructor() {}
  decals = new FormControl('');

  decalList: string[] = ['Brown', 'Brown3', 'Commuter', 'Gated', 'Gold', 'Silver']

  ngOnInit(): void { 
}
}
