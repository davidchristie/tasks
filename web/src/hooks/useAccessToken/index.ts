import { parse } from "query-string";
import { useHistory, useLocation } from "react-router-dom";

const QUERY_STRING_KEY = "access_token";
const LOCAL_STORAGE_KEY = "access_token";

export default function useAccessToken(): string | null {
  const history = useHistory();
  const location = useLocation();
  const { [QUERY_STRING_KEY]: accessToken } = parse(location.search);

  if (typeof accessToken === "string") {
    localStorage.setItem(LOCAL_STORAGE_KEY, accessToken);
    history.push("/");
  }

  return localStorage.getItem(LOCAL_STORAGE_KEY);
}
