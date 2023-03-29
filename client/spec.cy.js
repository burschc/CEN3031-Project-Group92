describe('Parking App Test', () => {
  it('Will apply one of the filters to the map', () => {
    
    cy.visit('localhost:4200')
    
    cy.get('App-Select').click()
    //cy.contains('Decals').click()

    //cy.visit('localhost:4200/api/filter/decals') //populates the dropdown menu with an array of decals
    //localhost:4200/api/filter/decal/{"Brown3"} //filters out the options and returns those in Brown3

  })
})