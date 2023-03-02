# Sprint 2

### Frontend:



-------
### Backend:
##### Work completed:

With the pivot from doing a notetaking app to doing a flashcard/study app, our main goal this sprint was to find a way to create the data of a flashcard and posting the data of the flashcard to an https address. We also wanted to work on requesting a list of the flashcards and flashcard IDs from the http address. 

This sprint we were able to create the struct of the flashcard. We then used Fiber to create functions that:
1. Detects when the golang api is up and running.
2. Creates a flash card and posts the data on an https address.
3. Gets the list of flashcards from an https address.
4. Get a certain flashcard at a certain ID number.

Unit tests created:
1. Test if the server is up and running.
2. Test if the flashcard has successfully been posted. "func TestCreateFlashCard(t *testing.T)"
3. Test if the flashcard data on the server matches the data in the test. "func TestGetFlashCards(t *testing.T)"
4. Test if the flashcard data at a certain ID on the server matches the data in the test. "func TestGetFlashCardsID(t *testing.T)"
