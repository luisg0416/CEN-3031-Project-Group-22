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
    it("Should visit flashcards", () =>  {
      cy.get('#flashcards').click()
      cy.url().should('include', '/flashcards')
    })
  })
  it("Should acknowledge user wants quizzes", () =>  {
    cy.get('#quizzesButton').click()
    cy.url().should('include', '/quizzes')
  })
})

describe('Interaction: Flashcards', () => {
  it("Select a course...", () =>  {
    cy.visit('http://127.0.0.1:3000/flashcards')
    cy.get('#cop4600Button').click()
    cy.url().should('include', '/flashcards/cop4600')
  })
  it("Should flip one of the cards", () =>  {
    cy.visit('http://127.0.0.1:3000/flashcards/cop4600')
    cy.get('#card').click()
    cy.get('#cardflip').should('exist');
  })
})

describe('Interaction: Quizzes', () => {
  it("Select a course...", () =>  {
    cy.visit('http://127.0.0.1:3000/quizzes')
    cy.get('#cop4600Button').click()
    cy.url().should('include', '/quizzes/cop4600')
  })
})

describe('Sign in', () => {
  beforeEach(() => {
    cy.visit('http://127.0.0.1:3000/sign-in')
  })
  it("Submit on sign-in page", () =>  {
    cy.get('#signinButtonSub').click()
    cy.on('window:alert',(t)=>{
      expect(t).to.contains('clicked sign-in');
    })
  })
  it("Navigates to sign-up page", () =>  {
    cy.get('#signupButton').click()
    cy.url().should('include', '/sign-in/create-acc')
  })
})

describe('Sign up', () => {
  beforeEach(() => {
    cy.visit('http://127.0.0.1:3000/sign-in/create-acc')
  })
  it("Account creation submission", () =>  {
    cy.get('#signupButtonSub').click()
    cy.on('window:alert',(t)=>{
      expect(t).to.contains('clicked sign-up');
    })
  })
  it("Navigates to sign-in pages", () =>  {
    cy.get('#signinButton').click()
    cy.url().should('include', '/sign-in')
  })
})