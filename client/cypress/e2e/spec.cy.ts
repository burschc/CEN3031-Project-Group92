describe('Parking App Test', () => {  
  it('Will apply one of the flters to the map', () => {


    cy.intercept(
      {
        method: 'GET',
        url: 'http://localhost:8080/api/*'
      }
    ).as("api_requests")

    cy.visit('http://localhost:4200')//populates the dropdown menu with an array of decals
    cy.get('App-Select').click()
    cy.get('#mat-option-2 > .mat-pseudo-checkbox').click()
    cy.get('#mat-option-2 > .mdc-list-item__primary-text').contains('Orange')
    //cy.contains('Decals').click()
    //localhost:4200/api/filter/decal/{"Brown3"} //filters out the options and returns those in Brown3

  })
})