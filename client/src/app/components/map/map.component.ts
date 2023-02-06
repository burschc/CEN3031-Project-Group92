import { Component, OnInit } from '@angular/core';
import * as Leaflet from 'leaflet'; 

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})
export class MapComponent implements OnInit{
  private map!: Leaflet.Map;
  private centroid: Leaflet.LatLngExpression = [29.643946, -82.355659];

  private initMap(): void {
    this.map = Leaflet.map('map', {
      center: this.centroid,
      zoom: 16
    });

    const tiles = Leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 18,
    minZoom: 10,
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    });

    tiles.addTo(this.map);
  }
  constructor() {}

  ngOnInit(): void {
    this.initMap();
  }
}
