import { Routes, Route } from "react-router-dom";
import { BrowserRouter } from "react-router-dom";
import './App.css';
import Home from './components/home/home.js'
import Article from "./components/article/article";
function App() {
  return (
      <BrowserRouter>
          <Routes>
              <Route path="/fp/home" element={<Home />} />
              <Route path="/fp/article" element={<Article />} />
          </Routes>
      </BrowserRouter>
  );
}

export default App;
