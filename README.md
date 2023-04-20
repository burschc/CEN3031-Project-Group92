# CEN3031 Project - Group 92
## UF Parking Map Plus
A project that aims to make a UF parking map which is understandable and easy to use in Golang and Typescript. For those unfamiliar with the current parking structure, it becomes inconvenient for them to find the next available parking spot. Few resources provide users with details on the parking situation and the places where they can park, thus making it more difficult for every driver trying to park on campus. To address this issue, the team plans to add filtering options to the map such as defining the type of permit the user has, the time of day when certain parking spots are available, road closures, etc. It is also essential to understand the users’ needs, so the team will achieve this by collecting data and using feedback from others to refine the program. This application will make finding parking at UF quicker and easier for students, faculty, staff, and visitors.


## Team Members - 
#### Front End - 
- Natalie Valcin
- Samantha Barthelemy (sbarthelemy01)
#### Back End - 
- Christopher Bursch (burschc)
- Yovany Molina (yomole)

## Web App Requirements
- [Node.js](https://nodejs.org/en)
- [Go](https://go.dev/)
- A computer
- A web browser
- An Internet connection
- A local python installation

## Using the Web App
To start the frontend and backend of the web app and open the site:
1. Open two terminal/command prompt instances.
2. In one instance, navigate to the `client` folder of the application.
3. Enter `npm i` to install all frontend dependencies.
4. Enter `npm start` or `ng serve --open` to start up the front end.
5. In the other instance navigate to the `api` folder of the application.
6. Enter `go mod tidy` to install all backend dependencies.
7. Enter `go run platform/main.go` to start up the backend.
8. (If the site is not already open) go to a new tab in your favorite web browser and go to `localhost:4200`.

From this site you can select the decal you wish to see and type in the building you wish to locate. You can then use the map to figure out the closest legal parking spot to a target location or sit back and admire the awfulness of UF's parking system.

## Troubleshooting
If the backend of the web app throws errors that mention Python on startup, there are a few things you can do to fix them:
1. Ensure that you have Python installed by opening a new terminal/command prompt and typing `python` or `python3`. You should launch into the python command prompt view. You can exit afterwards with `exit()`.
2. If `python3` is what got you to the python command prompt view, then use the flag `-python-cmd` followed by `python3` (i.e. `go run platform/main.go -python-cmd python3`)
3. Install the `python-is-python3` package. This may cause issues with other Python-related applications, so this is not advised.
4. Use the `-ignore-python` flag. This may result in issues viewing the parking spots on the web app's map.

### A Note on Python Usage
The web application uses Python only for the `gjf` or `geojsonfixer` script. The geojson file the web application uses for the parking spots comes straight from UF servers and contain many errors according to some online geojson verification utilities. These are errors that we are not able to fix without extensive research into the geojson format. Instead, we use this Python script which goes through geojson files and fixes any issues that they might have.
