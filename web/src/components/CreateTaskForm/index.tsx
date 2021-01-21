import { Button, TextField } from "@material-ui/core"
import { ChangeEventHandler, FormEventHandler, useState } from "react"
import { TasksDocument, useCreateTaskMutation } from "../../generated/graphql"

export default function CreateTaskForm() {
  const [createTask, { loading }] = useCreateTaskMutation({
    refetchQueries: [
      { query: TasksDocument },
    ],
  })
  const [text, setText] = useState("")

  const handleFormSubmit: FormEventHandler = async (event) => {
    event.preventDefault()
    await createTask({
      variables: {
        input: {
          text,
        }
      }
    })
    setText("")
  }

  const handleTextChange: ChangeEventHandler = (event) => {
    if (event.target instanceof HTMLInputElement) {
      setText(event.target.value)
    }
  }

  return (
    <form onSubmit={handleFormSubmit}>
      <TextField onChange={handleTextChange} placeholder="New task" value={text} />
      <Button disabled={loading || text.trim() === ""} type="submit">
        {loading ? "Creating" : "Create"}
      </Button>
    </form>
  )
}
