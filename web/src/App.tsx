import Container from '@material-ui/core/Container'
import CreateTaskForm from './components/CreateTaskForm'
import TaskList from './components/TaskList'
import Topbar from './components/Topbar'
import { useLoggedInUserQuery } from './generated/graphql'
import useTokenFromQueryString from './hooks/useTokenFromQueryString';

function App() {
  useTokenFromQueryString()

  const loggedInUser = useLoggedInUserQuery()

  if (loggedInUser.error) {
    return <>Error: {loggedInUser.error}</>
  }

  if (loggedInUser.loading) {
    return <>Loading</>
  }

  return (
    <div className="App">
      <Topbar loggedInUser={loggedInUser.data?.loggedInUser || null} />
      {loggedInUser.data?.loggedInUser!! && (
        <Container>
          <CreateTaskForm />
          <TaskList />
        </Container>
      )}
    </div>
  );
}

export default App;
