import {
  Backdrop,
  CircularProgress,
  Container,
  CssBaseline,
  Fab,
} from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import AddIcon from "@material-ui/icons/Add";
import CreateTaskForm from "./components/CreateTaskForm";
import TaskList from "./components/TaskList";
import Topbar from "./components/Topbar";
import { useLoggedInUserQuery } from "./generated/graphql";
import useTokenFromQueryString from "./hooks/useTokenFromQueryString";

const useStyles = makeStyles((theme) => ({
  backdrop: {
    background: theme.palette.background.default,
  },
  fab: {
    bottom: theme.spacing(4),
    position: "absolute",
    right: theme.spacing(4),
  },
}));

function App() {
  const classes = useStyles();

  useTokenFromQueryString();

  const loggedInUser = useLoggedInUserQuery();

  return (
    <div>
      <CssBaseline />
      {!loggedInUser.loading && (
        <>
          <Topbar loggedInUser={loggedInUser.data?.loggedInUser || null} />
          {loggedInUser.data?.loggedInUser && (
            <Container>
              <CreateTaskForm />
              <TaskList />
              <Fab aria-label="add" className={classes.fab} color="default">
                <AddIcon />
              </Fab>
            </Container>
          )}
        </>
      )}
      <Backdrop className={classes.backdrop} open={loggedInUser.loading}>
        <CircularProgress color="inherit" size={100} />
      </Backdrop>
    </div>
  );
}

export default App;
