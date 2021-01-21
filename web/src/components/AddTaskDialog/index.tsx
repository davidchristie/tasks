import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
} from "@material-ui/core";
import {
  ChangeEventHandler,
  FormEventHandler,
  useEffect,
  useState,
} from "react";
import useCreateTask from "../../hooks/useCreateTask";

interface Props {
  onClose: () => void;
  open: boolean;
}

export default function AddTaskDialog({ onClose, open }: Props) {
  const createTask = useCreateTask();

  const [text, setText] = useState("");

  useEffect(() => setText(""), [open]);

  const handleFormSubmit: FormEventHandler = async (event) => {
    event.preventDefault();
    onClose();
    await createTask({
      text,
    });
    setText("");
  };

  const handleTextChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    setText(event.target.value);
  };

  return (
    <Dialog
      aria-labelledby="add-task-dialog"
      fullWidth
      onClose={onClose}
      open={open}
    >
      <form onSubmit={handleFormSubmit}>
        <DialogTitle>Add Task</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            fullWidth
            label="Text"
            onChange={handleTextChange}
            value={text}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose} color="default">
            Cancel
          </Button>
          <Button color="default" type="submit">
            Save
          </Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}
