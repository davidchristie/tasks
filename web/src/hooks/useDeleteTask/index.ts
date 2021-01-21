import {
  Task,
  TasksDocument,
  useDeleteTaskMutation,
} from "../../generated/graphql";

export default function useDeleteTask(): (
  task: Pick<Task, "id">
) => Promise<void> {
  const [deleteTask] = useDeleteTaskMutation();

  return async (task: Pick<Task, "id">) => {
    await deleteTask({
      refetchQueries: [{ query: TasksDocument }],
      variables: {
        id: task.id,
      },
    });
  };
}
