import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Lot } from '../Lot';
import { Building } from '../Building';

@Injectable({
  providedIn: 'root'
})
export class DecalService {
  private decalUrl = 'http://localhost:4200/api/filter/decals'
  private apiUrl = 'http://localhost:4200/api/filter/decal/'
  private listBuildingURL = 'http://localhost:4200/api/search/offline/*'
  private buildingUrl = 'http://localhost:4200/api/search/offline/'

  constructor(private http: HttpClient) {}

  getDecals(): Observable<any> {
    return this.http.get(this.decalUrl);
  }

  getData(): Observable<Building[]>{
    return this.http.get<Building[]>(this.listBuildingURL);
  }
  
  getParkingLots(): Observable<Lot[]> {
    return this.http.get<Lot[]>(this.apiUrl);
  }
  
  getBuildings(): Observable<Building[]> {
    return this.http.get<Building[]>(this.listBuildingURL);
  }
 
}
