# Sprint 3

### Frontend:
For the frontend we have began working on delivering the actual purpose of our application which is to provide a light-weight studying application for UF students 

We have added the flash cards studying resource.
Additionally, we have updated the homepage so that clicking the homepage buttons perform actual navigation to the features of the website.

Routing, Interaction, Sign in are the three suite of tests we performed using cypress
* The *routing* test ensures proper navigation of each element in navigation bar by inspecting inspecting the URL for the intended address
* The *interaction* test catches the input of the homepage buttons which navigate to the different features of the website. We have provided additional testing for flashcards in which interactivity is simply tested by finding the card element, clicking it, and asserting that the card has in fact been flipped.
* The *sign in* test catches the input of the sign in submission button

-------
### Backend:
##### Work completed:

For this sprint, our goal was to refine the functions we made in the last sprint and add a couple of needed functionalities.  

Functions that we currently have:
1. Detects when the golang api is up and running. Send GET request to /api
2. Creates a flash card and posts the data on an https address. Send POST to /api/flashCards
3. Gets the list of flashcards from an https address. Send GET request to /api/flashCards
4. Get a certain flashcard at a certain ID number. Send GET request to /api/flashCards/:id
5. Deletes a certain flashcard at a certain ID number. Send DELETE request to /api/flashCards/:id

Unit tests:
1. Test if the server is up and running.
2. Test if the flashcard has successfully been posted. "func TestCreateFlashCard(t *testing.T)"
3. Test if the flashcard data on the server matches the data in the test. "func TestGetFlashCards(t *testing.T)"
4. Test if the flashcard data at a certain ID on the server matches the data in the test. "func TestGetFlashCardsID(t *testing.T)"
5. Test if the flashcard data was successfully deleted.
