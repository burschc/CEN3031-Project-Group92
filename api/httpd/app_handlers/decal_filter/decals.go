package decal_filter

//The information for this file is based on https://taps.ufl.edu/permits. They are statically defined because the JSON
//file can only describe SO much and TAPS can only be SO accurate and consistent with their information.

var ParkingDecals = map[string]Decal{
	"Green": {
		ParkingOptions: []string{
			"Green",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},
	},

	"Park and Ride": {
		ParkingOptions: []string{
			"Park and Ride",
			"Any Decal",
			"All Decals (No Red)",
		},
	},

	"Red 1": {
		ParkingOptions: []string{
			"Red One",
			"Red",
			"Any Decal",
			"All Decals (No Park and Ride)",
		},
	},

	"Red 3": {
		ParkingOptions: []string{
			"Red",
			"Any Decal",
			"All Decals (No Park and Ride)",
		},
	},

	"Brown 2": {
		ParkingOptions: []string{
			"Brown",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},
	},

	"Brown 3": {
		ParkingOptions: []string{
			"Brown 3",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
			"Brown 3 (No Official Business)",
		},
	},

	"Disabled Student": {
		ParkingOptions: []string{
			"Orange",
			"Blue",
			"Orange/Blue",
			"Green",
			"Student Green",
			"Red One",
			"Red",
			"Brown",
			"Brown 3",
			"Brown 3 (No Official Business)",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},
	},

	"Gold": {ParkingOptions: []string{
		"Gated",
		"Orange",
		"Blue",
		"Orange/Blue",
		"Green",
		"Red One",
		"Red",
		"Brown",
		"Brown 3",
		"Brown 3 (No Official Business)",
		"Any Decal",
		"All Decals (No Park and Ride)",
		"All Decals (No Red)",
	},
		Special: []SpecialRestriction{
			{
				ParkingOptions: []string{
					"Brown",
					"Brown 3",
					"Brown 3 (No Official Business)",
				},
				Description: "Except where prohibited by signage.",
			},
		},
	},

	"Silver": {
		ParkingOptions: []string{
			"Gated",
		},
	},

	"Official Business": {
		ParkingOptions: []string{
			"Orange",
			"Blue",
			"Orange/Blue",
			"Green",
			"Red One",
			"Red",
			"Brown",
			"Brown 3",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},

		Special: []SpecialRestriction{
			{
				ParkingOptions: []string{
					"Brown",
					"Brown 3",
				},
				Description: "Except where prohibited by signage.",
			},
		},
	},

	"Orange": {
		ParkingOptions: []string{
			"Orange",
			"Green",
			"Orange/Blue",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},
	},

	"Blue": {
		ParkingOptions: []string{
			"Blue",
			"Green",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},
	},

	"Medical Resident": {
		ParkingOptions: []string{
			"Orange",
			"Orange/Blue",
			"Green",
			"Red One",
			"Red",
			"Brown",
			"Brown 3",
			"Brown 3 (No Official Business)",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
			"Medical Resident",
		},
		Special: []SpecialRestriction{
			{
				ParkingOptions: []string{
					"Brown",
					"Brown 3",
					"Brown 3 (No Official Business)",
				},
				Description: "Except where prohibited by signage.",
			},
		},
	},

	"Staff Commuter": {
		ParkingOptions: []string{
			"Green",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},
		Special: []SpecialRestriction{
			{
				ParkingOptions: []string{"Green"},
				Description: "Excluding Garage VII at the O’Connell Center, excluding the Green sections in Garage 14, " +
					"and where prohibited by signage."},
		},
	},

	"Disabled Employee": {
		ParkingOptions: []string{
			"Orange",
			"Blue",
			"Orange/Blue",
			"Green",
			"Red One",
			"Red",
			"Brown",
			"Brown 3",
			"Brown 3 (No Official Business)",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},
	},

	"Carpool": {
		ParkingOptions: []string{
			"Carpool",
			"Orange",
			"Blue",
			"Orange/Blue",
			"Green",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
		},
		Special: []SpecialRestriction{
			{
				ParkingOptions: []string{
					"Orange",
					"Blue",
					"Green",
					"Any Decal",
					"All Decals (No Park and Ride)",
					"All Decals (No Red)",
				},
				Description: "After 9:30 AM.",
			},
		},
	},

	"Commercial": {
		ParkingOptions: []string{
			"Orange",
			"Blue",
			"Orange/Blue",
			"Green",
			"Red One",
			"Red",
			"Brown",
			"Brown 3",
			"Brown 3 (No Official Business)",
			"Any Decal",
			"All Decals (No Park and Ride)",
			"All Decals (No Red)",
			"Service",
			"Service (No Official Business)",
		},
		Special: []SpecialRestriction{
			{
				ParkingOptions: []string{
					"Brown",
					"Brown 3",
					"Brown 3 (No Official Business)",
				},
				Description: "Except where prohibited by signage.",
			},

			{
				ParkingOptions: []string{
					"Service",
					"Service (No Official Business)",
				},
				Description: "Unmarked contractor vehicles have a time limit of 1 hour.",
			},
		},
	},

	"Visitor": {
		ParkingOptions: []string{
			"Visitor",
			"Visitor (30 Minute Limit)",
		},
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//													STRUCTURES												  		  //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Decal is the basic structure of a decal entry that is connected to the mapped decals through the ParkingOptions object.
type Decal struct {
	ParkingOptions []string             // ParkingOptions are the JSON property names that this decal can park in.
	Special        []SpecialRestriction // Special is an array of any special restrictions the parking decal may have.
}

// SpecialRestriction describes any restrictions that may exist for a decal's parking option that annoyingly isn't in
// the json file or anywhere else.
type SpecialRestriction struct {
	ParkingOptions []string // ParkingOptions is the parking spot(s) that this restriction affects.
	Description    string   // Description is what TAPS describes the restriction as.
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//											PUBLIC UTILITY FUNCTIONS												  //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetNames() []string {
	var names []string
	for k, _ := range ParkingDecals {
		names = append(names, k)
	}

	return names
}
