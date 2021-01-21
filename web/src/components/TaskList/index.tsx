import { CircularProgress, List, } from '@material-ui/core';
import { useTasksQuery } from '../../generated/graphql'
import TaskListItem from '../TaskListItem'

export default function TaskList() {
  const tasks = useTasksQuery()

  return (
    <div>
      {tasks.loading && (
        <CircularProgress />
      )}
      {tasks.data && (
        <List>
          {tasks.data.tasks.map(task => (
            <TaskListItem key={task.id} task={task} />
          ))}
        </List>
      )}
    </div>
  )
}
