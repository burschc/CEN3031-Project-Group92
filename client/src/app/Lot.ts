
export interface Lot 
{
    type: "FeatureCollection";
    features: [
        {
            type: "Feature";
            geometry: {
                type?: any;
                coordinates: [[[],[]]];
            }
            properties: {
                Lot_Class: string;
                Lot_Name: string;
            }
        }
    ]
}