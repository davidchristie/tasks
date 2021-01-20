import Checkbox from '@material-ui/core/Checkbox';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';
import { Task, TasksDocument, useDeleteTaskMutation } from '../../generated/graphql'
import { MouseEventHandler } from 'react';

interface Props {
  task: Pick<Task, 'createdAt' | 'id' | 'text'>
}

export default function TaskListItem({ task }: Props) {
  const [deleteTask] = useDeleteTaskMutation({
    refetchQueries: [
      { query: TasksDocument },
    ],
    variables: {
      id: task.id,
    },
  })

  const handleDeleteButtonClick: MouseEventHandler = () => {
    deleteTask()
  }

  return (
    <ListItem className="TaskListItem">
      <ListItemAvatar>
        <Checkbox color="default" />
      </ListItemAvatar>
      <ListItemText
        primary={task.text}
        secondary={new Date(task.createdAt).toLocaleString()}
      />
      <ListItemSecondaryAction onClick={handleDeleteButtonClick}>
        <IconButton aria-label="delete" edge="end">
          <DeleteIcon />
        </IconButton>
      </ListItemSecondaryAction>
    </ListItem>
  )
}
