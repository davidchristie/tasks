import { DialogTitle, IconButton, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import { Delete as DeleteIcon } from "@material-ui/icons";
import { Task } from "../../generated/graphql";
import useDeleteTask from "../../hooks/useDeleteTask";

const useStyles = makeStyles((theme) => ({
  deleteButton: {
    position: "absolute",
    right: theme.spacing(1),
    top: theme.spacing(1),
  },
}));

interface Props {
  children: React.ReactNode;
  id: string;
  onClose: () => void;
  task: Pick<Task, "id">;
}

export default function EditTaskDialog({
  children,
  onClose,
  task,
  ...otherProps
}: Props) {
  const classes = useStyles();

  const deleteTask = useDeleteTask();

  const handleDeleteButtonClick = async () => {
    onClose();
    await deleteTask(task);
  };

  return (
    <DialogTitle disableTypography {...otherProps}>
      <Typography variant="h6">{children}</Typography>
      <IconButton
        aria-label="delete"
        className={classes.deleteButton}
        onClick={handleDeleteButtonClick}
      >
        <DeleteIcon />
      </IconButton>
    </DialogTitle>
  );
}
