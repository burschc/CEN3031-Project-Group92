# Project Group 92 Sprint 2

## Relevant Links

- [Video](link)

## Work Completed this Sprint

### [Issue 10: Setup database](https://github.com/burschc/CEN3031-Project-Group92/issues/10)
This will allow us to setup a login system which will let users sign in and store their parking pass for future usage of the site, and eventually pin parking areas they prefer to use.

### [Issue 12: Implement caching of application json files](https://github.com/burschc/CEN3031-Project-Group92/issues/12)
This makes it possible for the program to cache json files that are commonly used, with it checking for updates on a given schedule, if internet access is available.

### Notable Progress on [Issue 4: Filter based on parking pass type](https://github.com/burschc/CEN3031-Project-Group92/issues/4)
This allows for users enter a given parking pass type and then be shown only the parking areas that are relevant to them. 

## Unit Tests for Front End

## Cypress Tests

## Backend Tests
- TestGitHubJSON: Pulls a json file from the api on Github. Checks for json content and verifies that the file was downloaded properly. It then cleans up by deleting the file from the cache.
- TestNBPXML: Pulls an XML file from the National Bank of Poland. As it is an XML file, the test should fail.
- TestGoogleHTML: Pulls the Google homepage as an html file. This test additionally should fail.
- TestLotsFC: Tests the FeatureCollection conversion function on the parking lots json file present on the UF api Github.
- TestLotsFCNoExist: Tests FeatureCollection on a file that does not exist. 
- TestIsFresh: Tests the IsFresh function on two copies of the same json file, with the only difference being one of them has had its metadata edited to appear older than the default update time.
- cleanup: Simply cleans the cache.

## Documentation
- "/api/filter/decal/{decal}": Displays a specific decal, represented in the URL via {decal}
- "/api/filter/decals": Displays all decals. 
- "/api/test": Base handler.
