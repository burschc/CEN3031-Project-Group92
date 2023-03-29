describe('handle dropdowns', () => {
    
    it('Dropdown with select', () => {

        cy.visit('localhost:4200')
        cy.get('app-select#filter')
        cy.get('#select-form').click()
        cy.get('#select-decal').click()
        // cy.get('mat-mdc-select-min-line.ng-tns-c90-1.ng-star-inserted').contains('Reserved').click()
        // cy.get('.mdc-list-item__primary-text').contains('Reserved').click()
        // cy.get('#select-decal').click('Reserved')
        // mdc-list-item__primary-text
        // <span class="mat-mdc-select-min-line ng-tns-c90-1 ng-star-inserted">Reserved</span>
        // <span class="mat-mdc-select-min-line ng-tns-c90-1 ng-star-inserted">Reserved</span>

    })
})


