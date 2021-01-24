import { Avatar, IconButton, MenuItem, Menu } from "@material-ui/core";
import { useState } from "react";
import { User } from "../../generated/graphql";
import useLogout from "../../hooks/useLogout";

interface Props {
  loggedInUser: User;
}

const ACCOUNT_MENU_ID = "account-menu";

export default function AccountMenu({ loggedInUser }: Props) {
  const logout = useLogout();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogoutClick = () => {
    logout();
    handleClose();
  };

  return (
    <div>
      <IconButton
        aria-controls={ACCOUNT_MENU_ID}
        aria-haspopup="true"
        aria-label="account menu"
        color="inherit"
        data-test="account-menu"
        onClick={handleOpen}
      >
        <Avatar alt={loggedInUser.name} src={loggedInUser.avatar} />
      </IconButton>
      <Menu
        anchorEl={anchorEl}
        anchorOrigin={{
          horizontal: "right",
          vertical: "top",
        }}
        id={ACCOUNT_MENU_ID}
        keepMounted
        onClose={handleClose}
        open={open}
        transformOrigin={{
          horizontal: "right",
          vertical: "bottom",
        }}
      >
        <MenuItem
          data-test="account-menu-logout"
          onClick={handleLogoutClick}
        >
          Logout
        </MenuItem>
      </Menu>
    </div>
  );
}
