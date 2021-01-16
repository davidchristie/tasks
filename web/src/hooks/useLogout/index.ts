export default function useLogout() {
  return () => {
    localStorage.clear();
    document.location.href = "/";
  };
}
