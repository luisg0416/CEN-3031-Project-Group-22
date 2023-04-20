import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import About from "./pages/About";
import Home from "./pages/Home";
import Navbar from "./Navbar";
import SignIn from "./pages/SignIn";
import FlashcardsMenu from "./pages/FlashcardsMenu";
import Flashcards from "./pages/Flashcards";
import QuizzesMenu from "./pages/QuizzesMenu";
import CreateAccount from "./pages/CreateAccount"
import Quiz from "./pages/Quiz";
import './Navbar.css';
import './Home.css';
import './About.css';
import './Quiz.css';

const root = ReactDOM.createRoot(document.querySelector("#application")!);
const PHYSICS_FLASHCARDS = [
  {
      id: 1,
      question: 'The slope of a velocity versus time graph represents:' ,
      answer: 'acceleration',
      options: [
          'acceleration',
          'velocity',
          'position',
          'displacement',
      ]
  },
  {
      id: 2,
      question: 'Which of the following is an EXTENSIVE property of matter:' ,
      answer: 'mass',
      options: [
          'temperature',
          'density',
          'mass',
          'viscosity',
      ]
  },
  {
      id: 3,
      question: 'Which of the following is closest to the conversion of 900 millibars into atmospheres:' ,
      answer: '0.9',
      options: [
          '0.09',
          '0.9',
          '9.0',
          '90.0',
      ]
  }
]

const OS_FLASHCARDS = [
  {
      id: 1,
      question: 'In Operating Systems, which of the following is/are CPU scheduling algorithms?' ,
      answer: 'all of the above',
      options: [
          'round robin',
          'priority',
          'shortest job first',
          'all of the mentioned',
      ]
  },
  {
      id: 2,
      question: 'CPU scheduling is the basis of' ,
      answer: 'multiprogramming operating systems',
      options: [
          'multiprogramming operating systems',
          'larger memory sized systems',
          'multiprocessor systems',
          'none of the mentioned',
      ]
  },
  {
      id: 3,
      question: 'Which one of the following errors will be handle by the operating system?' ,
      answer: 'all of the mentioned',
      options: [
          'lack of paper in printer',
          'connection failure in the network',
          'power failure',
          'all of the mentioned',
      ]
  }
]

const OS_QUESTIONS = [
  {
    questionText: 'Where is the operating system placed in the memory?',
    answerOptions: [
      { answerText: 'in the low memory', isCorrect: false },
      { answerText: 'in the high memory', isCorrect: false },
      { answerText: 'either low or high memory (depending on the location of interrupt vector)', isCorrect: true },
      { answerText: 'none of the above', isCorrect: false },
    ],
  },
  {
    questionText: 'If a process fails, most operating system write the error information to a ______',
    answerOptions: [
      { answerText: 'new file', isCorrect: false },
      { answerText: 'log file', isCorrect: true },
      { answerText: 'another running process', isCorrect: false },
      { answerText: 'none of the above', isCorrect: false },
    ],
  },
  {
    questionText: 'Which one of the following is not a real time operating system?',
    answerOptions: [
      { answerText: 'Palm OS', isCorrect: true },
      { answerText: 'RTLinux', isCorrect: false },
      { answerText: 'QNX', isCorrect: false },
      { answerText: 'VxWorks', isCorrect: false },
    ],
  },
  {
    questionText: 'What does OS X has?',
    answerOptions: [
      { answerText: 'monolithic kernel with modules', isCorrect: false },
      { answerText: 'microkernel', isCorrect: false },
      { answerText: 'monolithic kernel', isCorrect: false },
      { answerText: 'hybrid kernel', isCorrect: true },
    ],
  },
];

const Physics_Questions = [
  {
    questionText: 'For the hydrogen atom, which series describes electron transitions to the N=1 orbit, the lowest energy electron orbit? Is it the:',
    answerOptions: [
      { answerText: 'Balmer series', isCorrect: false },
      { answerText: 'Paschen series', isCorrect: false },
      { answerText: 'Lyman series', isCorrect: true },
      { answerText: 'Pfund series', isCorrect: false },
    ],
  },
  {
    questionText: 'The work done by a friction force is:',
    answerOptions: [
      { answerText: 'always positive', isCorrect: false },
      { answerText: 'always negative', isCorrect: true },
      { answerText: 'always zero', isCorrect: false },
      { answerText: 'depends on situation', isCorrect: false },
    ],
  },
  {
    questionText: 'As defined in physics, work is:',
    answerOptions: [
      { answerText: 'scalar quantity', isCorrect: true },
      { answerText: 'always a positive quantity', isCorrect: false },
      { answerText: 'a vector quantity', isCorrect: false },
      { answerText: 'always zero', isCorrect: false },
    ],
  },
  {
    questionText: 'Two forces have magnitudes of 11 newtons and 5 newtons. The magnitude of their sum could NOT be equal to which of the following values?',
    answerOptions: [
      { answerText: '16 newtons', isCorrect: false },
      { answerText: '9 newtons', isCorrect: false },
      { answerText: '7 newtons', isCorrect: false },
      { answerText: '5 newtons', isCorrect: true },
    ],
  },
];

root.render(
  <BrowserRouter>
    <Navbar />
    <div className="content">
      <Routes>
        <Route index element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/sign-in" element={<SignIn />} />
        <Route path="/sign-in/create-acc" element={<CreateAccount />} />
        <Route path="/flashcards" element={<FlashcardsMenu/>} />
        <Route path="/flashcards/cop4600" element={<Flashcards props={OS_FLASHCARDS}/>} />
        <Route path="/flashcards/phy2053" element={<Flashcards props={PHYSICS_FLASHCARDS}/>} />
        <Route path="/quizzes" element={<QuizzesMenu/>} />
        <Route path="/quizzes/cop4600" element={<Quiz props={OS_QUESTIONS}/>} />
        <Route path="/quizzes/phy2053" element={<Quiz props={Physics_Questions}/>} />
      </Routes>
    </div>
  </BrowserRouter>
);
