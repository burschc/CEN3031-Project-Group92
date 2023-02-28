import { Component, OnInit } from '@angular/core';
import * as Leaflet from 'leaflet'; 

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})
export class MapComponent implements OnInit{
  private map!: Leaflet.Map;
  private centroid: Leaflet.LatLngExpression = [29.64833, -82.34944];

  private initMap(): void {
    this.map = Leaflet.map('map').setView(this.centroid, 16);

  // Making a marker with a icon
  var myIcon = Leaflet.icon({
    iconUrl: 'marker-icon.png',
    iconSize: [25, 32],
    iconAnchor: [25, 32],

  });

  const marker = Leaflet.marker([0,0], {icon: myIcon}).addTo(this.map);
  const attribution = '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>';
  const tileUrl = 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png';
  const tiles = Leaflet.tileLayer(tileUrl, {attribution});

  tiles.addTo(this.map);
  marker.setLatLng(this.centroid);
  }
  constructor() {}

  ngOnInit(): void {
    this.initMap();
  }
}
