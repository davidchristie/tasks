import { useRef } from "react";
import { debounce } from "underscore";
import { UpdateTask, useUpdateTaskMutation } from "../../generated/graphql";
import useUpdateTaskInCache from "./useUpdateTaskInCache";

const DEBOUNCE_WAIT = 500;

export default function useUpdateTask(): (input: UpdateTask) => Promise<void> {
  const [updateTask] = useUpdateTaskMutation({
    ignoreResults: true,
  });
  const updateTaskInCache = useUpdateTaskInCache();

  const debouncedUpdateTask = useRef(debounce(updateTask, DEBOUNCE_WAIT))
    .current;

  return async (input: UpdateTask): Promise<void> => {
    updateTaskInCache(input);
    await debouncedUpdateTask({
      variables: {
        input,
      },
    });
  };
}
