# Sprint 2

### Frontend:

For the frontend we build a navigation bar, added frontend routing to seamlessly switch through different pages

Our homepage displays cards using the Material UI package with React

Routing, Interaction, Sign in are the three suite of tests we performed using cypress
* The *routing* test ensures proper navigation of each element in navigation bar by inspecting inspecting the URL for the intended address
* The *interaction* test catches the input of the homepage buttons which eventually will be used to access the main features of our project
* The *sign in* test catches the input of the sign in submition button

-------
### Backend:
##### Work completed:

With the pivot from doing a notetaking app to doing a flashcard/study app, our main goal this sprint was to find a way to create the data of a flashcard and posting the data of the flashcard to an https address. We also wanted to work on requesting a list of the flashcards and flashcard IDs from the http address. 

This sprint we were able to create the struct of the flashcard. We then used Fiber to create functions that:
1. Detects when the golang api is up and running.
2. Creates a flash card and posts the data on an https address. Send POST to /api/flashCards
3. Gets the list of flashcards from an https address. Send GET request to /api/flashCards
4. Get a certain flashcard at a certain ID number. Send GET request to /api/flashCards/:id

Unit tests created:
1. Test if the server is up and running.
2. Test if the flashcard has successfully been posted. "func TestCreateFlashCard(t *testing.T)"
3. Test if the flashcard data on the server matches the data in the test. "func TestGetFlashCards(t *testing.T)"
4. Test if the flashcard data at a certain ID on the server matches the data in the test. "func TestGetFlashCardsID(t *testing.T)"
