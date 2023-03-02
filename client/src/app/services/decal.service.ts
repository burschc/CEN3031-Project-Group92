import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Decal } from '../Decal';

@Injectable({
  providedIn: 'root'
})
export class DecalService {
  private apiUrl = 'http://localhost:4200/api/filter/decals'

  constructor(private http: HttpClient) {}

  getDecals(): Observable<Decal[]> {
    return this.http.get<Decal[]>(this.apiUrl)
  }
}
