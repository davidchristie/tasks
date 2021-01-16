import Topbar from './components/Topbar'
import useLoggedInUser from './hooks/useLoggedInUser'
import './App.css';
import logo from './logo.svg';

function App() {
  const loggedInUser = useLoggedInUser()

  if (loggedInUser.error) {
    return <>Error: {loggedInUser.error}</>
  }

  if (loggedInUser.loading) {
    return <>Loading</>
  }

  return (
    <div className="App">
      <Topbar loggedInUser={loggedInUser.data} />
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
      </header>
    </div>
  );
}

export default App;
