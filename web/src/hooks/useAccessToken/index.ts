import { parse } from "query-string";
import { useLocation } from "react-router-dom";

export default function useAccessToken() {
  const location = useLocation();
  const { access_token: accessToken } = parse(location.search);
  return accessToken;
}
