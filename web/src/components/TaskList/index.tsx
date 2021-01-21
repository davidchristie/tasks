import { useTasksQuery } from "../../generated/graphql";
import TaskListAccordion from "./TaskListAccordion";

export default function TaskList() {
  const tasks = useTasksQuery();

  const currentTasks = (tasks.data?.tasks || []).filter((task) => !task.done);
  const completeTasks = (tasks.data?.tasks || []).filter((task) => task.done);

  return (
    <div>
      <TaskListAccordion
        defaultExpanded
        label="Current"
        loading={tasks.loading}
        tasks={currentTasks}
      />
      <TaskListAccordion
        label="Completed"
        loading={tasks.loading}
        tasks={completeTasks}
      />
    </div>
  );
}
