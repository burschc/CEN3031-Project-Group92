import { Component, EventEmitter,OnInit, Output } from '@angular/core';
import { FormControl } from '@angular/forms';
import { DecalService } from '../../services/decal.service';
import { Building } from 'src/app/Building';
import { startWith,map } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { filter } from 'rxjs/operators';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit{
  formcontrol = new FormControl('');
  buildingList: Building[] = [];
  filteroptions!: Observable<Building[]>;
  buildingSelected!: string;
  @Output() buildingEvent = new EventEmitter<string>();

  constructor(private data: DecalService) {}

  ngOnInit(): void { 
    this.getAllData();
    this.filteroptions=this.formcontrol.valueChanges.pipe(
      startWith(''),map(value=>this._FILTER(value ||''))
    )
  }

  private _FILTER(value:string):Building[]{
    const searchvalue=value.toLocaleLowerCase();
    return this.buildingList.filter(option=>option.NAME.toLocaleLowerCase().includes(searchvalue) || 
    option.ABBREV.toLocaleLowerCase().includes(searchvalue) || option.BLDG.toLocaleLowerCase().includes(searchvalue));
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
