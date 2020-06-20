import { signout } from './api-auth';

const auth = {
  isAuthenticated(): boolean {
    if (typeof window == "undefined")
      return false;
    if (sessionStorage.getItem("cookie"))
      return JSON.parse(sessionStorage.getItem("cookie"));
    else
      return false;
  },
  authenticate(cookie, cb): void {
    if (typeof window !== "undefined")
      sessionStorage.setItem("cookie", JSON.stringify(cookie));
    cb();
  },
  clearCookie(cb): void {
    if (typeof window !== "undefined")
      sessionStorage.removeItem("cookie");
    cb();
    //optional
    signout().then((data) => {
      document.cookie = "t=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;"
    });
  },
};

export default auth
