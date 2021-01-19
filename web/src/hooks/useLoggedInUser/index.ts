import { gql, useQuery } from "@apollo/client";
import { User } from "../../types";

interface LoggedInUserQuery {
  loggedInUser: {
    avatar: string;
    id: string;
    name: string;
  };
}

interface Result {
  data: User | null;
  error: Error | null;
  loading: boolean;
}

const LOGGED_IN_USER_QUERY = gql`
  query LoggedInUser {
    loggedInUser {
      avatar
      id
      name
    }
  }
`;

export default function useLoggedInUser(): Result {
  const loggedInUser = useQuery<LoggedInUserQuery>(LOGGED_IN_USER_QUERY);
  return {
    data: loggedInUser.data?.loggedInUser || null,
    error: loggedInUser.error || null,
    loading: loggedInUser.loading,
  };
}
