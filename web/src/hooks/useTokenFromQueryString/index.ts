import { parse } from "query-string";
import { useLocation } from "react-router-dom";
import { setToken } from "../../storage/token";

const QUERY_STRING_KEY = "token";

export default function useTokenFromQueryString(): void {
  const location = useLocation();
  const { [QUERY_STRING_KEY]: token } = parse(location.search);

  if (typeof token === "string") {
    setToken(token);
    window.location.href = "/";
  }
}
