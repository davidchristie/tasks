import { Checkbox, CheckboxProps } from "@material-ui/core";
import { ChangeEventHandler, useEffect, useState } from "react";
import { Task } from "../../generated/graphql";
import useUpdateTask from "../../hooks/useUpdateTask";

type Props = CheckboxProps & {
  task: Pick<Task, "done" | "id">;
};

export default function TaskDoneCheckbox({ task, ...otherProps }: Props) {
  const [done, setDone] = useState(task.done);

  useEffect(() => setDone(task.done), [task.done]);

  const updateTask = useUpdateTask();

  const handleChange: ChangeEventHandler<HTMLInputElement> = async (event) => {
    const nextDone = event.target.checked;
    setDone(nextDone);
    await updateTask({
      done: nextDone,
      id: task.id,
    });
  };

  return (
    <Checkbox
      {...otherProps}
      checked={done}
      data-test="task-done-checkbox"
      onChange={handleChange}
    />
  );
}
