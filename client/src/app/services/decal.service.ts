import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DecalService {

  constructor(private http: HttpClient) { }

  getDecal() {
    return this.http.get('http://localhost:8080/api/filter/decals');
    // return this.http.get('http://localhost:8080/api/filter/decals').toPromise
  }
}
