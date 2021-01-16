const GATEWAY = process.env.REACT_APP_GATEWAY || "";

export default function useLogin() {
  return () => {
    document.location.href = `${GATEWAY}/login/github`;
  };
}
