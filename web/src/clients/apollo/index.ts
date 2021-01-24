import { ApolloClient, InMemoryCache, createHttpLink } from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { getToken } from "../../storage/token";

const GATEWAY = process.env.REACT_APP_GATEWAY || "";

export function newApolloClient() {
  const httpLink = createHttpLink({
    uri: `${GATEWAY}/query`,
  });
  const authLink = setContext((_, { headers }) => {
    const token = getToken();
    return {
      headers: {
        ...headers,
        authorization: token ? `token ${token}` : undefined,
      },
    };
  });

  return new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });
}
