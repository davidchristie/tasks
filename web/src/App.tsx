import { BrowserRouter } from 'react-router-dom'
import LoginButton from './components/LoginButton'
import LogoutButton from './components/LogoutButton'
import UserData from './components/UserData'
import './App.css';
import logo from './logo.svg';

function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.tsx</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
          <LoginButton />
          <UserData />
          <LogoutButton />
        </header>
      </div>
    </BrowserRouter>
  );
}

export default App;
