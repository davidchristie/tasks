import { TextField, TextFieldProps } from "@material-ui/core"
import { ChangeEventHandler, useEffect, useState } from "react"
import { Task } from "../../generated/graphql";
import useUpdateTask from '../../hooks/useUpdateTask'

type Props = TextFieldProps & {
  task: Pick<Task, "id" | "text">
}

export default function TaskTextField({ task, ...otherProps }: Props) {
  const [text, setText] = useState(task.text)

  useEffect(() => setText(task.text), [task.text])

  const updateTask = useUpdateTask()

  const handleChange: ChangeEventHandler<HTMLInputElement> = async (event) => {
    const { value: newText } = event.target
    console.log('change', newText)
    setText(newText)
    await updateTask({
      id: task.id,
      text: newText,
    })
  }

  return (
    <TextField
      onChange={handleChange}
      value={text}
      {...otherProps}
    />
  )
}
