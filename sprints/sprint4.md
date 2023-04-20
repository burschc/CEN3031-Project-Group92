# Project Group 92 Sprint 4

## Relevant Links

- [Video](https://youtu.be/G2GV9QhWlpQ)

## Work Completed this Sprint

### [Issue 8: Allow searching for buildings](https://github.com/burschc/CEN3031-Project-Group92/issues/8)
Allow the search bar to look for specific buildings, and thus see what parking areas are around them.

### [Pull Request 26: Adds cookies for login, tests for accounts, and more features for accounts](https://github.com/burschc/CEN3031-Project-Group92/pull/26)
This fixes some mistakes with the original login system and adds test cases for it. 
Additionally adds cookies.go, which allows for the browser to remember if a user is logged in, and as such the basic database has been updated with the appropriate functions to call for setting and expiring cookies.
User decal preferences can also be set with /api/account/set/passtype/{passtype} and can be gotten at /api/account/get/passtype

## Unit Tests for Front End
- it 'should create the app': Simple test which just ensures the app is properly created.
- it 'should have as title 'UFParkingMap'': Ensures the app title is UFParkingMap.
- it 'should render title': Checks to see if the app is running.
- it 'should be created': This checks to see if the DecalService is properly created.
- it 'should create': This checks if the ButtonComponent was properly made. There is another test with the same name also testing to check if the SelectComponent is made, another for the HeaderComponent, and another for the MapComponent. 

## Cypress Tests
- Dropdowns.cy.js visits the default page (currently localhost:4200), and attempts to open the dropdown menu and check a random parking type option box, ensuring it the positioning matches what was that type should be.
- HandleSearchBuildingByName enters a building name and makes sure it is shown
- HandleSearchBuildingByNumber enters a building number and makes sure it is shown

## Backend Tests
- TestGitHubJSON: Pulls a json file from the api on Github. Checks for json content and verifies that the file was downloaded properly. It then cleans up by deleting the file from the cache.
- TestNBPXML: Pulls an XML file from the National Bank of Poland. As it is an XML file, the test should fail.
- TestGoogleHTML: Pulls the Google homepage as an html file. This test additionally should fail.
- TestLotsFC: Tests the FeatureCollection conversion function on the parking lots json file present on the UF api Github.
- TestLotsFCNoExist: Tests FeatureCollection on a file that does not exist. 
- TestIsFresh: Tests the IsFresh function on two copies of the same json file, with the only difference being one of them has had its metadata edited to appear older than the default update time.
- TestDecalDevTypesHandler: Tests the decal types using the list of decal types in the UF api Github's parking_polys.json file. Should result in a list containing the expected number of lots.
- TestFindDecalHandlerPresent: Tests the find decal handler on decals that exist. Should return a feacture collection that is not empty. 
- TestFindDecalHandlerAbsent: Tests the find decal handler on decals that do not exist. Should return an empty feature collection.
- TestSignup: Tests sigining in as two different users. 
- TestSignin: Tests logging in as those same two users.
- TestPreexistingUser: Attempts to create a duplicate user with identical credentials to the first user, and create a duplicate user of the second user using the wrong password. Both should fail and return a StatusConflict http error.
- TestInvalidCredentials: Attempts to sign in to the first user with a slightly wrong password, then attempts to log into the second user's account with the first user's password. Both should fail and reutrn a StatusUnauthorized http error.

## Documentation
- "/api/filter/decal/{decal}": Returns a feature collection of all lots matching the given decal
- "/api/filter/decals": Returns a list of all entered decal types.
- "/api/filter/dev/decals": Returns as list of all decal types listed in the parkingJSON file as a json array to the requester. Kept for development purposes.
- "/api/signup": Given a form, attempts to create a user in the database with the given username and password, with the default parking pass type -1 (any/unspecified). Returns an error if it was unable.
- "/api/login": Requests a username and password, and if given a valid pair, returns that login was successful. 
- "/api/version": Returns some information about the current version and the authors.
- "/api/account/set/passtype/{passtype}": Given the name of a passtype and with a logged in user, sets their internal passtype to whatever was selected. 
- "/api/account/get/passtype": Returns the passtype stored in a users account.
