import List from '@material-ui/core/List';
import CircularProgress from '@material-ui/core/CircularProgress';
import { useTasksQuery } from '../../generated/graphql'
import TaskListItem from '../TaskListItem'

export default function TaskList() {
  const tasks = useTasksQuery()

  return (
    <div className="TaskList">
      {tasks.loading && (
        <CircularProgress />
      )}
      {tasks.data && (
        <List dense>
          {tasks.data.tasks.map(task => (
            <TaskListItem key={task.id} task={task} />
          ))}
        </List>
      )}
    </div>
  )
}
