import { Component, EventEmitter,OnInit, Output } from '@angular/core';
import { FormControl } from '@angular/forms';
import { DecalService } from '../../services/decal.service';
import { Building } from 'src/app/Building';
import { startWith,map } from 'rxjs/operators';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit{
  buildings = new FormControl('');
  buildingList: Building[] = [];
  buildingSelected!: string;
  @Output() buildingEvent = new EventEmitter<string>();

  constructor(private data: DecalService) {}

  ngOnInit(): void { 
    this.getAllData();
  }

  getAllData() {
    this.data.getData().subscribe((res: any) => {
      this.buildingList = res;
      console.log(res);
    })
  }

  onSelected() {
    this.buildingEvent.emit(this.buildingSelected);
  }
}
