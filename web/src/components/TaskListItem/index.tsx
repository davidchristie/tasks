import {
  IconButton,
  ListItem,
  ListItemIcon,
  ListItemSecondaryAction,
  ListItemText,
  Paper,
} from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import { Edit as EditIcon } from "@material-ui/icons";
import { FocusEventHandler, MouseEventHandler, useRef, useState } from "react";
import { Task } from "../../generated/graphql";
import EditTaskDialog from "../EditTaskDialog";
import TaskDoneCheckbox from "../TaskDoneCheckbox";
import TaskTextField from "../TaskTextField";

interface Props {
  divider?: boolean;
  task: Pick<Task, "done" | "id" | "text">;
}

const stopPropagation: MouseEventHandler = (event) => {
  event.stopPropagation();
};

const useStyles = makeStyles((theme) => ({
  primary: {
    marginRight: theme.spacing(4),
  },
  root: {
    background: "transparent",
  },
}));

export default function TaskListItem({ divider, task }: Props) {
  const classes = useStyles();

  const [editTaskDialogOpen, setEditTaskDialogOpen] = useState(false);

  const textInputRef = useRef<HTMLInputElement | null>(null);

  const [focused, setFocused] = useState(false);
  const [hovered, setHovered] = useState(false);

  const shadow = focused || hovered;

  const handleEditTaskDialogClose = () => {
    setEditTaskDialogOpen(false);
  };

  const handleRootBlur: FocusEventHandler<HTMLDivElement> = async () => {
    setFocused(false);
  };

  const handleRootClick: MouseEventHandler = (event) => {
    if (textInputRef.current) {
      textInputRef.current.focus();
    }
  };

  const handleRootDoubleClick: MouseEventHandler = (event) => {
    if (textInputRef.current && textInputRef.current !== event.target) {
      textInputRef.current.select();
    }
  };

  const handleRootFocus: FocusEventHandler<HTMLDivElement> = () => {
    setFocused(true);
  };

  const handleRootMouseEnter: MouseEventHandler<HTMLDivElement> = () => {
    setHovered(true);
  };

  const handleRootMouseLeave: MouseEventHandler<HTMLDivElement> = () => {
    setHovered(false);
  };

  const handleSecondaryActionClick = () => {
    setEditTaskDialogOpen(true);
  };

  return (
    <>
      <Paper
        className={classes.root}
        data-test="task-list-item"
        elevation={shadow ? 3 : 0}
        onBlur={handleRootBlur}
        onClick={handleRootClick}
        onDoubleClick={handleRootDoubleClick}
        onFocus={handleRootFocus}
        onMouseEnter={handleRootMouseEnter}
        onMouseLeave={handleRootMouseLeave}
      >
        <ListItem divider={divider}>
          <ListItemIcon>
            <TaskDoneCheckbox
              color="default"
              onClick={stopPropagation}
              onDoubleClick={stopPropagation}
              task={task}
            />
          </ListItemIcon>
          <ListItemText
            primary={
              <div className={classes.primary}>
                <TaskTextField
                  data-test="task-list-item-text"
                  fullWidth
                  InputProps={{
                    disableUnderline: true,
                  }}
                  inputRef={textInputRef}
                  task={task}
                />
              </div>
            }
          />
          <ListItemSecondaryAction
            data-test="task-list-item-edit"
            onClick={handleSecondaryActionClick}
          >
            <IconButton aria-label="edit">
              <EditIcon />
            </IconButton>
          </ListItemSecondaryAction>
        </ListItem>
      </Paper>
      <EditTaskDialog
        onClose={handleEditTaskDialogClose}
        open={editTaskDialogOpen}
        task={task}
      />
    </>
  );
}
