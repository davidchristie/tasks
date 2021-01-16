import { useEffect, useState } from "react";
import { User } from "../../types";
import useAccessToken from "../useAccessToken";

interface Result {
  data: User | null;
  error: Error | null;
  loading: boolean;
}

export default function useLoggedInUser(): Result {
  const accessToken = useAccessToken();
  const [data, setData] = useState<User | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    if (accessToken) {
      setLoading(true);
      fetch("https://api.github.com/user", {
        headers: {
          Authorization: "token " + accessToken,
        },
      })
        .then((response) => response.json())
        .then((data) => ({
          avatarUrl: data.avatar_url,
          name: data.name,
        }))
        .then((data) => {
          setData(data);
          setError(null);
          setLoading(false);
        })
        .catch((error) => {
          setData(null);
          setError(error);
          setLoading(false);
        });
    }
  }, [accessToken]);

  return { data, error, loading };
}
