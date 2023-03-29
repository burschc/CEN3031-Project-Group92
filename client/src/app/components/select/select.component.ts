import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormControl } from '@angular/forms';
import { DecalService } from '../../services/decal.service';
import { Decal } from 'src/app/Decal';

@Component({
  selector: 'app-select',
  templateUrl: './select.component.html',
  styleUrls: ['./select.component.css']
})
export class SelectComponent implements OnInit {
  decals = new FormControl('');
  decalList: Decal[] = [];
  decalSelected!: string;
  @Output() decalEvent = new EventEmitter<string>();
  
  constructor(private data: DecalService) {}

  ngOnInit(): void { 
    this.data.getDecals().subscribe(decals => this.decalList = decals);
  }

  onSelected() {
    this.decalEvent.emit(this.decalSelected);
  }

}
