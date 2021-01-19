import React from 'react';
import Avatar from '@material-ui/core/Avatar'
import IconButton from '@material-ui/core/IconButton';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import useLogout from '../../hooks/useLogout';
import { User } from '../../types';

interface Props {
  loggedInUser: User
}

export default function AccountMenu({ loggedInUser }: Props) {
  const logout = useLogout()
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div>
      <IconButton
        aria-controls="account-menu"
        aria-haspopup="true"
        aria-label="account of logged in user"
        color="inherit"
        onClick={handleOpen}
      >
        <Avatar alt={loggedInUser.name} src={loggedInUser.avatar} />
      </IconButton>
      <Menu
        anchorEl={anchorEl}
        anchorOrigin={{
          horizontal: 'right',
          vertical: 'top',
        }}
        id="account-menu"
        keepMounted
        onClose={handleClose}
        open={open}
        transformOrigin={{
          horizontal: 'right',
          vertical: 'bottom',
        }}
      >
        <MenuItem onClick={logout}>Logout</MenuItem>
      </Menu>
    </div>
  );
}
