# Project Group 92 Sprint 2

## Relevant Links

- [Video](link)
- [Video](link)

## Work Completed this Sprint

### [Issue 10: Setup database](https://github.com/burschc/CEN3031-Project-Group92/issues/10)
This will allow us to setup a login system which will let users sign in and store their parking pass for future usage of the site, and eventually pin parking areas they prefer to use.

### [Issue 12: Implement caching of application json files](https://github.com/burschc/CEN3031-Project-Group92/issues/12)
This makes it possible for the program to cache json files that are commonly used, with it checking for updates on a given schedule, if internet access is available.

### Notable Progress on [Issue 4: Filter based on parking pass type](https://github.com/burschc/CEN3031-Project-Group92/issues/4)
This allows for users enter a given parking pass type and then be shown only the parking areas that are relevant to them. 
So far, there are handlers which list all the parking decals, and handlers which filter the json parking lot file for specific decals a user would be searching for. It additionally has required backend supporting functions to allow this to work.

## Unit Tests for Front End
- it 'should create the app': Simple test which just ensures the app is properly created.
- it 'should have as title 'UFParkingMap'': Ensures the app title is UFParkingMap.
- it 'should render title': Checks to see if the app is running.
- it 'should be created': This checks to see if the DecalService is properly created.
- it 'should create': This checks if the ButtonComponent was properly made. There is another test with the same name also testing to check if the SelectComponent is made, another for the HeaderComponent, and another for the MapComponent. 

## Cypress Tests
- Test 1 just visits the default page (currently localhost:4200), and attempts to open the dropdown menu and check the Orange option box, ensuring it the positioning matches what was expected.

## Backend Tests
- TestGitHubJSON: Pulls a json file from the api on Github. Checks for json content and verifies that the file was downloaded properly. It then cleans up by deleting the file from the cache.
- TestNBPXML: Pulls an XML file from the National Bank of Poland. As it is an XML file, the test should fail.
- TestGoogleHTML: Pulls the Google homepage as an html file. This test additionally should fail.
- TestLotsFC: Tests the FeatureCollection conversion function on the parking lots json file present on the UF api Github.
- TestLotsFCNoExist: Tests FeatureCollection on a file that does not exist. 
- TestIsFresh: Tests the IsFresh function on two copies of the same json file, with the only difference being one of them has had its metadata edited to appear older than the default update time.
- cleanup: Simply cleans the cache.

## Documentation
- "/api/filter/decal/{decal}": Returns a specific decal, represented in the URL via {decal} (ex. Red One} or "/api/filter/decal/any" which returns all of them
- "/api/filter/decals": Returns a list of all parking types.
- "/api/test": Base handler. Returns "Hello" and "World"
