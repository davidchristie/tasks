import { useEffect, useState } from 'react'
import useAccessToken from '../../hooks/useAccessToken'

export default function UserData() {
  const accessToken = useAccessToken()
  const [userData, setUserData] = useState<unknown>()

  useEffect(() => {
    if (accessToken) {
      fetch("https://api.github.com/user", {
        headers: {
          Authorization: "token " + accessToken,
        },
      })
        .then(response => response.json())
        .then(setUserData)
        .catch(() => {
          setUserData(undefined)
        })
    }
  }, [accessToken])

  if (!userData) {
    return null
  }

  return <div>{JSON.stringify(userData, null, 2)}</div>
}
