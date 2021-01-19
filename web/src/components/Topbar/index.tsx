import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import { User } from '../../generated/graphql';
import useLogin from '../../hooks/useLogin';
import AccountMenu from '../AccountMenu'

interface Props {
  loggedInUser: User | null
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  }),
);

export default function Topbar({ loggedInUser }: Props) {
  const classes = useStyles();
  const login = useLogin()

  return (
    <AppBar color="default" position="fixed">
      <Toolbar>
        <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
          <MenuIcon />
        </IconButton>
        <Typography variant="h6" className={classes.title}>
          Tasks
          </Typography>
        {loggedInUser && (
          <AccountMenu loggedInUser={loggedInUser} />
        )}
        {!loggedInUser && (
          <Button color="inherit" onClick={login}>Login</Button>
        )}
      </Toolbar>
    </AppBar>
  );
}
