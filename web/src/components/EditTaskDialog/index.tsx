import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
} from "@material-ui/core";
import { Task } from "../../generated/graphql";
import TaskTextField from "../TaskTextField";
import EditTaskDialogTitle from "./EditTaskDialogTitle";

interface Props {
  onClose: () => void;
  open: boolean;
  task: Pick<Task, "id" | "text">;
}

export default function EditTaskDialog({ onClose, open, task }: Props) {
  return (
    <Dialog
      aria-labelledby="edit-task-dialog"
      fullWidth
      onClose={onClose}
      open={open}
    >
      <EditTaskDialogTitle id="edit-task-dialog" onClose={onClose} task={task}>
        Edit Task
      </EditTaskDialogTitle>
      <DialogContent>
        <TaskTextField
          autoFocus
          fullWidth
          id="text"
          label="Text"
          margin="dense"
          task={task}
        />
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose} color="default">
          Close
        </Button>
      </DialogActions>
    </Dialog>
  );
}
