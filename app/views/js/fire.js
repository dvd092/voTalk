// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyALNGoisGCGZvDVKXFH-zlvVjK04zqo4Xk",
  authDomain: "votalk-6466c.firebaseapp.com",
  projectId: "votalk-6466c",
  storageBucket: "votalk-6466c.appspot.com",
  messagingSenderId: "401812555021",
  appId: "1:401812555021:web:3dd9e678097b8a05425c44",
  measurementId: "G-R5LGCF9N4R"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);