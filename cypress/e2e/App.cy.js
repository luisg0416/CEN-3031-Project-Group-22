describe('Routing', () => {
  beforeEach(() => {
    cy.visit('http://127.0.0.1:3000')
  })
  it("Should visit the about page", () =>  {
    cy.get('#about').click()
    cy.url().should('include', '/about')
  })
  it("Should visit the sign-in page", () =>  {
    cy.get('#sign-in').click()
    cy.url().should('include', '/sign-in')
  })
  it("Should visit the home page", () =>  {
    cy.get('#home').click()
    cy.url().should('include', '/')
  })
})

describe('Interaction', () => {
  beforeEach(() => {
    cy.visit('http://127.0.0.1:3000')
  })
  it("Should acknowledge user wants flashcards", () =>  {
    cy.get('#flashcardsButton').click()
    cy.on('window:alert',(t)=>{
      expect(t).to.contains('clicked flashcards');
   })
  })
  it("Should acknowledge user wants games", () =>  {
    cy.get('#gamesButton').click()
    cy.on('window:alert',(t)=>{
      expect(t).to.contains('clicked games');
   })
  })
})

describe('Sign in', () => {
  beforeEach(() => {
    cy.visit('http://127.0.0.1:3000/sign-in')
  })
  it("Submit on sign-in page", () =>  {
    cy.get('#signinButton').click()
    cy.on('window:alert',(t)=>{
      expect(t).to.contains('clicked sign-in');
   })
  })
})