import './App.css';
import Home from './components/home/home.js'
import Article from "./components/article/article";
function App() {
  return (
    <div className="App">
      <div className="AppHome">
        <Article />
      </div>
      <div className="AppNavigation">
          底部
      </div>
    </div>
  );
}

export default App;
