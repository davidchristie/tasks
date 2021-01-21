import { useApolloClient } from "@apollo/client";
import { TaskFragmentDoc, UpdateTask } from "../../generated/graphql";

export default function useUpdateTaskInCache(): (
  input: UpdateTask
) => Promise<void> {
  const client = useApolloClient();

  return async (input: UpdateTask): Promise<void> => {
    const id = `Task:${input.id}`;

    const data = client.readFragment({
      fragment: TaskFragmentDoc,
      id,
    });

    console.log("data", data);

    client.writeFragment({
      data: {
        text: input.text || data.text,
      },
      fragment: TaskFragmentDoc,
      id,
    });
  };
}
