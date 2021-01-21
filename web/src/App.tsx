import {
  Backdrop,
  CircularProgress,
  Container,
  CssBaseline,
  Fab,
} from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import AddIcon from "@material-ui/icons/Add";
import { useState } from "react";
import AddTaskDialog from "./components/AddTaskDialog";
import TaskList from "./components/TaskList";
import Topbar from "./components/Topbar";
import { useLoggedInUserQuery } from "./generated/graphql";
import useTokenFromQueryString from "./hooks/useTokenFromQueryString";

const useStyles = makeStyles((theme) => ({
  backdrop: {
    background: theme.palette.background.default,
  },
  container: {
    marginTop: theme.spacing(4),
  },
  fab: {
    bottom: theme.spacing(4),
    position: "absolute",
    right: theme.spacing(4),
  },
}));

export default function App() {
  useTokenFromQueryString();

  const classes = useStyles();

  const [addTaskDialogOpen, setAddTaskDialogOpen] = useState(false);

  const loggedInUser = useLoggedInUserQuery();

  const handleAddTaskButtonClick = () => {
    setAddTaskDialogOpen(true);
  };

  const handleAddTaskDialogClose = () => {
    setAddTaskDialogOpen(false);
  };

  return (
    <div>
      <CssBaseline />
      {!loggedInUser.loading && (
        <>
          <Topbar loggedInUser={loggedInUser.data?.loggedInUser || null} />
          {loggedInUser.data?.loggedInUser && (
            <Container className={classes.container} maxWidth="sm">
              <TaskList />
              <Fab
                aria-label="add"
                className={classes.fab}
                color="default"
                onClick={handleAddTaskButtonClick}
              >
                <AddIcon />
              </Fab>
            </Container>
          )}
        </>
      )}
      <AddTaskDialog
        onClose={handleAddTaskDialogClose}
        open={addTaskDialogOpen}
      />
      <Backdrop className={classes.backdrop} open={loggedInUser.loading}>
        <CircularProgress color="inherit" size={100} />
      </Backdrop>
    </div>
  );
}
