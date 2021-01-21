import { AppBar, Button, Toolbar, Typography } from "@material-ui/core";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import { User } from "../../generated/graphql";
import useLogin from "../../hooks/useLogin";
import AccountMenu from "../AccountMenu";

interface Props {
  loggedInUser: User | null;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  })
);

export default function Topbar({ loggedInUser }: Props) {
  const classes = useStyles();
  const login = useLogin();

  return (
    <AppBar color="default" position="static">
      <Toolbar>
        <Typography variant="h6" className={classes.title}>
          Tasks
        </Typography>
        {loggedInUser && <AccountMenu loggedInUser={loggedInUser} />}
        {!loggedInUser && (
          <Button color="inherit" onClick={login}>
            Login
          </Button>
        )}
      </Toolbar>
    </AppBar>
  );
}
