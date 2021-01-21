import {
  CreateTask,
  TasksDocument,
  useCreateTaskMutation,
} from "../../generated/graphql";

export default function useCreateTask() {
  const [createTask] = useCreateTaskMutation({
    refetchQueries: [{ query: TasksDocument }],
  });

  return async (input: CreateTask): Promise<void> => {
    await createTask({
      variables: {
        input,
      },
    });
  };
}
