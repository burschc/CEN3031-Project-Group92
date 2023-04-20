describe('handle search building by building name autocomplete', () => {
  it('passes', () => {
    cy.visit('localhost:4200')
    cy.get('app-search')
    cy.get('#select-form').click({force: true})
    cy.get('#mat-input-0').type("Peabody Hall")
    cy.get('#mat-option-18 > .mdc-list-item__primary-text').click()
  })
})

describe('handle search building by number autocomplete', () => {
  it('passes', () => {
    cy.visit('localhost:4200')
    cy.get('app-search')
    cy.get('#select-form').click({force: true})
    cy.get('#mat-input-0').type("0099")
    cy.get('#mat-option-214 > .mdc-list-item__primary-text').click()
  })
})