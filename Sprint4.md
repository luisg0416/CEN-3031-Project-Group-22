# Sprint 4
### Frontend:
##### Work completed:
- Accomplished final formatting
- Added sign-in and sign-up pages that follow one another
- Fixed footer to the bottom of the homepage

Unit tests:
Routing
1. Should visit the about page
2. Should visit the sign-in page
3. Should visit the home page
Interaction
1. Acknowledge user wants flashcards
2. Acknowledge user wants quizzes
Interaction: Flashcards
1. Select a course...
2. Flip one of the cards
Interaction: Quizzes
1. Select a course...
Sign in
1. Submit on sign-in page
2. Navigates to sign-up page
Sign up
1. Account creation submission
2. Navigates to sign-in pages

-------
### Backend:
##### Work completed:

For this sprint, our goal was to refine the functions we made in the last sprint and add a couple of functionalities for the database.  

Functions that we currently have:
1. Detects when the golang api is up and running. Send GET request to /api
2. Creates a flash card and posts the data on an https address. Send POST to /api/flashCards
3. Gets the list of flashcards from an https address. Send GET request to /api/flashCards
4. Get a certain flashcard at a certain ID number. Send GET request to /api/flashCards/:id
5. Deletes a certain flashcard at a certain ID number. Send DELETE request to /api/flashCards/:id
6. Logs in a user. send GET request to /api/login
7. Initializes previous created cards from user.
8. Creates a user. sent POST request to /api/signup

Unit tests:
1. Test if the server is up and running.
2. Test if the flashcard has successfully been posted. "func TestCreateFlashCard(t *testing.T)"
3. Test if the flashcard data on the server matches the data in the test. "func TestGetFlashCards(t *testing.T)"
4. Test if the flashcard data at a certain ID on the server matches the data in the test. "func TestGetFlashCardsID(t *testing.T)"
5. Test if the flashcard data was successfully deleted.
6. Test if user exists.
7. Test if cards are properly initialized.
8. Test if user is created.
