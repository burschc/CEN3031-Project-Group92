import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Lot } from '../Lot';

@Injectable({
  providedIn: 'root'
})
export class DecalService {
  private decalUrl = 'http://localhost:4200/api/filter/decals'
  private apiUrl = 'http://localhost:4200/api/filter/decal/'

  constructor(private http: HttpClient) {}

  getDecals(): Observable<any> {
    return this.http.get(this.decalUrl);
  }
  
  getParkingLots(): Observable<Lot[]> {
    return this.http.get<Lot[]>(this.apiUrl)
  }
  
 
}
