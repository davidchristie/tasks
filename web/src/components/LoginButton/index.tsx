export default function GithubLoginButton() {
  const gateway = process.env.REACT_APP_GATEWAY || ""
  return <a href={`${gateway}/login/github`}>Login with GitHub</a>
}
