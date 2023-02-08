# Project Group 92 Sprint 1

## Relevant Links

- [Backend Video](https://youtu.be/GyUN9wWJ888)
- [Frontend Video](https://youtu.be/2tVi8m7m4zs)

## User Stories

Use the links below to go to the issue for each user story. Alternatively, use the [projects page](https://github.com/users/burschc/projects/1).

### [Issue 1: Setup a basic Go enviornment](https://github.com/burschc/CEN3031-Project-Group92/issues/1)

As a security-conscious user, I want to have simple centralized source code access to make sure the application will not do any unwanted or nefarious actions.

### [Issue 2: Setup a basic web server](https://github.com/burschc/CEN3031-Project-Group92/issues/2)

As an app user, I can start the application through launching a command or clicking a shortcut so I have easy access to the application.

### [Issue 3: Incorporate Leaflet into application](https://github.com/burschc/CEN3031-Project-Group92/issues/3)

As an app user, I would like to view the University of Florida Parking Map upon running the application.

### [Issue 4: Filter based on parking pass type](https://github.com/burschc/CEN3031-Project-Group92/issues/4)

As a user, I can select what parking pass I have and only see relevant spots that I could actually park in, so I have a faster and easier time finding an appropriate spot to park.

### [Issue 5: Provide a link to where users can purchase parking passes and view information about TAPS](https://github.com/burschc/CEN3031-Project-Group92/issues/5)

As a user, I can easily see which parking pass would be most applicable to me and purchase it, as well as better understand the specifics of UF parking, so planning my trip is easier and I am less likely to accidentally break a rule.

### [Issue 8: Allow searching for buildings](https://github.com/burschc/CEN3031-Project-Group92/issues/8)

As a user, I can type in the name of my dorm/student's dorm and easily see parking around it. Additionally, as a student, I can search for buildings my classes will be in and find their parking.

### [Issue 9: Sidebar for filters and user login](https://github.com/burschc/CEN3031-Project-Group92/issues/9)

As a UF student, I want to have access to an account so that I can store my preferences for easier navigation.

As a frequent visitor of UF, I want to know when parking restrictions are lifted so that I don’t waste time and money paying for parking.

## Issues Planned to Address

We planned for sprint one to essentially allow for us to get our bearings properly with Go and Typescript, and generally get everything set up for future sprints.
As such, we primarily hoped to get Issues 1, 2, and 3 done, as they act as set up for both front and back end, and would connect them together. Other issues past the first three are mostly future plans and functionality we would hope to get to in a future sprint.

## Issues Completed

The issues that were successfully completed were (1) setting up the Go environment, (2) setting up a basic web server, (3) implementing a map into the application using Leaflet, and (9) implementing a sidebar menu with Angular.

## Explanation for Issues not Completed

Regarding the project’s frontend development, there were plans to create popup windows with relevant information for their respective buttons. This would include the user’s settings, account, UF schedules, and an About tab. However, from the libraries and/or packages that were used, none were successful in completing the task as some of the modules imported issued errors during runtime. For now, the buttons generate a popup window, but they all contain the same text in the same window size, which was difficult to adjust. Finally, with the map filters, we haven’t proposed what functions we wanted to add and where the filter section would be placed on the map. So, in the future, we will continue to search for other libraries or figure a way to build multiple modal dialog boxes in the same class structure of the TypeScript file. ![image](https://user-images.githubusercontent.com/80710960/217671081-7bc25f74-7aab-4e01-96a6-757992070b2c.png)

