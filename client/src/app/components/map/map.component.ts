import { Component, OnInit } from '@angular/core';
import * as Leaflet from 'leaflet'; 

@Component({
  selector: 'app-map', 
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})

export class MapComponent implements OnInit{
  constructor() {}

  serviceData="";
  private map!: Leaflet.Map;
  private centroid: Leaflet.LatLngExpression = [29.64833, -82.34944];
  geojson: Leaflet.GeoJSON<any> | null = null;


  private initMap(): void {
    
    // Making a map and tiles
    this.map = Leaflet.map('map').setView(this.centroid, 16);
    const attribution = '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>';

    const tileUrl = 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png';
    const OSM = Leaflet.tileLayer(tileUrl, {attribution});
    this.map.addLayer(OSM);
    // tiles.addTo(this.map);
    
    // Making a marker with an icon
    var myIcon = Leaflet.icon({
      iconUrl: 'marker-icon.png',
      iconSize: [25, 32],
      iconAnchor: [25, 32],
    });
    const marker = Leaflet.marker([0,0], {icon: myIcon}).addTo(this.map);
    marker.setLatLng(this.centroid);
  }

  
  getParkingLots($event: any) {
    this.serviceData = $event
    const url = 'http://localhost:8080/api/filter/decal/' + this.serviceData;

    fetch(url, {
      method: 'GET'
    })
    .then(response => response.json())
    .then(json => {
      if (this.geojson) {
        this.geojson.clearLayers();
        this.geojson.addData(json);
        this.map.fitBounds(this.geojson.getBounds());
      }
      else {
        this.geojson = Leaflet.geoJSON(json, {
          style: function(feature) {
            switch (feature?.properties.Lot_Class)
            {
              case 'Brown': return {color: 'brown', fillColor: 'brown', fillOpacity: 0.2};
              case 'Brown 3': return {color: 'brown', fillColor: 'brown', fillOpacity: 0.2};
              case 'Orange': return {color: 'orange', fillColor: 'orange', fillOpacity: 0.2};
              case 'Red': return {color: 'red', fillColor: 'red', fillOpacity: 0.2}
              case 'Red One': return {color: 'red', fillColor: 'red', fillOpacity: 0.2}
              case 'Blue': return {color: 'blue', fillColor: 'blue', fillOpacity: 0.2}
              case 'Green': return {color: 'green', fillColor: 'green', fillOpacity: 0.2}
              case 'Student Green': return {color: 'green', fillColor: 'green', fillOpacity: 0.2}
            }
            return {
              color: 'black',
              fillColor: 'black',
              fillOpacity: 0.2
            };
          },
          onEachFeature: function(feature, layer) {
            if (feature.properties.Lot_Name == null) {
              layer.bindPopup('Lot: I have no name :(')
            }
            else {
              layer.bindPopup('Lot: ' + feature.properties.Lot_Name);
            }  
            // layer.on('mouseover', function (e) {
            //   layer.openPopup();
            // })
            // layer.on('mouseout', function (e) {
            //   layer.closePopup();
            // })
          }
        })
        .addTo(this.map);
      }
    })
    .catch(error => {
      console.log('error!')
      console.error(error)
    });
  }
  
  ngOnInit(): void {  
    this.initMap();
  }
  
}
 
    // this.dataService.share.subscribe(x => this.serviceData = x);
  
    // const url = 'http://localhost:8080/api/filter/decal/' + this.serviceData; 

    // async function load_shapefile() {
    //   const response = await fetch(url)
    //   const shape_obj = await response.json();
    //   console.log(shape_obj);
    //   return shape_obj;
    // }

    // load_shapefile().then(Leaflet.geoJSON).then(this.map.addLayer.bind(this.map));
   
    // fetch(url, {
    //   method: 'GET'
    // })
    // .then(response => response.json())
    // .then(json => {
    //   console.log(json)
    //   var geojson = Leaflet.geoJSON(json, {
    //     style: function (feature) {
    //       return {
    //         fillOpacity: 0
    //       };
    //     }
    //   }).addTo(this.map)
    // })
    // .catch(error => console.log(error.message));