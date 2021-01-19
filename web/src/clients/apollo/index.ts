import {
  ApolloClient,
  InMemoryCache,
  NormalizedCacheObject,
} from "@apollo/client";
import { getToken } from "../../storage/token";

const GATEWAY = process.env.REACT_APP_GATEWAY || "";

export function newApolloClient(): ApolloClient<NormalizedCacheObject> {
  return new ApolloClient({
    cache: new InMemoryCache(),
    headers: {
      Authorization: `token ${getToken()}`,
    },
    uri: `${GATEWAY}/query`,
  });
}
