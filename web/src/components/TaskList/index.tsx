import Checkbox from '@material-ui/core/Checkbox';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';
import { useTasksQuery } from '../../generated/graphql'

export default function TaskList() {
  const tasks = useTasksQuery()

  if (tasks.loading || !tasks.data) {
    return <>Loading</>
  }

  return (
    <div className="TaskList">
      <List dense>
        {tasks.data.tasks.map(task => (
          <ListItem key={task.id}>
            <ListItemAvatar>
              <Checkbox color="default" />
            </ListItemAvatar>
            <ListItemText
              primary={task.text}
              secondary={new Date(task.createdAt).toLocaleString()}
            />
            <ListItemSecondaryAction>
              <IconButton edge="end" aria-label="delete">
                <DeleteIcon />
              </IconButton>
            </ListItemSecondaryAction>
          </ListItem>
        ))}
      </List>
    </div>
  )
}
