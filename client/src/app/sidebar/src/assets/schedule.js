import Map from 'ol/Map';
import Tile from 'ol/layer/Tile';
import OSM from 'ol/source/OSM';
import View from 'ol/View';
import { fromLonLat, transform } from 'ol/proj';

import Feature from 'ol/Feature';
import {Icon, Style, Fill, Stroke} from 'ol/style';
import Point from 'ol/geom/Point';
import VectorSource from 'ol/source/Vector';
import {Vector as VectorLayer} from 'ol/layer';
import { Coordinate } from 'ol/coordinate';
import { LineString } from 'ol/geom';
import Overlay from 'ol/Overlay';


 function initializeMap() {
      //creating the map and its bounds
      var map = new Map({
        target: 'map',
        layers: [
          new Tile({
            source: new OSM()
          })
        ],
        view: new View({
          center: fromLonLat([-82.3479, 29.6465]), //Reitz Union coord.
          zoom: 15,
          extent: [
            -9168178.673614483, 3456856.972914327, 
            -9165651.953933213, 3458994.9664907856
          ]
        })
        
      });

      //getting the bounds of the map to restrict it to a certain extent
      //console.log(map.getView().calculateExtent(map.getSize()));
      /*function MapExtent() { 
        console.log(map.getView().calculateExtent(map.getSize()));
        window.setTimeout(MapExtent, 5000);
      }
      MapExtent(); */



      //adding markers
      /* using this to help find the coordinates of the starting and ending points of each road closure
        map.on('click', function(evt){
          console.log(transform(evt.coordinate, 'EPSG:3857', 'EPSG:4326'));
        });
      */
      
      //drawing lines between two points
      var points = [ [-82.35362489258719, 29.64492532536866], [-82.35049811897056, 29.64488927690509] ];        
      for (var i = 0; i < points.length; i++) {
          points[i] = transform(points[i], 'EPSG:4326', 'EPSG:3857');
      }

      var featureLine = new Feature({
          geometry: new LineString(points)
      });

      var vectorLine = new VectorSource({});
      vectorLine.addFeature(featureLine);

      var vectorLineLayer = new VectorLayer({
          source: vectorLine,
          style: new Style({
              fill: new Fill({ color: '#FF0000'}),
              stroke: new Stroke({ color: '#FF0000', width: 4 })
          })
      });
      map.addLayer(vectorLineLayer);

      points = [ [-82.35027370723854, 29.64488564633872], [-82.34863481433405, 29.6448975015971] ];        
      for (var i = 0; i < points.length; i++) {
          points[i] = transform(points[i], 'EPSG:4326', 'EPSG:3857');
      }

      featureLine = new Feature({
          geometry: new LineString(points)
      });

      vectorLine = new VectorSource({});
      vectorLine.addFeature(featureLine);

      vectorLineLayer = new VectorLayer({
          source: vectorLine,
          style: new Style({
              fill: new Fill({ color: '#FF0000'}),
              stroke: new Stroke({ color: '#FF0000', width: 4 })
          })
      });
      map.addLayer(vectorLineLayer);

      points = [ [-82.34350595357903, 29.643894684525307], [-82.34362219666214, 29.642689549913456] ];        
      for (var i = 0; i < points.length; i++) {
          points[i] = transform(points[i], 'EPSG:4326', 'EPSG:3857');
      }

      featureLine = new Feature({
          geometry: new LineString(points)
      });

      vectorLine = new VectorSource({});
      vectorLine.addFeature(featureLine);

      vectorLineLayer = new VectorLayer({
          source: vectorLine,
          style: new Style({
              fill: new Fill({ color: '#FF0000'}),
              stroke: new Stroke({ color: '#FF0000', width: 4 })
          })
      });
      map.addLayer(vectorLineLayer);

      points = [ [-82.34344767151374, 29.647530584255605], [-82.33949050885526, 29.64731145646985] ];        
      for (var i = 0; i < points.length; i++) {
          points[i] = transform(points[i], 'EPSG:4326', 'EPSG:3857');
      }

      featureLine = new Feature({
          geometry: new LineString(points)
      });

      vectorLine = new VectorSource({});
      vectorLine.addFeature(featureLine);

      vectorLineLayer = new VectorLayer({
          source: vectorLine,
          style: new Style({
              fill: new Fill({ color: '#FF0000'}),
              stroke: new Stroke({ color: '#FF0000', width: 4 })
          })
      });
      map.addLayer(vectorLineLayer);
  
      //placing markers
      const pts = new VectorLayer({
        source: new VectorSource({
            features: [
            new Feature({
                geometry: new Point(fromLonLat([-82.35362489258719, 29.64492532536866])),
                name : 'Closed Until 05/02/23', //hume
                html: '<b>Closed Until 05/02/23</b>'
            }),
            new Feature({
              geometry: new Point(fromLonLat([-82.35049811897056, 29.64488927690509])),
              name : 'Closed Until 05/02/23', //graham
              html: '<b>Closed Until 05/02/23</b>'
            }),
            new Feature({
              geometry: new Point(fromLonLat([-82.35027370723854, 29.64488564633872])),
              name : 'Closed Until 05/01/23', //gale
              html: '<b>Closed Until 05/01/23</b>'
            }),
            new Feature({
              geometry: new Point(fromLonLat([-82.34863481433405, 29.6448975015971])),
              name : 'Closed Until 05/01/23', //reitz union dr
              html: '<b>Closed Until 05/01/23</b>'
            }),
            new Feature({
              geometry: new Point(fromLonLat([-82.34350595357903, 29.643894684525307])),
              name : 'Closed Until 05/01/23', //dickinson
              html: '<b>Closed Until 05/01/23</b>'
            }),
            new Feature({
              geometry: new Point(fromLonLat([-82.34362649945504, 29.642666724696838])),
              name : 'Closed Until 05/01/23', //harrell
              html: '<b>Closed Until 05/01/23</b>'
            }),
            new Feature({
              geometry: new Point(fromLonLat([-82.34344767151374, 29.647530584255605])),
              name : 'Closed Until 08/01/24', //newell
              html: '<b>Closed Until 08/01/24</b>'
            }),
            new Feature({
              geometry: new Point(fromLonLat([-82.33949050885526, 29.64731145646985])),
              name : 'Closed Until 08/01/24', //inner & 13th
              html: '<b>Closed Until 08/01/24</b>'
            })
         
            ]
        }),
        style: new Style({ //all markers will have the same style and image
            image: new Icon({
            anchor: [0.5, 1],
            crossOrigin: 'anonymous',
            src: '../assets/images/road_icon.png',
            size: [512,512],
            scale: 0.1
          })
        })
        });
        map.addLayer(pts);

    
        //popup function
        //help from https://stackoverflow.com/questions/66501146/adding-more-than-one-popup-marker-openlayers-map
        var container = document.getElementById('popup');
        var content = document.getElementById('popup-content');
        var closer = document.getElementById('popup-closer');

        var overlay = new Overlay({
          element: container,
          autoPan: true,
          autoPanAnimation: {
            duration: 250
          }
        });
        map.addOverlay(overlay);

        closer.onclick = function() {
          overlay.setPosition(undefined);
          closer.blur();
          return false;
        };

        map.on('singleclick', function(event) {
          if (map.hasFeatureAtPixel(event.pixel) === true) {
            var coordinate = event.coordinate;
            var features = map.getFeaturesAtPixel(event.pixel);
            content.innerHTML = features[0].getProperties().html;
            overlay.setPosition(coordinate);
            overlay.setPositioning('center-center');
          } else {
            overlay.setPosition(undefined);
            closer.blur();
          }
        }); 

        map.getView().fit(marker.getSource().getExtent(), {
            padding: [40, 16, 40, 16]
        });

      

  }

  export {initializeMap};
  

